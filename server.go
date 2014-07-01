package main

import (
	"bytes"
	"time"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

var server = &command{
	name:        "server",
	handler:     startServer,
	requireAuth: true,
	usage:       `sl server [--port]`,
}

func startServer(_ *command, _ []string) {
	fmt.Printf("Listening on localhost%s (^C to exit)\n", ":9292")

	mux := http.NewServeMux()
	mux.Handle("/", &renderer{})

	http.ListenAndServe(":9292", mux)
}

// TODO: Move this stuff into a seperate file.
// Run the old test suite on it.
// Clean up these things. They are old.

type renderer struct{}

func (r *renderer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	tmpl := &template{req.URL.Path}
	tmpl.ServeHTTP(rw, req)
}

var liquidRegex = regexp.MustCompile(`{%\s*include\s*['"](.*)['"]\s*%}`)

type template struct {
	path string
}

func (t *template) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	filename, err := t.filename()

	if err != nil {
		printError("No file found for: " + t.path)
		http.Error(rw, err.Error(), 404)
		return
	}

	if ext := path.Ext(filename); ext == ".liquid" || ext == ".html" {
		var rendered bytes.Buffer

		template := t.renderLiquid(filename)

		params := struct {
			Url      string
			Template string
		}{t.path, template}

		// TODO: load the site id before getting here
		err := client.Post("sites/SITE_ID_GOES_HERE/preview", &rendered, &params)

		if err != nil {
			printFatal(err.Error())
		}

		// These branches return almost the same thing. Simplify it.
		http.ServeContent(rw, req, t.path, time.Now(), strings.NewReader(rendered.String()))
	} else {
		http.ServeContent(rw, req, t.path, time.Now(), strings.NewReader(readFile(filename)))
	}
}

func readFile(filename string) string {
	if body, err := ioutil.ReadFile(filename); err != nil {
		return err.Error()
	} else {
		return string(body[:])
	}
}

func (t *template) filename() (str string, err error) {
	for _, filename := range t.potentialFiles() {
		fs, err := os.Stat(filename)
		if !os.IsNotExist(err) || (fs != nil && !fs.IsDir()) {
			return filename, nil
		}
	}
	err = &templateError{t.potentialFiles()}
	return
}

func (t *template) renderLiquid(filename string) (template string) {
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		printFatal(err.Error())
	}

	template = string(body[:])
	matched, _ := regexp.MatchString(`{{.*}}|{%.*%}`, template)

	if matched {
		return liquidRegex.ReplaceAllStringFunc(template, func(match string) string {
			filename := liquidRegex.FindStringSubmatch(match)
			return t.renderLiquid("_" + filename[1] + ".html")
		})
	}

	return
}

func (t *template) potentialFiles() (files []string) {

	// Strip leading/trailing slashes
	tpath := strings.TrimPrefix(path.Join(t.path), "/")

	// Catch static assets and liquid templates
	// Also the top level named .html file
	if tpath != "" {
		files = append(files, tpath)
		files = append(files, tpath+".liquid")
		files = append(files, tpath+".html")
	}

	// Add the top level index and default templates
	files = append(files, path.Join(tpath, "index.html"))
	files = append(files, path.Join(tpath, "default.html"))

	// Split path into remaining parent directories and
	// loop through them in reverse to add default templates
	dir, _ := path.Split(tpath)
	dirs := strings.Split(dir, "/")
	for index := range dirs {
		// We added the nested default.html already
		if index == 0 {
			continue
		}
		element := append(dirs[:len(dirs)-index], "default.html")
		files = append(files, path.Join(element...))
	}

	// Add the base default.html if were not on the index
	if tpath != "" {
		files = append(files, "default.html")
	}

	return
}

type templateError struct {
	Filenames []string
}

func (e *templateError) Error() string {
	return strings.Join(e.messages(), "\n")
}

func (e *templateError) messages() (messages []string) {
	messages = append(messages, "No file found in:\n"+strings.Join(e.Filenames, "\n"))
	return
}

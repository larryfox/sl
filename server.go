package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/larryfox/sl/template"
)

var cmdServer = &command{
	Name:        "server",
	handler:     runServer,
	requireAuth: true,
	requireSite: true,
	usage:       `sl server [--port]`,
}

func runServer(_ *command, _ []string) {
	fmt.Printf("Listening on localhost%s (ctrl+C to exit)\n", ":9292")

	mux := http.NewServeMux()
	mux.Handle("/", serveLocalFile(templateHandler))
	http.ListenAndServe(":9292", mux)
}

func serveLocalFile(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		f := strings.Trim(req.URL.Path, "/")

		if !isLiquid(f) && localFileExists(f) {
			http.ServeFile(w, req, f)
		} else {
			next(w, req)
		}
	}
}

func templateHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.New(req.URL.Path)
	if err != nil {
		printWarning(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	body, err := renderTemplate(req.URL.Path, tmpl)
	if err != nil {
		printWarning(err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	if !strings.HasPrefix(req.URL.Path, "/assets/") {
		fmt.Printf("    %s -> %s\n", req.URL.Path, tmpl.Filename())
	}

	body.WriteTo(w)
}

func renderTemplate(path string, tmpl *template.Template) (bytes.Buffer, error) {
	var rendered bytes.Buffer

	params := struct {
		Url      string
		Template string
	}{path, tmpl.String()}

	// TODO: load the site id before getting here
	// FIXME: increase the timeout in the client
	err := client.Post("sites/"+currentSite.Id+"/preview", &rendered, &params)

	return rendered, err
}

func localFileExists(f string) bool {
	fs, err := os.Stat(f)
	return !os.IsNotExist(err) && (fs != nil && !fs.IsDir())
}

func isLiquid(f string) bool {
	ext := path.Ext(f)
	return ext == ".liquid" || ext == ".html"
}

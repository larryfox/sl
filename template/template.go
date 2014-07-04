package template

import (
	"bytes"
	"io/ioutil"
	"log"
	"path"
	"regexp"
)

var liquidRegex = regexp.MustCompile(`{%\s*include\s*['"](.*)['"]\s*%}`)

// A Template represents a resource to be sent to a browser.
//
// It implements io.ReadSeeker so it can be passed directly
// to an http response.
type Template struct {
	Filename string
	content  *bytes.Reader
}

// New returns a new Template given a path string.
func New(p string) (*Template, error) {
	filename, err := ResolvePath(p)

	log.Println(filename)

	if err != nil {
		return nil, err
	}

	tmpl := &Template{
		Filename: filename,
		content:  newReader(filename),
	}

	return tmpl, nil
}

// Read implements the io.Reader interface by
// delegating to the content field.
func (t *Template) Read(b []byte) (int, error) {
	return t.content.Read(b)
}

// Seek implements the io.Seeker interface by
// delegating to the content field.
func (t *Template) Seek(offset int64, whence int) (int64, error) {
	return t.content.Seek(offset, whence)
}

func newReader(filename string) *bytes.Reader {
	var content []byte

	switch path.Ext(filename) {
	case ".liquid", ".html":
		content = readLiquid(filename)
	default:
		content = readFile(filename)
	}

	return bytes.NewReader(content)
}

func readLiquid(filename string) []byte {
	body := readFile(filename)
	return liquidRegex.ReplaceAllFunc(body, replaceIncludes)
}

func replaceIncludes(match []byte) []byte {
	filename := liquidRegex.FindSubmatch(match)[1]
	return readLiquid("_" + string(filename[:]) + ".html")
}

func readFile(filename string) []byte {
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Println(err.Error())
	}

	return body
}

package template

import (
	"bytes"
	"io/ioutil"
	"log"
	"regexp"
)

var liquidRegex = regexp.MustCompile(`{%\s*include\s*['"](.*)['"]\s*%}`)

// A Template represents a resource to be sent to a browser.
//
// It implements io.ReadSeeker so it can be passed directly
// to an http response.
type Template struct {
	*bytes.Buffer
	filename string
}

// New returns a new Template given a path string.
func New(p string) (*Template, error) {
	filename, err := ResolvePath(p)

	if err != nil {
		return nil, err
	}

	tmpl := &Template{
		newBuffer(filename),
		filename,
	}

	return tmpl, nil
}

func (t *Template) Filename() string {
	return t.filename
}

func newBuffer(filename string) *bytes.Buffer {
	return bytes.NewBuffer(readLiquid(filename))
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
		log.Panicln(err.Error())
	}

	return body
}

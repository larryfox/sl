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
	filename string
	content  *bytes.Reader
}

// New returns a new Template given a path string.
func New(p string) (*Template, error) {
	filename, err := ResolvePath(p)

	if err != nil {
		return nil, err
	}

	tmpl := &Template{
		filename: filename,
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

func (t *Template) String() string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(t)
	return buf.String()
}

func (t *Template) Filename() string {
	return t.filename
}

func newReader(filename string) *bytes.Reader {
	return bytes.NewReader(readLiquid(filename))
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

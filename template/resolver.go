package template

import (
	"os"
	"path"
	"strings"
)

// ResolvePath takes a path string and returns the
// coresponding filename of the template on the
// system or an error if no template is found.
func ResolvePath(path string) (string, error) {
	files := potentialFiles(path)

	existing := func(f string) bool {
		fs, err := os.Stat(f)
		return !os.IsNotExist(err) && (fs != nil && !fs.IsDir())
	}

	return find(existing, files)
}

type NoTemplate struct {
	Files []string
}

func (e *NoTemplate) Error() string {
	return "No template found in:\n" + strings.Join(e.Files, "\n")
}

func potentialFiles(p string) (files []string) {
	// Strip leading/trailing slashes
	tpath := strings.TrimPrefix(path.Join(p), "/")

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

func find(found func(string) bool, arr []string) (str string, err error) {
	for _, x := range arr {
		if found(x) {
			return x, nil
		}
	}
	err = &NoTemplate{arr}
	return
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/larryfox/sl/template"
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
	mux.HandleFunc("/", previewServer)
	http.ListenAndServe(":9292", mux)
}

func previewServer(rw http.ResponseWriter, req *http.Request) {
	tmpl, err := template.New(req.URL.Path)

	if err != nil {
		printWarning(err.Error())
		http.Error(rw, err.Error(), 500)
	} else {
		http.ServeContent(rw, req, req.URL.Path, time.Now(), tmpl)
	}
}

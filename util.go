package main

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"
)

func printError(message string, args ...interface{}) {
	log.Println(prefixWith("error:", fmt.Sprintf(message, args...)))
}

func printFatal(message string, args ...interface{}) {
	log.Fatalln(prefixWith("fatal:", fmt.Sprintf(message, args...)))
}

func printWarning(message string, args ...interface{}) {
	log.Println(prefixWith("warning:", fmt.Sprintf(message, args...)))
}

func prefixWith(a, b string) string {
	return a + " " + b
}

func currentUser() (usr *user.User) {
	usr, err := user.Current()

	if err != nil {
		printFatal("%v", err)
	}

	return
}

func absPath(paths ...string) string {
	path, _ := filepath.Abs(filepath.Join(paths...))
	return path
}

package main

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"
)

func printError(message string, args ...interface{}) {
	log.Println("error: " + fmt.Sprintf(message, args...))
}

func printFatal(message string, args ...interface{}) {
	log.Fatalln("fatal: " + fmt.Sprintf(message, args...))
}

func printWarning(message string, args ...interface{}) {
	log.Println("warning: " + fmt.Sprintf(message, args...))
}

func systemUser() (usr *user.User) {
	usr, err := user.Current()
	checkFatal(err)
	return
}

func absPath(paths ...string) string {
	path, _ := filepath.Abs(filepath.Join(paths...))
	return path
}

func checkFatal(err error) {
	if err != nil {
		printFatal("%v", err)
	}
}

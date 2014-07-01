package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/bgentry/speakeasy"
	"github.com/larryfox/prompt"
)

var auth = &command{
	name:        "auth",
	handler:     authenticateUser,
	requireAuth: false,
}

var siteleafrc = absPath(currentUser().HomeDir, ".siteleafrc")

func authenticateUser(_ *command, _ []string) {
	user := prompt.Ask("Enter user email: ")
	pass, _ := speakeasy.Ask("Enter user password: ")

	err := client.Auth(user, pass)

	if err != nil {
		printError("Could not authenticate :(")
	}

	log.Println("Authorized!")
	saveAuthFile()
}

func saveAuthFile() {
	file, err := os.Create(siteleafrc)

	if err != nil {
		printFatal("Could not open or create %v for writing.", siteleafrc)
	}

	file.WriteString(client.ApiKey + "\n" + client.ApiSecret)
}

func loadAuthFile() {
	content, err := ioutil.ReadFile(siteleafrc)

	if err != nil {
		printError("Could not open %v, have you ran 'sl auth' yet?", siteleafrc)
	}

	lines := strings.Split(string(content), "\n")

	if len(lines) < 2 || len(lines[0]) == 0 || len(lines[1]) == 0 {
		printError("Not authorized, have you ran 'sl auth' yet?")
	}

	client.ApiKey = lines[0]
	client.ApiSecret = lines[1]
}

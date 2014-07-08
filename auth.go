package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/bgentry/speakeasy"
	"github.com/larryfox/go-prompt"
	"gopkg.in/yaml.v1"
)

var cmdAuth = &command{
	Name:        "auth",
	handler:     runAuth,
	requireAuth: false,
}

func runAuth(_ *command, _ []string) {
	email, _ := prompt.Ask("Enter user email: ")
	pass, _ := speakeasy.Ask("Enter user password: ")

	err := client.Auth(email, pass)

	if err != nil {
		printError("Could not authenticate :(")
	}

	writeAuthFile()

	log.Println("Authorized!")
}

type auth struct {
	ApiKey    string `api_key`
	ApiSecret string `api_secret`
}

func writeAuthFile() {
	file, err := os.Create(siteleafYAMLPath)
	if err != nil {
		printFatal("Could not open or create %v for writing.", siteleafYAMLPath)
	}

	yml, err := yaml.Marshal(&auth{client.ApiKey, client.ApiSecret})
	checkFatal(err)
	file.Write(yml)
}

func readAuthFile() {
	var conf auth

	content, err := ioutil.ReadFile(siteleafYAMLPath)
	if err != nil {
		printFatal("Could not open %v, have you ran 'sl auth' yet?", siteleafYAMLPath)
	}

	err = yaml.Unmarshal(content, &conf)
	checkFatal(err)

	client.ApiKey = conf.ApiKey
	client.ApiSecret = conf.ApiSecret
}

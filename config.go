package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v1"
)

var cmdConfig = &command{
	Name:        "config",
	handler:     runConfig,
	requireAuth: true,
	usage:       `sl config <domain> [<dir>]`,
}

func runConfig(cmd *command, args []string) {
	var path string

	if len(args) < 2 {
		printError("Please provide a domain name.")
		cmd.PrintUsage()
		os.Exit(2)
	}

	site, err := client.Sites().FindByDomain(args[1])

	if err != nil {
		printError("No site found with domain: %s", args[1])
		os.Exit(1)
	}

	if len(args) > 2 {
		path = absPath(args[2], configYAML)
	} else {
		path = absPath(args[1], configYAML)
	}

	writeSiteConfigFile(site.Id, path)
}

type site struct {
	Id   string `site_id`
	Port string `port,omitempty`
}

func writeSiteConfigFile(id, path string) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		printFatal("Could not create directory %v.", path)
	}

	file, err := os.Create(path)
	if err != nil {
		printFatal("Could not open or create %v for writing.", path)
	}

	yml, err := yaml.Marshal(&site{id, ""})
	checkFatal(err)
	file.Write(yml)
}

func readSiteConfigFile() {
	filepath := absPath(configYAML)
	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		printFatal("Could not open %v, have you ran 'sl config' yet?", filepath)
	}

	err = yaml.Unmarshal(content, &currentSite)
	checkFatal(err)
}

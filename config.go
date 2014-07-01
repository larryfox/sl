package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var config = &command{
	name:        "config",
	handler:     configureSite,
	requireAuth: true,
	usage:       `sl config <domain> [<dir>]`,
}

func configureSite(cmd *command, args []string) {
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
		path = absPath(args[2], ".siteleaf")
	} else {
		path = absPath(args[1], ".siteleaf")
	}

	fmt.Println(path)
	saveSiteConfigFile(site.Id, path)
}

func saveSiteConfigFile(id, path string) {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		printFatal("Could not create directory %v.", path)
	}

	file, err := os.Create(path)

	if err != nil {
		printFatal("Could not open or create %v for writing.", path)
	}

	file.WriteString(id)
}

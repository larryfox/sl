package main

import (
	"fmt"
	"os"
	"strings"
)

var cmdNew = &command{
	Name:        "new",
	handler:     runNew,
	requireAuth: true,
	usage:       `sl new [--title] <domain> [<dir>]`,
}

func runNew(cmd *command, args []string) {
	var path string
	var title string

	if len(args) < 2 {
		printError("Please provide a domain name.")
		cmd.PrintUsage()
		os.Exit(2)
	}

	// TODO: Let users specify a title `sl new --title "My Site" example.com`
	// Also need to come up with a better form of error handling from the api
	// client, casting the error to a string, and matching on that message
	// feels gross but works for nowish. ¯\_(ツ)_/¯
	site, err := client.CreateSite(title, args[1])
	if err != nil {
		if strings.Contains(err.Error(), "domain is taken") {
			printError("Domain is already taken. Try 'sl config %s'.", args[1])
		} else {
			printFatal("%v", err)
		}
		os.Exit(1)
	}

	if len(args) > 2 {
		path = absPath(args[2], ".siteleaf")
	} else {
		path = absPath(args[1], ".siteleaf")
	}

	fmt.Println(path)
	writeSiteConfigFile(site.Id, path)
}

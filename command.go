package main

import (
	"log"
)

type command struct {
	Name        string
	handler     func(*command, []string)
	requireAuth bool
	requireSite bool
	usage       string
	summary     string
}

func (cmd *command) Handler(args []string) {
	if cmd.requireAuth {
		readAuthFile()
	}

	if cmd.requireSite {
		readSiteConfigFile()
	}

	cmd.handler(cmd, args)
}

func (cmd *command) PrintUsage() {
	log.Printf("usage: %s", cmd.usage)
}

func (cmd *command) PrintSummary() {
	cmd.PrintUsage()
	log.Println(cmd.summary)
}

package main

import (
	"log"
)

type command struct {
	name        string
	handler     func(*command, []string)
	requireAuth bool
	usage       string
	summary     string
}

func (cmd *command) Handler(args []string) {
	if cmd.requireAuth {
		loadAuthFile()
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

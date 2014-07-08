package main

import (
	"fmt"
)

var cmdWhoami = &command{
	Name:        "whoami",
	handler:     runWhoami,
	requireAuth: true,
}

func runWhoami(_ *command, _ []string) {
	user := client.Users().Me()
	fmt.Printf("You are %s <%s>", user.FullName, user.Email)
}

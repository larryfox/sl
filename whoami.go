package main

import (
	"fmt"
)

var whoami = &command{
	name:        "whoami",
	handler:     printUserInfo,
	requireAuth: true,
}

func printUserInfo(_ *command, _ []string) {
	user := client.Users().Me()
	fmt.Printf("You are %s <%s>", user.FullName, user.Email)
}

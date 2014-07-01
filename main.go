package main

import (
	"log"
	"os"

	"github.com/larryfox/siteleaf-go"
)

var client = siteleaf.NewClient()

func main() {
	log.SetFlags(0)

	args := os.Args[1:]

	if len(args) == 0 {
		help.PrintSummary()
		os.Exit(2)
	}

	switch args[0] {
	default:
		log.Printf("Unknown option: %s. See 'sl help'.", args[0])
		os.Exit(2)
	case help.name:
		help.Handler(args)
	case auth.name:
		auth.Handler(args)
	case config.name:
		config.Handler(args)
	case newCmd.name:
		newCmd.Handler(args)
	case server.name:
		server.Handler(args)
	case "push":
		printWarning("TODO Theme Push")
	case "pull":
		printWarning("TODO Theme Pull")
	case "open":
		printWarning("TODO Open")
	case whoami.name:
		whoami.Handler(args)
	}
}

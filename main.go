package main

import (
	"log"
	"os"

	"github.com/larryfox/siteleaf-go"
)

const (
	configYAML   = ".config.yml"
	siteleafYAML = ".siteleaf.yml"
)

var (
	client           = siteleaf.NewClient()
	siteleafYAMLPath = absPath(systemUser().HomeDir, siteleafYAML)
	currentSite      site
)

func main() {
	log.SetFlags(0)

	args := os.Args[1:]

	if len(args) == 0 {
		cmdHelp.PrintSummary()
		os.Exit(2)
	}

	switch args[0] {
	default:
		log.Printf("Unknown option: %s. See 'sl help'.", args[0])
		os.Exit(2)
	case cmdHelp.Name:
		cmdHelp.Handler(args)
	case cmdAuth.Name:
		cmdAuth.Handler(args)
	case cmdConfig.Name:
		cmdConfig.Handler(args)
	case cmdNew.Name:
		cmdNew.Handler(args)
	case cmdServer.Name:
		cmdServer.Handler(args)
	case "push":
		printWarning("TODO Theme Push")
	case "pull":
		printWarning("TODO Theme Pull")
	case "open":
		printWarning("TODO Open")
	case cmdWhoami.Name:
		cmdWhoami.Handler(args)
	}
}

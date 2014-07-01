package main

var help = &command{
	name:        "help",
	handler:     printHelp,
	requireAuth: false,
	usage:       `sl [--version] <command> [<args>]`,
	summary: `
Commonly used sl commands are:
    auth    Authorize sl with your Siteleaf credentials
    config  Setup an existing Siteleaf site in a directory
    server  Start a local Siteleaf server
    help    Prints this message

Use 'sl help <command>' to read about specific subcommands.
`,
}

func printHelp(cmd *command, _ []string) {
	cmd.PrintSummary()
}

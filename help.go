package main

var cmdHelp = &command{
	Name:        "help",
	handler:     runHelp,
	requireAuth: false,
	usage:       `sl [--version] <command> [<args>]`,
	summary: `
Commonly used sl commands are:
    auth    Authorize sl with your Siteleaf credentials
    new     Create and configure a new site
    config  Setup an existing Siteleaf site in a directory
    server  Start a local Siteleaf server
    help    Prints this message

Use 'sl help <command>' to read about specific subcommands.
`,
}

func runHelp(cmd *command, _ []string) {
	cmd.PrintSummary()
}

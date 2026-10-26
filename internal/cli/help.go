package cli

import "fmt"

var helpText = `Usage: tsd [flags] <config_file>
       tsd init <config_file>

Commands:
    init <config_file>
        Create a new configuration file.
        usage: tsd init <filename.hcl>

Flags:
    --show
        Show the passed configuration.
        usage: tsd --show <config_file>

    --attach <session>
        Attach to a session after setup.
        usage: tsd --attach <session-name> <config_file>

    --single
        Attach to the session after setup.
        Works only when exactly one session is defined.
        usage: tsd --single <config_file>

    --help
        Print help.
        usage: tsd --help`

func PrintHelp() {
	fmt.Println(helpText)
}

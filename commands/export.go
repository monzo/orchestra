package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/monzo/orchestra/config"
	"github.com/wsxiaoys/terminal"
)

var ExportCommand = &cli.Command{
	Name:         "export",
	Usage:        "Export those *#%&! env vars ",
	Action:       BeforeAfterWrapper(ExportAction),
	BashComplete: ServicesBashComplete,
}

func ExportAction(c *cli.Context) error {
	for key, value := range config.GetBaseEnvVars() {
		terminal.Stdout.Print(fmt.Sprintf("export %s=%s\n", key, value))
	}
	return nil
}

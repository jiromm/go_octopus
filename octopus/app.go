package octopus

import (
	"os"
	"github.com/urfave/cli"
)

func Run() {

	app := cli.NewApp()
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:    "run",
			Description: "Run test job",
			Aliases: []string{"r"},
			Usage:   "run, r",
			Action:  func(c *cli.Context) error {
				RunTestJob()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
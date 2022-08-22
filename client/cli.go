package client

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Client() {
	app := &cli.App{
		Name:  "HeLang",
		Usage: "The Next-Generation Cyber Programming Language from Li Tang",
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run your code",
				Action: runCodes,
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build your code",
				ArgsUsage: "*.he",
				Action: buildCodes,
				Flags: []cli.Flag {
					&cli.StringFlag{
						Name: "output",
						Usage: "output file",
						Aliases: []string { "o" },
					},
				},
			},
			{
				Name:    "shell",
				Aliases: []string{"s"},
				Usage:   "Run HeLang shell",
				Action: runShell,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
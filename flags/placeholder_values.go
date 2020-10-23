package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//NAME:
//placeholder_values - A new cli application
//
//USAGE:
//placeholder_values [global options] command [command options] [arguments...]
//
//COMMANDS:
//help, h  Shows a list of comma

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

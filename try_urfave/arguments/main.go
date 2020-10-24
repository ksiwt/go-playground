package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//NAME:
//   main - A new cli application
//
//USAGE:
//   main [global options] command [command options] [arguments...]
//
//COMMANDS:
//   help, h  Shows a list of commands or help for one command
//
//GLOBAL OPTIONS:
//   --help, -h  show help (default: false)

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello %q", c.Args().Get(0))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//NAME:
//main - A new cli application
//
//USAGE:
//main [global options] command [command options] [arguments...]
//
//COMMANDS:
//help, h  Shows a list of commands or help for one command
//
//GLOBAL OPTIONS:
//--lang value  language for the greeting (default: "English")
//--help, -h    show help (default: false)

func main() {
	// You can also set a destination variable for a flag, to which the content will be scanned.
	//var language string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "English",
				Usage: "language for the greeting",
				//Destination: &language,
			},
		},
		Action: func(c *cli.Context) error {
			name := "Nefertiti"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

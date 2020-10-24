package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	numFlagStr string = "n"
	isShowNum  bool   = false
)

func main() {
	app := cli.NewApp()
	app.Flags = setFlags()
	app.Action = setAction()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setFlags() []cli.Flag {
	Flags := []cli.Flag{
		&cli.StringSliceFlag{
			Name:  numFlagStr,
			Usage: "Read line number from `files`",
		},
	}

	return Flags
}

func setAction() cli.ActionFunc {
	return func(c *cli.Context) error {
		args := c.Args().Slice()

		if c.IsSet(numFlagStr) {
			isShowNum = true
			args = c.StringSlice(numFlagStr)
		}

		files, err := findFiles(args)
		if err != nil {
			return err
		}

		for _, file := range files {
			readFile(file, isShowNum)
		}

		return nil
	}
}

func findFiles(args []string) ([]*os.File, error) {
	var files = make([]*os.File, 0, len(args))
	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

func readFile(
	file *os.File,
	isShowNum bool,
) {
	fileScanner := bufio.NewScanner(file)

	line := 0
	if isShowNum {
		for fileScanner.Scan() {
			line += 1
			fmt.Printf("%d: %s\n", line, fileScanner.Text())
		}
	} else {
		for fileScanner.Scan() {
			fmt.Println(fileScanner.Text())
		}
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	defer file.Close()
}

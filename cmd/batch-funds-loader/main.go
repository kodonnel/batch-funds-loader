package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "batch funds loader",
		Usage: "Load funds from a batch file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "Load funds from `FILE`",
				Value:   "input.txt",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output load results to `FILE`",
				Value:   "output.txt",
			},
		},

		Action: func(c *cli.Context) error {
			fmt.Println("Hello World")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

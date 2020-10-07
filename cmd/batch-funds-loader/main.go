package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	data "github.com/kodonnel/batch-funds-loader/internal/data"
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
			// read the lines
			input := c.String("input")
			output := c.String("output")
			processFile(input, output)

			// process the lines
			// write the lines
			fmt.Println(hello())
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func processFile(fname string, output string) {
	f, err := os.Open(fname)
	defer f.Close()

	if err != nil {
		// handle error
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		var v data.Load
		if err := json.Unmarshal(s.Bytes(), &v); err != nil {
			//handle error
		}
		// do something with v
		writeOutput(output, v)
		fmt.Println(v.CustomerID)
		fmt.Println(v.LoadAmount)
	}
	if s.Err() != nil {
		// handle scan error
	}
}

func writeOutput(fname string, load data.Load) {
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	var loadResult *data.LoadResult
	loadResult = new(data.LoadResult)
	loadResult.CustomerID = load.CustomerID
	loadResult.ID = load.ID
	loadResult.Accepted = true

	b, err := json.Marshal(loadResult)
	_, _ = datawriter.Write(b)
	datawriter.WriteString("\n")

	datawriter.Flush()
	file.Close()
}

func hello() string {
	return "Hello World"
}

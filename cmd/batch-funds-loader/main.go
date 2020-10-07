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

			// channel to communicate between goroutines
			msg := make(chan data.LoadResult)
			done := make(chan bool)

			go processFile(input, msg)
			go writeOutput(output, msg, done)

			// wait until writing is done to exit
			<-done
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

func processFile(fname string, msg chan<- data.LoadResult) {
	f, err := os.Open(fname)

	if err != nil {
		// handle error
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var v data.Load
		if err := json.Unmarshal(s.Bytes(), &v); err != nil {
			//handle error
		}
		fmt.Println(v.CustomerID)
		fmt.Println(v.LoadAmount)

		// MOVE TO HANDLER CODE

		// check if duplicate
		exists := data.CheckLoadIDExistsForCustomer(v.ID, v.CustomerID)

		if exists == false {

			// check new load validity
			// data.CheckDailyLimitExceededForCustomer
			// data.CheckWeeklyLimitExceededForCustomer
			// data.CheckMaxLoadsExceededForCustomer

			var vl data.VelocityLimit
			vl.CustomerID = v.CustomerID
			vl.DailyAmount = 1000
			vl.WeeklyAmount = 1000
			vl.DailyLoads = 1
			vl.LoadIDs = append(vl.LoadIDs, v.ID)

			data.AddVelocityLimit(vl)

			// if valid add VL and return success
			var loadResult *data.LoadResult
			loadResult = new(data.LoadResult)
			loadResult.CustomerID = v.CustomerID
			loadResult.ID = v.ID
			loadResult.Accepted = true

			// else return decline

			msg <- *loadResult
		}

	}

	close(msg)

	if s.Err() != nil {
		// handle scan error
	}
}

func writeOutput(fname string, msg <-chan data.LoadResult, done chan<- bool) {
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for {
		loadResult, more := <-msg
		if more {
			fmt.Println("received result")
			b, _ := json.Marshal(loadResult)
			_, _ = datawriter.Write(b)
			datawriter.WriteString("\n")
		} else {
			fmt.Println("done receiving results")
			datawriter.Flush()
			file.Close()
			done <- true
			close(done)
			return
		}
	}

}

func hello() string {
	return "Hello World"
}

package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	data "github.com/kodonnel/batch-funds-loader/internal/data"
	"github.com/kodonnel/batch-funds-loader/internal/handlers"
	"github.com/sirupsen/logrus"
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
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Enable verbose logging output to stdout",
			},
		},

		Action: func(c *cli.Context) error {
			// get flags from context
			input := c.String("input")
			output := c.String("output")
			verbose := c.Bool("verbose")

			// channel to communicate between goroutines
			msg := make(chan data.LoadResult)
			done := make(chan bool)

			// use the standard logger
			//logger := log.New(os.Stdout, "batch-funds-loader", log.LstdFlags)
			logger := &logrus.Logger{
				Out:   os.Stdout,
				Level: logrus.DebugLevel,
				Formatter: &logrus.TextFormatter{
					DisableColors:   true,
					TimestampFormat: "2006-01-02 15:04:05",
					FullTimestamp:   true,
				},
			}

			// only show logs if verbose is set
			if !verbose {
				logger.SetOutput(ioutil.Discard)
			}

			// create db instance
			db := data.NewLoadsDB(logger)

			// req handlers
			loadsHandler := handlers.NewLoads(logger, db)

			go processFile(input, msg, loadsHandler)
			go writeOutput(output, msg, done)

			// wait until writing is done to exit
			<-done
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func processFile(fname string, msg chan<- data.LoadResult, lh *handlers.Loads) {
	f, err := os.Open(fname)

	if err != nil {
		// handle error
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var l data.Load
		if err := json.Unmarshal(s.Bytes(), &l); err != nil {
			//handle error
		}

		lh.ProcessLoadRequest(l, msg)
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
			// received a new LoadRequest
			b, _ := json.Marshal(loadResult)
			_, _ = datawriter.Write(b)
			datawriter.WriteString("\n")
		} else {

			// finished receiving LoadRequests
			datawriter.Flush()
			file.Close()
			done <- true
			close(done)
			return
		}
	}

}

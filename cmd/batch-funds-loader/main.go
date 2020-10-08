package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/go-playground/validator"
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
			msg := make(chan data.Load)
			done := make(chan bool)

			// use the standard logger
			//logger := log.New(os.Stdout, "batch-funds-loader", log.LstdFlags)
			logger := &logrus.Logger{
				Out:   os.Stdout,
				Level: logrus.InfoLevel,
				Formatter: &logrus.TextFormatter{
					DisableColors:   true,
					TimestampFormat: "2006-01-02 15:04:05",
					FullTimestamp:   true,
				},
			}

			// only show logs if verbose is set
			if !verbose {
				logger.SetLevel(logrus.FatalLevel)
			}

			// create db instance
			db := data.NewLoadsDB(logger)

			// create validator
			v := validator.New()
			v.RegisterValidation("loadAmount", data.ValidateLoadAmount)
			v.RegisterValidation("identifier", data.ValidateID)

			// req handlers
			loadsHandler := handlers.NewLoads(logger, db, v)

			go processFile(logger, input, msg, loadsHandler)
			go writeOutput(logger, output, msg, done)

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

func processFile(log *logrus.Logger, fname string, msg chan<- data.Load, lh *handlers.Loads) {
	f, err := os.Open(fname)

	if err != nil {
		// handle error
		log.Fatalln("unable to process file,", err)
	}

	// make sure filehandle is closed
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var l data.Load
		if err := json.Unmarshal(s.Bytes(), &l); err != nil {
			// log error and continue to the next item
			log.Errorln("unable to process load request,", err)
			continue
		}

		loadResult, err := lh.ProcessLoadRequest(l)
		if err != nil {
			// log error and continue to the next item
			log.Errorln("unable to process load request", l, err)
			continue
		}

		// write the loadResult to the lrChannel
		msg <- *loadResult
	}

	close(msg)

	if s.Err() != nil {
		// handle scan error
		log.Fatalln("unable to process file,", err)
	}
}

func writeOutput(log *logrus.Logger, fname string, msg <-chan data.Load, done chan<- bool) {
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	datawriter := bufio.NewWriter(file)

	// ensure resources are freed, even if errors occured
	defer file.Close()
	defer close(done)

	if err != nil {
		log.Fatalln("unable to write output", err)
	}

	for {
		loadResult, more := <-msg
		if more {
			// received a new LoadRequest
			b, err := data.MarshalJSON(loadResult)
			if err != nil {
				log.Errorln("unable to write load result", err)
			}

			_, err = datawriter.Write(b)
			if err != nil {
				log.Errorln("unable to write load result", err)
				continue
			}

			_, err = datawriter.WriteString("\n")
			if err != nil {
				log.Errorln("unable to write load result", err)
				continue
			}

		} else {

			// finished receiving LoadRequests
			// free resources
			err = datawriter.Flush()
			if err != nil {
				log.Errorln("unable to flush data writer", err)
			}
			done <- true
			return
		}
	}
}

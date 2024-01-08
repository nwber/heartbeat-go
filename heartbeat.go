package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/urfave/cli/v2"
)

func heartbeat(url string, duration int) {
	secondsElapsed := 0

	// gets timestamp in ISO 8601 (in UTC, of course)
	// saves output to timestamp.csv
	now := time.Now().Format(time.RFC3339)
	outfile := now + ".csv"

	file, err := os.Create(outfile)

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	fmt.Println("Writing output to", outfile)

	cron := gocron.NewScheduler(time.UTC)
	cron.Every(1).Seconds().Do(func() {
		if secondsElapsed == duration {
			fmt.Println("Ran for specified duration", duration, "seconds, closing")
			cron.StopBlockingChan()
		}

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		Status := strconv.Itoa(resp.StatusCode)
		fmt.Println("Status code:", Status)

		message := []string{Status, url, time.Now().Format(time.RFC3339)}
		w.Write(message)
		secondsElapsed++
	})
	cron.StartBlocking()
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "endpoint",
				Usage: "checks heartbeat of this FQDN",
			},
			&cli.StringFlag{
				Name:  "duration",
				Usage: "duration in seconds to check heartbeat of this endpoint",
			},
		},
		Name:  "heartbeat",
		Usage: "runs GET for a specified endpoint",
		Action: func(cCtx *cli.Context) error {
			var duration int
			var err error
			var endpoint string = cCtx.String("endpoint")

			if cCtx.String("duration") != "" {
				fmt.Println("Setting duration to:", cCtx.String("duration"))

				duration, err = strconv.Atoi(cCtx.String("duration"))
				if err != nil {
					fmt.Println("Error during conversion")
					return err
				}

			} else {
				duration = 600
			}

			fmt.Println("Checking heartbeat of:", endpoint)
			heartbeat(endpoint, duration)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

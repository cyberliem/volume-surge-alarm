package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
	"go.uber.org/zap"

	libapp "github.com/cyberliem/volume-surge-alarm/app"
	"github.com/cyberliem/volume-surge-alarm/binance"
)

const (
	retryDelayFlag        = "retry-delay"
	tickerIntervalFlag    = "ticker-interval"
	attemptFlag           = "attempt"
	defaultRetryDelay     = 2 * time.Minute
	defaultAttempt        = 4
	defaultTickerInterval = 15 * time.Second
)

func main() {
	app := libapp.NewApp()
	app.Name = "Accounting binance trades fetcher"
	app.Usage = "Fetch and store trades history from binance"
	app.Action = run

	app.Flags = append(app.Flags,
		cli.DurationFlag{
			Name:   retryDelayFlag,
			Usage:  "delay time when do a retry",
			EnvVar: "RETRY_DELAY",
			Value:  defaultRetryDelay,
		},
		cli.DurationFlag{
			Name:   tickerIntervalFlag,
			Usage:  "ticker interval to fetch data",
			EnvVar: "TICKER_INTERVAL",
			Value:  defaultTickerInterval,
		},
		cli.IntFlag{
			Name:   attemptFlag,
			Usage:  "number of time doing retry",
			EnvVar: "ATTEMPT",
			Value:  defaultAttempt,
		},
	)

	app.Flags = append(app.Flags, binance.NewCliFlags()...)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	sugar, flusher, err := libapp.NewSugaredLogger(c)
	if err != nil {
		return err
	}

	defer flusher()

	sugar.Info("initiate fetcher")

	binanceClient, err := binance.NewClientFromContext(c, sugar)
	if err != nil {
		return err
	}

	retryDelay := c.Duration(retryDelayFlag)
	attempt := c.Int(attemptFlag)
	tickerInterval := c.Duration(tickerIntervalFlag)

	ticker := time.NewTicker(tickerInterval)
	for t := range ticker.C {
		sugar.Debugw("getting Bookticker", "time", t)
		data, err := fetch(sugar, binanceClient, attempt, retryDelay)
		if err != nil {
			return err
		}
		sugar.Debugw("bookTicker result", "time", time.Now(), "bookTicker", data)

	}
	return nil

}

func fetch(sugar *zap.SugaredLogger, binanceClient binance.Interface, attempt int, retryDelay time.Duration) ([]binance.BookTicker, error) {
	var (
		result []binance.BookTicker
		err    error
	)
	for at := 0; at < attempt; at++ {
		result, err = binanceClient.GetBookTicker()
		if err == nil {
			return result, nil
		}
		sugar.Debugw("bookTicker error", "error", err)
		time.Sleep(retryDelay)
	}
	return result, err
}

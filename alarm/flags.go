package alarm

import (
	"log"

	"github.com/urfave/cli"
)

const (
	changeThresholdFlag = "change-threshold"
	basesFlag           = "bases"
)

//NewCliFlags return cli flags to configure cex-trade client
func NewCliFlags() []cli.Flag {
	return []cli.Flag{
		cli.Float64Flag{
			Name:   changeThresholdFlag,
			Usage:  "change Threshold for price step checker",
			EnvVar: "CHANGE_THRESHOLD",
		},
		cli.StringSliceFlag{
			Name:   basesFlag,
			Usage:  "base for checker.Default is [ETH,USDT,BTC]",
			EnvVar: "BASES",
		},
	}
}

//NewPriceStepCheckerFromContext return pricestepchecker
//Return nil if there is no price threshold
func NewPriceStepCheckerFromContext(c *cli.Context) *PriceStepChecker {
	threshold := c.Float64(changeThresholdFlag)
	if threshold == 0.0 {
		return nil
	}
	base := c.StringSlice(basesFlag)
	if len(base) == 0 {
		base = []string{
			"ETH",
			"BTC",
			"USDT",
		}
	}
	log.Printf("Threshold is %.2f", threshold)
	return &PriceStepChecker{
		threshold: threshold,
		bases:     base,
	}
}

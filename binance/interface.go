package binance

import (
	"context"

	"github.com/cyberliem/volume-surge-alarm/common"
)

//Interface is interface for binance api client
type Interface interface {
	GetBookTicker() (common.PriceList, error)
}

// Limiter is the resource limiter for accessing Binance API.
type Limiter interface {
	WaitN(context.Context, int) error
}

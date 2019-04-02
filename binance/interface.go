package binance

import (
	"context"
)

//Interface is interface for binance api client
type Interface interface {
	GetBookTicker() ([]BookTicker, error)
}

// Limiter is the resource limiter for accessing Binance API.
type Limiter interface {
	WaitN(context.Context, int) error
}

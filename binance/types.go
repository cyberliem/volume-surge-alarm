package binance

import (
	"strconv"

	"github.com/cyberliem/volume-surge-alarm/common"
)

//BookTicker is the struct to hold bookTicker from server
type BookTicker struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BigQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

func toPriceList(bks []BookTicker) (common.PriceList, error) {
	var result = make(common.PriceList)
	for _, bk := range bks {
		bidPrice, err := strconv.ParseFloat(bk.BidPrice, 64)
		if err != nil {
			return result, err
		}
		askPrice, err := strconv.ParseFloat(bk.AskPrice, 64)
		if err != nil {
			return result, err
		}
		result[bk.Symbol] = (bidPrice + askPrice) / 2.0
	}
	return result, nil
}

package binance

//BookTicker is the struct to hold bookTicker from server
type BookTicker struct {
	Symbol   string  `json:"symbol"`
	BidPrice float64 `json:"bidPrice"`
	BigQty   float64 `json:"bidQty"`
	AskPrice float64 `json:"askPrice"`
	AskQty   float64 `json:"askQty"`
}

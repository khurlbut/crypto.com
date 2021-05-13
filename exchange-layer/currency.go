package main

type currency struct {
	Code      string
	Label     string
	Symbol    string
	Precision int
}

const (
	AAVE = iota
	ADA  = iota
	BTC  = iota
	ETH  = iota
	LINK = iota
	LTC  = iota
	UNI  = iota
	USD  = iota
	XLM  = iota
)

var currencies = map[int]currency{
	ADA:  {"ADA", "Cardano", "", 8},
	AAVE: {"AAVE", "Aave", "", 5},
	BTC:  {"BTC", "Bitcoin", "", 8},
	ETH:  {"ETH", "Etherium", "", 8},
	LINK: {"LINK", "Chainlink", "", 8},
	LTC:  {"LTC", "Litecoin", "", 8},
	UNI:  {"UNI", "Uniswap", "", 8},
	USD:  {"USD", "Dollar", "$", 2},
	XLM:  {"XLM", "Stellar", "", 8},
}

var currencyMap = map[string]int{
	"AAVE": AAVE,
	"ADA":  ADA,
	"BTC":  BTC,
	"ETH":  ETH,
	"LINK": LINK,
	"LTC":  LTC,
	"UNI":  UNI,
	"USD":  USD,
	"XLM":  XLM,
}

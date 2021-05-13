package main

type exchange interface {
	spotPrice() spotPrice
}

const (
	CoinBase = iota
	Bitfinix = iota
)

var exchanges = map[int]string{
	CoinBase: "CoinBase",
	Bitfinix: "Bitfinix",
}

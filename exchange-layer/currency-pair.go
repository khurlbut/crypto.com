package main

import "strings"

var currencyPairSet = map[string]struct{}{
	"BTC-USD":  {},
	"ETH-USD":  {},
	"LTC-USD":  {},
	"ADA-USD":  {},
	"AAVE-USD": {},
	"XLM-USD":  {},
	"UNI-USD":  {},
	"LINK-USD": {},
}

func currencyPair(currencyPair string) (Currency, Currency) {
	_, ok := currencyPairSet[currencyPair]
	if !ok {
		currencyPair = "BTC-USD"
	}
	p := strings.Split(currencyPair, "-")
	return currencies[currencyMap[p[0]]], currencies[currencyMap[p[1]]]
}

func currencyPairs() []string {
	keys := make([]string, len(currencyPairSet))
	i := 0
	for k := range currencyPairSet {
		keys[i] = k
		i++
	}
	return keys
}

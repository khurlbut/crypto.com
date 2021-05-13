package main

import (
	"encoding/json"
	"fmt"
)

func newCoinBaseSpotPriceRequest(pair string) spotPriceRequest {
	coin, quote := currencyPair(pair)
	return spotPriceRequest{
		coin,
		quote,
		fmt.Sprintf("http://api.coinbase.com/v2/prices/%s/spot", pair),
	}
}

type spotPriceResponse struct {
	data `json:"data"`
}

type data struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

func (req spotPriceRequest) spotPrice() spotPrice {
	coin := req.coinCurrency
	quote := req.quoteCurrency
	spr := new(spotPriceResponse)

	json.Unmarshal(req.httpGet(), &spr)

	return spotPrice{
		coin,
		newPrice(spr.Amount, quote),
		exchanges[CoinBase],
	}
}

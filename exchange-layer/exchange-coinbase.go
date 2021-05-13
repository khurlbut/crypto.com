package main

import (
	"encoding/json"
	"fmt"
)

func newCoinBaseSpotPriceRequest(pair string) spotPriceRequest {
	coin, quote := currencyPair(pair)
	return spotPriceRequest{
		CoinCurrency:  coin,
		QuoteCurrency: quote,
		Url:           fmt.Sprintf("http://api.coinbase.com/v2/prices/%s/spot", pair),
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
	coin := req.CoinCurrency
	quote := req.QuoteCurrency
	spr := new(spotPriceResponse)

	json.Unmarshal(req.httpGet(), &spr)

	return spotPrice{
		CoinCurrency: coin,
		Price:        newPrice(spr.Amount, quote),
		Exchange:     exchanges[CoinBase],
	}
}

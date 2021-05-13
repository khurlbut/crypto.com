package main

import (
	"encoding/json"
	"fmt"
)

func newCoinBaseSpotPriceRequest(pair string) SpotPriceRequest {
	coin, quote := currencyPair(pair)
	return SpotPriceRequest{
		CoinCurrency:  coin,
		QuoteCurrency: quote,
		Url:           fmt.Sprintf("http://api.coinbase.com/v2/prices/%s/spot", pair),
	}
}

type spotPriceResponse struct {
	Data `json:"data"`
}

type Data struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

func (req SpotPriceRequest) spotPrice() SpotPrice {
	coin := req.CoinCurrency
	quote := req.QuoteCurrency
	spr := new(spotPriceResponse)

	json.Unmarshal(req.httpGet(), &spr)

	return SpotPrice{
		CoinCurrency: coin,
		Price:        newPrice(spr.Amount, quote),
		Exchange:     exchanges[CoinBase],
	}
}

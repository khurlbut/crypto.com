package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SpotPrice struct {
	CoinCurrency Currency
	Exchange     string
	Price        Price
}

type SpotPriceRequest struct {
	CoinCurrency  Currency
	QuoteCurrency Currency
	Url           string
}

func (sp SpotPrice) format() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("\t%s", sp.CoinCurrency.Code))
	b.WriteString(fmt.Sprintf("\t%s", sp.Price.FormattedPrice))
	b.WriteString(fmt.Sprintf("\t%s", sp.Exchange))

	return b.String()
}

func (spr SpotPriceRequest) httpGet() []byte {
	data, e := ioutil.ReadAll(spr.get())
	if e != nil {
		log.Fatal(e)
	}

	return data
}

func (spr SpotPriceRequest) get() io.ReadCloser {
	res, e := http.Get(spr.Url)
	if e != nil {
		fmt.Print(e.Error())
		os.Exit(1)
	}

	return res.Body
}
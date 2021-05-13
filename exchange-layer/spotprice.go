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

type spotPrice struct {
	coinCurrency currency
	price        price
	exchange     string
}

type spotPriceRequest struct {
	coinCurrency  currency
	quoteCurrency currency
	url           string
}

func (sp spotPrice) format() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("\t%s", sp.coinCurrency.Code))
	b.WriteString(fmt.Sprintf("\t%s", sp.price.formattedPrice))
	b.WriteString(fmt.Sprintf("\t%s", sp.exchange))

	return b.String()
}

func (spr spotPriceRequest) httpGet() []byte {
	data, e := ioutil.ReadAll(spr.get())
	if e != nil {
		log.Fatal(e)
	}

	return data
}

func (spr spotPriceRequest) get() io.ReadCloser {
	res, e := http.Get(spr.url)
	if e != nil {
		fmt.Print(e.Error())
		os.Exit(1)
	}

	return res.Body
}

package main

import (
	"fmt"
)

func main() {

	for _, currencyPair := range currencyPairs() {
		sp := exchange(coinBaseSpotPriceRequest(currencyPair))
		fmt.Println(sp.spotPrice().format())
	}

}

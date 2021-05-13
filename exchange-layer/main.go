package main

import (
	"fmt"
)

func main() {

	for _, val := range currencyPairs() {
		sp := exchange(newCoinBaseSpotPriceRequest(val))
		fmt.Println(sp.spotPrice().format())
	}

}

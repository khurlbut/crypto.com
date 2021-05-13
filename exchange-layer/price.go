package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func newPrice(p string, c currency) price {
	a, _ := strconv.ParseFloat(p, c.Precision)

	return price{
		amount:         a,
		currency:       c,
		formattedPrice: formatPrice(a, c.Symbol),
	}
}

type price struct {
	amount         float64
	currency       currency
	formattedPrice string
}

func formatPrice(amount float64, symbol string) string {
	var b bytes.Buffer

	wholePart := int64(amount)
	decimalPart := int64((amount - float64(wholePart)) * 100)
	priceWithCommas := formatThousandsAmerican(wholePart)

	b.WriteString(fmt.Sprintf("%s%s.%d", symbol, priceWithCommas, decimalPart))

	if wholePart < 1000 {
		b.WriteString("\t")
	}

	return b.String()
}

func formatThousandsAmerican(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

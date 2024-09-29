package main

import "strings"

type productDetails struct {
	productType  string
	manufacturer string
	model        string
}

func getProductInfoParts(product Product) productDetails {
	before, _, _ := strings.Cut(product.info, ",")
	words := strings.Fields(before)

	result := productDetails{
		productType:  strings.Join(words[0:2], " "),
		manufacturer: words[2],
		model:        strings.Join(words[3:], " "),
	}

	return result
}

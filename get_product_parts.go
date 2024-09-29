package main

import "strings"

type productDetails struct {
	productType  string
	manufacturer string
	model        string
}

type ComponentType int

const (
	CPU ComponentType = iota
	GPU
	Motherboard
	RAM
	Case
	PSU
)

func getProductParts(product Product, cType ComponentType) productDetails {
	var result productDetails
	before, _, _ := strings.Cut(product.info, ",")
	words := strings.Fields(before)

	switch cType {
	case CPU, RAM, Case, PSU:
		result = productDetails{
			productType:  words[0],
			manufacturer: words[1],
			model:        strings.Join(words[2:], " "),
		}
	case GPU:
		result = productDetails{
			productType:  strings.Join(words[0:2], " "),
			manufacturer: words[2],
			model:        strings.Join(words[3:], " "),
		}
	}

	return result
}

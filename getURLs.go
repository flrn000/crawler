package main

import (
	"strings"

	"golang.org/x/net/html"
)

type Product struct {
	info  string
	price string
}

var currentProduct Product

func findProductInfo(n *html.Node, results *[]Product) {
	if n.Type == html.ElementNode {
		// Find product title in the <a> tag with class "card-v2-title"
		if n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, "card-v2-title") {
					currentProduct.info = n.FirstChild.Data
				}
			}
		}

		// Find product price in the <p> tag with class "product-new-price"
		if n.Data == "p" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, "product-new-price") {
					if n.FirstChild != nil {
						price := n.FirstChild.Data
						price = strings.Replace(price, ".", "", 1)

						if n.FirstChild.NextSibling != nil && n.FirstChild.NextSibling.LastChild != nil {
							price += "." + n.FirstChild.NextSibling.LastChild.Data
						}

						currentProduct.price = price
					}
				}
			}
		}

		// If both title and price are found, add the product to the list
		if currentProduct.info != "" && currentProduct.price != "" {
			*results = append(*results, currentProduct)
			// Reset current product for the next item
			currentProduct = Product{}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findProductInfo(c, results)
	}
}

func getProductsFromHTML(htmlBody string) ([]Product, error) {
	r := strings.NewReader(htmlBody)
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var results []Product

	findProductInfo(doc, &results)

	return results, nil
}

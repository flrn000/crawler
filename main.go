package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

type Component struct {
	Name         string
	Type         string
	Manufacturer string
	Model        string
	Price        string
}

func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}
	writer := csv.NewWriter(f)
	return writer, f, nil
}

func writeCSVRecord(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	if err != nil {
		fmt.Println("Error writing record to CSV:", err)
	}
}

func main() {
	cmdArgs := os.Args[1:]

	// if len(cmdArgs) < 3 {
	// 	fmt.Println("not enough arguments provided")
	// 	os.Exit(1)
	// } else if len(cmdArgs) > 3 {
	// 	fmt.Println("too many arguments provided")
	// 	os.Exit(1)
	// }

	baseURL := cmdArgs[0]
	// maxConcurrent, err := strconv.Atoi(cmdArgs[1])
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// maxPages, err := strconv.Atoi(cmdArgs[2])
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// parsedURL, err := url.Parse(baseURL)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// cfg := config{
	// 	pages:              make(map[string]int),
	// 	baseURL:            parsedURL,
	// 	mu:                 &sync.Mutex{},
	// 	concurrencyControl: make(chan struct{}, maxConcurrent),
	// 	wg:                 &sync.WaitGroup{},
	// 	maxPages:           maxPages,
	// }

	fmt.Printf("starting crawl\n%v\n", baseURL)

	doc, err := getHTML(baseURL)
	if err != nil {
		log.Fatalf("getHTML: %v", err)
	}

	products, err := getProductsFromHTML(doc)
	if err != nil {
		log.Fatalf("getProductsFromHTML: %v", err)
	}

	w, file, err := createCSVWriter("components.csv")
	if err != nil {
		log.Fatalf("creating CSV writer: %v", err)
	}

	defer file.Close()

	var results [][]string

	for _, product := range products {
		result := getProductInfoParts(product)
		results = append(results, []string{
			product.info,
			result.productType,
			result.manufacturer,
			result.model,
			product.price,
		})
	}

	headers := []string{"name", "type", "manufacturer", "model", "price"}
	writeCSVRecord(w, headers)
	//  \copy components (name, type, manufacturer, model, price) from './components.csv' WITH DELIMITER ',' CSV HEADER;
	for _, r := range results {
		writeCSVRecord(w, r)
	}

	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Println("flushing csv writer:", err)
	}

	// cfg.wg.Add(1)
	// go cfg.crawlPage(baseURL)
	// cfg.wg.Wait()

	// printReport(cfg.pages, baseURL)
}

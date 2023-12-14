package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nurcholisnanda/gocli-fetchurl/fetcher"
)

func main() {
	// Parse command-line options
	var printMetadata bool
	flag.BoolVar(&printMetadata, "metadata", false, "Print metadata")
	flag.Parse()

	// Check for other unexpected flags
	if flag.NFlag() > 1 {
		fmt.Println("Error: Only --metadata flag is allowed")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check if at least one URL is provided
	if flag.NArg() < 1 {
		fmt.Println("Usage: ./fetch [--metadata] <url1> <url2> ...")
		os.Exit(1)
	}

	// Process each URL provided as a command-line argument
	for _, url := range flag.Args() {
		// Create a new fetcher and get response from url
		fetcher := fetcher.New(url)

		switch printMetadata {
		case true:
			filePath := "./" + fetcher.URL + ".html"
			//check whether the file is exist in disk
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				// Save fetched URL to disk if the file is not in the disk
				if err := fetcher.Save(); err != nil {
					fmt.Printf("Error fetching %s: %v\n", url, err)
				}
			}
			// Print metadata from stored file in disk
			fetcher.PrintMetadata(filePath)
		case false:
			// Save fetched URL to disk
			if err := fetcher.Save(); err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
			}
		}
	}
}

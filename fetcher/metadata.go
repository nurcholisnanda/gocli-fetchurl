package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/nurcholisnanda/gocli-fetchurl/models"
	"golang.org/x/net/html"
)

// PrintMetadata prints metadata associated to the response for the given fetcher
func (f *Fetcher) PrintMetadata(filePath string) {
	file, _ := os.Open(filePath)
	info, _ := file.Stat()
	defer file.Close()

	metadata, err := f.extractMetadata(file)
	if err != nil {
		log.Fatal(err)
	}
	metadata.LastFetch = info.ModTime()

	// MarshalIndent formats the JSON with indentation for better readability
	jsonData, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// %s is used to print the string representation of the JSON
	fmt.Printf("%s\n", jsonData)
}

// extractMetadata extracts metadata from the HTML body
func (f *Fetcher) extractMetadata(body io.Reader) (*models.Metadata, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	// Use recursive traversal to count the number of links and images
	numLinks, numImages := countNodes(doc, "a"), countNodes(doc, "img")

	return &models.Metadata{
		Sites:     f.URL,
		NumLinks:  numLinks,
		NumImages: numImages,
	}, nil
}

// countNodes recursively counts the number of nodes with the specified tag name
func countNodes(node *html.Node, tagName string) int {
	count := 0
	if node.Type == html.ElementNode && node.Data == tagName {
		count++
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		count += countNodes(c, tagName)
	}
	return count
}

package fetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Fetcher struct {
	URL  string
	resp *http.Response
}

// NewFetcher creates a new Fetcher instance
func New(url string) *Fetcher {
	r := fetchResponse(url)
	return &Fetcher{
		URL:  sanitize(r.Request.URL.Host, r.Request.URL.Path),
		resp: r,
	}
}

// Save fetched url to disk,
func (f *Fetcher) Save() error {
	// Create a file to save the web page
	file, err := os.Create("./" + f.URL + ".html")
	if err != nil {
		return err
	}
	defer file.Close()

	// Save web page to file
	_, err = io.Copy(file, f.resp.Body)
	if err != nil {
		return err
	}

	//TODO: create mirror folder to store all asssets from the url
	//will try to achieve it in another branch

	return nil
}

func fetchResponse(url string) *http.Response {
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

// sanitize cleans a url string so that it can be used as a file or directory name
func sanitize(host, path string) string {
	if len(path) > 0 {
		if path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}
	}
	s := fmt.Sprintf("%s%s", host, path)
	s = strings.ReplaceAll(s, "/", "_")
	return s
}

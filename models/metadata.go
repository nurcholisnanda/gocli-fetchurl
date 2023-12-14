package models

import "time"

// Metadata struct holds information about the fetched web page
type Metadata struct {
	Sites     string    `json:"sites"`
	NumLinks  int       `json:"num_links"`
	NumImages int       `json:"num_images"`
	LastFetch time.Time `json:"last_fetch"`
}

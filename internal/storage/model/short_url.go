package model

// LinkData describes data model representing a short URL in the storage
type LinkData struct {
	LongURL     string `json:"LongURL"`
	ShortURL    string `json:"ShortURL"`
	Statistics  int64  `json:"Statistics"`
	Description string `json:"Description"`
}

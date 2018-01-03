package models

// Song - holds song metadata
type Song struct {
	Artist       string
	Year         uint16
	Title        string
	Album        string
	StreamingURL string
	Label        string
	ArtWork      string
}

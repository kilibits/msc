package models

//PlayList playlist of songs
type PlayList struct {
	ID          string
	Key         string
	PlaylistURL string
	Songs       []Song
	Creator     string
}

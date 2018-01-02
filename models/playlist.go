package models

//PlayList playlist of songs
type PlayList struct {
	Key         string
	PlaylistURL string
	Songs       []Song
}

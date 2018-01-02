package models

//PlayList playlist of songs
type PlayList struct {
	ID          string
	Key         string
	PlaylistURL string
	Songs       []Song
	Creator     string
}

//AddSongToPlaylist add new song to playlist
func (playlist *PlayList) AddSongToPlaylist(songID string) {

}

//RemoveFromPlaylist remove song from playlist
func (playlist *PlayList) RemoveFromPlaylist(songID string) {

}

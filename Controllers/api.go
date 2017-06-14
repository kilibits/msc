package api

import (
	"net/http"
)

//GetSong retrieve a song by id
func GetSong(res http.ResponseWriter, req *http.Request) {
	RetrieveSong("song-id")
}

//GetPlaylist retrieve playlist by id
func GetPlaylist(res http.ResponseWriter, req *http.Request) {
	//get playlist
}

//PostSong post a song
func PostSong(res http.ResponseWriter, req *http.Request) {
	//save song
}

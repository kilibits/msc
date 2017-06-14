package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

//GetSong retrieve a song by id
func GetSong(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	songID := vars["songId"]
	streamingURL := RetrieveSong(songID)

	if streamingURL == nil {
		http.Error(res, "Error message", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(streamingURL.String()))
}

//GetPlaylist retrieve playlist by id
func GetPlaylist(res http.ResponseWriter, req *http.Request) {
	//get playlist
}

//PostSong post a song
func PostSong(res http.ResponseWriter, req *http.Request) {
	//save song
}

package main

import (
	"net/http"

	api "../Music/Controllers"
	mux "github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/song/{song_id}", api.GetSong).Methods("GET")
	router.HandleFunc("/playlist/playlist_id", api.GetPlaylist).Methods("GET")
	router.HandleFunc("/song", api.PostSong).Methods("POST")

	http.ListenAndServe(":8080", router)
}

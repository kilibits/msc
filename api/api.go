package api

import (
	"fmt"
	"net/http"
	"os"

	utils "../utils"
	"github.com/gorilla/mux"
)

//GetSong retrieve a song by id
func GetSong(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	songID := vars["song_id"]
	streamingURL := utils.MscClient.RetrieveSong(songID)

	if streamingURL == nil {
		http.Error(res, "Error message", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(streamingURL.String()))
}

//GetSongsList retrieve playlist by id
func GetSongsList(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	prefix := vars["prefix"]

	listJSON := utils.MscClient.RetrieveSongList(prefix)

	if listJSON == nil {
		http.Error(res, "Error message", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(listJSON)

}

//PostSong post a song
func PostSong(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	filePath := vars["file_path"]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	utils.MscClient.UploadSong(file)
}

package utils

import (
	"encoding/json"
	"io/ioutil"

	models "../models"
)

//RetrievePlaylist return playlist with given id
func RetrievePlaylist(id string) []string {
	return []string{"song1"}
}

//RetrieveSongList return songs with given prefix
func (client *Client) RetrieveSongList(prefix string) []byte {
	bucketName := client.Bucket
	minioClient := client.MinioClient
	isRecursive := true
	doneCh := make(chan struct{})
	defer close(doneCh)

	var playListEntries []models.PlayList

	for objectInfo := range minioClient.ListObjects(bucketName, prefix, isRecursive, doneCh) {
		if objectInfo.Err != nil {
			return nil
		}
		objectName := objectInfo.Key
		playListEntry := models.PlayList{
			Key: objectName,
		}
		playListEntries = append(playListEntries, playListEntry)
	}
	playListEntriesJSON, err := json.Marshal(playListEntries)
	if err != nil {
		return nil
	}

	return playListEntriesJSON
}

//GetPlaylist return playlist with given id
func (client *Client) GetPlaylist(playlistID string) []byte {
	bucketName := client.Bucket
	minioClient := client.MinioClient

	playlist, err := minioClient.GetObject(bucketName, playlistID)

	if err != nil {
		return nil
	}

	defer playlist.Close()

	res, err := ioutil.ReadAll(playlist)

	if err != nil {
		return nil
	}

	playlistJSON, err := json.Marshal(res)

	if err != nil {
		return nil
	}

	return playlistJSON
}

//AddSongToPlaylist add new song to playlist
func (client *Client) AddSongToPlaylist(playlistID, songID string) {

}

//RemoveFromPlaylist remove song from playlist
func (client *Client) RemoveFromPlaylist(playlistID, songID string) {

}

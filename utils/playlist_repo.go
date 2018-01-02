package utils

import (
	"encoding/json"

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

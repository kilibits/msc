package api

import (
	"encoding/json"

	common "../Common"
)

type mediaPlayList struct {
	Key string
	URL string
}

//RetrievePlaylist return playlist with given id
func RetrievePlaylist(id string) []string {
	return []string{"song1"}
}

//RetrieveSongList return songs with given prefix
func RetrieveSongList(prefix string) []byte {
	client := common.GetClient()
	bucketName := client.Bucket
	minioClient := client.MinioClient
	isRecursive := true
	doneCh := make(chan struct{})
	defer close(doneCh)

	var playListEntries []mediaPlayList

	for objectInfo := range minioClient.ListObjects(bucketName, prefix, isRecursive, doneCh) {
		if objectInfo.Err != nil {
			return nil
		}
		objectName := objectInfo.Key
		playListEntry := mediaPlayList{
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

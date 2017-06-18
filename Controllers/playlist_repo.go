package api

import (
	"encoding/json"

	common "../../Music/Common"
)

type mediaPlayList struct {
	Key string
	URL string
}

func RetrievePlaylist(id string) []string {
	return []string{"song1"}
}

func RetrieveSongList(prefix string) []byte {

	client := common.GetSharedClient()
	bucketName := common.GetBucketName()
	isRecursive := true
	doneCh := make(chan struct{})
	defer close(doneCh)

	var playListEntries []mediaPlayList

	for objectInfo := range client.ListObjects(bucketName, prefix, isRecursive, doneCh) {
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

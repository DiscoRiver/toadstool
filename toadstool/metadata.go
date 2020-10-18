package toadstool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// extract UUID value from $extDir/metadata.json
func getMetaFromExtention(extensionDir string) (string, error) {
	metadata := fmt.Sprintf("%s%s", extensionDir, "/metadata.json")

	jsonFile, err := os.Open(metadata)
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()

	type UUID struct {
		UUID string `json:"UUID"`
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var uuid UUID
	if err = json.Unmarshal(byteValue, &uuid); err != nil {
		return "", err
	}

	return uuid.UUID, nil
}

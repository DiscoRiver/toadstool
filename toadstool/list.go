package toadstool

import (
	"fmt"
	"io/ioutil"
)

func getInstalledUserExtensions() (extensions []string) {
	files, err := ioutil.ReadDir(extensionsDirectory)
	if err != nil {
		fmt.Println("ERROR: Couldn't list extensions from directory ", extensionsDirectory)
		return nil
	}

	for _, file:= range files {
		extensions = append(extensions, file.Name())
	}

	return
}

func ListInstalledUserExtensions() {
	extensions := getInstalledUserExtensions()

	if len(extensions) == 0 {
		fmt.Println("No extensions installed")
		return
	}
	for i := range extensions {
		fmt.Println(extensions[i])
	}

}
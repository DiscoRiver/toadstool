package toadstool

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var extensionsDirectory string

func init() {
	determineExtensionsDir()
	if extensionsDirectory == "" {
		fmt.Println("unable to determine extensions directory, check flags and/or config file.")
		os.Exit(1)
	}
	fmt.Printf("Referencing extensions in %s", extensionsDirectory)
}

func determineExtensionsDir() {
	extensionsDirectory = viper.GetString("gnome.extensionsDirectory")
}
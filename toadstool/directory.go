package toadstool

import (
	"fmt"
	"os"
)

func getCurrentUserHome() (string, error) {
	if homeDir, _ := os.LookupEnv("HOME"); homeDir != "" {
		return "", fmt.Errorf("user home not found")
	} else {
		return homeDir, nil
	}
}

func doesDirectoryExist(dir string) bool {
	// perform some checks to see if a temp directory exists.
	if _, err := os.Stat(dir); err != nil {
		fmt.Printf("%s does not exist\n", dir)
		return false
	}
	return true
}

func makeDirectory(dir string, perm os.FileMode) (err error) {
	if err = os.Mkdir(dir, perm); err == nil {
		fmt.Printf("created directory: %s\n", dir)
		return nil
	}
	return err
}

func removeDirectory(dir string) (err error) {
	if err = os.RemoveAll(dir); err == nil {
		fmt.Printf("removed directory: %s\n", dir)
		return nil
	}
	return err
}

func renameDirectory(from string, to string) (err error) {
	if err = os.Rename(from, to); err == nil {
		fmt.Printf("renamed directory: %s --> %s\n", from, to)
		return nil
	}
	return err
}

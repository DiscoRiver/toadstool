// Automates installation of Gnome extensions
package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
	"io"
	"io/ioutil"
	"bufio"
	"log"
	"archive/zip"
	"path/filepath"
	"encoding/json"
)

var homeDir, _ = os.LookupEnv("HOME")
var extDir = flag.String("d", fmt.Sprintf("%v/.local/share/gnome-shell/extensions", homeDir), "Override default gnome extensions path.")
var fPath = flag.String("f", " ", "Sets Gnome extension to install (.zip file)" )

func main() {
	flag.Parse()
	setup()
	install()
}

func setup() {
	reader := bufio.NewReader(os.Stdin)

	if *fPath == " " {
		fmt.Print("Gnome extension(s) to install not specified. Please enter path: ")
		*fPath, _ = reader.ReadString('\n')
	}
	if i := strings.TrimRight(*fPath, "\n"); strings.HasSuffix(i, ".zip") != true {
		log.Fatalf("Extension '%v' isn't a valid .zip directory. \n", i)
	}

	fmt.Println("Setup complete.")
	fmt.Printf("Extensions to install: %v\n", *fPath)
	fmt.Printf("Install directory: %v \n", *extDir)

}

func install() {
	fPath := strings.TrimRight(*fPath, "\n")
	extDir := strings.TrimRight(*extDir, "\n")
	tmp := fmt.Sprintf("%s%s", extDir, "/tmp")
	var prompt bool

	// perform some checks to see if a temp directory exists.
	if _, err := os.Stat(tmp); err != nil {
		os.Mkdir(tmp, 0775)
		fmt.Printf("Temp directory created: %s\n", tmp)
	} else {
		fmt.Printf("Temp directory not created, already exists: %s\n", tmp)
		fmt.Printf("The directory %v should not exist. Perhaps this program previously terminated unexpectedly? Would you like to delete this directory and proceed with the installation? (yes/no): ", tmp)		   
		
		prompt = askForConfirmation()
		if prompt == true {
			os.RemoveAll(tmp)
			os.Mkdir(tmp, 0775)
			fmt.Printf("Temp directory created: %s\n", tmp)
		} else if prompt == false {
			log.Fatal("Program cannot continue as temp directory not set. Thanks for playing.")
		} else {
			log.Fatal("There was a fatal error with the confirmation, it's not safe to continue.")
		}
	}

	// unzip out file
	files, err := Unzip(fPath, tmp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

	metadata := fmt.Sprintf("%s%s", tmp, "/metadata.json")
	metaUUID := getMeta(metadata)
	fmt.Printf("Set metadata UUID: %s\n", metaUUID)
	finalDir := fmt.Sprintf("%s%s%s", extDir, "/", metaUUID)

	if _, err := os.Stat(finalDir); err != nil {
		fmt.Printf("Renaming temp directory to %v\n", finalDir)
		os.Rename(tmp, finalDir)
	} else {
		fmt.Printf("WARNING: Attempting to install extension at %v, but directory already exists. Continuing will overwrite this directory. Would you like to continue? (yes/no): ", finalDir)
		prompt = askForConfirmation()
		if prompt == true {
			fmt.Printf("Renaming temp directory to %v\n", finalDir)
			os.RemoveAll(finalDir)
			os.Rename(tmp, finalDir)
		} else if prompt == false {
			os.RemoveAll(tmp)
			log.Fatal("System was not changed. Thanks for playing.")
		} else {
			log.Fatal("There was a fatal error with the confirmation, it's not safe to continue.")
		}
	}
}

func askForConfirmation() bool {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	response = strings.Replace(response, "\n", "", -1)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"yes", "Yes", "YES"}
	nokayResponses := []string{"no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation()
	}
}

func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func getMeta(src string) (string) {
	jsonFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Found metadata.json.")
	defer jsonFile.Close()

	type UUID struct {
		UUID string `json:"UUID"`
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var uuid UUID
	json.Unmarshal(byteValue, &uuid)

	extUUID := uuid.UUID
	return extUUID
}

func Unzip(src string, dest string) ([]string, error) {

    var filenames []string

    r, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer r.Close()

    for _, f := range r.File {

        // Store filename/path for returning and using later on
        fpath := filepath.Join(dest, f.Name)

        // Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
        if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
            return filenames, fmt.Errorf("%s: illegal file path", fpath)
        }

        filenames = append(filenames, fpath)

        if f.FileInfo().IsDir() {
            // Make Folder
            os.MkdirAll(fpath, os.ModePerm)
            continue
        }

        // Make File
        if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
            return filenames, err
        }

        outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil {
            return filenames, err
        }

        rc, err := f.Open()
        if err != nil {
            return filenames, err
        }

        _, err = io.Copy(outFile, rc)

        // Close the file without defer to close before next iteration of loop
        outFile.Close()
        rc.Close()

        if err != nil {
            return filenames, err
        }
    }
    return filenames, nil
}
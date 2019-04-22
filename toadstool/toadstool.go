// Automates installation of Gnome extensions
package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"bufio"
	"log"
	"encoding/json"
	"github.com/discoriver/toadstool/util"
)

var homeDir, _ = os.LookupEnv("HOME")
var extDir = flag.String("d", fmt.Sprintf("%v/.local/share/gnome-shell/extensions", homeDir), "Override default gnome extensions path.")
var fPath = flag.String("f", " ", "Sets Gnome extension to install (.zip file)" )

func main() {
	flag.Parse()
	setup()
	install()
}

// Make sure we have the information needed to install.
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
		
		prompt = util.AskForConfirmation()
		if prompt == true {
			os.RemoveAll(tmp)
			// umask value is required for Mkdir
			os.Mkdir(tmp, 0775)
			fmt.Printf("Temp directory created: %s\n", tmp)
		} else if prompt == false {
			log.Fatal("Program cannot continue as temp directory not set. Thanks for playing.")
		} else {
			log.Fatal("There was a fatal error with the confirmation, it's not safe to continue.")
		}
	}

	// unzip extension file
	files, err := util.Unzip(fPath, tmp)
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
		prompt = util.AskForConfirmation()
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

// extract UUID value from $extDir/metadata.json
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

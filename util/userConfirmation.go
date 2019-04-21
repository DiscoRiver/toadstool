// asks user for confirmation
package util

import (
	"bufio"
	"log"
	"strings"
	"fmt"
	"os"
)

// Simple function to handle user confirmation tasks
func AskForConfirmation() bool {
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
		return AskForConfirmation()
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
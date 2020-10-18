package toadstool

import (
	"strings"
)

func trimNewLine(s string) string {
	return strings.TrimRight(s, "\n")
}

func isValidZipFile(path string) bool {
	if i := strings.TrimRight(path, "\n"); strings.HasSuffix(i, ".zip") != true {
		return false
	}
	return true
}

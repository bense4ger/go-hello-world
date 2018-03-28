package helper

import (
	"os"
)

// DirExists checks whether a directory exists
func DirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

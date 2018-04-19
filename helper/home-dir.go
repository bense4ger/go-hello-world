package helper

import (
	"os"
	"path"
	"runtime"
)

// GetHomeDir returns the user's home directory. Cross platform
func GetHomeDir() string {
	if runtime.GOOS == "windows" {
		home := path.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"))
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}

		return home
	}

	return os.Getenv("HOME")
}

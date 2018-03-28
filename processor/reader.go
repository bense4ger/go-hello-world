package processor

import (
	"fmt"
	"io/ioutil"
)

// Do processes each of the files in the passed directory
func Do(dir string) (int, error) {
	count := 0
	files, err := getFileList(dir)
	if err != nil {
		return -1, fmt.Errorf("Do failed to get files: %s", err.Error())
	}

	for range files {
		count++
	}

	return count, nil
}

func getFileList(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("ReadDir failed: %s", err.Error())
	}

	fileNames := make([]string, len(files))
	for ix, file := range files {
		fileNames[ix] = file.Name()
	}

	return fileNames, nil
}

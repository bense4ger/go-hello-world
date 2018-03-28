package processor

import (
	"fmt"
	"io/ioutil"

	"github.com/bense4ger/go-hello-world/model"
)

type counter int

func (c counter) Log() {
	fmt.Printf("count: %d\n", c)
}

// Do processes each of the files in the passed directory
func Do(dir string) (int, error) {
	var count counter

	files, err := getFileList(dir)
	if err != nil {
		return -1, fmt.Errorf("Do failed to get files: %s", err.Error())
	}

	for range files {
		count++
		count.Log()
	}

	return int(count), nil
}

func getFileList(dir string) ([]model.FileReference, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("ReadDir failed: %s", err.Error())
	}

	fileNames := make([]model.FileReference, len(files))
	for ix, file := range files {
		f := &model.File{
			Path: dir,
			Name: file.Name(),
		}

		fileNames[ix] = f
	}

	return fileNames, nil
}

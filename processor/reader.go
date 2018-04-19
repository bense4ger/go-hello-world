package processor

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/bense4ger/go-hello-world/model"
)

// Do processes each of the files in the passed directory
func Do(dir string) error {

	files, err := getFileList(dir)
	if err != nil {
		return fmt.Errorf("Do failed to get files: %s", err.Error())
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(files))
	c := make(chan string, len(files))
	defer close(c)

	for ix, f := range files {
		go f.ProcessContent(ix, wg, c)
	}
	wg.Wait()

	for range files {
		log.Println(<-c)
	}

	return nil
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

package model

import (
	"fmt"
	"io/ioutil"
	"path"
	"time"
)

// File represents a filename
type File struct {
	Path string
	Name string
}

func (f *File) String() string {
	return path.Join(f.Path, f.Name)
}

// ProcessContent takes the reciever's content and appends the date
func (f *File) ProcessContent() (string, error) {
	bytes, err := ioutil.ReadFile(f.String())
	if err != nil {
		return "", fmt.Errorf("Error processing content for %s: %s", f, err.Error())
	}

	t := time.Now()
	formattedT := t.Format("2006-02-01")

	return fmt.Sprintf("%s - [%s]", bytes, formattedT), nil
}

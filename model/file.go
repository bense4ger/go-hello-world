package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"sync"
	"time"

	"github.com/bense4ger/go-hello-world/helper"
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
func (f *File) ProcessContent(id int, wg *sync.WaitGroup, out chan string) {
	d := helper.RandomDuration()
	time.Sleep(d)

	bytes, err := ioutil.ReadFile(f.String())
	if err != nil {
		log.Println(fmt.Errorf("Error processing content for %s: %s", f, err.Error()))
		wg.Done()
		return
	}

	t := time.Now()
	formattedT := t.Format("2006-02-01")

	out <- fmt.Sprintf("%s - [%d] [%s] - Delayed: %s", bytes, id, formattedT, d)
	wg.Done()
}

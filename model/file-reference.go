package model

import (
	"sync"
)

// FileReference describes the shape of types which can process file content
type FileReference interface {
	ProcessContent(id int, wg *sync.WaitGroup, out chan string)
}

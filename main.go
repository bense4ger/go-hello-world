package main

import (
	"flag"
	"log"
	"path"

	"github.com/bense4ger/go-hello-world/helper"
	"github.com/bense4ger/go-hello-world/processor"
)

var pathFlg string

func main() {
	if pathFlg == "" {
		log.Fatal("No path supplied")
	}

	if !helper.DirExists(pathFlg) {
		log.Fatalf("Input directory %s does not exist", pathFlg)
	}

	log.Printf("Reading from : %s", pathFlg)

	c, err := processor.Do(pathFlg)
	if err != nil {
		log.Printf("Error doing: %s", err.Error())
	}

	log.Println("Complete")
	if c > 0 {
		log.Printf("%d Records Processed", c)
	}
}

func init() {
	dPath := path.Join(helper.GetHomeDir(), "hello-world")
	flag.StringVar(&pathFlg, "path", dPath, "The path to read files from")

	flag.Parse()
}

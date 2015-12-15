package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dreamvids/go-webm"
)

func main() {
	fmt.Println("Go-WebM")

	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Usage: webminfo [options] <file>\n")
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	doc := webm.InitDocument(file)
	err = doc.ParseAll(func(el webm.Element) {
		fmt.Printf("Element %s - %d bytes\n", el.Name, el.Size)
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s at %d\n", err, doc.Cursor)
	}
}

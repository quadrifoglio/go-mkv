package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/quadrifoglio/go-mkv"
)

func main() {
	fmt.Println("go-mkv")

	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Usage: gomkvinfo [options] <file>\n")
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	doc := mkv.InitDocument(file)
	err = doc.ParseAll(func(el mkv.Element) {
		fmt.Printf("Element %s - %d bytes\n", el.Name, el.Size)
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s at %d\n", err, doc.Cursor)
	}
}

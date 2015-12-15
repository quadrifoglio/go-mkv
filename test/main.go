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
		fmt.Println("Usage: webminfo [options] <file>")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can not open file:", err)
		os.Exit(1)
	}

	doc := webm.InitDocument(data)
	doc.GetAllElements()
	fmt.Println("Done loading")

	/*for _, el := range elements {
		fmt.Printf("Element %s - %d bytes\n", el.Name, el.Size)
	}*/
}

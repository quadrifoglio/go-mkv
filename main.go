package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dreamvids/webm-info/webm"
)

const (
	Name    = "WebM Info Parser"
	Version = "0.1"
)

func main() {
	fmt.Println(Name, "-", Version)

	if len(os.Args) == 2 {
		doc, err := webm.ParseFile(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			return
		}

		doc.Cursor = 0
		data, err := webm.ReadUntilDataStart(&doc)
		if err != nil {
			fmt.Println("Can not get cluster data:", err)
		}

		ioutil.WriteFile("mdr.webm", data, os.ModeAppend)

	} else {
		fmt.Fprintf(os.Stderr, "Usage: webm-info <file>\n")
	}
}

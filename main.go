package main

import (
	"fmt"
	"os"

	"github.com/dreamvids/go-webm/webm"
)

const (
	Name    = "go-webm info parser"
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
		f, _ := os.Create("mdr.webm")

		header, err := webm.ReadHeader(&doc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			return
		}

		f.Write(header)

		for {
			clu, err := webm.ReadCluster(&doc)
			if err != nil {
				fmt.Printf("Cluster error: %s\n", err)
				break
			}

			f.Write(clu)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Usage: webm-info <file>\n")
	}
}

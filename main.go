package main

import (
	"fmt"
	"io"
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
		f, _ := os.Create("mdr.webm")

		header, err := webm.ReadHeader(&doc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			return
		}

		f.Write(header)

		nClu := 0
		nBl := 0
		for {
			clusterData, err := webm.ReadClusterData(&doc)
			if err != nil {
				fmt.Println("Cluster %d error:", nClu, err)
				break
			}

			f.Write(clusterData)

			for {
				block, err := webm.ReadBlock(&doc)
				f.Write(block)
				//fmt.Printf("Wrote %d for block %d\n", len(block), nBl)
				nBl++

				if err != nil && err == io.EOF {
					fmt.Fprintf(os.Stderr, "Block %d EOF\n", nBl)
					return
				} else if err != nil {
					fmt.Fprintf(os.Stderr, "Block %d error: %s\n", nBl, err)
				}
			}

			nClu++
		}
	} else {
		fmt.Fprintf(os.Stderr, "Usage: webm-info <file>\n")
	}
}

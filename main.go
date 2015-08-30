package main

import (
	"fmt"
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
		err := ebml.Parse(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Usage: webm-info <file>\n")
	}
}

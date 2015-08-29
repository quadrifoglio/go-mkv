package main

import (
	"fmt"
)

const (
	Name    = "WebM Info Parser"
	Version = "0.1"
)

func main() {
	fmt.Println(Name, "-", Version)
	fmt.Println("Usage: webm-info <file>")
}

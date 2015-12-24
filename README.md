# go-mkv

Matroska/WebM parsing for the Go Programming Language.

## Example

```go
file, err := ioutil.ReadFile("video.webm")
if err != nil {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}

doc := mkv.InitDocument(file)
err = doc.ParseAll(func(el mkv.Element) {
	// Found an EBML/WebM element
	fmt.Printf("Element %s - %d bytes\n", el.Name, el.Size)
})

if err != nil {
	fmt.Fprintf(os.Stderr, "%s at %d\n", err, doc.Cursor)
}
```

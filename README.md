# go-mkv

Matroska/WebM parsing for the Go Programming Language.

## Example

```go
file, err := os.Open("video.webm")
if err != nil {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	return
}

defer file.Close()

doc := mkv.InitDocument(file)
err = doc.ParseAll(func(el mkv.Element) {
	fmt.Printf("Element %s - %d bytes\n", el.Name, el.Size)
})

if err != nil {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	return
}
```

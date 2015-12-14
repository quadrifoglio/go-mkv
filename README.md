# Go-WebM

WebM parsing for the Go Programming Language.

## Example

```go
data, err := ioutil.ReadFile(os.Args[1])
if err != nil {
	fmt.Println("Can not open file:", err)
	os.Exit(1)
}

doc := webm.InitDocument(data)
elements := doc.GetAllElements()

for _, el := range elements {
	fmt.Printf("Element %s - %d bytes\n", el.Name, el.Size)
}
```

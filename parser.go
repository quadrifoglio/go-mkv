package webm

func InitDocument(data []byte) *Document {
	doc := new(Document)

	doc.Data = data
	doc.Cursor = 0

	return doc
}

func ReadElement(doc *Document) Element {
	var el Element
	return el
}

package webm

type Document struct {
	Data   []byte
	Cursor int64
}

type Element struct {
	ID    uint32
	Type  uint8
	Name  string
	Level int32
	Size  int32
	Data  []byte
}

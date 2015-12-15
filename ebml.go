package webm

type Document struct {
	Data   []byte
	Length uint64
	Cursor uint64
}

type ElementRegister struct {
	ID   uint32
	Type uint8
	Name string
}

type Element struct {
	ElementRegister

	Parent *Element // May be nil
	Level  int32
	Index  uint64
	Size   uint64
	Data   []byte
}

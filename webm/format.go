package webm

const (
	TypeInt           = 0x0
	TypeUint          = 0x1
	TypeFloat         = 0x2
	TypeString        = 0x3
	TypeUnicode       = 0x4
	TypeDate          = 0x5
	TypeMasterElement = 0x6
	TypeBinary        = 0x7
)

type Document struct {
	Data     []byte
	Length   uint64
	Cursor   uint64
	Elements []Element
}

type Element struct {
	ID       uint32
	Level    uint8
	Multiple bool
	Type     uint8
	Size     uint64
	Data     []byte
}

func (e *Element) GetFullSize() uint64 {
	clID := getElementIDClass(e.ID)
	clSize := getElementSizeClass(e.Size)

	return e.Size + uint64(clID) + uint64(clSize)
}

package ebml

import (
	"fmt"
	"io"
	"io/ioutil"
)

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

type Element struct {
	Level uint8
	Type  uint8
	Size  uint
	Data  []byte
}

func Parse(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if pack32(data[0:4]) != ElementEBML {
		return fmt.Errorf("Invalid WebM file: EBML root element not found")
	}

	var cursor int = 0
	var i int = 0

	for cursor < len(data) {
		_, err := getNextElement(data, &cursor)
		if err != nil {
			cursor++
			continue
		}

		if i >= 16 {
			break
		}

		i++
	}

	return nil
}

func getNextElement(data []byte, cursor *int) (Element, error) {
	var res Element
	var b = data[*cursor]

	if ((b & 0x80) >> 7) == 1 { // Class A ID (on 1 byte)
		*cursor += 1
		b = data[*cursor]

		fmt.Println("Class A of size", getElementSize(data, *cursor))
		return res, nil
	}
	if ((b & 0x40) >> 6) == 1 {
		*cursor += 2
		b = data[*cursor]

		fmt.Println("Class B of size", getElementSize(data, *cursor))
		return res, nil

	}
	if ((b & 0x20) >> 5) == 1 {
		*cursor += 3
		b = data[*cursor]

		fmt.Println("Class C of size", getElementSize(data, *cursor))
		return res, nil

	}
	if ((b & 0x10) >> 4) == 1 {
		*cursor += 4
		b = data[*cursor]

		fmt.Println("Class D of size", getElementSize(data, *cursor))
		return res, nil
	}

	return res, io.EOF
}

func getElementSize(data []byte, at int) uint8 {
	return 0
}

func pack16(b []byte) uint16 {
	return (uint16(b[2]) << 8) | (uint16(b[3]) << 0)
}

func pack32(b []byte) uint32 {
	return (uint32(b[0]) << 24) | (uint32(b[1]) << 16) | (uint32(b[2]) << 8) | (uint32(b[3]) << 0)
}

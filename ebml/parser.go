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

	var cursor uint64 = 0
	var i int = 0

	for cursor < uint64(len(data)) {
		_, err := getNextElement(data, &cursor)
		if err != nil {
			cursor++
			continue
		}

		i++
	}

	return nil
}

func getNextElement(data []byte, cursor *uint64) (Element, error) {
	var res Element
	var b = data[*cursor]

	if ((b & 0x80) >> 7) == 1 { // Class A ID (on 1 byte)
		*cursor += 1
		b = data[*cursor]

		size := getElementSize(data, *cursor)
		*cursor += size

		fmt.Println("Class A of size", size)
		return res, nil
	}
	if ((b & 0x40) >> 6) == 1 {
		*cursor += 2
		b = data[*cursor]

		size := getElementSize(data, *cursor)
		*cursor += size

		fmt.Println("Class B of size", size)
		return res, nil

	}
	if ((b & 0x20) >> 5) == 1 {
		*cursor += 3
		b = data[*cursor]

		size := getElementSize(data, *cursor)
		*cursor += size

		fmt.Println("Class C of size", size)
		return res, nil

	}
	if ((b & 0x10) >> 4) == 1 {
		*cursor += 4
		b = data[*cursor]

		size := getElementSize(data, *cursor)
		*cursor += size

		fmt.Println("Class D of size", size)
		return res, nil
	}

	return res, io.EOF
}

func getElementSize(data []byte, at uint64) uint64 {
	b := data[at]

	if ((b & 0x80) >> 7) == 1 { // Size coded on 1 byte
		return uint64(b & 0x7f)
	}
	if ((b & 0x40) >> 6) == 1 { // Size coded on 2 byte
		return uint64(pack16([]byte{0x3f & b, data[at+1]}))
	}
	if ((b & 0x20) >> 5) == 1 { // Size coded on 3 byte
		return uint64(pack24([]byte{0x1f & b, data[at+1], data[at+2]}))
	}
	if ((b & 0x10) >> 4) == 1 { // Size coded on 4 byte
		return uint64(pack32([]byte{0xf & b, data[at+1], data[at+2], data[at+3]}))
	}
	if ((b & 0x8) >> 3) == 1 { // Size coded on 5 byte
		return uint64(pack40([]byte{0x7 & b, data[at+1], data[at+2], data[at+3], data[at+4]}))
	}
	if ((b & 0x4) >> 2) == 1 { // Size coded on 6 byte
		return uint64(pack48([]byte{0x3 & b, data[at+1], data[at+2], data[at+3], data[at+4], data[at+5]}))
	}
	if ((b & 0x2) >> 1) == 1 { // Size coded on 7 byte
		return uint64(pack56([]byte{0x1 & b, data[at+1], data[at+2], data[at+3], data[at+4], data[at+5], data[at+6]}))
	}
	if ((b & 0x1) >> 0) == 1 { // Size coded on 8 byte
		return uint64(pack64([]byte{0x0 & b, data[at+1], data[at+2], data[at+3], data[at+4], data[at+5], data[at+6], data[at+7]}))
	}

	return 0
}

func pack16(b []byte) uint16 {
	return (uint16(b[0]) << 8) | (uint16(b[1]) << 0)
}

func pack24(b []byte) uint32 {
	return (uint32(b[0]) << 16) | (uint32(b[1]) << 8) | (uint32(b[2]) << 0)
}

func pack32(b []byte) uint32 {
	return (uint32(b[0]) << 24) | (uint32(b[1]) << 16) | (uint32(b[2]) << 8) | (uint32(b[3]) << 0)
}

func pack40(b []byte) uint64 {
	return (uint64(b[0]) << 32) | (uint64(b[1]) << 24) | (uint64(b[2]) << 16) | (uint64(b[3]) << 8) | (uint64(b[4]) << 0)
}

func pack48(b []byte) uint64 {
	return (uint64(b[0]) << 40) | (uint64(b[1]) << 32) | (uint64(b[2]) << 24) | (uint64(b[3]) << 16) | (uint64(b[4]) << 8) | (uint64(b[5]) << 0)
}

func pack56(b []byte) uint64 {
	return (uint64(b[0]) << 48) | (uint64(b[1]) << 40) | (uint64(b[2]) << 32) | (uint64(b[3]) << 24) | (uint64(b[4]) << 16) | (uint64(b[5]) << 8) | (uint64(b[6]) << 0)
}

func pack64(b []byte) uint64 {
	return (uint64(b[0]) << 56) | (uint64(b[1]) << 48) | (uint64(b[2]) << 40) | (uint64(b[3]) << 32) | (uint64(b[4]) << 24) | (uint64(b[5]) << 16) | (uint64(b[6]) << 8) | (uint64(b[7]) << 0)
}

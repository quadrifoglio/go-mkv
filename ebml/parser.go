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
	ID    uint32
	Level uint8
	Type  uint8
	Size  uint64
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

	var level = 0
	var levelEnd uint64 = 0

	for cursor < uint64(len(data)) {
		el, err := getNextElement(data, &cursor)
		if err != nil {
			fmt.Printf("Finished at index %d (value 0x%x): %s\n", cursor, data[cursor], err)
			break
		}

		if levelEnd != 0 && cursor == levelEnd {
			level--
			levelEnd = 0
		}

		el.Level = uint8(level)

		if el.Type == TypeMasterElement {
			level++
			levelEnd = cursor + el.Size
		}

		fmt.Printf("%d: %s (0x%d) of size %d - Value: %v\n", el.Level, GetElementName(el.ID), el.ID, el.Size, el.Data)
		i++
	}

	return nil
}

func getNextElement(data []byte, cursor *uint64) (Element, error) {
	var res Element
	if *cursor >= uint64(len(data)) {
		return res, io.EOF
	}

	var b = data[*cursor]

	if ((b & 0x80) >> 7) == 1 { // Class A ID (on 1 byte)
		t, err := getElementID(1, data, cursor)
		switch t {
		}

		_, err = getElementSize(data, cursor)
		if err != nil {
			return res, err
		}

		return res, nil
	}
	if ((b & 0x40) >> 6) == 1 { // Class B ID (on 2 byte)
		id, err := getElementID(2, data, cursor)
		if err != nil {
			return res, err
		}

		size, err := getElementSize(data, cursor)
		if err != nil {
			return res, err
		}

		switch id {
		case ElementEBMLVersion:
			res.Type = TypeUint
			break
		case ElementEBMLReadVersion:
			res.Type = TypeUint
			break
		case ElementEBMLMaxIDLength:
			res.Type = TypeUint
			break
		case ElementEBMLMaxSizeLength:
			res.Type = TypeUint
			break
		case ElementDocType:
			res.Type = TypeString
			break
		case ElementDocTypeVersion:
			res.Type = TypeUint
			break
		case ElementDocTypeReadVersion:
			res.Type = TypeUint
			break
		}

		d, err := getElementData(size, data, cursor)
		if err != nil {
			return res, nil
		}

		res.ID = id
		res.Data = d
		res.Size = size
		return res, nil
	}
	if ((b & 0x20) >> 5) == 1 { // Class C ID (on 4 bytes)
		id, err := getElementID(3, data, cursor)

		switch id {
		}

		size, err := getElementSize(data, cursor)
		if err != nil {
			return res, err
		}

		res.ID = id
		res.Size = size
		return res, nil
	}
	if ((b & 0x10) >> 4) == 1 { // Class D ID (on 4 bytes)
		id, err := getElementID(4, data, cursor)

		switch id {
		case ElementEBML:
			res.Type = TypeMasterElement
			break
		case ElementSegment:
			res.Type = TypeMasterElement
			break
		}

		size, err := getElementSize(data, cursor)
		if err != nil {
			return res, err
		}

		res.ID = id
		res.Size = size
		return res, nil
	}

	return res, fmt.Errorf("Failed to identify tag")
}

func getElementID(class uint8, data []byte, at *uint64) (uint32, error) {
	if class == 1 {
		b := data[*at]

		*at++
		return uint32(b), nil
	}
	if class == 2 {
		b := data[*at : *at+2]

		*at += 2
		return uint32(pack16(b)), nil
	}
	if class == 3 {
		b := data[*at : *at+3]

		*at += 3
		return uint32(pack24(b)), nil
	}
	if class == 4 {
		b := data[*at : *at+4]

		*at += 4
		return uint32(pack32(b)), nil
	}

	return 0, fmt.Errorf("Unknown element")
}

func getElementSize(data []byte, at *uint64) (uint64, error) {
	b := data[*at]

	if ((b & 0x80) >> 7) == 1 { // Size coded on 1 byte
		v := uint64(b & 0x7f)

		*at++
		return v, nil
	}
	if ((b & 0x40) >> 6) == 1 { // Size coded on 2 byte
		v := uint64(pack16([]byte{0x3f & b, data[*at+1]}))

		*at += 2
		return v, nil
	}
	if ((b & 0x20) >> 5) == 1 { // Size coded on 3 byte
		v := uint64(pack24([]byte{0x1f & b, data[*at+1], data[*at+2]}))

		*at += 3
		return v, nil
	}
	if ((b & 0x10) >> 4) == 1 { // Size coded on 4 byte
		v := uint64(pack32([]byte{0xf & b, data[*at+1], data[*at+2], data[*at+3]}))

		*at += 4
		return v, nil
	}
	if ((b & 0x8) >> 3) == 1 { // Size coded on 5 byte
		v := uint64(pack40([]byte{0x7 & b, data[*at+1], data[*at+2], data[*at+3], data[*at+4]}))

		*at += 5
		return v, nil
	}
	if ((b & 0x4) >> 2) == 1 { // Size coded on 6 byte
		v := uint64(pack48([]byte{0x3 & b, data[*at+1], data[*at+2], data[*at+3], data[*at+4], data[*at+5]}))

		*at += 6
		return v, nil
	}
	if ((b & 0x2) >> 1) == 1 { // Size coded on 7 byte
		v := uint64(pack56([]byte{0x1 & b, data[*at+1], data[*at+2], data[*at+3], data[*at+4], data[*at+5], data[*at+6]}))

		*at += 7
		return v, nil
	}
	if ((b & 0x1) >> 0) == 1 { // Size coded on 8 byte
		v := uint64(pack64([]byte{0x0 & b, data[*at+1], data[*at+2], data[*at+3], data[*at+4], data[*at+5], data[*at+6], data[*at+7]}))

		*at += 8
		return v, nil
	}

	return 0, fmt.Errorf("Invalid size format")
}

func getElementData(size uint64, b []byte, at *uint64) ([]byte, error) {
	if uint64(len(b[*at:*at+size])) != size {
		return nil, io.EOF
	}

	var i uint64
	var buffer = make([]byte, size)

	for i = 0; i < size; i++ {
		buffer[i] = b[i]
	}

	*at += size
	return buffer, nil
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

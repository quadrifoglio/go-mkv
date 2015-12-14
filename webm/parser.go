package webm

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

var (
	NoHeader   = errors.New("No header")
	NoCluster  = errors.New("No cluster")
	EndOfBlock = errors.New("EOB")
)

func ParseFile(filename string) (Document, error) {
	var doc Document

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return doc, err
	}

	if pack32(data[0:4]) != ElementEBML {
		return doc, fmt.Errorf("Invalid WebM file: EBML root element not found")
	}

	doc.Data = data
	doc.Length = uint64(len(doc.Data))

	for doc.Cursor < doc.Length {
		el, err := getNextElement(&doc)
		if err != nil {
			fmt.Printf("Finished at index %d (value 0x%x): %s\n", doc.Cursor, doc.Data[doc.Cursor], err)
			break
		}

		fmt.Printf("%s (0x%x) containing %d bytes (total: %d)\n", GetElementName(el.ID), el.ID, el.Size, el.GetFullSize())
		doc.Elements = append(doc.Elements, el)
	}

	return doc, nil
}

func ReadHeader(doc *Document) ([]byte, error) {
	for doc.Cursor < doc.Length {
		if doc.Cursor+4 >= doc.Length {
			return nil, io.EOF
		}

		if pack32(doc.Data[doc.Cursor:doc.Cursor+4]) == ElementCluster {
			return doc.Data[0:doc.Cursor], nil
		}

		doc.Cursor++
	}

	return nil, NoHeader
}

func ReadClusterHeader(doc *Document) ([]byte, error) {
	if doc.Cursor >= doc.Length {
		return nil, io.EOF
	}

	if doc.Cursor+4 < doc.Length && pack32(doc.Data[doc.Cursor:doc.Cursor+4]) == ElementCluster {
		start := doc.Cursor
		doc.Cursor += 4

		_, err := getElementSize(doc)
		if err != nil {
			return nil, err
		}

		var id uint32 = 0

		for doc.Cursor < doc.Length {
			id, err = getElementID(doc)
			if err != nil {
				return nil, err
			}

			cl := uint64(getElementIDClass(id))

			if id == ElementBlock || id == ElementSimpleBlock {
				doc.Cursor -= cl
				return doc.Data[start:doc.Cursor], nil
			} else {
				doc.Cursor -= cl
				getNextElement(doc)
			}
		}
	}

	return nil, io.EOF
}

func ReadCluster(doc *Document) ([]byte, error) {
	if doc.Cursor >= doc.Length {
		return nil, io.EOF
	}

	for doc.Cursor < doc.Length {
		start := doc.Cursor
		el, err := getNextElement(doc)
		if err != nil {
			return nil, err
		}

		if el.ID == ElementCluster {
			return doc.Data[start:doc.Cursor], nil
		}
	}

	return nil, NoCluster
}

func ReadClusterFrame(doc *Document) ([]byte, error) {
	if doc.Cursor >= doc.Length {
		return nil, io.EOF
	}

	for doc.Cursor < doc.Length {
		var c = doc.Data[doc.Cursor]
		var cl = getElementIDClass(uint32(c))
		var id, err = getElementID(doc)

		if err != nil {
			fmt.Println("Frame err", err)
			doc.Cursor++
			continue
		}

		if id == ElementBlock || id == ElementSimpleBlock {
			doc.Cursor -= uint64(cl)
			return ReadBlock(doc)
		}
	}

	return nil, io.EOF
}

func ReadBlock(doc *Document) ([]byte, error) {
	if doc.Cursor >= doc.Length {
		return nil, io.EOF
	}

	for doc.Cursor < doc.Length {
		start := doc.Cursor
		el, err := getNextElement(doc)
		if err != nil {
			return nil, err
		}

		if el.ID == ElementBlock || el.ID == ElementSimpleBlock {
			return doc.Data[start:doc.Cursor], nil
		} else if el.ID == ElementCluster {
			doc.Cursor = start
			return nil, EndOfBlock
		}
	}

	return nil, fmt.Errorf("Invalid data")
}

func getNextElement(doc *Document) (Element, error) {
	var res Element
	if doc.Cursor >= doc.Length {
		return res, io.EOF
	}

	id, err := getElementID(doc)
	if err != nil {
		return res, err
	}

	res.ID = id

	switch id {
	case ElementEBML:
		res.Type = TypeMasterElement
		res.Multiple = false
		break
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
	case ElementSeek:
		res.Type = TypeMasterElement
		break
	case ElementSeekID:
		res.Type = TypeBinary
		break
	case ElementSeekPosition:
		res.Type = TypeUint
		break
	case ElementSegment:
		res.Type = TypeMasterElement
		res.Multiple = true
		break
	case ElementSeekHead:
		res.Type = TypeMasterElement
		res.Multiple = true
		break
	case ElementCluster:
		res.Type = TypeMasterElement
		res.Multiple = true
		break
	}

	size, err := getElementSize(doc)
	if err != nil {
		return res, err
	}

	res.Size = size

	d, err := getElementData(size, doc)
	if err != nil {
		return res, err
	}

	res.Data = d
	return res, nil
}

func getElementID(doc *Document) (uint32, error) {
	class := getElementIDClass(uint32(doc.Data[doc.Cursor]))

	if class == 1 {
		b := doc.Data[doc.Cursor]

		doc.Cursor++
		return uint32(b), nil
	} else if class == 2 {
		b := doc.Data[doc.Cursor : doc.Cursor+2]

		doc.Cursor += 2
		return uint32(pack16(b)), nil
	} else if class == 3 {
		b := doc.Data[doc.Cursor : doc.Cursor+3]

		doc.Cursor += 3
		return uint32(pack24(b)), nil
	} else if class == 4 {
		b := doc.Data[doc.Cursor : doc.Cursor+4]

		doc.Cursor += 4
		return uint32(pack32(b)), nil
	} else {
		return 0, fmt.Errorf("Unknown ID class")
	}

	return 0, fmt.Errorf("Unknown element (ID class %d)", class)
}

func getElementSize(doc *Document) (uint64, error) {
	if doc.Cursor >= doc.Length {
		return 0, io.EOF
	}

	b := doc.Data[doc.Cursor]
	length := 0
	mask := byte(0)

	if b >= 0x80 {
		length = 1
		mask = 0x7f
	} else if b >= 0x40 {
		length = 2
		mask = 0x3f
	} else if b >= 0x20 {
		length = 3
		mask = 0x1f
	} else if b >= 0x10 {
		length = 4
		mask = 0xf
	} else if b >= 0x8 {
		length = 5
		mask = 0x7
	} else if b >= 0x4 {
		length = 6
		mask = 0x3
	} else if b >= 0x2 {
		length = 7
		mask = 0x1
	} else if b >= 0x1 {
		length = 8
		mask = 0x0
	} else {
		return 0, fmt.Errorf("Invalid size format")
	}

	if doc.Cursor+uint64(length) >= doc.Length {
		return 0, io.EOF
	}

	var v uint64 = 0
	var s = doc.Cursor
	for i, l := doc.Cursor, uint64(length); i < s+l; i++ {
		bb := doc.Data[i]

		if i == s {
			bb &= mask
		}

		v <<= 8
		v += uint64(bb)

		doc.Cursor++
	}

	return v, nil
}

func getElementData(size uint64, doc *Document) ([]byte, error) {
	if doc.Cursor+size > doc.Length {
		return nil, io.EOF
	}
	if uint64(len(doc.Data[doc.Cursor:doc.Cursor+size])) != size {
		return nil, io.EOF
	}

	var i uint64
	var buffer = make([]byte, size)

	for i = 0; i < size; i++ {
		buffer[i] = doc.Data[doc.Cursor]
		doc.Cursor++
	}

	//doc.Cursor += size
	return buffer, nil
}

func getElementIDClass(id uint32) uint {
	if ((id & 0x80) >> 7) == 1 { // Class A ID (on 1 byte)
		return 1
	}
	if ((id & 0x40) >> 6) == 1 { // Class B ID (on 2 byte)
		return 2
	}
	if ((id & 0x20) >> 5) == 1 { // Class C ID (on 3 bytes)
		return 3
	}
	if ((id & 0x10) >> 4) == 1 { // Class D ID (on 4 bytes)
		return 4
	} else {
		return 0
	}
}

func getElementSizeClass(size uint64) uint {
	if size >= 0x80 {
		return 1
	} else if size >= 0x40 {
		return 2
	} else if size >= 0x20 {
		return 3
	} else if size >= 0x10 {
		return 4
	} else if size >= 0x8 {
		return 5
	} else if size >= 0x4 {
		return 6
	} else if size >= 0x2 {
		return 7
	} else if size >= 0x1 {
		return 8
	} else {
		return 0
	}
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

func unpack(n int, v uint64) []byte {
	var bytes []byte

	for i := uint(n); i > 0; i-- {
		bytes = append(bytes, byte(v>>(8*i))&0xff)
	}

	return bytes
}

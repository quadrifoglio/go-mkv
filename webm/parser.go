package webm

import (
	"fmt"
	"io"
	"io/ioutil"
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

		//fmt.Printf("%d: %s (0x%x) containing %d bytes\n", el.Level, GetElementName(el.ID), el.ID, el.Size)
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

	return nil, fmt.Errorf("Invalid data")
}

func ReadClusterData(doc *Document) ([]byte, error) {
	for doc.Cursor < doc.Length {
		if doc.Cursor+4 >= doc.Length {
			fmt.Printf("mdr: cur %d len %d\n", doc.Cursor, doc.Length)
			return nil, io.EOF
		}

		if pack32(doc.Data[doc.Cursor:doc.Cursor+4]) == ElementCluster {

			start := doc.Cursor
			getNextElement(doc)

			i := 0
			for {
				idClass := uint64(getElementIDClass(uint32(doc.Data[doc.Cursor])))
				elId, err := getElementID(uint8(idClass), doc)
				if err != nil {
					return nil, err
				}

				if elId == ElementSimpleBlock || elId == ElementBlock {
					end := doc.Cursor - idClass
					doc.Cursor -= idClass

					return doc.Data[start:end], nil
				}

				size, err := getElementSize(doc)
				if err != nil {
					return nil, err
				}

				doc.Cursor += size
				i++

				if i >= 5 {
					return nil, fmt.Errorf("Did not find begining of video data")
				}
			}
		}

		doc.Cursor++
	}

	return nil, fmt.Errorf("Invalid data")
}

func ReadBlock(doc *Document) ([]byte, error) {
	if doc.Cursor >= doc.Length {
		return nil, io.EOF
	}

	for doc.Cursor < doc.Length {
		el, err := getNextElement(doc)
		if err != nil {
			return nil, err
		}

		if el.ID == ElementBlock || el.ID == ElementSimpleBlock {
			bytes, err := el.GetRawBytes()
			if err != nil {
				return nil, err
			}

			return bytes, nil
		}
	}

	return nil, fmt.Errorf("Invalid data")
}

func getNextElement(doc *Document) (Element, error) {
	var res Element
	if doc.Cursor >= doc.Length {
		return res, io.EOF
	}

	var b = doc.Data[doc.Cursor]
	var c = getElementIDClass(uint32(b))

	if c == 1 { // Class A ID (on 1 byte)
		id, err := getElementID(1, doc)
		switch id {
		}

		size, err := getElementSize(doc)
		if err != nil {
			return res, err
		}

		d, err := getElementData(size, doc)
		if err != nil {
			return res, nil
		}

		res.ID = id
		res.Size = size
		res.Data = d
		return res, nil
	}
	if c == 2 { // Class B ID (on 2 byte)
		id, err := getElementID(2, doc)
		if err != nil {
			return res, err
		}

		size, err := getElementSize(doc)
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
		case ElementSeek:
			res.Type = TypeMasterElement
			break
		case ElementSeekID:
			res.Type = TypeBinary
			break
		case ElementSeekPosition:
			res.Type = TypeUint
			break
		}

		d, err := getElementData(size, doc)
		if err != nil {
			return res, nil
		}

		res.ID = id
		res.Data = d
		res.Size = size
		return res, nil
	}
	if c == 3 { // Class C ID (on 3 bytes)
		id, err := getElementID(3, doc)

		switch id {
		}

		size, err := getElementSize(doc)
		if err != nil {
			return res, err
		}

		d, err := getElementData(size, doc)
		if err != nil {
			return res, nil
		}

		res.ID = id
		res.Size = size
		res.Data = d
		return res, nil
	}
	if c == 4 { // Class D ID (on 4 bytes)
		id, err := getElementID(4, doc)

		switch id {
		case ElementEBML:
			res.Type = TypeMasterElement
			res.Multiple = false
			break
		case ElementSegment:
			res.Type = TypeMasterElement
			res.Multiple = true
			break
		case ElementSeekHead:
			res.Type = TypeMasterElement
			res.Multiple = true
			break
		}

		size, err := getElementSize(doc)
		if err != nil {
			return res, err
		}

		res.ID = id
		res.Size = size
		return res, nil
	}

	return res, fmt.Errorf("Failed to identify element ID 0x%x (class %d)", b, c)
}

func getElementID(class uint8, doc *Document) (uint32, error) {
	if class == 1 {
		b := doc.Data[doc.Cursor]

		doc.Cursor++
		return uint32(b), nil
	}
	if class == 2 {
		b := doc.Data[doc.Cursor : doc.Cursor+2]

		doc.Cursor += 2
		return uint32(pack16(b)), nil
	}
	if class == 3 {
		b := doc.Data[doc.Cursor : doc.Cursor+3]

		doc.Cursor += 3
		return uint32(pack24(b)), nil
	}
	if class == 4 {
		b := doc.Data[doc.Cursor : doc.Cursor+4]

		doc.Cursor += 4
		return uint32(pack32(b)), nil
	}

	return 0, fmt.Errorf("Unknown element (ID class %d)", class)
}

func getElementSize(doc *Document) (uint64, error) {
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
	if doc.Cursor >= doc.Length {
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

package webm

import (
	"fmt"
	"math"
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

func (e *Element) GetRawBytes() ([]byte, error) {
	var bytes []byte

	if e.ID <= uint32(math.Pow(2, 7)-2) {
		bytes = append(bytes, byte(e.ID))
	} else if e.ID <= uint32(math.Pow(2, 14)-2) {
		bytes = append(bytes, []byte(unpack(2, uint64(e.ID)))...)
	} else if e.ID <= uint32(math.Pow(2, 21)-2) {
		bytes = append(bytes, []byte(unpack(3, uint64(e.ID)))...)
	} else if e.ID <= uint32(math.Pow(2, 28)-2) {
		bytes = append(bytes, []byte(unpack(4, uint64(e.ID)))...)
	} else {
		return nil, fmt.Errorf("Invalid element (invalid id)")
	}

	if e.Size <= uint64(math.Pow(2, 7)-2) {
		bytes = append(bytes, byte(e.Size))
	} else if e.Size <= uint64(math.Pow(2, 14)-2) {
		bytes = append(bytes, []byte(unpack(2, uint64(e.Size)))...)
	} else if e.Size <= uint64(math.Pow(2, 21)-2) {
		bytes = append(bytes, []byte(unpack(3, uint64(e.Size)))...)
	} else if e.Size <= uint64(math.Pow(2, 28)-2) {
		bytes = append(bytes, []byte(unpack(4, uint64(e.Size)))...)
	} else if e.Size <= uint64(math.Pow(2, 35)-2) {
		bytes = append(bytes, []byte(unpack(5, uint64(e.Size)))...)
	} else if e.Size <= uint64(math.Pow(2, 42)-2) {
		bytes = append(bytes, []byte(unpack(6, uint64(e.Size)))...)
	} else if e.Size <= uint64(math.Pow(2, 49)-2) {
		bytes = append(bytes, []byte(unpack(7, uint64(e.Size)))...)
	} else if e.Size <= uint64(math.Pow(2, 56)-2) {
		bytes = append(bytes, []byte(unpack(8, uint64(e.Size)))...)
	} else {
		return nil, fmt.Errorf("Invalid element (invalid size)")
	}

	bytes = append(bytes, e.Data...)
	return bytes, nil
}

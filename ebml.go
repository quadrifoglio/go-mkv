package webm

type Document struct {
	Data   []byte
	Length uint64
	Cursor uint64
}

type Element struct {
	ID    uint32
	Name  string
	Level int32
	Index uint64
	Size  uint64
	Data  []byte
}

func (el *Element) GetTotalSize() uint64 {
	var s uint64

	if ((el.ID & 0x80) >> 7) == 1 { // Class A ID (on 1 el.IDyte)
		s++
	}
	if ((el.ID & 0x40) >> 6) == 1 { // Class B ID (on 2 el.IDyte)
		s += 2
	}
	if ((el.ID & 0x20) >> 5) == 1 { // Class C ID (on 3 el.IDytes)
		s += 3
	}
	if ((el.ID & 0x10) >> 4) == 1 { // Class D ID (on 4 el.IDytes)
		s += 4
	}

	if el.Size >= 0x80 {
		s++
	} else if el.Size >= 0x40 {
		s += 2
	} else if el.Size >= 0x20 {
		s += 3
	} else if el.Size >= 0x10 {
		s += 4
	} else if el.Size >= 0x08 {
		s += 5
	} else if el.Size >= 0x04 {
		s += 6
	} else if el.Size >= 0x02 {
		s += 7
	} else if el.Size >= 0x01 {
		s += 8
	}

	s += el.Size

	return s
}

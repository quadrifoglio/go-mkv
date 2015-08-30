package ebml

const (
	ElementEBML               uint32 = 0x1a45dfa3
	ElementEBMLVersion        uint32 = 0x4286
	ElementEBMLReadVersion    uint32 = 0x42f7
	ElementEBMLMaxIDLength    uint32 = 0x42f2
	ElementEBMLMaxSizeLength  uint32 = 0x42f3
	ElementDocType            uint32 = 0x4282
	ElementDocTypeVersion     uint32 = 0x4287
	ElementDocTypeReadVersion uint32 = 0x4285

	ElementVoid uint32 = 0xec

	ElementSegment uint32 = 0x18538067
)

func GetElementName(id uint32) string {
	switch id {
	case ElementEBML:
		return "EBML"
	case ElementEBMLVersion:
		return "EBMLVersion"
	case ElementEBMLReadVersion:
		return "EBMLReadVersion"
	case ElementEBMLMaxIDLength:
		return "EBMLMaxIDLength"
	case ElementEBMLMaxSizeLength:
		return "EBMLMaxSizeLength"
	case ElementDocType:
		return "DocType"
	case ElementDocTypeVersion:
		return "DocTypeVersion"
	case ElementDocTypeReadVersion:
		return "DocTypeReadVersion"
	case ElementVoid:
		return "Void"
	case ElementSegment:
		return "Segment"
	default:
		return "Unknown"
	}
}

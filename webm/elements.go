package webm

const (
	ElementUnknown uint32 = 0x0

	ElementEBML               uint32 = 0x1a45dfa3
	ElementEBMLVersion        uint32 = 0x4286
	ElementEBMLReadVersion    uint32 = 0x42f7
	ElementEBMLMaxIDLength    uint32 = 0x42f2
	ElementEBMLMaxSizeLength  uint32 = 0x42f3
	ElementDocType            uint32 = 0x4282
	ElementDocTypeVersion     uint32 = 0x4287
	ElementDocTypeReadVersion uint32 = 0x4285

	ElementVoid    uint32 = 0xec
	ElementSegment uint32 = 0x18538067

	ElementSeekHead     uint32 = 0x114d9b74
	ElementSeek         uint32 = 0x4dbb
	ElementSeekID       uint32 = 0x53ab
	ElementSeekPosition uint32 = 0x53ac

	ElementInfo          uint32 = 0x1549a966
	ElementTimecodeScale uint32 = 0x2ad7b1
	ElementDuration      uint32 = 0x4489
	ElementDateUTC       uint32 = 0x4461
	ElementTitle         uint32 = 0x7ba9
	ElementMuxingApp     uint32 = 0x4d80
	ElementWritingApp    uint32 = 0x5741

	ElementCluster         uint32 = 0x1f43b675
	ElementTimecode        uint32 = 0xe7
	ElementPrevSize        uint32 = 0xab
	ElementSimpleBlock     uint32 = 0xa3
	ElementBlockGroup      uint32 = 0xa0
	ElementBlock           uint32 = 0xa1
	ElementBlockAdditions  uint32 = 0x75a1
	ElementBlockMore       uint32 = 0xa6
	ElementBlockAddID      uint32 = 0xee
	ElementBlockAdditional uint32 = 0xa5
	ElementBlockDuration   uint32 = 0x9b
	ElementReferenceBlock  uint32 = 0xfb
	ElementDiscardPadding  uint32 = 0x75a2
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

	case ElementSeekHead:
		return "SeekHead"
	case ElementSeek:
		return "Seek"
	case ElementSeekID:
		return "SeekID"
	case ElementSeekPosition:
		return "SeekPosition"

	case ElementInfo:
		return "Info"
	case ElementTimecodeScale:
		return "TimecodeScale"
	case ElementDuration:
		return "Duration"
	case ElementDateUTC:
		return "DateUTC"
	case ElementTitle:
		return "Title"
	case ElementMuxingApp:
		return "MuxingApp"
	case ElementWritingApp:
		return "WritingApp"

	case ElementCluster:
		return "Cluster"
	case ElementTimecode:
		return "Timecode"
	case ElementPrevSize:
		return "PrevSize"
	case ElementSimpleBlock:
		return "SimpleBlock"
	case ElementBlockGroup:
		return "BlockGroup"
	case ElementBlock:
		return "Block"
	case ElementBlockAdditions:
		return "BlockAdditions"
	case ElementBlockMore:
		return "BlockMore"
	case ElementBlockAddID:
		return "BlockAddID"
	case ElementBlockAdditional:
		return "BlockAdditional"
	case ElementBlockDuration:
		return "BlockDuration"
	case ElementReferenceBlock:
		return "ReferenceBlock"
	case ElementDiscardPadding:
		return "DiscardPadding"

	case ElementUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

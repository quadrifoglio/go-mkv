package webm

const (
	ElementTypeUnknown uint8 = 0x0
	ElementTypeMaster  uint8 = 0x1
	ElementTypeUint    uint8 = 0x2
	ElementTypeInt     uint8 = 0x3
	ElementTypeString  uint8 = 0x4
	ElementTypeUnicode uint8 = 0x5
	ElementTypeBinary  uint8 = 0x6
	ElementTypeFloat   uint8 = 0x7
	ElementTypeDate    uint8 = 0x8

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

	ElementTracks                  uint32 = 0x1654ae6b
	ElementTrackEntry              uint32 = 0xae
	ElementTrackNumber             uint32 = 0xd7
	ElementTrackUID                uint32 = 0x73c5
	ElementTrackType               uint32 = 0x83
	ElementFlagEnabled             uint32 = 0xb9
	ElementFlagDefault             uint32 = 0x88
	ElementFlagForced              uint32 = 0x55aa
	ElementFlagLacing              uint32 = 0x9c
	ElementDefaultDuration         uint32 = 0x23e383
	ElementName                    uint32 = 0x536e
	ElementLanguage                uint32 = 0x22b59c
	ElementCodecID                 uint32 = 0x86
	ElementCodecPrivate            uint32 = 0x63a2
	ElementCodecName               uint32 = 0x258688
	ElementCodecDelay              uint32 = 0x56aa
	ElementSeekPreRoll             uint32 = 0x56bb
	ElementVideo                   uint32 = 0xe0
	ElementFlagInterlaced          uint32 = 0x9a
	ElementStereoMode              uint32 = 0x53b8
	ElementAlphaMode               uint32 = 0x53c0
	ElementPixelWidth              uint32 = 0xb0
	ElementPixelHeight             uint32 = 0xba
	ElementPixelCropBottom         uint32 = 0x54aa
	ElementPixelCropTop            uint32 = 0x54bb
	ElementPixelCropLeft           uint32 = 0x54cc
	ElementPixelCropRight          uint32 = 0x54dd
	ElementDisplayWidth            uint32 = 0x54b0
	ElementDisplayHeight           uint32 = 0x54ba
	ElementDisplayUint             uint32 = 0x54b2
	ElementAspectRatioType         uint32 = 0x54b3
	ElementAudio                   uint32 = 0xe1
	ElementSamplingFrequency       uint32 = 0xb5
	ElementOutputSamplingFrequency uint32 = 0x78b5
	ElementChannels                uint32 = 0x9f
	ElementBitDepth                uint32 = 0x6264
	ElementContentEncoding         uint32 = 0x6240
	ElementContentEncodingOrder    uint32 = 0x5031
	ElementContentEncodingScope    uint32 = 0x5032
	ElementContentEncodingType     uint32 = 0x5033
	ElementContentEncryption       uint32 = 0x5035
	ElementContentEncAlgo          uint32 = 0x47e1
	ElementContentEncKeyID         uint32 = 0x47e2
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

	case ElementTracks:
		return "ElementTracks"
	case ElementTrackEntry:
		return "ElementTrackEntry"
	case ElementTrackNumber:
		return "ElementTrackNumber"
	case ElementTrackUID:
		return "ElementTrackUID"
	case ElementTrackType:
		return "ElementTrackType"
	case ElementFlagEnabled:
		return "ElementFlagEnabled"
	case ElementFlagDefault:
		return "ElementFlagDefault"
	case ElementFlagForced:
		return "ElementFlagForced"
	case ElementFlagLacing:
		return "ElementFlagLacing"
	case ElementDefaultDuration:
		return "ElementDefaultDuration"
	case ElementName:
		return "ElementName"
	case ElementLanguage:
		return "ElementLanguage"
	case ElementCodecID:
		return "ElementCodecID"
	case ElementCodecPrivate:
		return "ElementCodecPrivate"
	case ElementCodecName:
		return "ElementCodecName"
	case ElementCodecDelay:
		return "ElementCodecDelay"
	case ElementSeekPreRoll:
		return "ElementSeekPreRoll"
	case ElementVideo:
		return "ElementVideo"
	case ElementFlagInterlaced:
		return "ElementFlagInterlaced"
	case ElementStereoMode:
		return "ElementStereoMode"
	case ElementAlphaMode:
		return "ElementAlphaMode"
	case ElementPixelWidth:
		return "ElementPixelWidth"
	case ElementPixelHeight:
		return "ElementPixelHeight"
	case ElementPixelCropBottom:
		return "ElementPixelCropBottom"
	case ElementPixelCropTop:
		return "ElementPixelCropTop"
	case ElementPixelCropLeft:
		return "ElementPixelCropLeft"
	case ElementPixelCropRight:
		return "ElementPixelCropRight"
	case ElementDisplayWidth:
		return "ElementDisplayWidth"
	case ElementDisplayHeight:
		return "ElementDisplayHeight"
	case ElementDisplayUint:
		return "ElementDisplayUint"
	case ElementAspectRatioType:
		return "ElementAspectRatioType"
	case ElementAudio:
		return "ElementAudio"
	case ElementSamplingFrequency:
		return "ElementSamplingFrequency"
	case ElementOutputSamplingFrequency:
		return "ElementOutputSamplingFrequency"
	case ElementChannels:
		return "ElementChannels"
	case ElementBitDepth:
		return "ElementBitDepth"
	case ElementContentEncoding:
		return "ElementContentEncoding"
	case ElementContentEncodingOrder:
		return "ElementContentEncodingOrder"
	case ElementContentEncodingScope:
		return "ElementContentEncodingScope"
	case ElementContentEncodingType:
		return "ElementContentEncodingType"
	case ElementContentEncryption:
		return "ElementContentEncryption"
	case ElementContentEncAlgo:
		return "ElementContentEncAlgo"
	case ElementContentEncKeyID:
		return "ElementContentEncKeyID"

	case ElementUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

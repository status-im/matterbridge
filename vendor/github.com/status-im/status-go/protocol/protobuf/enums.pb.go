// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: enums.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageType int32

const (
	MessageType_UNKNOWN_MESSAGE_TYPE MessageType = 0
	MessageType_ONE_TO_ONE           MessageType = 1
	MessageType_PUBLIC_GROUP         MessageType = 2
	MessageType_PRIVATE_GROUP        MessageType = 3
	// Only local
	MessageType_SYSTEM_MESSAGE_PRIVATE_GROUP MessageType = 4
	MessageType_COMMUNITY_CHAT               MessageType = 5
	// Only local
	MessageType_SYSTEM_MESSAGE_GAP MessageType = 6
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "UNKNOWN_MESSAGE_TYPE",
		1: "ONE_TO_ONE",
		2: "PUBLIC_GROUP",
		3: "PRIVATE_GROUP",
		4: "SYSTEM_MESSAGE_PRIVATE_GROUP",
		5: "COMMUNITY_CHAT",
		6: "SYSTEM_MESSAGE_GAP",
	}
	MessageType_value = map[string]int32{
		"UNKNOWN_MESSAGE_TYPE":         0,
		"ONE_TO_ONE":                   1,
		"PUBLIC_GROUP":                 2,
		"PRIVATE_GROUP":                3,
		"SYSTEM_MESSAGE_PRIVATE_GROUP": 4,
		"COMMUNITY_CHAT":               5,
		"SYSTEM_MESSAGE_GAP":           6,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{0}
}

type ImageFormat int32

const (
	ImageFormat_UNKNOWN_IMAGE_FORMAT ImageFormat = 0
	// Raster image files is payload data that can be read as a raster image
	ImageFormat_PNG  ImageFormat = 1
	ImageFormat_JPEG ImageFormat = 2
	ImageFormat_WEBP ImageFormat = 3
	ImageFormat_GIF  ImageFormat = 4
)

// Enum value maps for ImageFormat.
var (
	ImageFormat_name = map[int32]string{
		0: "UNKNOWN_IMAGE_FORMAT",
		1: "PNG",
		2: "JPEG",
		3: "WEBP",
		4: "GIF",
	}
	ImageFormat_value = map[string]int32{
		"UNKNOWN_IMAGE_FORMAT": 0,
		"PNG":                  1,
		"JPEG":                 2,
		"WEBP":                 3,
		"GIF":                  4,
	}
)

func (x ImageFormat) Enum() *ImageFormat {
	p := new(ImageFormat)
	*p = x
	return p
}

func (x ImageFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[1].Descriptor()
}

func (ImageFormat) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[1]
}

func (x ImageFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageFormat.Descriptor instead.
func (ImageFormat) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{1}
}

type CommunityTokenType int32

const (
	CommunityTokenType_UNKNOWN_TOKEN_TYPE CommunityTokenType = 0
	CommunityTokenType_ERC20              CommunityTokenType = 1
	CommunityTokenType_ERC721             CommunityTokenType = 2
	CommunityTokenType_ENS                CommunityTokenType = 3
)

// Enum value maps for CommunityTokenType.
var (
	CommunityTokenType_name = map[int32]string{
		0: "UNKNOWN_TOKEN_TYPE",
		1: "ERC20",
		2: "ERC721",
		3: "ENS",
	}
	CommunityTokenType_value = map[string]int32{
		"UNKNOWN_TOKEN_TYPE": 0,
		"ERC20":              1,
		"ERC721":             2,
		"ENS":                3,
	}
)

func (x CommunityTokenType) Enum() *CommunityTokenType {
	p := new(CommunityTokenType)
	*p = x
	return p
}

func (x CommunityTokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommunityTokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[2].Descriptor()
}

func (CommunityTokenType) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[2]
}

func (x CommunityTokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommunityTokenType.Descriptor instead.
func (CommunityTokenType) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{2}
}

var File_enums_proto protoreflect.FileDescriptor

var file_enums_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2a, 0xaa, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x4f, 0x4e, 0x45, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x5f, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d,
	0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4d, 0x4d,
	0x55, 0x4e, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x48, 0x41, 0x54, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12,
	0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x47,
	0x41, 0x50, 0x10, 0x06, 0x2a, 0x4d, 0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x49,
	0x4d, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x50, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x50, 0x45, 0x47, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x57, 0x45, 0x42, 0x50, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x49,
	0x46, 0x10, 0x04, 0x2a, 0x4c, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x43, 0x32, 0x30, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x45, 0x52, 0x43, 0x37, 0x32, 0x31, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4e, 0x53, 0x10,
	0x03, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_proto_rawDescOnce sync.Once
	file_enums_proto_rawDescData = file_enums_proto_rawDesc
)

func file_enums_proto_rawDescGZIP() []byte {
	file_enums_proto_rawDescOnce.Do(func() {
		file_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_proto_rawDescData)
	})
	return file_enums_proto_rawDescData
}

var file_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_enums_proto_goTypes = []interface{}{
	(MessageType)(0),        // 0: protobuf.MessageType
	(ImageFormat)(0),        // 1: protobuf.ImageFormat
	(CommunityTokenType)(0), // 2: protobuf.CommunityTokenType
}
var file_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_proto_init() }
func file_enums_proto_init() {
	if File_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_proto_goTypes,
		DependencyIndexes: file_enums_proto_depIdxs,
		EnumInfos:         file_enums_proto_enumTypes,
	}.Build()
	File_enums_proto = out.File
	file_enums_proto_rawDesc = nil
	file_enums_proto_goTypes = nil
	file_enums_proto_depIdxs = nil
}

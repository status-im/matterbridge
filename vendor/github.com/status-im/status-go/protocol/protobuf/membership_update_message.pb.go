// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: membership_update_message.proto

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

type MembershipUpdateEvent_EventType int32

const (
	MembershipUpdateEvent_UNKNOWN        MembershipUpdateEvent_EventType = 0
	MembershipUpdateEvent_CHAT_CREATED   MembershipUpdateEvent_EventType = 1
	MembershipUpdateEvent_NAME_CHANGED   MembershipUpdateEvent_EventType = 2
	MembershipUpdateEvent_MEMBERS_ADDED  MembershipUpdateEvent_EventType = 3
	MembershipUpdateEvent_MEMBER_JOINED  MembershipUpdateEvent_EventType = 4
	MembershipUpdateEvent_MEMBER_REMOVED MembershipUpdateEvent_EventType = 5
	MembershipUpdateEvent_ADMINS_ADDED   MembershipUpdateEvent_EventType = 6
	MembershipUpdateEvent_ADMIN_REMOVED  MembershipUpdateEvent_EventType = 7
	MembershipUpdateEvent_COLOR_CHANGED  MembershipUpdateEvent_EventType = 8
	MembershipUpdateEvent_IMAGE_CHANGED  MembershipUpdateEvent_EventType = 9
)

// Enum value maps for MembershipUpdateEvent_EventType.
var (
	MembershipUpdateEvent_EventType_name = map[int32]string{
		0: "UNKNOWN",
		1: "CHAT_CREATED",
		2: "NAME_CHANGED",
		3: "MEMBERS_ADDED",
		4: "MEMBER_JOINED",
		5: "MEMBER_REMOVED",
		6: "ADMINS_ADDED",
		7: "ADMIN_REMOVED",
		8: "COLOR_CHANGED",
		9: "IMAGE_CHANGED",
	}
	MembershipUpdateEvent_EventType_value = map[string]int32{
		"UNKNOWN":        0,
		"CHAT_CREATED":   1,
		"NAME_CHANGED":   2,
		"MEMBERS_ADDED":  3,
		"MEMBER_JOINED":  4,
		"MEMBER_REMOVED": 5,
		"ADMINS_ADDED":   6,
		"ADMIN_REMOVED":  7,
		"COLOR_CHANGED":  8,
		"IMAGE_CHANGED":  9,
	}
)

func (x MembershipUpdateEvent_EventType) Enum() *MembershipUpdateEvent_EventType {
	p := new(MembershipUpdateEvent_EventType)
	*p = x
	return p
}

func (x MembershipUpdateEvent_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MembershipUpdateEvent_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_membership_update_message_proto_enumTypes[0].Descriptor()
}

func (MembershipUpdateEvent_EventType) Type() protoreflect.EnumType {
	return &file_membership_update_message_proto_enumTypes[0]
}

func (x MembershipUpdateEvent_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MembershipUpdateEvent_EventType.Descriptor instead.
func (MembershipUpdateEvent_EventType) EnumDescriptor() ([]byte, []int) {
	return file_membership_update_message_proto_rawDescGZIP(), []int{0, 0}
}

type MembershipUpdateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lamport timestamp of the event
	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	// List of public keys of objects of the action
	Members []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	// Name of the chat for the CHAT_CREATED/NAME_CHANGED event types
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the event
	Type MembershipUpdateEvent_EventType `protobuf:"varint,4,opt,name=type,proto3,enum=protobuf.MembershipUpdateEvent_EventType" json:"type,omitempty"`
	// Color of the chat for the CHAT_CREATED/COLOR_CHANGED event types
	Color string `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	// Chat image
	Image []byte `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *MembershipUpdateEvent) Reset() {
	*x = MembershipUpdateEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_membership_update_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MembershipUpdateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipUpdateEvent) ProtoMessage() {}

func (x *MembershipUpdateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_membership_update_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembershipUpdateEvent.ProtoReflect.Descriptor instead.
func (*MembershipUpdateEvent) Descriptor() ([]byte, []int) {
	return file_membership_update_message_proto_rawDescGZIP(), []int{0}
}

func (x *MembershipUpdateEvent) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *MembershipUpdateEvent) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *MembershipUpdateEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MembershipUpdateEvent) GetType() MembershipUpdateEvent_EventType {
	if x != nil {
		return x.Type
	}
	return MembershipUpdateEvent_UNKNOWN
}

func (x *MembershipUpdateEvent) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *MembershipUpdateEvent) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

// MembershipUpdateMessage is a message used to propagate information
// about group membership changes.
// For more information, see https://github.com/status-im/specs/blob/master/status-group-chats-spec.md.
type MembershipUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The chat id of the private group chat
	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	// A list of events for this group chat, first x bytes are the signature, then is a
	// protobuf encoded MembershipUpdateEvent
	Events [][]byte `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	// An optional chat message
	//
	// Types that are assignable to ChatEntity:
	//
	//	*MembershipUpdateMessage_Message
	//	*MembershipUpdateMessage_EmojiReaction
	ChatEntity isMembershipUpdateMessage_ChatEntity `protobuf_oneof:"chat_entity"`
}

func (x *MembershipUpdateMessage) Reset() {
	*x = MembershipUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_membership_update_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MembershipUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipUpdateMessage) ProtoMessage() {}

func (x *MembershipUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_membership_update_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembershipUpdateMessage.ProtoReflect.Descriptor instead.
func (*MembershipUpdateMessage) Descriptor() ([]byte, []int) {
	return file_membership_update_message_proto_rawDescGZIP(), []int{1}
}

func (x *MembershipUpdateMessage) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *MembershipUpdateMessage) GetEvents() [][]byte {
	if x != nil {
		return x.Events
	}
	return nil
}

func (m *MembershipUpdateMessage) GetChatEntity() isMembershipUpdateMessage_ChatEntity {
	if m != nil {
		return m.ChatEntity
	}
	return nil
}

func (x *MembershipUpdateMessage) GetMessage() *ChatMessage {
	if x, ok := x.GetChatEntity().(*MembershipUpdateMessage_Message); ok {
		return x.Message
	}
	return nil
}

func (x *MembershipUpdateMessage) GetEmojiReaction() *EmojiReaction {
	if x, ok := x.GetChatEntity().(*MembershipUpdateMessage_EmojiReaction); ok {
		return x.EmojiReaction
	}
	return nil
}

type isMembershipUpdateMessage_ChatEntity interface {
	isMembershipUpdateMessage_ChatEntity()
}

type MembershipUpdateMessage_Message struct {
	Message *ChatMessage `protobuf:"bytes,3,opt,name=message,proto3,oneof"`
}

type MembershipUpdateMessage_EmojiReaction struct {
	EmojiReaction *EmojiReaction `protobuf:"bytes,4,opt,name=emoji_reaction,json=emojiReaction,proto3,oneof"`
}

func (*MembershipUpdateMessage_Message) isMembershipUpdateMessage_ChatEntity() {}

func (*MembershipUpdateMessage_EmojiReaction) isMembershipUpdateMessage_ChatEntity() {}

var File_membership_update_message_proto protoreflect.FileDescriptor

var file_membership_update_message_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x12, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x03, 0x0a, 0x15, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0xc1,
	0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x48, 0x41,
	0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53, 0x5f, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x52, 0x45,
	0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x53, 0x5f, 0x41, 0x44, 0x44, 0x45, 0x44, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x08, 0x12,
	0x11, 0x0a, 0x0d, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44,
	0x10, 0x09, 0x22, 0xce, 0x01, 0x0a, 0x17, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x5f, 0x72, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_membership_update_message_proto_rawDescOnce sync.Once
	file_membership_update_message_proto_rawDescData = file_membership_update_message_proto_rawDesc
)

func file_membership_update_message_proto_rawDescGZIP() []byte {
	file_membership_update_message_proto_rawDescOnce.Do(func() {
		file_membership_update_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_membership_update_message_proto_rawDescData)
	})
	return file_membership_update_message_proto_rawDescData
}

var file_membership_update_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_membership_update_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_membership_update_message_proto_goTypes = []interface{}{
	(MembershipUpdateEvent_EventType)(0), // 0: protobuf.MembershipUpdateEvent.EventType
	(*MembershipUpdateEvent)(nil),        // 1: protobuf.MembershipUpdateEvent
	(*MembershipUpdateMessage)(nil),      // 2: protobuf.MembershipUpdateMessage
	(*ChatMessage)(nil),                  // 3: protobuf.ChatMessage
	(*EmojiReaction)(nil),                // 4: protobuf.EmojiReaction
}
var file_membership_update_message_proto_depIdxs = []int32{
	0, // 0: protobuf.MembershipUpdateEvent.type:type_name -> protobuf.MembershipUpdateEvent.EventType
	3, // 1: protobuf.MembershipUpdateMessage.message:type_name -> protobuf.ChatMessage
	4, // 2: protobuf.MembershipUpdateMessage.emoji_reaction:type_name -> protobuf.EmojiReaction
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_membership_update_message_proto_init() }
func file_membership_update_message_proto_init() {
	if File_membership_update_message_proto != nil {
		return
	}
	file_chat_message_proto_init()
	file_emoji_reaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_membership_update_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MembershipUpdateEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_membership_update_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MembershipUpdateMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_membership_update_message_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*MembershipUpdateMessage_Message)(nil),
		(*MembershipUpdateMessage_EmojiReaction)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_membership_update_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_membership_update_message_proto_goTypes,
		DependencyIndexes: file_membership_update_message_proto_depIdxs,
		EnumInfos:         file_membership_update_message_proto_enumTypes,
		MessageInfos:      file_membership_update_message_proto_msgTypes,
	}.Build()
	File_membership_update_message_proto = out.File
	file_membership_update_message_proto_rawDesc = nil
	file_membership_update_message_proto_goTypes = nil
	file_membership_update_message_proto_depIdxs = nil
}

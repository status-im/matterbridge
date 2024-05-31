// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: community_privileged_user_sync_message.proto

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

type CommunityPrivilegedUserSyncMessage_EventType int32

const (
	CommunityPrivilegedUserSyncMessage_UNKNOWN                                CommunityPrivilegedUserSyncMessage_EventType = 0
	CommunityPrivilegedUserSyncMessage_CONTROL_NODE_ACCEPT_REQUEST_TO_JOIN    CommunityPrivilegedUserSyncMessage_EventType = 1
	CommunityPrivilegedUserSyncMessage_CONTROL_NODE_REJECT_REQUEST_TO_JOIN    CommunityPrivilegedUserSyncMessage_EventType = 2
	CommunityPrivilegedUserSyncMessage_CONTROL_NODE_ALL_SYNC_REQUESTS_TO_JOIN CommunityPrivilegedUserSyncMessage_EventType = 3
)

// Enum value maps for CommunityPrivilegedUserSyncMessage_EventType.
var (
	CommunityPrivilegedUserSyncMessage_EventType_name = map[int32]string{
		0: "UNKNOWN",
		1: "CONTROL_NODE_ACCEPT_REQUEST_TO_JOIN",
		2: "CONTROL_NODE_REJECT_REQUEST_TO_JOIN",
		3: "CONTROL_NODE_ALL_SYNC_REQUESTS_TO_JOIN",
	}
	CommunityPrivilegedUserSyncMessage_EventType_value = map[string]int32{
		"UNKNOWN":                                0,
		"CONTROL_NODE_ACCEPT_REQUEST_TO_JOIN":    1,
		"CONTROL_NODE_REJECT_REQUEST_TO_JOIN":    2,
		"CONTROL_NODE_ALL_SYNC_REQUESTS_TO_JOIN": 3,
	}
)

func (x CommunityPrivilegedUserSyncMessage_EventType) Enum() *CommunityPrivilegedUserSyncMessage_EventType {
	p := new(CommunityPrivilegedUserSyncMessage_EventType)
	*p = x
	return p
}

func (x CommunityPrivilegedUserSyncMessage_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommunityPrivilegedUserSyncMessage_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_community_privileged_user_sync_message_proto_enumTypes[0].Descriptor()
}

func (CommunityPrivilegedUserSyncMessage_EventType) Type() protoreflect.EnumType {
	return &file_community_privileged_user_sync_message_proto_enumTypes[0]
}

func (x CommunityPrivilegedUserSyncMessage_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommunityPrivilegedUserSyncMessage_EventType.Descriptor instead.
func (CommunityPrivilegedUserSyncMessage_EventType) EnumDescriptor() ([]byte, []int) {
	return file_community_privileged_user_sync_message_proto_rawDescGZIP(), []int{0, 0}
}

type CommunityPrivilegedUserSyncMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock              uint64                                       `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Type               CommunityPrivilegedUserSyncMessage_EventType `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.CommunityPrivilegedUserSyncMessage_EventType" json:"type,omitempty"`
	CommunityId        []byte                                       `protobuf:"bytes,3,opt,name=community_id,json=communityId,proto3" json:"community_id,omitempty"`
	RequestToJoin      map[string]*CommunityRequestToJoin           `protobuf:"bytes,4,rep,name=request_to_join,json=requestToJoin,proto3" json:"request_to_join,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SyncRequestsToJoin []*SyncCommunityRequestsToJoin               `protobuf:"bytes,5,rep,name=sync_requests_to_join,json=syncRequestsToJoin,proto3" json:"sync_requests_to_join,omitempty"`
}

func (x *CommunityPrivilegedUserSyncMessage) Reset() {
	*x = CommunityPrivilegedUserSyncMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_privileged_user_sync_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunityPrivilegedUserSyncMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunityPrivilegedUserSyncMessage) ProtoMessage() {}

func (x *CommunityPrivilegedUserSyncMessage) ProtoReflect() protoreflect.Message {
	mi := &file_community_privileged_user_sync_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunityPrivilegedUserSyncMessage.ProtoReflect.Descriptor instead.
func (*CommunityPrivilegedUserSyncMessage) Descriptor() ([]byte, []int) {
	return file_community_privileged_user_sync_message_proto_rawDescGZIP(), []int{0}
}

func (x *CommunityPrivilegedUserSyncMessage) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *CommunityPrivilegedUserSyncMessage) GetType() CommunityPrivilegedUserSyncMessage_EventType {
	if x != nil {
		return x.Type
	}
	return CommunityPrivilegedUserSyncMessage_UNKNOWN
}

func (x *CommunityPrivilegedUserSyncMessage) GetCommunityId() []byte {
	if x != nil {
		return x.CommunityId
	}
	return nil
}

func (x *CommunityPrivilegedUserSyncMessage) GetRequestToJoin() map[string]*CommunityRequestToJoin {
	if x != nil {
		return x.RequestToJoin
	}
	return nil
}

func (x *CommunityPrivilegedUserSyncMessage) GetSyncRequestsToJoin() []*SyncCommunityRequestsToJoin {
	if x != nil {
		return x.SyncRequestsToJoin
	}
	return nil
}

var File_community_privileged_user_sync_message_proto protoreflect.FileDescriptor

var file_community_privileged_user_sync_message_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x76,
	0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x61, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x04, 0x0a, 0x22, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x4a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c,
	0x65, 0x67, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x67, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x74, 0x6f, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x4a, 0x6f, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x4a, 0x6f, 0x69, 0x6e, 0x12,
	0x58, 0x0a, 0x15, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x5f, 0x74, 0x6f, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x54,
	0x6f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x12, 0x73, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x54, 0x6f, 0x4a, 0x6f, 0x69, 0x6e, 0x1a, 0x62, 0x0a, 0x12, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x4a, 0x6f, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x96, 0x01,
	0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x4f, 0x4e, 0x54,
	0x52, 0x4f, 0x4c, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x10,
	0x01, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x4e, 0x4f, 0x44,
	0x45, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x5f, 0x54, 0x4f, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x4f,
	0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x53,
	0x59, 0x4e, 0x43, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x53, 0x5f, 0x54, 0x4f, 0x5f,
	0x4a, 0x4f, 0x49, 0x4e, 0x10, 0x03, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_community_privileged_user_sync_message_proto_rawDescOnce sync.Once
	file_community_privileged_user_sync_message_proto_rawDescData = file_community_privileged_user_sync_message_proto_rawDesc
)

func file_community_privileged_user_sync_message_proto_rawDescGZIP() []byte {
	file_community_privileged_user_sync_message_proto_rawDescOnce.Do(func() {
		file_community_privileged_user_sync_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_community_privileged_user_sync_message_proto_rawDescData)
	})
	return file_community_privileged_user_sync_message_proto_rawDescData
}

var file_community_privileged_user_sync_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_community_privileged_user_sync_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_community_privileged_user_sync_message_proto_goTypes = []interface{}{
	(CommunityPrivilegedUserSyncMessage_EventType)(0), // 0: protobuf.CommunityPrivilegedUserSyncMessage.EventType
	(*CommunityPrivilegedUserSyncMessage)(nil),        // 1: protobuf.CommunityPrivilegedUserSyncMessage
	nil,                                 // 2: protobuf.CommunityPrivilegedUserSyncMessage.RequestToJoinEntry
	(*SyncCommunityRequestsToJoin)(nil), // 3: protobuf.SyncCommunityRequestsToJoin
	(*CommunityRequestToJoin)(nil),      // 4: protobuf.CommunityRequestToJoin
}
var file_community_privileged_user_sync_message_proto_depIdxs = []int32{
	0, // 0: protobuf.CommunityPrivilegedUserSyncMessage.type:type_name -> protobuf.CommunityPrivilegedUserSyncMessage.EventType
	2, // 1: protobuf.CommunityPrivilegedUserSyncMessage.request_to_join:type_name -> protobuf.CommunityPrivilegedUserSyncMessage.RequestToJoinEntry
	3, // 2: protobuf.CommunityPrivilegedUserSyncMessage.sync_requests_to_join:type_name -> protobuf.SyncCommunityRequestsToJoin
	4, // 3: protobuf.CommunityPrivilegedUserSyncMessage.RequestToJoinEntry.value:type_name -> protobuf.CommunityRequestToJoin
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_community_privileged_user_sync_message_proto_init() }
func file_community_privileged_user_sync_message_proto_init() {
	if File_community_privileged_user_sync_message_proto != nil {
		return
	}
	file_communities_proto_init()
	file_pairing_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_community_privileged_user_sync_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunityPrivilegedUserSyncMessage); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_community_privileged_user_sync_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_community_privileged_user_sync_message_proto_goTypes,
		DependencyIndexes: file_community_privileged_user_sync_message_proto_depIdxs,
		EnumInfos:         file_community_privileged_user_sync_message_proto_enumTypes,
		MessageInfos:      file_community_privileged_user_sync_message_proto_msgTypes,
	}.Build()
	File_community_privileged_user_sync_message_proto = out.File
	file_community_privileged_user_sync_message_proto_rawDesc = nil
	file_community_privileged_user_sync_message_proto_goTypes = nil
	file_community_privileged_user_sync_message_proto_depIdxs = nil
}

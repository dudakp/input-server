// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: Typing.proto

package grpc

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

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Typing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Typing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_Typing_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSessionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSessionRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type JoinSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *JoinSessionRequest) Reset() {
	*x = JoinSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Typing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinSessionRequest) ProtoMessage() {}

func (x *JoinSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Typing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinSessionRequest.ProtoReflect.Descriptor instead.
func (*JoinSessionRequest) Descriptor() ([]byte, []int) {
	return file_Typing_proto_rawDescGZIP(), []int{1}
}

func (x *JoinSessionRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Users  []*User `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	Region string  `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Typing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_Typing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_Typing_proto_rawDescGZIP(), []int{2}
}

func (x *Session) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Session) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Session) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Session) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type ListSessionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *ListSessionsRequest) Reset() {
	*x = ListSessionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Typing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSessionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSessionsRequest) ProtoMessage() {}

func (x *ListSessionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Typing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSessionsRequest.ProtoReflect.Descriptor instead.
func (*ListSessionsRequest) Descriptor() ([]byte, []int) {
	return file_Typing_proto_rawDescGZIP(), []int{3}
}

func (x *ListSessionsRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Typing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_Typing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_Typing_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_Typing_proto protoreflect.FileDescriptor

var file_Typing_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x54, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x74, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x12, 0x4a, 0x6f,
	0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x69, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x74, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xcc, 0x01, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x2e, 0x74, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x74, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x2e, 0x74, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x74, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x40, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1b, 0x2e, 0x74, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x74, 0x79, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x64, 0x61, 0x6b, 0x70, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Typing_proto_rawDescOnce sync.Once
	file_Typing_proto_rawDescData = file_Typing_proto_rawDesc
)

func file_Typing_proto_rawDescGZIP() []byte {
	file_Typing_proto_rawDescOnce.Do(func() {
		file_Typing_proto_rawDescData = protoimpl.X.CompressGZIP(file_Typing_proto_rawDescData)
	})
	return file_Typing_proto_rawDescData
}

var file_Typing_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Typing_proto_goTypes = []interface{}{
	(*CreateSessionRequest)(nil), // 0: typing.CreateSessionRequest
	(*JoinSessionRequest)(nil),   // 1: typing.JoinSessionRequest
	(*Session)(nil),              // 2: typing.Session
	(*ListSessionsRequest)(nil),  // 3: typing.ListSessionsRequest
	(*User)(nil),                 // 4: typing.User
}
var file_Typing_proto_depIdxs = []int32{
	4, // 0: typing.Session.users:type_name -> typing.User
	0, // 1: typing.Typing.CreateSession:input_type -> typing.CreateSessionRequest
	1, // 2: typing.Typing.JoinSession:input_type -> typing.JoinSessionRequest
	3, // 3: typing.Typing.ListSessions:input_type -> typing.ListSessionsRequest
	2, // 4: typing.Typing.CreateSession:output_type -> typing.Session
	2, // 5: typing.Typing.JoinSession:output_type -> typing.Session
	2, // 6: typing.Typing.ListSessions:output_type -> typing.Session
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Typing_proto_init() }
func file_Typing_proto_init() {
	if File_Typing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Typing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
		file_Typing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinSessionRequest); i {
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
		file_Typing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_Typing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSessionsRequest); i {
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
		file_Typing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_Typing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Typing_proto_goTypes,
		DependencyIndexes: file_Typing_proto_depIdxs,
		MessageInfos:      file_Typing_proto_msgTypes,
	}.Build()
	File_Typing_proto = out.File
	file_Typing_proto_rawDesc = nil
	file_Typing_proto_goTypes = nil
	file_Typing_proto_depIdxs = nil
}

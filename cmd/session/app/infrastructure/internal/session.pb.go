// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: session.proto

package infrastructure

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Region  string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	LevelId string `protobuf:"bytes,3,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[0]
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
	return file_session_proto_rawDescGZIP(), []int{0}
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

func (x *CreateSessionRequest) GetLevelId() string {
	if x != nil {
		return x.LevelId
	}
	return ""
}

type JoinSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//
	//	*JoinSessionRequest_Join
	//	*JoinSessionRequest_Ping
	Event isJoinSessionRequest_Event `protobuf_oneof:"event"`
}

func (x *JoinSessionRequest) Reset() {
	*x = JoinSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinSessionRequest) ProtoMessage() {}

func (x *JoinSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[1]
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
	return file_session_proto_rawDescGZIP(), []int{1}
}

func (m *JoinSessionRequest) GetEvent() isJoinSessionRequest_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *JoinSessionRequest) GetJoin() *JoinEvent {
	if x, ok := x.GetEvent().(*JoinSessionRequest_Join); ok {
		return x.Join
	}
	return nil
}

func (x *JoinSessionRequest) GetPing() *PingEvent {
	if x, ok := x.GetEvent().(*JoinSessionRequest_Ping); ok {
		return x.Ping
	}
	return nil
}

type isJoinSessionRequest_Event interface {
	isJoinSessionRequest_Event()
}

type JoinSessionRequest_Join struct {
	Join *JoinEvent `protobuf:"bytes,1,opt,name=join,proto3,oneof"`
}

type JoinSessionRequest_Ping struct {
	Ping *PingEvent `protobuf:"bytes,2,opt,name=ping,proto3,oneof"`
}

func (*JoinSessionRequest_Join) isJoinSessionRequest_Event() {}

func (*JoinSessionRequest_Ping) isJoinSessionRequest_Event() {}

type JoinEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *JoinEvent) Reset() {
	*x = JoinEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinEvent) ProtoMessage() {}

func (x *JoinEvent) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinEvent.ProtoReflect.Descriptor instead.
func (*JoinEvent) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{2}
}

func (x *JoinEvent) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type PingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelState *LevelState          `protobuf:"bytes,1,opt,name=level_state,json=levelState,proto3" json:"level_state,omitempty"`
	Timestamp  *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PingEvent) Reset() {
	*x = PingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingEvent) ProtoMessage() {}

func (x *PingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingEvent.ProtoReflect.Descriptor instead.
func (*PingEvent) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{3}
}

func (x *PingEvent) GetLevelState() *LevelState {
	if x != nil {
		return x.LevelState
	}
	return nil
}

func (x *PingEvent) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type LevelState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Progress  int32  `protobuf:"varint,2,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *LevelState) Reset() {
	*x = LevelState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelState) ProtoMessage() {}

func (x *LevelState) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelState.ProtoReflect.Descriptor instead.
func (*LevelState) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{4}
}

func (x *LevelState) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *LevelState) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level      *Level `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	NumPlayers int32  `protobuf:"varint,4,opt,name=numPlayers,proto3" json:"numPlayers,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[5]
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
	return file_session_proto_rawDescGZIP(), []int{5}
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

func (x *Session) GetLevel() *Level {
	if x != nil {
		return x.Level
	}
	return nil
}

func (x *Session) GetNumPlayers() int32 {
	if x != nil {
		return x.NumPlayers
	}
	return 0
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
		mi := &file_session_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSessionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSessionsRequest) ProtoMessage() {}

func (x *ListSessionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[6]
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
	return file_session_proto_rawDescGZIP(), []int{6}
}

func (x *ListSessionsRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type CreateLevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateLevelRequest) Reset() {
	*x = CreateLevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLevelRequest) ProtoMessage() {}

func (x *CreateLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLevelRequest.ProtoReflect.Descriptor instead.
func (*CreateLevelRequest) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{7}
}

func (x *CreateLevelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLevelRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Level struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Difficulty int32  `protobuf:"varint,3,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
}

func (x *Level) Reset() {
	*x = Level{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Level) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Level) ProtoMessage() {}

func (x *Level) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Level.ProtoReflect.Descriptor instead.
func (*Level) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{8}
}

func (x *Level) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Level) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Level) GetDifficulty() int32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

type ListLevelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListLevelsRequest) Reset() {
	*x = ListLevelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_session_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLevelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLevelsRequest) ProtoMessage() {}

func (x *ListLevelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_session_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLevelsRequest.ProtoReflect.Descriptor instead.
func (*ListLevelsRequest) Descriptor() ([]byte, []int) {
	return file_session_proto_rawDescGZIP(), []int{9}
}

var File_session_proto protoreflect.FileDescriptor

var file_session_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a,
	0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x09, 0x4a,
	0x6f, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x47, 0x0a, 0x0a, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x73, 0x0a,
	0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x22, 0x2d, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x22, 0x3c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x4b, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x22, 0x13, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0xd3, 0x01, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1b, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x32, 0x83, 0x01, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x3c, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x1b, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x1a, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x00, 0x30, 0x01, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x64, 0x61,
	0x6b, 0x70, 0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x63, 0x6d, 0x64, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_session_proto_rawDescOnce sync.Once
	file_session_proto_rawDescData = file_session_proto_rawDesc
)

func file_session_proto_rawDescGZIP() []byte {
	file_session_proto_rawDescOnce.Do(func() {
		file_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_session_proto_rawDescData)
	})
	return file_session_proto_rawDescData
}

var file_session_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_session_proto_goTypes = []interface{}{
	(*CreateSessionRequest)(nil), // 0: session.CreateSessionRequest
	(*JoinSessionRequest)(nil),   // 1: session.JoinSessionRequest
	(*JoinEvent)(nil),            // 2: session.JoinEvent
	(*PingEvent)(nil),            // 3: session.PingEvent
	(*LevelState)(nil),           // 4: session.LevelState
	(*Session)(nil),              // 5: session.Session
	(*ListSessionsRequest)(nil),  // 6: session.ListSessionsRequest
	(*CreateLevelRequest)(nil),   // 7: session.CreateLevelRequest
	(*Level)(nil),                // 8: session.Level
	(*ListLevelsRequest)(nil),    // 9: session.ListLevelsRequest
	(*timestamp.Timestamp)(nil),  // 10: google.protobuf.Timestamp
}
var file_session_proto_depIdxs = []int32{
	2,  // 0: session.JoinSessionRequest.join:type_name -> session.JoinEvent
	3,  // 1: session.JoinSessionRequest.ping:type_name -> session.PingEvent
	4,  // 2: session.PingEvent.level_state:type_name -> session.LevelState
	10, // 3: session.PingEvent.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 4: session.Session.level:type_name -> session.Level
	0,  // 5: session.session.CreateSession:input_type -> session.CreateSessionRequest
	1,  // 6: session.session.JoinSession:input_type -> session.JoinSessionRequest
	6,  // 7: session.session.ListSessions:input_type -> session.ListSessionsRequest
	7,  // 8: session.level.UploadLevel:input_type -> session.CreateLevelRequest
	9,  // 9: session.level.ListLevels:input_type -> session.ListLevelsRequest
	5,  // 10: session.session.CreateSession:output_type -> session.Session
	5,  // 11: session.session.JoinSession:output_type -> session.Session
	5,  // 12: session.session.ListSessions:output_type -> session.Session
	8,  // 13: session.level.UploadLevel:output_type -> session.Level
	8,  // 14: session.level.ListLevels:output_type -> session.Level
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_session_proto_init() }
func file_session_proto_init() {
	if File_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinEvent); i {
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
		file_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingEvent); i {
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
		file_session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelState); i {
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
		file_session_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_session_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_session_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLevelRequest); i {
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
		file_session_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Level); i {
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
		file_session_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLevelsRequest); i {
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
	file_session_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*JoinSessionRequest_Join)(nil),
		(*JoinSessionRequest_Ping)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_session_proto_goTypes,
		DependencyIndexes: file_session_proto_depIdxs,
		MessageInfos:      file_session_proto_msgTypes,
	}.Build()
	File_session_proto = out.File
	file_session_proto_rawDesc = nil
	file_session_proto_goTypes = nil
	file_session_proto_depIdxs = nil
}

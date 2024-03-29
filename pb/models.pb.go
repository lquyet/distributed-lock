// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.11.4
// source: models.proto

package proto

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

type AcquireLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockId int32 `protobuf:"varint,1,opt,name=lockId,proto3" json:"lockId,omitempty"`
}

func (x *AcquireLockRequest) Reset() {
	*x = AcquireLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcquireLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcquireLockRequest) ProtoMessage() {}

func (x *AcquireLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcquireLockRequest.ProtoReflect.Descriptor instead.
func (*AcquireLockRequest) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{0}
}

func (x *AcquireLockRequest) GetLockId() int32 {
	if x != nil {
		return x.LockId
	}
	return 0
}

type AcquireLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockId int32 `protobuf:"varint,1,opt,name=lockId,proto3" json:"lockId,omitempty"`
}

func (x *AcquireLockResponse) Reset() {
	*x = AcquireLockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcquireLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcquireLockResponse) ProtoMessage() {}

func (x *AcquireLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcquireLockResponse.ProtoReflect.Descriptor instead.
func (*AcquireLockResponse) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{1}
}

func (x *AcquireLockResponse) GetLockId() int32 {
	if x != nil {
		return x.LockId
	}
	return 0
}

type ReleaseLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockId int32 `protobuf:"varint,1,opt,name=lockId,proto3" json:"lockId,omitempty"`
}

func (x *ReleaseLockRequest) Reset() {
	*x = ReleaseLockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseLockRequest) ProtoMessage() {}

func (x *ReleaseLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseLockRequest.ProtoReflect.Descriptor instead.
func (*ReleaseLockRequest) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{2}
}

func (x *ReleaseLockRequest) GetLockId() int32 {
	if x != nil {
		return x.LockId
	}
	return 0
}

type ReleaseLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockId int32 `protobuf:"varint,1,opt,name=lockId,proto3" json:"lockId,omitempty"`
}

func (x *ReleaseLockResponse) Reset() {
	*x = ReleaseLockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseLockResponse) ProtoMessage() {}

func (x *ReleaseLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseLockResponse.ProtoReflect.Descriptor instead.
func (*ReleaseLockResponse) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{3}
}

func (x *ReleaseLockResponse) GetLockId() int32 {
	if x != nil {
		return x.LockId
	}
	return 0
}

type AppendEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         uint32   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	LeaderId     uint32   `protobuf:"varint,2,opt,name=leaderId,proto3" json:"leaderId,omitempty"`
	PrevLogIndex uint32   `protobuf:"varint,3,opt,name=prevLogIndex,proto3" json:"prevLogIndex,omitempty"`
	PrevLogTerm  uint32   `protobuf:"varint,4,opt,name=prevLogTerm,proto3" json:"prevLogTerm,omitempty"`
	Entries      []string `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	LeaderCommit uint32   `protobuf:"varint,6,opt,name=leaderCommit,proto3" json:"leaderCommit,omitempty"`
}

func (x *AppendEntriesRequest) Reset() {
	*x = AppendEntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesRequest) ProtoMessage() {}

func (x *AppendEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesRequest.ProtoReflect.Descriptor instead.
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{4}
}

func (x *AppendEntriesRequest) GetTerm() uint32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesRequest) GetLeaderId() uint32 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *AppendEntriesRequest) GetPrevLogIndex() uint32 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntriesRequest) GetPrevLogTerm() uint32 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntriesRequest) GetEntries() []string {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *AppendEntriesRequest) GetLeaderCommit() uint32 {
	if x != nil {
		return x.LeaderCommit
	}
	return 0
}

type AppendEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    uint32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AppendEntriesResponse) Reset() {
	*x = AppendEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesResponse) ProtoMessage() {}

func (x *AppendEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesResponse.ProtoReflect.Descriptor instead.
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{5}
}

func (x *AppendEntriesResponse) GetTerm() uint32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RequestVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         uint32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	CandidateId  uint32 `protobuf:"varint,2,opt,name=candidateId,proto3" json:"candidateId,omitempty"`
	LastLogIndex uint32 `protobuf:"varint,3,opt,name=lastLogIndex,proto3" json:"lastLogIndex,omitempty"`
	LastLogTerm  uint32 `protobuf:"varint,4,opt,name=lastLogTerm,proto3" json:"lastLogTerm,omitempty"`
}

func (x *RequestVoteRequest) Reset() {
	*x = RequestVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteRequest) ProtoMessage() {}

func (x *RequestVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteRequest.ProtoReflect.Descriptor instead.
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{6}
}

func (x *RequestVoteRequest) GetTerm() uint32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteRequest) GetCandidateId() uint32 {
	if x != nil {
		return x.CandidateId
	}
	return 0
}

func (x *RequestVoteRequest) GetLastLogIndex() uint32 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *RequestVoteRequest) GetLastLogTerm() uint32 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

type RequestVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        uint32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=voteGranted,proto3" json:"voteGranted,omitempty"`
}

func (x *RequestVoteResponse) Reset() {
	*x = RequestVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteResponse) ProtoMessage() {}

func (x *RequestVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteResponse.ProtoReflect.Descriptor instead.
func (*RequestVoteResponse) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{7}
}

func (x *RequestVoteResponse) GetTerm() uint32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteResponse) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

type JoinClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *JoinClusterRequest) Reset() {
	*x = JoinClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinClusterRequest) ProtoMessage() {}

func (x *JoinClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinClusterRequest.ProtoReflect.Descriptor instead.
func (*JoinClusterRequest) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{8}
}

func (x *JoinClusterRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JoinClusterRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type JoinClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ids   []uint32 `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Addrs []string `protobuf:"bytes,3,rep,name=addrs,proto3" json:"addrs,omitempty"`
}

func (x *JoinClusterResponse) Reset() {
	*x = JoinClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinClusterResponse) ProtoMessage() {}

func (x *JoinClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinClusterResponse.ProtoReflect.Descriptor instead.
func (*JoinClusterResponse) Descriptor() ([]byte, []int) {
	return file_models_proto_rawDescGZIP(), []int{9}
}

func (x *JoinClusterResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JoinClusterResponse) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *JoinClusterResponse) GetAddrs() []string {
	if x != nil {
		return x.Addrs
	}
	return nil
}

var File_models_proto protoreflect.FileDescriptor

var file_models_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x12, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x13, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x22, 0x2d, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x22,
	0xca, 0x01, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76,
	0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x45, 0x0a, 0x15,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54,
	0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x4b, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x76, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x64, 0x22, 0x38, 0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x4d, 0x0a,
	0x13, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x42, 0x2a, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x71, 0x75, 0x79, 0x65,
	0x74, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2d, 0x6c, 0x6f,
	0x63, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_proto_rawDescOnce sync.Once
	file_models_proto_rawDescData = file_models_proto_rawDesc
)

func file_models_proto_rawDescGZIP() []byte {
	file_models_proto_rawDescOnce.Do(func() {
		file_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_proto_rawDescData)
	})
	return file_models_proto_rawDescData
}

var file_models_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_models_proto_goTypes = []interface{}{
	(*AcquireLockRequest)(nil),    // 0: proto.AcquireLockRequest
	(*AcquireLockResponse)(nil),   // 1: proto.AcquireLockResponse
	(*ReleaseLockRequest)(nil),    // 2: proto.ReleaseLockRequest
	(*ReleaseLockResponse)(nil),   // 3: proto.ReleaseLockResponse
	(*AppendEntriesRequest)(nil),  // 4: proto.AppendEntriesRequest
	(*AppendEntriesResponse)(nil), // 5: proto.AppendEntriesResponse
	(*RequestVoteRequest)(nil),    // 6: proto.RequestVoteRequest
	(*RequestVoteResponse)(nil),   // 7: proto.RequestVoteResponse
	(*JoinClusterRequest)(nil),    // 8: proto.JoinClusterRequest
	(*JoinClusterResponse)(nil),   // 9: proto.JoinClusterResponse
}
var file_models_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_proto_init() }
func file_models_proto_init() {
	if File_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcquireLockRequest); i {
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
		file_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcquireLockResponse); i {
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
		file_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseLockRequest); i {
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
		file_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseLockResponse); i {
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
		file_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesRequest); i {
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
		file_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesResponse); i {
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
		file_models_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteRequest); i {
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
		file_models_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteResponse); i {
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
		file_models_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinClusterRequest); i {
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
		file_models_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinClusterResponse); i {
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
			RawDescriptor: file_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_proto_goTypes,
		DependencyIndexes: file_models_proto_depIdxs,
		MessageInfos:      file_models_proto_msgTypes,
	}.Build()
	File_models_proto = out.File
	file_models_proto_rawDesc = nil
	file_models_proto_goTypes = nil
	file_models_proto_depIdxs = nil
}

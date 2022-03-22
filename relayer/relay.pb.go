// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: relayer/proto/relay.proto

package relayer

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

type RelayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpecId      uint32 `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	ApiId       uint32 `protobuf:"varint,2,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
	SessionId   uint64 `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	CuSum       uint64 `protobuf:"varint,4,opt,name=cu_sum,json=cuSum,proto3" json:"cu_sum,omitempty"` // total compute unit used including this relay
	Data        []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Sig         []byte `protobuf:"bytes,6,opt,name=sig,proto3" json:"sig,omitempty"`
	Servicer    []byte `protobuf:"bytes,7,opt,name=servicer,proto3" json:"servicer,omitempty"`
	BlockHeight int64  `protobuf:"varint,8,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (x *RelayRequest) Reset() {
	*x = RelayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relayer_proto_relay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayRequest) ProtoMessage() {}

func (x *RelayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relayer_proto_relay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayRequest.ProtoReflect.Descriptor instead.
func (*RelayRequest) Descriptor() ([]byte, []int) {
	return file_relayer_proto_relay_proto_rawDescGZIP(), []int{0}
}

func (x *RelayRequest) GetSpecId() uint32 {
	if x != nil {
		return x.SpecId
	}
	return 0
}

func (x *RelayRequest) GetApiId() uint32 {
	if x != nil {
		return x.ApiId
	}
	return 0
}

func (x *RelayRequest) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *RelayRequest) GetCuSum() uint64 {
	if x != nil {
		return x.CuSum
	}
	return 0
}

func (x *RelayRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RelayRequest) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *RelayRequest) GetServicer() []byte {
	if x != nil {
		return x.Servicer
	}
	return nil
}

func (x *RelayRequest) GetBlockHeight() int64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

type RelayReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Sig  []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *RelayReply) Reset() {
	*x = RelayReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relayer_proto_relay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayReply) ProtoMessage() {}

func (x *RelayReply) ProtoReflect() protoreflect.Message {
	mi := &file_relayer_proto_relay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayReply.ProtoReflect.Descriptor instead.
func (*RelayReply) Descriptor() ([]byte, []int) {
	return file_relayer_proto_relay_proto_rawDescGZIP(), []int{1}
}

func (x *RelayReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RelayReply) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

var File_relayer_proto_relay_proto protoreflect.FileDescriptor

var file_relayer_proto_relay_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x22, 0xd9, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x70, 0x65, 0x63, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x61, 0x70, 0x69, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x75, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x75, 0x53, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69,
	0x67, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x32, 0x0a, 0x0a, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x73, 0x69, 0x67, 0x32, 0x40, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x35, 0x0a, 0x05, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relayer_proto_relay_proto_rawDescOnce sync.Once
	file_relayer_proto_relay_proto_rawDescData = file_relayer_proto_relay_proto_rawDesc
)

func file_relayer_proto_relay_proto_rawDescGZIP() []byte {
	file_relayer_proto_relay_proto_rawDescOnce.Do(func() {
		file_relayer_proto_relay_proto_rawDescData = protoimpl.X.CompressGZIP(file_relayer_proto_relay_proto_rawDescData)
	})
	return file_relayer_proto_relay_proto_rawDescData
}

var file_relayer_proto_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_relayer_proto_relay_proto_goTypes = []interface{}{
	(*RelayRequest)(nil), // 0: relayer.RelayRequest
	(*RelayReply)(nil),   // 1: relayer.RelayReply
}
var file_relayer_proto_relay_proto_depIdxs = []int32{
	0, // 0: relayer.Relayer.Relay:input_type -> relayer.RelayRequest
	1, // 1: relayer.Relayer.Relay:output_type -> relayer.RelayReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relayer_proto_relay_proto_init() }
func file_relayer_proto_relay_proto_init() {
	if File_relayer_proto_relay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relayer_proto_relay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelayRequest); i {
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
		file_relayer_proto_relay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelayReply); i {
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
			RawDescriptor: file_relayer_proto_relay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relayer_proto_relay_proto_goTypes,
		DependencyIndexes: file_relayer_proto_relay_proto_depIdxs,
		MessageInfos:      file_relayer_proto_relay_proto_msgTypes,
	}.Build()
	File_relayer_proto_relay_proto = out.File
	file_relayer_proto_relay_proto_rawDesc = nil
	file_relayer_proto_relay_proto_goTypes = nil
	file_relayer_proto_relay_proto_depIdxs = nil
}
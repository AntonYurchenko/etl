//
//Compilation:
//protoc --go_out=. \
//--go_opt=paths=source_relative \
//--go-grpc_out=. \
//--go-grpc_opt=paths=source_relative \
//contract/contract.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: contract/contract.proto

package contract

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

// This is a container with data batch and a target entity.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of target entity for saving a data batch.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Batch  *Batch `protobuf:"bytes,2,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Message) GetBatch() *Batch {
	if x != nil {
		return x.Batch
	}
	return nil
}

// Any batch of data.
// Examples
//     - A column with name = names[0] has a type = types[0] and value = values[0].
//     - A column with name = names[1] has a type = types[1] and value = values[len(names)].
type Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of all column names in batch.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// List of all column types in batch.
	// Available types:
	//   - int[8, 16, 32, 64]
	//   - uint[8, 16, 32, 64]
	//   - float[34, 64]
	//   - date (in format 2006-01-02)
	//   - datetime (in format 2006-01-02T15:04:05.999Z07:00)
	//   - string
	Types []string `protobuf:"bytes,2,rep,name=types,proto3" json:"types,omitempty"`
	// List of all column values in batch.
	Values [][]byte `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Batch) Reset() {
	*x = Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Batch) ProtoMessage() {}

func (x *Batch) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Batch.ProtoReflect.Descriptor instead.
func (*Batch) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{1}
}

func (x *Batch) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *Batch) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *Batch) GetValues() [][]byte {
	if x != nil {
		return x.Values
	}
	return nil
}

// Status is a response from DataConsumer.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text of error.
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// count of rows which have been success written.
	Success int32 `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_contract_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_contract_contract_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Status) GetSuccess() int32 {
	if x != nil {
		return x.Success
	}
	return 0
}

var File_contract_contract_proto protoreflect.FileDescriptor

var file_contract_contract_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x05,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x22, 0x4b, 0x0a, 0x05, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0x33, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x12, 0x23, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x08, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_contract_proto_rawDescOnce sync.Once
	file_contract_contract_proto_rawDescData = file_contract_contract_proto_rawDesc
)

func file_contract_contract_proto_rawDescGZIP() []byte {
	file_contract_contract_proto_rawDescOnce.Do(func() {
		file_contract_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_contract_proto_rawDescData)
	})
	return file_contract_contract_proto_rawDescData
}

var file_contract_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_contract_contract_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: Message
	(*Batch)(nil),   // 1: Batch
	(*Status)(nil),  // 2: Status
}
var file_contract_contract_proto_depIdxs = []int32{
	1, // 0: Message.batch:type_name -> Batch
	0, // 1: DataConsumer.Exchange:input_type -> Message
	2, // 2: DataConsumer.Exchange:output_type -> Status
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contract_contract_proto_init() }
func file_contract_contract_proto_init() {
	if File_contract_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_contract_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Batch); i {
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
		file_contract_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_contract_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_contract_proto_goTypes,
		DependencyIndexes: file_contract_contract_proto_depIdxs,
		MessageInfos:      file_contract_contract_proto_msgTypes,
	}.Build()
	File_contract_contract_proto = out.File
	file_contract_contract_proto_rawDesc = nil
	file_contract_contract_proto_goTypes = nil
	file_contract_contract_proto_depIdxs = nil
}
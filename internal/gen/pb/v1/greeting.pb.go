// edition = "2023";

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: com/acme/schema/v1/greeting.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Language int32

const (
	Language_UNKNOWN Language = 0 // Default value when language is not set
	Language_EN      Language = 1 // Generic English
	Language_EN_GB   Language = 2 // British English
	Language_ES      Language = 3 // Spanish
	Language_FR      Language = 4 // French
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "UNKNOWN",
		1: "EN",
		2: "EN_GB",
		3: "ES",
		4: "FR",
	}
	Language_value = map[string]int32{
		"UNKNOWN": 0,
		"EN":      1,
		"EN_GB":   2,
		"ES":      3,
		"FR":      4,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_com_acme_schema_v1_greeting_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_com_acme_schema_v1_greeting_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_com_acme_schema_v1_greeting_proto_rawDescGZIP(), []int{0}
}

type GreetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Preferred language for the greeting (e.g., EN, ES, FR).
	// If not set, the server may return a default greeting.
	Language Language `protobuf:"varint,1,opt,name=language,proto3,enum=com.acme.schema.v1.Language" json:"language,omitempty"`
}

func (x *GreetingRequest) Reset() {
	*x = GreetingRequest{}
	mi := &file_com_acme_schema_v1_greeting_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GreetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingRequest) ProtoMessage() {}

func (x *GreetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_acme_schema_v1_greeting_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingRequest.ProtoReflect.Descriptor instead.
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return file_com_acme_schema_v1_greeting_proto_rawDescGZIP(), []int{0}
}

func (x *GreetingRequest) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_UNKNOWN
}

type GreetingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Actual language of the greeting.
	Language  Language               `protobuf:"varint,1,opt,name=language,proto3,enum=com.acme.schema.v1.Language" json:"language,omitempty"`
	Greeting  string                 `protobuf:"bytes,2,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Metadata  map[string]string      `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GreetingResponse) Reset() {
	*x = GreetingResponse{}
	mi := &file_com_acme_schema_v1_greeting_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GreetingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingResponse) ProtoMessage() {}

func (x *GreetingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_acme_schema_v1_greeting_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingResponse.ProtoReflect.Descriptor instead.
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return file_com_acme_schema_v1_greeting_proto_rawDescGZIP(), []int{1}
}

func (x *GreetingResponse) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_UNKNOWN
}

func (x *GreetingResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

func (x *GreetingResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *GreetingResponse) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_com_acme_schema_v1_greeting_proto protoreflect.FileDescriptor

var file_com_acme_schema_v1_greeting_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x0f, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x22, 0xaf, 0x02, 0x0a, 0x10, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x2a, 0x3a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a,
	0x02, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e, 0x5f, 0x47, 0x42, 0x10, 0x02,
	0x12, 0x06, 0x0a, 0x02, 0x45, 0x53, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x46, 0x52, 0x10, 0x04,
	0x32, 0x60, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x08, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63,
	0x6d, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x69, 0x74, 0x6f, 0x70, 0x73, 0x2d, 0x63, 0x69, 0x2d, 0x63, 0x64, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_acme_schema_v1_greeting_proto_rawDescOnce sync.Once
	file_com_acme_schema_v1_greeting_proto_rawDescData = file_com_acme_schema_v1_greeting_proto_rawDesc
)

func file_com_acme_schema_v1_greeting_proto_rawDescGZIP() []byte {
	file_com_acme_schema_v1_greeting_proto_rawDescOnce.Do(func() {
		file_com_acme_schema_v1_greeting_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_acme_schema_v1_greeting_proto_rawDescData)
	})
	return file_com_acme_schema_v1_greeting_proto_rawDescData
}

var file_com_acme_schema_v1_greeting_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_acme_schema_v1_greeting_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_acme_schema_v1_greeting_proto_goTypes = []any{
	(Language)(0),                 // 0: com.acme.schema.v1.Language
	(*GreetingRequest)(nil),       // 1: com.acme.schema.v1.GreetingRequest
	(*GreetingResponse)(nil),      // 2: com.acme.schema.v1.GreetingResponse
	nil,                           // 3: com.acme.schema.v1.GreetingResponse.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_com_acme_schema_v1_greeting_proto_depIdxs = []int32{
	0, // 0: com.acme.schema.v1.GreetingRequest.language:type_name -> com.acme.schema.v1.Language
	0, // 1: com.acme.schema.v1.GreetingResponse.language:type_name -> com.acme.schema.v1.Language
	4, // 2: com.acme.schema.v1.GreetingResponse.timestamp:type_name -> google.protobuf.Timestamp
	3, // 3: com.acme.schema.v1.GreetingResponse.metadata:type_name -> com.acme.schema.v1.GreetingResponse.MetadataEntry
	1, // 4: com.acme.schema.v1.Greeter.Greeting:input_type -> com.acme.schema.v1.GreetingRequest
	2, // 5: com.acme.schema.v1.Greeter.Greeting:output_type -> com.acme.schema.v1.GreetingResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_com_acme_schema_v1_greeting_proto_init() }
func file_com_acme_schema_v1_greeting_proto_init() {
	if File_com_acme_schema_v1_greeting_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_acme_schema_v1_greeting_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_acme_schema_v1_greeting_proto_goTypes,
		DependencyIndexes: file_com_acme_schema_v1_greeting_proto_depIdxs,
		EnumInfos:         file_com_acme_schema_v1_greeting_proto_enumTypes,
		MessageInfos:      file_com_acme_schema_v1_greeting_proto_msgTypes,
	}.Build()
	File_com_acme_schema_v1_greeting_proto = out.File
	file_com_acme_schema_v1_greeting_proto_rawDesc = nil
	file_com_acme_schema_v1_greeting_proto_goTypes = nil
	file_com_acme_schema_v1_greeting_proto_depIdxs = nil
}

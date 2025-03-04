// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc1
// source: slide/v1/errors.proto

package slideschema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Code int32

const (
	Code_err_unspecified                   Code = 0
	Code_err_endpoint_not_found            Code = 1
	Code_err_entity_not_found              Code = 2
	Code_err_validation_error              Code = 3
	Code_err_missing_authentication        Code = 4
	Code_err_unauthorized                  Code = 5
	Code_err_internal_server_error         Code = 6
	Code_err_rate_limit_exceeded           Code = 7
	Code_err_agent_not_connected_to_device Code = 8
	Code_err_device_not_connected_to_cloud Code = 9
	Code_err_backup_already_running        Code = 10
	Code_err_client_not_found              Code = 11
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:  "err_unspecified",
		1:  "err_endpoint_not_found",
		2:  "err_entity_not_found",
		3:  "err_validation_error",
		4:  "err_missing_authentication",
		5:  "err_unauthorized",
		6:  "err_internal_server_error",
		7:  "err_rate_limit_exceeded",
		8:  "err_agent_not_connected_to_device",
		9:  "err_device_not_connected_to_cloud",
		10: "err_backup_already_running",
		11: "err_client_not_found",
	}
	Code_value = map[string]int32{
		"err_unspecified":                   0,
		"err_endpoint_not_found":            1,
		"err_entity_not_found":              2,
		"err_validation_error":              3,
		"err_missing_authentication":        4,
		"err_unauthorized":                  5,
		"err_internal_server_error":         6,
		"err_rate_limit_exceeded":           7,
		"err_agent_not_connected_to_device": 8,
		"err_device_not_connected_to_cloud": 9,
		"err_backup_already_running":        10,
		"err_client_not_found":              11,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_slide_v1_errors_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_slide_v1_errors_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_slide_v1_errors_proto_rawDescGZIP(), []int{0}
}

type ErrorResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Codes         []Code                 `protobuf:"varint,1,rep,packed,name=codes,enum=slide.v1.Code" json:"codes,omitempty"`
	Details       []string               `protobuf:"bytes,2,rep,name=details" json:"details,omitempty"`
	Message       *string                `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	mi := &file_slide_v1_errors_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slide_v1_errors_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_slide_v1_errors_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorResponse) GetCodes() []Code {
	if x != nil {
		return x.Codes
	}
	return nil
}

func (x *ErrorResponse) GetDetails() []string {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_slide_v1_errors_proto protoreflect.FileDescriptor

var file_slide_v1_errors_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0x69, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0xe5, 0x02, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x65, 0x72, 0x72, 0x5f, 0x75, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x65, 0x72,
	0x72, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x65, 0x72, 0x72, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x65, 0x72, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x65, 0x72,
	0x72, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x65, 0x72,
	0x72, 0x5f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x10, 0x05,
	0x12, 0x1d, 0x0a, 0x19, 0x65, 0x72, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x06, 0x12,
	0x1b, 0x0a, 0x17, 0x65, 0x72, 0x72, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x07, 0x12, 0x25, 0x0a, 0x21,
	0x65, 0x72, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x10, 0x08, 0x12, 0x25, 0x0a, 0x21, 0x65, 0x72, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x6f, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x65, 0x72,
	0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14, 0x65, 0x72,
	0x72, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x10, 0x0b, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x67, 0x69, 0x62, 0x73, 0x6f, 0x6e, 0x2f,
	0x73, 0x6c, 0x69, 0x64, 0x65, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x08,
	0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var (
	file_slide_v1_errors_proto_rawDescOnce sync.Once
	file_slide_v1_errors_proto_rawDescData []byte
)

func file_slide_v1_errors_proto_rawDescGZIP() []byte {
	file_slide_v1_errors_proto_rawDescOnce.Do(func() {
		file_slide_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_slide_v1_errors_proto_rawDesc), len(file_slide_v1_errors_proto_rawDesc)))
	})
	return file_slide_v1_errors_proto_rawDescData
}

var file_slide_v1_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_slide_v1_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_slide_v1_errors_proto_goTypes = []any{
	(Code)(0),             // 0: slide.v1.Code
	(*ErrorResponse)(nil), // 1: slide.v1.ErrorResponse
}
var file_slide_v1_errors_proto_depIdxs = []int32{
	0, // 0: slide.v1.ErrorResponse.codes:type_name -> slide.v1.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_slide_v1_errors_proto_init() }
func file_slide_v1_errors_proto_init() {
	if File_slide_v1_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_slide_v1_errors_proto_rawDesc), len(file_slide_v1_errors_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_slide_v1_errors_proto_goTypes,
		DependencyIndexes: file_slide_v1_errors_proto_depIdxs,
		EnumInfos:         file_slide_v1_errors_proto_enumTypes,
		MessageInfos:      file_slide_v1_errors_proto_msgTypes,
	}.Build()
	File_slide_v1_errors_proto = out.File
	file_slide_v1_errors_proto_goTypes = nil
	file_slide_v1_errors_proto_depIdxs = nil
}

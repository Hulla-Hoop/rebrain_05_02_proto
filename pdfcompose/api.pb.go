// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.13.0
// source: server/api/api.proto

package pdfcompose

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

type Upfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	One   []byte `protobuf:"bytes,1,opt,name=One,proto3" json:"One,omitempty"`
	Two   []byte `protobuf:"bytes,2,opt,name=Two,proto3" json:"Two,omitempty"`
	Three []byte `protobuf:"bytes,3,opt,name=Three,proto3" json:"Three,omitempty"`
}

func (x *Upfile) Reset() {
	*x = Upfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Upfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Upfile) ProtoMessage() {}

func (x *Upfile) ProtoReflect() protoreflect.Message {
	mi := &file_server_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Upfile.ProtoReflect.Descriptor instead.
func (*Upfile) Descriptor() ([]byte, []int) {
	return file_server_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Upfile) GetOne() []byte {
	if x != nil {
		return x.One
	}
	return nil
}

func (x *Upfile) GetTwo() []byte {
	if x != nil {
		return x.Two
	}
	return nil
}

func (x *Upfile) GetThree() []byte {
	if x != nil {
		return x.Three
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pdf []byte `protobuf:"bytes,1,opt,name=Pdf,proto3" json:"Pdf,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_server_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_server_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetPdf() []byte {
	if x != nil {
		return x.Pdf
	}
	return nil
}

var File_server_api_api_proto protoreflect.FileDescriptor

var file_server_api_api_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x65, 0x22, 0x42, 0x0a, 0x06, 0x55, 0x70, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x4f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x4f, 0x6e, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x54, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x54, 0x77, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x68, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x54, 0x68, 0x72, 0x65, 0x65, 0x22, 0x18, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x50, 0x64, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x50, 0x64, 0x66,
	0x32, 0x43, 0x0a, 0x11, 0x50, 0x64, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x2e,
	0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x55, 0x70, 0x66, 0x69, 0x6c,
	0x65, 0x1a, 0x10, 0x2e, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_api_api_proto_rawDescOnce sync.Once
	file_server_api_api_proto_rawDescData = file_server_api_api_proto_rawDesc
)

func file_server_api_api_proto_rawDescGZIP() []byte {
	file_server_api_api_proto_rawDescOnce.Do(func() {
		file_server_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_api_api_proto_rawDescData)
	})
	return file_server_api_api_proto_rawDescData
}

var file_server_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_api_api_proto_goTypes = []interface{}{
	(*Upfile)(nil), // 0: pdfcompose.Upfile
	(*File)(nil),   // 1: pdfcompose.File
}
var file_server_api_api_proto_depIdxs = []int32{
	0, // 0: pdfcompose.PdfComposeService.Send:input_type -> pdfcompose.Upfile
	1, // 1: pdfcompose.PdfComposeService.Send:output_type -> pdfcompose.File
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_api_api_proto_init() }
func file_server_api_api_proto_init() {
	if File_server_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Upfile); i {
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
		file_server_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_server_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_api_api_proto_goTypes,
		DependencyIndexes: file_server_api_api_proto_depIdxs,
		MessageInfos:      file_server_api_api_proto_msgTypes,
	}.Build()
	File_server_api_api_proto = out.File
	file_server_api_api_proto_rawDesc = nil
	file_server_api_api_proto_goTypes = nil
	file_server_api_api_proto_depIdxs = nil
}

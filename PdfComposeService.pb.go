// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: PdfComposeService.proto

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

type Images struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data1:
	//
	//	*Images_Upfile1
	Data1 isImages_Data1 `protobuf_oneof:"data1"`
	// Types that are assignable to Data2:
	//
	//	*Images_Upfile2
	Data2 isImages_Data2 `protobuf_oneof:"data2"`
	// Types that are assignable to Data3:
	//
	//	*Images_Upfile3
	Data3 isImages_Data3 `protobuf_oneof:"data3"`
}

func (x *Images) Reset() {
	*x = Images{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PdfComposeService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Images) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Images) ProtoMessage() {}

func (x *Images) ProtoReflect() protoreflect.Message {
	mi := &file_PdfComposeService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Images.ProtoReflect.Descriptor instead.
func (*Images) Descriptor() ([]byte, []int) {
	return file_PdfComposeService_proto_rawDescGZIP(), []int{0}
}

func (m *Images) GetData1() isImages_Data1 {
	if m != nil {
		return m.Data1
	}
	return nil
}

func (x *Images) GetUpfile1() []byte {
	if x, ok := x.GetData1().(*Images_Upfile1); ok {
		return x.Upfile1
	}
	return nil
}

func (m *Images) GetData2() isImages_Data2 {
	if m != nil {
		return m.Data2
	}
	return nil
}

func (x *Images) GetUpfile2() []byte {
	if x, ok := x.GetData2().(*Images_Upfile2); ok {
		return x.Upfile2
	}
	return nil
}

func (m *Images) GetData3() isImages_Data3 {
	if m != nil {
		return m.Data3
	}
	return nil
}

func (x *Images) GetUpfile3() []byte {
	if x, ok := x.GetData3().(*Images_Upfile3); ok {
		return x.Upfile3
	}
	return nil
}

type isImages_Data1 interface {
	isImages_Data1()
}

type Images_Upfile1 struct {
	Upfile1 []byte `protobuf:"bytes,1,opt,name=upfile1,proto3,oneof"`
}

func (*Images_Upfile1) isImages_Data1() {}

type isImages_Data2 interface {
	isImages_Data2()
}

type Images_Upfile2 struct {
	Upfile2 []byte `protobuf:"bytes,2,opt,name=upfile2,proto3,oneof"`
}

func (*Images_Upfile2) isImages_Data2() {}

type isImages_Data3 interface {
	isImages_Data3()
}

type Images_Upfile3 struct {
	Upfile3 []byte `protobuf:"bytes,3,opt,name=upfile3,proto3,oneof"`
}

func (*Images_Upfile3) isImages_Data3() {}

type Pdf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*Pdf_Result
	Data isPdf_Data `protobuf_oneof:"data"`
}

func (x *Pdf) Reset() {
	*x = Pdf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PdfComposeService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pdf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pdf) ProtoMessage() {}

func (x *Pdf) ProtoReflect() protoreflect.Message {
	mi := &file_PdfComposeService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pdf.ProtoReflect.Descriptor instead.
func (*Pdf) Descriptor() ([]byte, []int) {
	return file_PdfComposeService_proto_rawDescGZIP(), []int{1}
}

func (m *Pdf) GetData() isPdf_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Pdf) GetResult() []byte {
	if x, ok := x.GetData().(*Pdf_Result); ok {
		return x.Result
	}
	return nil
}

type isPdf_Data interface {
	isPdf_Data()
}

type Pdf_Result struct {
	Result []byte `protobuf:"bytes,1,opt,name=result,proto3,oneof"`
}

func (*Pdf_Result) isPdf_Data() {}

var File_PdfComposeService_proto protoreflect.FileDescriptor

var file_PdfComposeService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x50, 0x64, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x64, 0x66, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x22, 0x77, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x07, 0x75, 0x70, 0x66, 0x69, 0x6c, 0x65, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x07, 0x75, 0x70, 0x66, 0x69, 0x6c, 0x65, 0x31, 0x12, 0x1a, 0x0a, 0x07, 0x75,
	0x70, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x07,
	0x75, 0x70, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x12, 0x1a, 0x0a, 0x07, 0x75, 0x70, 0x66, 0x69, 0x6c,
	0x65, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x02, 0x52, 0x07, 0x75, 0x70, 0x66, 0x69,
	0x6c, 0x65, 0x33, 0x42, 0x07, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x31, 0x42, 0x07, 0x0a, 0x05,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x42, 0x07, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x33, 0x22, 0x27,
	0x0a, 0x03, 0x50, 0x64, 0x66, 0x12, 0x18, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42,
	0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x39, 0x0a, 0x0a, 0x50, 0x64, 0x66, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x2e,
	0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x1a, 0x0f, 0x2e, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x50,
	0x64, 0x66, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x70, 0x64, 0x66, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PdfComposeService_proto_rawDescOnce sync.Once
	file_PdfComposeService_proto_rawDescData = file_PdfComposeService_proto_rawDesc
)

func file_PdfComposeService_proto_rawDescGZIP() []byte {
	file_PdfComposeService_proto_rawDescOnce.Do(func() {
		file_PdfComposeService_proto_rawDescData = protoimpl.X.CompressGZIP(file_PdfComposeService_proto_rawDescData)
	})
	return file_PdfComposeService_proto_rawDescData
}

var file_PdfComposeService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_PdfComposeService_proto_goTypes = []interface{}{
	(*Images)(nil), // 0: pdfcompose.Images
	(*Pdf)(nil),    // 1: pdfcompose.Pdf
}
var file_PdfComposeService_proto_depIdxs = []int32{
	0, // 0: pdfcompose.PdfCompose.Send:input_type -> pdfcompose.Images
	1, // 1: pdfcompose.PdfCompose.Send:output_type -> pdfcompose.Pdf
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PdfComposeService_proto_init() }
func file_PdfComposeService_proto_init() {
	if File_PdfComposeService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PdfComposeService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Images); i {
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
		file_PdfComposeService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pdf); i {
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
	file_PdfComposeService_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Images_Upfile1)(nil),
		(*Images_Upfile2)(nil),
		(*Images_Upfile3)(nil),
	}
	file_PdfComposeService_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Pdf_Result)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PdfComposeService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_PdfComposeService_proto_goTypes,
		DependencyIndexes: file_PdfComposeService_proto_depIdxs,
		MessageInfos:      file_PdfComposeService_proto_msgTypes,
	}.Build()
	File_PdfComposeService_proto = out.File
	file_PdfComposeService_proto_rawDesc = nil
	file_PdfComposeService_proto_goTypes = nil
	file_PdfComposeService_proto_depIdxs = nil
}

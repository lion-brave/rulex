// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: plugin.proto

package xproto

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type XPluginMetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version  string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Homepage string `protobuf:"bytes,3,opt,name=homepage,proto3" json:"homepage,omitempty"`
	HelpLink string `protobuf:"bytes,4,opt,name=helpLink,proto3" json:"helpLink,omitempty"`
	Author   string `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	Email    string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	License  string `protobuf:"bytes,7,opt,name=license,proto3" json:"license,omitempty"`
}

func (x *XPluginMetaInfo) Reset() {
	*x = XPluginMetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XPluginMetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XPluginMetaInfo) ProtoMessage() {}

func (x *XPluginMetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XPluginMetaInfo.ProtoReflect.Descriptor instead.
func (*XPluginMetaInfo) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *XPluginMetaInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *XPluginMetaInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *XPluginMetaInfo) GetHomepage() string {
	if x != nil {
		return x.Homepage
	}
	return ""
}

func (x *XPluginMetaInfo) GetHelpLink() string {
	if x != nil {
		return x.HelpLink
	}
	return ""
}

func (x *XPluginMetaInfo) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *XPluginMetaInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *XPluginMetaInfo) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

var File_plugin_proto protoreflect.FileDescriptor

var file_plugin_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x78, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0xbf, 0x01, 0x0a, 0x0f, 0x58, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x6c, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x6c, 0x70, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x32, 0xd7, 0x01, 0x0a, 0x07, 0x58, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x10, 0x2e, 0x78, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x78,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x78, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x78,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2d, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x10, 0x2e, 0x78, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x78, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x0e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x10, 0x2e, 0x78, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x78, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x58,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x42, 0x26, 0x0a, 0x07, 0x78, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x42, 0x0e, 0x58, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x00, 0x5a, 0x09, 0x2e,
	0x2f, 0x3b, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_proto_rawDescOnce sync.Once
	file_plugin_proto_rawDescData = file_plugin_proto_rawDesc
)

func file_plugin_proto_rawDescGZIP() []byte {
	file_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_proto_rawDescData)
	})
	return file_plugin_proto_rawDescData
}

var file_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_plugin_proto_goTypes = []interface{}{
	(*Request)(nil),         // 0: xplugin.Request
	(*Response)(nil),        // 1: xplugin.Response
	(*XPluginMetaInfo)(nil), // 2: xplugin.XPluginMetaInfo
}
var file_plugin_proto_depIdxs = []int32{
	0, // 0: xplugin.XPlugin.Init:input_type -> xplugin.Request
	0, // 1: xplugin.XPlugin.Start:input_type -> xplugin.Request
	0, // 2: xplugin.XPlugin.Stop:input_type -> xplugin.Request
	0, // 3: xplugin.XPlugin.PluginMetaInfo:input_type -> xplugin.Request
	1, // 4: xplugin.XPlugin.Init:output_type -> xplugin.Response
	1, // 5: xplugin.XPlugin.Start:output_type -> xplugin.Response
	1, // 6: xplugin.XPlugin.Stop:output_type -> xplugin.Response
	2, // 7: xplugin.XPlugin.PluginMetaInfo:output_type -> xplugin.XPluginMetaInfo
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plugin_proto_init() }
func file_plugin_proto_init() {
	if File_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XPluginMetaInfo); i {
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
			RawDescriptor: file_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_proto_depIdxs,
		MessageInfos:      file_plugin_proto_msgTypes,
	}.Build()
	File_plugin_proto = out.File
	file_plugin_proto_rawDesc = nil
	file_plugin_proto_goTypes = nil
	file_plugin_proto_depIdxs = nil
}

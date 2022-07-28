// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: proto/emoji.proto

package emoji

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

// Find an emoji by its alias e.g :beer:
type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the alias code e.g :beer:
	Alias string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_emoji_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_emoji_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.
func (*FindRequest) Descriptor() ([]byte, []int) {
	return file_proto_emoji_proto_rawDescGZIP(), []int{0}
}

func (x *FindRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type FindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the unicode emoji 🍺
	Emoji string `protobuf:"bytes,2,opt,name=emoji,proto3" json:"emoji,omitempty"`
}

func (x *FindResponse) Reset() {
	*x = FindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_emoji_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindResponse) ProtoMessage() {}

func (x *FindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_emoji_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindResponse.ProtoReflect.Descriptor instead.
func (*FindResponse) Descriptor() ([]byte, []int) {
	return file_proto_emoji_proto_rawDescGZIP(), []int{1}
}

func (x *FindResponse) GetEmoji() string {
	if x != nil {
		return x.Emoji
	}
	return ""
}

// Get the flag for a country. Requires country code e.g GB for great britain
type FlagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// country code e.g GB
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *FlagRequest) Reset() {
	*x = FlagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_emoji_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagRequest) ProtoMessage() {}

func (x *FlagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_emoji_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagRequest.ProtoReflect.Descriptor instead.
func (*FlagRequest) Descriptor() ([]byte, []int) {
	return file_proto_emoji_proto_rawDescGZIP(), []int{2}
}

func (x *FlagRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type FlagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the emoji flag
	Flag string `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *FlagResponse) Reset() {
	*x = FlagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_emoji_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagResponse) ProtoMessage() {}

func (x *FlagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_emoji_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagResponse.ProtoReflect.Descriptor instead.
func (*FlagResponse) Descriptor() ([]byte, []int) {
	return file_proto_emoji_proto_rawDescGZIP(), []int{3}
}

func (x *FlagResponse) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

// Print text and renders the emojis with aliases e.g
// let's grab a :beer: becomes let's grab a 🍺
type PrintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// text including any alias e.g let's grab a :beer:
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PrintRequest) Reset() {
	*x = PrintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_emoji_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintRequest) ProtoMessage() {}

func (x *PrintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_emoji_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintRequest.ProtoReflect.Descriptor instead.
func (*PrintRequest) Descriptor() ([]byte, []int) {
	return file_proto_emoji_proto_rawDescGZIP(), []int{4}
}

func (x *PrintRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type PrintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// text with rendered emojis
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PrintResponse) Reset() {
	*x = PrintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_emoji_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintResponse) ProtoMessage() {}

func (x *PrintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_emoji_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintResponse.ProtoReflect.Descriptor instead.
func (*PrintResponse) Descriptor() ([]byte, []int) {
	return file_proto_emoji_proto_rawDescGZIP(), []int{5}
}

func (x *PrintResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_proto_emoji_proto protoreflect.FileDescriptor

var file_proto_emoji_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x22, 0x23, 0x0a, 0x0b, 0x46, 0x69,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22,
	0x24, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x22, 0x21, 0x0a, 0x0b, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x22, 0x0a, 0x0c, 0x46, 0x6c, 0x61, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x22, 0x22, 0x0a, 0x0c,
	0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x23, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0xa3, 0x01, 0x0a, 0x05, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x12,
	0x31, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x65, 0x6d,
	0x6f, 0x6a, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x2e, 0x65, 0x6d, 0x6f,
	0x6a, 0x69, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x13,
	0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_emoji_proto_rawDescOnce sync.Once
	file_proto_emoji_proto_rawDescData = file_proto_emoji_proto_rawDesc
)

func file_proto_emoji_proto_rawDescGZIP() []byte {
	file_proto_emoji_proto_rawDescOnce.Do(func() {
		file_proto_emoji_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_emoji_proto_rawDescData)
	})
	return file_proto_emoji_proto_rawDescData
}

var file_proto_emoji_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_emoji_proto_goTypes = []interface{}{
	(*FindRequest)(nil),   // 0: emoji.FindRequest
	(*FindResponse)(nil),  // 1: emoji.FindResponse
	(*FlagRequest)(nil),   // 2: emoji.FlagRequest
	(*FlagResponse)(nil),  // 3: emoji.FlagResponse
	(*PrintRequest)(nil),  // 4: emoji.PrintRequest
	(*PrintResponse)(nil), // 5: emoji.PrintResponse
}
var file_proto_emoji_proto_depIdxs = []int32{
	0, // 0: emoji.Emoji.Find:input_type -> emoji.FindRequest
	2, // 1: emoji.Emoji.Flag:input_type -> emoji.FlagRequest
	4, // 2: emoji.Emoji.Print:input_type -> emoji.PrintRequest
	1, // 3: emoji.Emoji.Find:output_type -> emoji.FindResponse
	3, // 4: emoji.Emoji.Flag:output_type -> emoji.FlagResponse
	5, // 5: emoji.Emoji.Print:output_type -> emoji.PrintResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_emoji_proto_init() }
func file_proto_emoji_proto_init() {
	if File_proto_emoji_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_emoji_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRequest); i {
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
		file_proto_emoji_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindResponse); i {
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
		file_proto_emoji_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagRequest); i {
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
		file_proto_emoji_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagResponse); i {
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
		file_proto_emoji_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintRequest); i {
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
		file_proto_emoji_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintResponse); i {
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
			RawDescriptor: file_proto_emoji_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_emoji_proto_goTypes,
		DependencyIndexes: file_proto_emoji_proto_depIdxs,
		MessageInfos:      file_proto_emoji_proto_msgTypes,
	}.Build()
	File_proto_emoji_proto = out.File
	file_proto_emoji_proto_rawDesc = nil
	file_proto_emoji_proto_goTypes = nil
	file_proto_emoji_proto_depIdxs = nil
}
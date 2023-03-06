// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: lowcode.proto

package sgc7pb

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

// ComponentData
type ComponentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsedScenes      []int32 `protobuf:"varint,1,rep,packed,name=usedScenes,proto3" json:"usedScenes,omitempty"`
	UsedOtherScenes []int32 `protobuf:"varint,2,rep,packed,name=usedOtherScenes,proto3" json:"usedOtherScenes,omitempty"`
	UsedResults     []int32 `protobuf:"varint,3,rep,packed,name=usedResults,proto3" json:"usedResults,omitempty"`
	UsedPrizeScenes []int32 `protobuf:"varint,4,rep,packed,name=usedPrizeScenes,proto3" json:"usedPrizeScenes,omitempty"`
	CoinWin         int32   `protobuf:"varint,5,opt,name=coinWin,proto3" json:"coinWin,omitempty"`
	CashWin         int64   `protobuf:"varint,6,opt,name=cashWin,proto3" json:"cashWin,omitempty"`
	Name            string  `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ComponentData) Reset() {
	*x = ComponentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lowcode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentData) ProtoMessage() {}

func (x *ComponentData) ProtoReflect() protoreflect.Message {
	mi := &file_lowcode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentData.ProtoReflect.Descriptor instead.
func (*ComponentData) Descriptor() ([]byte, []int) {
	return file_lowcode_proto_rawDescGZIP(), []int{0}
}

func (x *ComponentData) GetUsedScenes() []int32 {
	if x != nil {
		return x.UsedScenes
	}
	return nil
}

func (x *ComponentData) GetUsedOtherScenes() []int32 {
	if x != nil {
		return x.UsedOtherScenes
	}
	return nil
}

func (x *ComponentData) GetUsedResults() []int32 {
	if x != nil {
		return x.UsedResults
	}
	return nil
}

func (x *ComponentData) GetUsedPrizeScenes() []int32 {
	if x != nil {
		return x.UsedPrizeScenes
	}
	return nil
}

func (x *ComponentData) GetCoinWin() int32 {
	if x != nil {
		return x.CoinWin
	}
	return 0
}

func (x *ComponentData) GetCashWin() int64 {
	if x != nil {
		return x.CashWin
	}
	return 0
}

func (x *ComponentData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// GameParam
type GameParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstComponent         string                    `protobuf:"bytes,1,opt,name=firstComponent,proto3" json:"firstComponent,omitempty"`
	NextStepFirstComponent string                    `protobuf:"bytes,2,opt,name=nextStepFirstComponent,proto3" json:"nextStepFirstComponent,omitempty"`
	MapComponents          map[string]*ComponentData `protobuf:"bytes,3,rep,name=mapComponents,proto3" json:"mapComponents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GameParam) Reset() {
	*x = GameParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lowcode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameParam) ProtoMessage() {}

func (x *GameParam) ProtoReflect() protoreflect.Message {
	mi := &file_lowcode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameParam.ProtoReflect.Descriptor instead.
func (*GameParam) Descriptor() ([]byte, []int) {
	return file_lowcode_proto_rawDescGZIP(), []int{1}
}

func (x *GameParam) GetFirstComponent() string {
	if x != nil {
		return x.FirstComponent
	}
	return ""
}

func (x *GameParam) GetNextStepFirstComponent() string {
	if x != nil {
		return x.NextStepFirstComponent
	}
	return ""
}

func (x *GameParam) GetMapComponents() map[string]*ComponentData {
	if x != nil {
		return x.MapComponents
	}
	return nil
}

var File_lowcode_proto protoreflect.FileDescriptor

var file_lowcode_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x77, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x73, 0x67, 0x63, 0x37, 0x70, 0x62, 0x22, 0xed, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x64, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x75,
	0x73, 0x65, 0x64, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65,
	0x64, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x64, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x64, 0x50, 0x72, 0x69,
	0x7a, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0f,
	0x75, 0x73, 0x65, 0x64, 0x50, 0x72, 0x69, 0x7a, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x57, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x57, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x73,
	0x68, 0x57, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x61, 0x73, 0x68,
	0x57, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x90, 0x02, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a,
	0x16, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x65, 0x70, 0x46, 0x69, 0x72, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6e,
	0x65, 0x78, 0x74, 0x53, 0x74, 0x65, 0x70, 0x46, 0x69, 0x72, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x0d, 0x6d, 0x61, 0x70, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73,
	0x67, 0x63, 0x37, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e,
	0x4d, 0x61, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x6d, 0x61, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x1a, 0x57, 0x0a, 0x12, 0x4d, 0x61, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x67, 0x63, 0x37, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x73, 0x30, 0x30, 0x37, 0x2f,
	0x73, 0x6c, 0x6f, 0x74, 0x73, 0x67, 0x61, 0x6d, 0x65, 0x63, 0x6f, 0x72, 0x65, 0x37, 0x2f, 0x73,
	0x67, 0x63, 0x37, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lowcode_proto_rawDescOnce sync.Once
	file_lowcode_proto_rawDescData = file_lowcode_proto_rawDesc
)

func file_lowcode_proto_rawDescGZIP() []byte {
	file_lowcode_proto_rawDescOnce.Do(func() {
		file_lowcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_lowcode_proto_rawDescData)
	})
	return file_lowcode_proto_rawDescData
}

var file_lowcode_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_lowcode_proto_goTypes = []interface{}{
	(*ComponentData)(nil), // 0: sgc7pb.ComponentData
	(*GameParam)(nil),     // 1: sgc7pb.GameParam
	nil,                   // 2: sgc7pb.GameParam.MapComponentsEntry
}
var file_lowcode_proto_depIdxs = []int32{
	2, // 0: sgc7pb.GameParam.mapComponents:type_name -> sgc7pb.GameParam.MapComponentsEntry
	0, // 1: sgc7pb.GameParam.MapComponentsEntry.value:type_name -> sgc7pb.ComponentData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lowcode_proto_init() }
func file_lowcode_proto_init() {
	if File_lowcode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lowcode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentData); i {
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
		file_lowcode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameParam); i {
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
			RawDescriptor: file_lowcode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lowcode_proto_goTypes,
		DependencyIndexes: file_lowcode_proto_depIdxs,
		MessageInfos:      file_lowcode_proto_msgTypes,
	}.Build()
	File_lowcode_proto = out.File
	file_lowcode_proto_rawDesc = nil
	file_lowcode_proto_goTypes = nil
	file_lowcode_proto_depIdxs = nil
}

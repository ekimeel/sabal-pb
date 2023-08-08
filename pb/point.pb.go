// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: point.proto

package pb

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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EquipId int64  `protobuf:"varint,2,opt,name=equipId,proto3" json:"equipId,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Point) GetEquipId() int64 {
	if x != nil {
		return x.EquipId
	}
	return 0
}

func (x *Point) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PointList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=Points,proto3" json:"Points,omitempty"`
}

func (x *PointList) Reset() {
	*x = PointList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointList) ProtoMessage() {}

func (x *PointList) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointList.ProtoReflect.Descriptor instead.
func (*PointList) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{1}
}

func (x *PointList) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

type PointId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PointId) Reset() {
	*x = PointId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_point_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointId) ProtoMessage() {}

func (x *PointId) ProtoReflect() protoreflect.Message {
	mi := &file_point_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointId.ProtoReflect.Descriptor instead.
func (*PointId) Descriptor() ([]byte, []int) {
	return file_point_proto_rawDescGZIP(), []int{2}
}

func (x *PointId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_point_proto protoreflect.FileDescriptor

var file_point_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x45, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x71,
	0x75, 0x69, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x09, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x06, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6b, 0x69, 0x6d, 0x65, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_point_proto_rawDescOnce sync.Once
	file_point_proto_rawDescData = file_point_proto_rawDesc
)

func file_point_proto_rawDescGZIP() []byte {
	file_point_proto_rawDescOnce.Do(func() {
		file_point_proto_rawDescData = protoimpl.X.CompressGZIP(file_point_proto_rawDescData)
	})
	return file_point_proto_rawDescData
}

var file_point_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_point_proto_goTypes = []interface{}{
	(*Point)(nil),     // 0: pb.Point
	(*PointList)(nil), // 1: pb.PointList
	(*PointId)(nil),   // 2: pb.PointId
}
var file_point_proto_depIdxs = []int32{
	0, // 0: pb.PointList.Points:type_name -> pb.Point
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_point_proto_init() }
func file_point_proto_init() {
	if File_point_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_point_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_point_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointList); i {
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
		file_point_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointId); i {
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
			RawDescriptor: file_point_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_point_proto_goTypes,
		DependencyIndexes: file_point_proto_depIdxs,
		MessageInfos:      file_point_proto_msgTypes,
	}.Build()
	File_point_proto = out.File
	file_point_proto_rawDesc = nil
	file_point_proto_goTypes = nil
	file_point_proto_depIdxs = nil
}

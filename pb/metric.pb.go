// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: metric.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PointId   int64                  `protobuf:"varint,1,opt,name=pointId,proto3" json:"pointId,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value     float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetPointId() int64 {
	if x != nil {
		return x.PointId
	}
	return 0
}

func (x *Metric) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Metric) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type MetricRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PointId int64                  `protobuf:"varint,1,opt,name=pointId,proto3" json:"pointId,omitempty"`
	From    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *MetricRequest) Reset() {
	*x = MetricRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricRequest) ProtoMessage() {}

func (x *MetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricRequest.ProtoReflect.Descriptor instead.
func (*MetricRequest) Descriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{1}
}

func (x *MetricRequest) GetPointId() int64 {
	if x != nil {
		return x.PointId
	}
	return 0
}

func (x *MetricRequest) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *MetricRequest) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

type MetricList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=Metrics,proto3" json:"Metrics,omitempty"`
}

func (x *MetricList) Reset() {
	*x = MetricList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricList) ProtoMessage() {}

func (x *MetricList) ProtoReflect() protoreflect.Message {
	mi := &file_metric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricList.ProtoReflect.Descriptor instead.
func (*MetricList) Descriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{2}
}

func (x *MetricList) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type MetricWriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

func (x *MetricWriteResponse) Reset() {
	*x = MetricWriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricWriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricWriteResponse) ProtoMessage() {}

func (x *MetricWriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metric_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricWriteResponse.ProtoReflect.Descriptor instead.
func (*MetricWriteResponse) Descriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{3}
}

func (x *MetricWriteResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

var File_metric_proto protoreflect.FileDescriptor

var file_metric_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x22,
	0x32, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x22, 0x31, 0x0a, 0x13, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6b, 0x69, 0x6d, 0x65, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metric_proto_rawDescOnce sync.Once
	file_metric_proto_rawDescData = file_metric_proto_rawDesc
)

func file_metric_proto_rawDescGZIP() []byte {
	file_metric_proto_rawDescOnce.Do(func() {
		file_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_metric_proto_rawDescData)
	})
	return file_metric_proto_rawDescData
}

var file_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_metric_proto_goTypes = []interface{}{
	(*Metric)(nil),                // 0: pb.Metric
	(*MetricRequest)(nil),         // 1: pb.MetricRequest
	(*MetricList)(nil),            // 2: pb.MetricList
	(*MetricWriteResponse)(nil),   // 3: pb.MetricWriteResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_metric_proto_depIdxs = []int32{
	4, // 0: pb.Metric.timestamp:type_name -> google.protobuf.Timestamp
	4, // 1: pb.MetricRequest.from:type_name -> google.protobuf.Timestamp
	4, // 2: pb.MetricRequest.to:type_name -> google.protobuf.Timestamp
	0, // 3: pb.MetricList.Metrics:type_name -> pb.Metric
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_metric_proto_init() }
func file_metric_proto_init() {
	if File_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_metric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricRequest); i {
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
		file_metric_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricList); i {
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
		file_metric_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricWriteResponse); i {
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
			RawDescriptor: file_metric_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metric_proto_goTypes,
		DependencyIndexes: file_metric_proto_depIdxs,
		MessageInfos:      file_metric_proto_msgTypes,
	}.Build()
	File_metric_proto = out.File
	file_metric_proto_rawDesc = nil
	file_metric_proto_goTypes = nil
	file_metric_proto_depIdxs = nil
}

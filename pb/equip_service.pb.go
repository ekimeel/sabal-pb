// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: equip_service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_equip_service_proto protoreflect.FileDescriptor

var file_equip_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x71, 0x75, 0x69, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf3, 0x01, 0x0a, 0x0c, 0x45, 0x71, 0x75, 0x69, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x12, 0x20, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x09,
	0x2e, 0x70, 0x62, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x49, 0x64, 0x1a, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x1a, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6b, 0x69, 0x6d, 0x65, 0x65, 0x6c,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_equip_service_proto_goTypes = []interface{}{
	(*EquipId)(nil),        // 0: pb.EquipId
	(*ListRequest)(nil),    // 1: pb.ListRequest
	(*Equip)(nil),          // 2: pb.Equip
	(*EquipList)(nil),      // 3: pb.EquipList
	(*DeleteResponse)(nil), // 4: pb.DeleteResponse
}
var file_equip_service_proto_depIdxs = []int32{
	0, // 0: pb.EquipService.Get:input_type -> pb.EquipId
	1, // 1: pb.EquipService.GetAll:input_type -> pb.ListRequest
	2, // 2: pb.EquipService.Create:input_type -> pb.Equip
	2, // 3: pb.EquipService.Update:input_type -> pb.Equip
	0, // 4: pb.EquipService.Delete:input_type -> pb.EquipId
	2, // 5: pb.EquipService.GetOrCreate:input_type -> pb.Equip
	2, // 6: pb.EquipService.Get:output_type -> pb.Equip
	3, // 7: pb.EquipService.GetAll:output_type -> pb.EquipList
	2, // 8: pb.EquipService.Create:output_type -> pb.Equip
	2, // 9: pb.EquipService.Update:output_type -> pb.Equip
	4, // 10: pb.EquipService.Delete:output_type -> pb.DeleteResponse
	2, // 11: pb.EquipService.GetOrCreate:output_type -> pb.Equip
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_equip_service_proto_init() }
func file_equip_service_proto_init() {
	if File_equip_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_equip_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_equip_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_equip_service_proto_goTypes,
		DependencyIndexes: file_equip_service_proto_depIdxs,
	}.Build()
	File_equip_service_proto = out.File
	file_equip_service_proto_rawDesc = nil
	file_equip_service_proto_goTypes = nil
	file_equip_service_proto_depIdxs = nil
}

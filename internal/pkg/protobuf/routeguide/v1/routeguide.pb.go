// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: routeguide/v1/routeguide.proto

package routeguidev1

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_routeguide_v1_routeguide_proto protoreflect.FileDescriptor

var file_routeguide_v1_routeguide_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc1, 0x03, 0x0a, 0x11,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x8a, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x20, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x12, 0x2f, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x7d, 0x2f, 0x7b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x7d, 0x12, 0x6f,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x22,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12,
	0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x30, 0x01, 0x12,
	0x58, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x21,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x54, 0x0a, 0x09, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x85, 0x02, 0xba, 0x47, 0x37, 0x12, 0x1c, 0x0a, 0x13, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x20, 0x47,
	0x75, 0x69, 0x64, 0x65, 0x20, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x05, 0x30, 0x2e,
	0x30, 0x2e, 0x31, 0x1a, 0x17, 0x0a, 0x15, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x31, 0x0a, 0x11, 0x63, 0x6f,
	0x6d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x0f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x61, 0x6b, 0x6f, 0x61, 0x2f, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_routeguide_v1_routeguide_proto_goTypes = []interface{}{
	(*GetFeatureRequest)(nil),    // 0: routeguide.v1.GetFeatureRequest
	(*ListFeaturesRequest)(nil),  // 1: routeguide.v1.ListFeaturesRequest
	(*RecordRouteRequest)(nil),   // 2: routeguide.v1.RecordRouteRequest
	(*RouteChatRequest)(nil),     // 3: routeguide.v1.RouteChatRequest
	(*GetFeatureResponse)(nil),   // 4: routeguide.v1.GetFeatureResponse
	(*ListFeaturesResponse)(nil), // 5: routeguide.v1.ListFeaturesResponse
	(*RecordRouteResponse)(nil),  // 6: routeguide.v1.RecordRouteResponse
	(*RouteChatResponse)(nil),    // 7: routeguide.v1.RouteChatResponse
}
var file_routeguide_v1_routeguide_proto_depIdxs = []int32{
	0, // 0: routeguide.v1.RouteGuideService.GetFeature:input_type -> routeguide.v1.GetFeatureRequest
	1, // 1: routeguide.v1.RouteGuideService.ListFeatures:input_type -> routeguide.v1.ListFeaturesRequest
	2, // 2: routeguide.v1.RouteGuideService.RecordRoute:input_type -> routeguide.v1.RecordRouteRequest
	3, // 3: routeguide.v1.RouteGuideService.RouteChat:input_type -> routeguide.v1.RouteChatRequest
	4, // 4: routeguide.v1.RouteGuideService.GetFeature:output_type -> routeguide.v1.GetFeatureResponse
	5, // 5: routeguide.v1.RouteGuideService.ListFeatures:output_type -> routeguide.v1.ListFeaturesResponse
	6, // 6: routeguide.v1.RouteGuideService.RecordRoute:output_type -> routeguide.v1.RecordRouteResponse
	7, // 7: routeguide.v1.RouteGuideService.RouteChat:output_type -> routeguide.v1.RouteChatResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_routeguide_v1_routeguide_proto_init() }
func file_routeguide_v1_routeguide_proto_init() {
	if File_routeguide_v1_routeguide_proto != nil {
		return
	}
	file_routeguide_v1_feature_proto_init()
	file_routeguide_v1_route_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_routeguide_v1_routeguide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routeguide_v1_routeguide_proto_goTypes,
		DependencyIndexes: file_routeguide_v1_routeguide_proto_depIdxs,
	}.Build()
	File_routeguide_v1_routeguide_proto = out.File
	file_routeguide_v1_routeguide_proto_rawDesc = nil
	file_routeguide_v1_routeguide_proto_goTypes = nil
	file_routeguide_v1_routeguide_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: handler.proto

package user

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

var File_handler_proto protoreflect.FileDescriptor

var file_handler_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xd4, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x12, 0x2c, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2f, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_handler_proto_goTypes = []interface{}{
	(*AddRequest)(nil),     // 0: user.AddRequest
	(*FindRequest)(nil),    // 1: user.FindRequest
	(*EditRequest)(nil),    // 2: user.EditRequest
	(*DeleteRequest)(nil),  // 3: user.DeleteRequest
	(*AddResponse)(nil),    // 4: user.AddResponse
	(*FindResponse)(nil),   // 5: user.FindResponse
	(*EditResponse)(nil),   // 6: user.EditResponse
	(*DeleteResponse)(nil), // 7: user.DeleteResponse
}
var file_handler_proto_depIdxs = []int32{
	0, // 0: user.UserHandler.Add:input_type -> user.AddRequest
	1, // 1: user.UserHandler.Find:input_type -> user.FindRequest
	2, // 2: user.UserHandler.Edit:input_type -> user.EditRequest
	3, // 3: user.UserHandler.Delete:input_type -> user.DeleteRequest
	4, // 4: user.UserHandler.Add:output_type -> user.AddResponse
	5, // 5: user.UserHandler.Find:output_type -> user.FindResponse
	6, // 6: user.UserHandler.Edit:output_type -> user.EditResponse
	7, // 7: user.UserHandler.Delete:output_type -> user.DeleteResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_handler_proto_init() }
func file_handler_proto_init() {
	if File_handler_proto != nil {
		return
	}
	file_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handler_proto_goTypes,
		DependencyIndexes: file_handler_proto_depIdxs,
	}.Build()
	File_handler_proto = out.File
	file_handler_proto_rawDesc = nil
	file_handler_proto_goTypes = nil
	file_handler_proto_depIdxs = nil
}

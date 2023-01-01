// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: enum.proto

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

type UserRole int32

const (
	UserRole_USER_UNKNOWN UserRole = 0
	UserRole_SUPER_ADMIN  UserRole = 1
	UserRole_ADMIN        UserRole = 2
	UserRole_SELLER       UserRole = 3
	UserRole_MANAGER      UserRole = 4
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "USER_UNKNOWN",
		1: "SUPER_ADMIN",
		2: "ADMIN",
		3: "SELLER",
		4: "MANAGER",
	}
	UserRole_value = map[string]int32{
		"USER_UNKNOWN": 0,
		"SUPER_ADMIN":  1,
		"ADMIN":        2,
		"SELLER":       3,
		"MANAGER":      4,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[0].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[0]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{0}
}

type ProjectStatus int32

const (
	ProjectStatus_ProjectStatus_NONE        ProjectStatus = 0
	ProjectStatus_ProjectStatus_IN_PROGRESS ProjectStatus = 1
	ProjectStatus_ProjectStatus_PAUSE       ProjectStatus = 2
	ProjectStatus_ProjectStatus_CLOSE       ProjectStatus = 3
	ProjectStatus_ProjectStatus_DRAFT       ProjectStatus = 4
)

// Enum value maps for ProjectStatus.
var (
	ProjectStatus_name = map[int32]string{
		0: "ProjectStatus_NONE",
		1: "ProjectStatus_IN_PROGRESS",
		2: "ProjectStatus_PAUSE",
		3: "ProjectStatus_CLOSE",
		4: "ProjectStatus_DRAFT",
	}
	ProjectStatus_value = map[string]int32{
		"ProjectStatus_NONE":        0,
		"ProjectStatus_IN_PROGRESS": 1,
		"ProjectStatus_PAUSE":       2,
		"ProjectStatus_CLOSE":       3,
		"ProjectStatus_DRAFT":       4,
	}
)

func (x ProjectStatus) Enum() *ProjectStatus {
	p := new(ProjectStatus)
	*p = x
	return p
}

func (x ProjectStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_enum_proto_enumTypes[1].Descriptor()
}

func (ProjectStatus) Type() protoreflect.EnumType {
	return &file_enum_proto_enumTypes[1]
}

func (x ProjectStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectStatus.Descriptor instead.
func (ProjectStatus) EnumDescriptor() ([]byte, []int) {
	return file_enum_proto_rawDescGZIP(), []int{1}
}

var File_enum_proto protoreflect.FileDescriptor

var file_enum_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x6e,
	0x75, 0x6d, 0x5f, 0x70, 0x62, 0x2a, 0x51, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x55, 0x50, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x45, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4d,
	0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x04, 0x2a, 0x91, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x04, 0x42, 0x08, 0x5a, 0x06,
	0x70, 0x62, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enum_proto_rawDescOnce sync.Once
	file_enum_proto_rawDescData = file_enum_proto_rawDesc
)

func file_enum_proto_rawDescGZIP() []byte {
	file_enum_proto_rawDescOnce.Do(func() {
		file_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_enum_proto_rawDescData)
	})
	return file_enum_proto_rawDescData
}

var file_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_enum_proto_goTypes = []interface{}{
	(UserRole)(0),      // 0: enum_pb.UserRole
	(ProjectStatus)(0), // 1: enum_pb.ProjectStatus
}
var file_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enum_proto_init() }
func file_enum_proto_init() {
	if File_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enum_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enum_proto_goTypes,
		DependencyIndexes: file_enum_proto_depIdxs,
		EnumInfos:         file_enum_proto_enumTypes,
	}.Build()
	File_enum_proto = out.File
	file_enum_proto_rawDesc = nil
	file_enum_proto_goTypes = nil
	file_enum_proto_depIdxs = nil
}

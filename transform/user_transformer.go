// Code generated by protoc-gen-struct-transformer, version: <dev>. DO NOT EDIT.
// source file: user.proto
// source package: user

package transform

import (
	"github.com/duyledat197/interview-hao/internal/models"
	"github.com/duyledat197/interview-hao/pb"
)

// "google.protobuf.MethodOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.StringRules": target: "", Omitted: true, OneofDecl: ""
// "user.User": target: "User", Omitted: false, OneofDecl: ""
// "google.protobuf.EnumValueOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MessageOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ServiceOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.FieldRules": target: "", Omitted: true, OneofDecl: ""
// "auth.ForgotPasswordResponse": target: "", Omitted: true, OneofDecl: ""
// "google.api.Http": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "validate.Int32Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.MapRules": target: "", Omitted: true, OneofDecl: ""
// "validate.DoubleRules": target: "", Omitted: true, OneofDecl: ""
// "auth.ForgotPasswordRequest": target: "", Omitted: true, OneofDecl: ""
// "user.CreateUserRequest": target: "", Omitted: true, OneofDecl: ""
// "user.CreateUserResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MethodDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FieldOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.GeneratedCodeInfo": target: "", Omitted: true, OneofDecl: ""
// "validate.UInt32Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.Fixed32Rules": target: "", Omitted: true, OneofDecl: ""
// "auth.VerifyOTPRequest": target: "", Omitted: true, OneofDecl: ""
// "google.api.CustomHttpPattern": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FieldDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ServiceDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.SourceCodeInfo": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.Timestamp": target: "", Omitted: true, OneofDecl: ""
// "validate.SInt32Rules": target: "", Omitted: true, OneofDecl: ""
// "user.GetUserByIDRequest": target: "", Omitted: true, OneofDecl: ""
// "google.api.HttpRule": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ExtensionRangeOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.SFixed32Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.RepeatedRules": target: "", Omitted: true, OneofDecl: ""
// "validate.UInt64Rules": target: "", Omitted: true, OneofDecl: ""
// "auth.ResetPasswordRequest": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileDescriptorSet": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.OneofOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.Duration": target: "", Omitted: true, OneofDecl: ""
// "validate.Int64Rules": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.DescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "validate.SFixed64Rules": target: "", Omitted: true, OneofDecl: ""
// "auth.VerifyOTPResponse": target: "", Omitted: true, OneofDecl: ""
// "user.GetListResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.UninterpretedOption": target: "", Omitted: true, OneofDecl: ""
// "validate.BoolRules": target: "", Omitted: true, OneofDecl: ""
// "validate.DurationRules": target: "", Omitted: true, OneofDecl: ""
// "auth.ChangePasswordRequest": target: "", Omitted: true, OneofDecl: ""
// "validate.MessageRules": target: "", Omitted: true, OneofDecl: ""
// "user.GetListRequest": target: "", Omitted: true, OneofDecl: ""
// "auth.ChangePasswordResponse": target: "", Omitted: true, OneofDecl: ""
// "user.UpdateUserRequest": target: "", Omitted: true, OneofDecl: ""
// "user.UpdateUserResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.AnyRules": target: "", Omitted: true, OneofDecl: ""
// "validate.TimestampRules": target: "", Omitted: true, OneofDecl: ""
// "auth.ResetPasswordResponse": target: "", Omitted: true, OneofDecl: ""
// "validate.Fixed64Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.EnumRules": target: "", Omitted: true, OneofDecl: ""
// "auth.LoginRequest": target: "", Omitted: true, OneofDecl: ""
// "auth.LoginResponse": target: "", Omitted: true, OneofDecl: ""
// "validate.SInt64Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.BytesRules": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.OneofDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumValueDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.FloatRules": target: "", Omitted: true, OneofDecl: ""
// message "GetUserByIDRequest" has no option "transformer.go_struct", skipped...

// Target struct fields:
// Field: "ID", Type: "int64", isPointer: false
// Field: "Name", Type: "pgtype.Text", isPointer: false
// Field: "Bio", Type: "sql.NullString", isPointer: false

// ===============================
// fdp.Name: "id", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"Int64", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "Id", gname: "ID"

// ===============================
// fdp.Name: "name", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"PgtypeText", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "Name", gname: "Name"

// ===============================
// fdp.Name: "bio", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"SqlNullString", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "Bio", gname: "Bio"
// message "CreateUserRequest" has no option "transformer.go_struct", skipped...
// message "CreateUserResponse" has no option "transformer.go_struct", skipped...
// message "GetListRequest" has no option "transformer.go_struct", skipped...
// message "GetListResponse" has no option "transformer.go_struct", skipped...
// message "UpdateUserRequest" has no option "transformer.go_struct", skipped...
// message "UpdateUserResponse" has no option "transformer.go_struct", skipped...
func PbToUserPtr(src *pb.User, opts ...TransformParam) *models.User {
	if src == nil {
		return nil
	}

	d := PbToUser(*src, opts...)
	return &d
}

func PbToUserPtrList(src []*pb.User, opts ...TransformParam) []*models.User {
	resp := make([]*models.User, len(src))

	for i, s := range src {
		resp[i] = PbToUserPtr(s, opts...)
	}

	return resp
}

func PbToUserPtrVal(src *pb.User, opts ...TransformParam) models.User {
	if src == nil {
		return models.User{}
	}

	return PbToUser(*src, opts...)
}

func PbToUserPtrValList(src []*pb.User, opts ...TransformParam) []models.User {
	resp := make([]models.User, len(src))

	for i, s := range src {
		resp[i] = PbToUser(*s)
	}

	return resp
}

// PbToUserList is DEPRECATED. Use PbToUserPtrValList instead.
func PbToUserList(src []*pb.User, opts ...TransformParam) []models.User {
	return PbToUserPtrValList(src)
}

func PbToUser(src pb.User, opts ...TransformParam) models.User {
	s := models.User{
		ID:   transformhelpers.StringToInt64(src.Id),
		Name: transformhelpers.StringToPgtypeText(src.Name),
		Bio:  transformhelpers.StringToSqlNullString(src.Bio),
	}

	applyOptions(opts...)

	return s
}

func PbToUserValPtr(src pb.User, opts ...TransformParam) *models.User {
	d := PbToUser(src, opts...)
	return &d
}

func PbToUserValList(src []pb.User, opts ...TransformParam) []models.User {
	resp := make([]models.User, len(src))

	for i, s := range src {
		resp[i] = PbToUser(s, opts...)
	}

	return resp
}

func UserToPbPtr(src *models.User, opts ...TransformParam) *pb.User {
	if src == nil {
		return nil
	}

	d := UserToPb(*src, opts...)
	return &d
}

func UserToPbPtrList(src []*models.User, opts ...TransformParam) []*pb.User {
	resp := make([]*pb.User, len(src))

	for i, s := range src {
		resp[i] = UserToPbPtr(s, opts...)
	}

	return resp
}

func UserToPbPtrVal(src *models.User, opts ...TransformParam) pb.User {
	if src == nil {
		return pb.User{}
	}

	return UserToPb(*src, opts...)
}

func UserToPbValPtrList(src []models.User, opts ...TransformParam) []*pb.User {
	resp := make([]*pb.User, len(src))

	for i, s := range src {
		g := UserToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// UserToPbList is DEPRECATED. Use UserToPbValPtrList instead.
func UserToPbList(src []models.User, opts ...TransformParam) []*pb.User {
	return UserToPbValPtrList(src)
}

func UserToPb(src models.User, opts ...TransformParam) pb.User {
	s := pb.User{
		Id:   transformhelpers.Int64ToString(src.ID),
		Name: transformhelpers.PgtypeTextToString(src.Name),
		Bio:  transformhelpers.SqlNullStringToString(src.Bio),
	}

	applyOptions(opts...)

	return s
}

func UserToPbValPtr(src models.User, opts ...TransformParam) *pb.User {
	d := UserToPb(src, opts...)
	return &d
}

func UserToPbValList(src []models.User, opts ...TransformParam) []pb.User {
	resp := make([]pb.User, len(src))

	for i, s := range src {
		resp[i] = UserToPb(s, opts...)
	}

	return resp
}

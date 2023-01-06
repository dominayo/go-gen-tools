// Code generated by protoc-gen-struct-transformer, version: <dev>. DO NOT EDIT.
// source file: hub.proto
// source package: hub

package transform

import (
	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/pb"
	"github.com/duyledat197/go-gen-tools/utils/transformhelpers"
)

// "google.protobuf.DescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ExtensionRangeOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.UInt64Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.BoolRules": target: "", Omitted: true, OneofDecl: ""
// "validate.RepeatedRules": target: "", Omitted: true, OneofDecl: ""
// "team.CreateTeamRequest": target: "", Omitted: true, OneofDecl: ""
// "search.SearchTeamHubRequest": target: "", Omitted: true, OneofDecl: ""
// "user.GetListUserResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ServiceDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MethodDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "validate.Fixed32Rules": target: "", Omitted: true, OneofDecl: ""
// "team.GetTeamByIDRequest": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileDescriptorSet": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.UninterpretedOption": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.Timestamp": target: "", Omitted: true, OneofDecl: ""
// "validate.StringRules": target: "", Omitted: true, OneofDecl: ""
// "validate.MessageRules": target: "", Omitted: true, OneofDecl: ""
// "hub.CreateHubResponse": target: "", Omitted: true, OneofDecl: ""
// "hub.GetListHubRequest": target: "", Omitted: true, OneofDecl: ""
// "google.api.CustomHttpPattern": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.ServiceOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.DoubleRules": target: "", Omitted: true, OneofDecl: ""
// "validate.SFixed32Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.AnyRules": target: "", Omitted: true, OneofDecl: ""
// "search.SearchTeamHubResponse": target: "", Omitted: true, OneofDecl: ""
// "user.GetListUserRequest": target: "", Omitted: true, OneofDecl: ""
// "google.api.HttpRule": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumValueDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "team.UpdateTeamRequest": target: "", Omitted: true, OneofDecl: ""
// "user.User": target: "User", Omitted: false, OneofDecl: ""
// "google.protobuf.SourceCodeInfo": target: "", Omitted: true, OneofDecl: ""
// "validate.SFixed64Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.DurationRules": target: "", Omitted: true, OneofDecl: ""
// "team.Team": target: "Team", Omitted: false, OneofDecl: ""
// "team.GetListTeamRequest": target: "", Omitted: true, OneofDecl: ""
// "user.GetUserByIDResponse": target: "", Omitted: true, OneofDecl: ""
// "user.UpdateUserRequest": target: "", Omitted: true, OneofDecl: ""
// "validate.FieldRules": target: "", Omitted: true, OneofDecl: ""
// "validate.FloatRules": target: "", Omitted: true, OneofDecl: ""
// "validate.MapRules": target: "", Omitted: true, OneofDecl: ""
// "hub.GetHubByIDRequest": target: "", Omitted: true, OneofDecl: ""
// "hub.CreateHubRequest": target: "", Omitted: true, OneofDecl: ""
// "hub.UpdateHubRequest": target: "", Omitted: true, OneofDecl: ""
// "user.CreateUserRequest": target: "", Omitted: true, OneofDecl: ""
// "team.GetTeamByIDResponse": target: "", Omitted: true, OneofDecl: ""
// "user.UpdateUserResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.OneofDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.OneofOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.SInt32Rules": target: "", Omitted: true, OneofDecl: ""
// "hub.GetHubByIDResponse": target: "", Omitted: true, OneofDecl: ""
// "hub.UpdateHubResponse": target: "", Omitted: true, OneofDecl: ""
// "team.CreateTeamResponse": target: "", Omitted: true, OneofDecl: ""
// "user.GetUserByIDRequest": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FieldOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.EnumValueOptions": target: "", Omitted: true, OneofDecl: ""
// "google.api.Http": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FileOptions": target: "", Omitted: true, OneofDecl: ""
// "validate.Int64Rules": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MethodOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.Duration": target: "", Omitted: true, OneofDecl: ""
// "validate.UInt32Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.EnumRules": target: "", Omitted: true, OneofDecl: ""
// "validate.TimestampRules": target: "", Omitted: true, OneofDecl: ""
// "team.UpdateTeamResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.MessageOptions": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.GeneratedCodeInfo": target: "", Omitted: true, OneofDecl: ""
// "hub.Hub": target: "Hub", Omitted: false, OneofDecl: ""
// "user.CreateUserResponse": target: "", Omitted: true, OneofDecl: ""
// "google.protobuf.FieldDescriptorProto": target: "", Omitted: true, OneofDecl: ""
// "validate.Fixed64Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.BytesRules": target: "", Omitted: true, OneofDecl: ""
// "validate.Int32Rules": target: "", Omitted: true, OneofDecl: ""
// "validate.SInt64Rules": target: "", Omitted: true, OneofDecl: ""
// "hub.GetListHubResponse": target: "", Omitted: true, OneofDecl: ""
// "team.GetListTeamResponse": target: "", Omitted: true, OneofDecl: ""
// message "GetHubByIDRequest" has no option "transformer.go_struct", skipped...
// message "GetHubByIDResponse" has no option "transformer.go_struct", skipped...

// Target struct fields:
// Field: "ID", Type: "pgtype.Text", isPointer: false
// Field: "Name", Type: "pgtype.Text", isPointer: false
// Field: "LocationID", Type: "pgtype.Text", isPointer: false
// Field: "CreatedAt", Type: "pgtype.Timestamptz", isPointer: false
// Field: "UpdatedAt", Type: "pgtype.Timestamptz", isPointer: false
// Field: "DeletedAt", Type: "pgtype.Timestamptz", isPointer: false

// ===============================
// fdp.Name: "id", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"PgtypeText", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "Id", gname: "ID"

// ===============================
// fdp.Name: "name", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"PgtypeText", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "Name", gname: "Name"

// ===============================
// fdp.Name: "location_id", mapAs: "", mapTo: ""
// sf: source.FieldInfo{Type:"PgtypeText", IsPointer:false}, pbType: "", goType: "string", ft: "TYPE_STRING", pname: "LocationId", gname: "LocationID"
// message "CreateHubRequest" has no option "transformer.go_struct", skipped...
// message "CreateHubResponse" has no option "transformer.go_struct", skipped...
// message "GetListHubRequest" has no option "transformer.go_struct", skipped...
// message "GetListHubResponse" has no option "transformer.go_struct", skipped...
// message "UpdateHubRequest" has no option "transformer.go_struct", skipped...
// message "UpdateHubResponse" has no option "transformer.go_struct", skipped...
func PbToHubPtr(src *pb.Hub, opts ...TransformParam) *models.Hub {
	if src == nil {
		return nil
	}

	d := PbToHub(*src, opts...)
	return &d
}

func PbToHubPtrList(src []*pb.Hub, opts ...TransformParam) []*models.Hub {
	resp := make([]*models.Hub, len(src))

	for i, s := range src {
		resp[i] = PbToHubPtr(s, opts...)
	}

	return resp
}

func PbToHubPtrVal(src *pb.Hub, opts ...TransformParam) models.Hub {
	if src == nil {
		return models.Hub{}
	}

	return PbToHub(*src, opts...)
}

func PbToHubPtrValList(src []*pb.Hub, opts ...TransformParam) []models.Hub {
	resp := make([]models.Hub, len(src))

	for i, s := range src {
		resp[i] = PbToHub(*s)
	}

	return resp
}

// PbToHubList is DEPRECATED. Use PbToHubPtrValList instead.
func PbToHubList(src []*pb.Hub, opts ...TransformParam) []models.Hub {
	return PbToHubPtrValList(src)
}

func PbToHub(src pb.Hub, opts ...TransformParam) models.Hub {
	s := models.Hub{
		ID:         transformhelpers.StringToPgtypeText(src.Id),
		Name:       transformhelpers.StringToPgtypeText(src.Name),
		LocationID: transformhelpers.StringToPgtypeText(src.LocationId),
	}

	applyOptions(opts...)

	return s
}

func PbToHubValPtr(src pb.Hub, opts ...TransformParam) *models.Hub {
	d := PbToHub(src, opts...)
	return &d
}

func PbToHubValList(src []pb.Hub, opts ...TransformParam) []models.Hub {
	resp := make([]models.Hub, len(src))

	for i, s := range src {
		resp[i] = PbToHub(s, opts...)
	}

	return resp
}

func HubToPbPtr(src *models.Hub, opts ...TransformParam) *pb.Hub {
	if src == nil {
		return nil
	}

	d := HubToPb(*src, opts...)
	return &d
}

func HubToPbPtrList(src []*models.Hub, opts ...TransformParam) []*pb.Hub {
	resp := make([]*pb.Hub, len(src))

	for i, s := range src {
		resp[i] = HubToPbPtr(s, opts...)
	}

	return resp
}

func HubToPbPtrVal(src *models.Hub, opts ...TransformParam) pb.Hub {
	if src == nil {
		return pb.Hub{}
	}

	return HubToPb(*src, opts...)
}

func HubToPbValPtrList(src []models.Hub, opts ...TransformParam) []*pb.Hub {
	resp := make([]*pb.Hub, len(src))

	for i, s := range src {
		g := HubToPb(s, opts...)
		resp[i] = &g
	}

	return resp
}

// HubToPbList is DEPRECATED. Use HubToPbValPtrList instead.
func HubToPbList(src []models.Hub, opts ...TransformParam) []*pb.Hub {
	return HubToPbValPtrList(src)
}

func HubToPb(src models.Hub, opts ...TransformParam) pb.Hub {
	s := pb.Hub{
		Id:         transformhelpers.PgtypeTextToString(src.ID),
		Name:       transformhelpers.PgtypeTextToString(src.Name),
		LocationId: transformhelpers.PgtypeTextToString(src.LocationID),
	}

	applyOptions(opts...)

	return s
}

func HubToPbValPtr(src models.Hub, opts ...TransformParam) *pb.Hub {
	d := HubToPb(src, opts...)
	return &d
}

func HubToPbValList(src []models.Hub, opts ...TransformParam) []pb.Hub {
	resp := make([]pb.Hub, len(src))

	for i, s := range src {
		resp[i] = HubToPb(s, opts...)
	}

	return resp
}

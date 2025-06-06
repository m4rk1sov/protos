// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: profile/profile.proto

package profilev1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserProfile struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// UUIDs for ID fields?
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Address       string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Email         string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	mi := &file_profile_profile_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_profile_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_profile_profile_proto_rawDescGZIP(), []int{0}
}

func (x *UserProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserProfile) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserProfile) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserProfile) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address       string                 `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProfileRequest) Reset() {
	*x = CreateProfileRequest{}
	mi := &file_profile_profile_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileRequest) ProtoMessage() {}

func (x *CreateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_profile_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_profile_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProfileRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProfileRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateProfileRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateProfileRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	mi := &file_profile_profile_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_profile_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_profile_proto_rawDescGZIP(), []int{2}
}

func (x *GetProfileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProfileByUserIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProfileByUserIDRequest) Reset() {
	*x = GetProfileByUserIDRequest{}
	mi := &file_profile_profile_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProfileByUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileByUserIDRequest) ProtoMessage() {}

func (x *GetProfileByUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_profile_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileByUserIDRequest.ProtoReflect.Descriptor instead.
func (*GetProfileByUserIDRequest) Descriptor() ([]byte, []int) {
	return file_profile_profile_proto_rawDescGZIP(), []int{3}
}

func (x *GetProfileByUserIDRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UpdateProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Phone         *string                `protobuf:"bytes,3,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Address       *string                `protobuf:"bytes,4,opt,name=address,proto3,oneof" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProfileRequest) Reset() {
	*x = UpdateProfileRequest{}
	mi := &file_profile_profile_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileRequest) ProtoMessage() {}

func (x *UpdateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_profile_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_profile_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProfileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProfileRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateProfileRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UpdateProfileRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

// Is it needed?
type DeleteProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProfileRequest) Reset() {
	*x = DeleteProfileRequest{}
	mi := &file_profile_profile_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileRequest) ProtoMessage() {}

func (x *DeleteProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_profile_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileRequest.ProtoReflect.Descriptor instead.
func (*DeleteProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_profile_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteProfileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListProfilesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProfilesRequest) Reset() {
	*x = ListProfilesRequest{}
	mi := &file_profile_profile_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProfilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfilesRequest) ProtoMessage() {}

func (x *ListProfilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_profile_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProfilesRequest.ProtoReflect.Descriptor instead.
func (*ListProfilesRequest) Descriptor() ([]byte, []int) {
	return file_profile_profile_proto_rawDescGZIP(), []int{6}
}

func (x *ListProfilesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProfilesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ProfileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Profile       *UserProfile           `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfileResponse) Reset() {
	*x = ProfileResponse{}
	mi := &file_profile_profile_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileResponse) ProtoMessage() {}

func (x *ProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_profile_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileResponse.ProtoReflect.Descriptor instead.
func (*ProfileResponse) Descriptor() ([]byte, []int) {
	return file_profile_profile_proto_rawDescGZIP(), []int{7}
}

func (x *ProfileResponse) GetProfile() *UserProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type ListProfilesResponse struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Profiles []*UserProfile         `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
	// is int32 enough?
	Total         int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProfilesResponse) Reset() {
	*x = ListProfilesResponse{}
	mi := &file_profile_profile_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProfilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfilesResponse) ProtoMessage() {}

func (x *ListProfilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_profile_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProfilesResponse.ProtoReflect.Descriptor instead.
func (*ListProfilesResponse) Descriptor() ([]byte, []int) {
	return file_profile_profile_proto_rawDescGZIP(), []int{8}
}

func (x *ListProfilesResponse) GetProfiles() []*UserProfile {
	if x != nil {
		return x.Profiles
	}
	return nil
}

func (x *ListProfilesResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_profile_profile_proto protoreflect.FileDescriptor

const file_profile_profile_proto_rawDesc = "" +
	"\n" +
	"\x15profile/profile.proto\x12\aprofile\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x90\x01\n" +
	"\vUserProfile\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x03R\x06userId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\x12\x18\n" +
	"\aaddress\x18\x05 \x01(\tR\aaddress\x12\x14\n" +
	"\x05email\x18\x06 \x01(\tR\x05email\"\x89\x01\n" +
	"\x14CreateProfileRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05phone\x18\x03 \x01(\tR\x05phone\x12\x18\n" +
	"\aaddress\x18\x04 \x01(\tR\aaddress\x12\x14\n" +
	"\x05email\x18\x05 \x01(\tR\x05email\"#\n" +
	"\x11GetProfileRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"4\n" +
	"\x19GetProfileByUserIDRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x03R\x06userId\"\x98\x01\n" +
	"\x14UpdateProfileRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\x04name\x18\x02 \x01(\tH\x00R\x04name\x88\x01\x01\x12\x19\n" +
	"\x05phone\x18\x03 \x01(\tH\x01R\x05phone\x88\x01\x01\x12\x1d\n" +
	"\aaddress\x18\x04 \x01(\tH\x02R\aaddress\x88\x01\x01B\a\n" +
	"\x05_nameB\b\n" +
	"\x06_phoneB\n" +
	"\n" +
	"\b_address\"&\n" +
	"\x14DeleteProfileRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"?\n" +
	"\x13ListProfilesRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\"A\n" +
	"\x0fProfileResponse\x12.\n" +
	"\aprofile\x18\x01 \x01(\v2\x14.profile.UserProfileR\aprofile\"^\n" +
	"\x14ListProfilesResponse\x120\n" +
	"\bprofiles\x18\x01 \x03(\v2\x14.profile.UserProfileR\bprofiles\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total2\xc5\x05\n" +
	"\x0eProfileService\x12m\n" +
	"\rCreateProfile\x12\x1d.profile.CreateProfileRequest\x1a\x18.profile.ProfileResponse\"#\x92A\t\n" +
	"\aProfile\x82\xd3\xe4\x93\x02\x11:\x01*\"\f/v1/profiles\x12i\n" +
	"\n" +
	"GetProfile\x12\x1a.profile.GetProfileRequest\x1a\x18.profile.ProfileResponse\"%\x92A\t\n" +
	"\aProfile\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/profiles/{id}\x12\x86\x01\n" +
	"\x12GetProfileByUserID\x12\".profile.GetProfileByUserIDRequest\x1a\x18.profile.ProfileResponse\"2\x92A\t\n" +
	"\aProfile\x82\xd3\xe4\x93\x02 \x12\x1e/v1/profiles/by-user/{user_id}\x12r\n" +
	"\rUpdateProfile\x12\x1d.profile.UpdateProfileRequest\x1a\x18.profile.ProfileResponse\"(\x92A\t\n" +
	"\aProfile\x82\xd3\xe4\x93\x02\x16:\x01*2\x11/v1/profiles/{id}\x12m\n" +
	"\rDeleteProfile\x12\x1d.profile.DeleteProfileRequest\x1a\x16.google.protobuf.Empty\"%\x92A\t\n" +
	"\aProfile\x82\xd3\xe4\x93\x02\x13*\x11/v1/profiles/{id}\x12m\n" +
	"\fListProfiles\x12\x1c.profile.ListProfilesRequest\x1a\x1d.profile.ListProfilesResponse\" \x92A\t\n" +
	"\aProfile\x82\xd3\xe4\x93\x02\x0e\x12\f/v1/profilesB\xbb\x01\x92A\x98\x01\x12o\n" +
	"\x13Profile Service API\x12SHandles user profile: creation, retrieving, updating, deleting and listing profiles2\x031.0*\x01\x012\x10application/json:\x10application/jsonZ\x1dm4rk1sov.profile.v1;profilev1b\x06proto3"

var (
	file_profile_profile_proto_rawDescOnce sync.Once
	file_profile_profile_proto_rawDescData []byte
)

func file_profile_profile_proto_rawDescGZIP() []byte {
	file_profile_profile_proto_rawDescOnce.Do(func() {
		file_profile_profile_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_profile_profile_proto_rawDesc), len(file_profile_profile_proto_rawDesc)))
	})
	return file_profile_profile_proto_rawDescData
}

var file_profile_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_profile_profile_proto_goTypes = []any{
	(*UserProfile)(nil),               // 0: profile.UserProfile
	(*CreateProfileRequest)(nil),      // 1: profile.CreateProfileRequest
	(*GetProfileRequest)(nil),         // 2: profile.GetProfileRequest
	(*GetProfileByUserIDRequest)(nil), // 3: profile.GetProfileByUserIDRequest
	(*UpdateProfileRequest)(nil),      // 4: profile.UpdateProfileRequest
	(*DeleteProfileRequest)(nil),      // 5: profile.DeleteProfileRequest
	(*ListProfilesRequest)(nil),       // 6: profile.ListProfilesRequest
	(*ProfileResponse)(nil),           // 7: profile.ProfileResponse
	(*ListProfilesResponse)(nil),      // 8: profile.ListProfilesResponse
	(*emptypb.Empty)(nil),             // 9: google.protobuf.Empty
}
var file_profile_profile_proto_depIdxs = []int32{
	0, // 0: profile.ProfileResponse.profile:type_name -> profile.UserProfile
	0, // 1: profile.ListProfilesResponse.profiles:type_name -> profile.UserProfile
	1, // 2: profile.ProfileService.CreateProfile:input_type -> profile.CreateProfileRequest
	2, // 3: profile.ProfileService.GetProfile:input_type -> profile.GetProfileRequest
	3, // 4: profile.ProfileService.GetProfileByUserID:input_type -> profile.GetProfileByUserIDRequest
	4, // 5: profile.ProfileService.UpdateProfile:input_type -> profile.UpdateProfileRequest
	5, // 6: profile.ProfileService.DeleteProfile:input_type -> profile.DeleteProfileRequest
	6, // 7: profile.ProfileService.ListProfiles:input_type -> profile.ListProfilesRequest
	7, // 8: profile.ProfileService.CreateProfile:output_type -> profile.ProfileResponse
	7, // 9: profile.ProfileService.GetProfile:output_type -> profile.ProfileResponse
	7, // 10: profile.ProfileService.GetProfileByUserID:output_type -> profile.ProfileResponse
	7, // 11: profile.ProfileService.UpdateProfile:output_type -> profile.ProfileResponse
	9, // 12: profile.ProfileService.DeleteProfile:output_type -> google.protobuf.Empty
	8, // 13: profile.ProfileService.ListProfiles:output_type -> profile.ListProfilesResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_profile_profile_proto_init() }
func file_profile_profile_proto_init() {
	if File_profile_profile_proto != nil {
		return
	}
	file_profile_profile_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_profile_profile_proto_rawDesc), len(file_profile_profile_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_profile_proto_goTypes,
		DependencyIndexes: file_profile_profile_proto_depIdxs,
		MessageInfos:      file_profile_profile_proto_msgTypes,
	}.Build()
	File_profile_profile_proto = out.File
	file_profile_profile_proto_goTypes = nil
	file_profile_profile_proto_depIdxs = nil
}

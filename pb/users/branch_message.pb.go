// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: users/branch_message.proto

package users

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	RegionId  string `protobuf:"bytes,3,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Code      string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	Address   string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	City      string `protobuf:"bytes,7,opt,name=city,proto3" json:"city,omitempty"`
	Province  string `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Npwp      string `protobuf:"bytes,9,opt,name=npwp,proto3" json:"npwp,omitempty"`
	Phone     string `protobuf:"bytes,10,opt,name=phone,proto3" json:"phone,omitempty"`
	Pic       string `protobuf:"bytes,11,opt,name=pic,proto3" json:"pic,omitempty"`
	PicPhone  string `protobuf:"bytes,12,opt,name=pic_phone,json=picPhone,proto3" json:"pic_phone,omitempty"`
	CreatedAt string `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy string `protobuf:"bytes,14,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt string `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy string `protobuf:"bytes,16,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *Branch) Reset() {
	*x = Branch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_branch_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_users_branch_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch.ProtoReflect.Descriptor instead.
func (*Branch) Descriptor() ([]byte, []int) {
	return file_users_branch_message_proto_rawDescGZIP(), []int{0}
}

func (x *Branch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Branch) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *Branch) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *Branch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Branch) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Branch) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Branch) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Branch) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Branch) GetNpwp() string {
	if x != nil {
		return x.Npwp
	}
	return ""
}

func (x *Branch) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Branch) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *Branch) GetPicPhone() string {
	if x != nil {
		return x.PicPhone
	}
	return ""
}

func (x *Branch) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Branch) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Branch) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Branch) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

var File_users_branch_message_proto protoreflect.FileDescriptor

var file_users_branch_message_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x69,
	0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x9b, 0x03, 0x0a,
	0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x70, 0x77, 0x70, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x70, 0x77, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x69, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69,
	0x63, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x69, 0x63, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x69, 0x63, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x42, 0x35, 0x0a, 0x21, 0x63, 0x6f,
	0x6d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72, 0x70, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x50,
	0x01, 0x5a, 0x0e, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_branch_message_proto_rawDescOnce sync.Once
	file_users_branch_message_proto_rawDescData = file_users_branch_message_proto_rawDesc
)

func file_users_branch_message_proto_rawDescGZIP() []byte {
	file_users_branch_message_proto_rawDescOnce.Do(func() {
		file_users_branch_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_branch_message_proto_rawDescData)
	})
	return file_users_branch_message_proto_rawDescData
}

var file_users_branch_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_users_branch_message_proto_goTypes = []interface{}{
	(*Branch)(nil), // 0: wiradata.users.Branch
}
var file_users_branch_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_users_branch_message_proto_init() }
func file_users_branch_message_proto_init() {
	if File_users_branch_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_branch_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch); i {
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
			RawDescriptor: file_users_branch_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_users_branch_message_proto_goTypes,
		DependencyIndexes: file_users_branch_message_proto_depIdxs,
		MessageInfos:      file_users_branch_message_proto_msgTypes,
	}.Build()
	File_users_branch_message_proto = out.File
	file_users_branch_message_proto_rawDesc = nil
	file_users_branch_message_proto_goTypes = nil
	file_users_branch_message_proto_depIdxs = nil
}
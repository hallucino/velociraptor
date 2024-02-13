// Code generated by protoc-gen-go. DO NOT EDIT.
// source: secrets.proto

package proto

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

type SecretDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeName    string   `protobuf:"bytes,1,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	Verifier    string   `protobuf:"bytes,2,opt,name=verifier,proto3" json:"verifier,omitempty"`
	SecretNames []string `protobuf:"bytes,3,rep,name=secret_names,json=secretNames,proto3" json:"secret_names,omitempty"`
}

func (x *SecretDefinition) Reset() {
	*x = SecretDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretDefinition) ProtoMessage() {}

func (x *SecretDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretDefinition.ProtoReflect.Descriptor instead.
func (*SecretDefinition) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{0}
}

func (x *SecretDefinition) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *SecretDefinition) GetVerifier() string {
	if x != nil {
		return x.Verifier
	}
	return ""
}

func (x *SecretDefinition) GetSecretNames() []string {
	if x != nil {
		return x.SecretNames
	}
	return nil
}

type SecretDefinitionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*SecretDefinition `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SecretDefinitionList) Reset() {
	*x = SecretDefinitionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretDefinitionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretDefinitionList) ProtoMessage() {}

func (x *SecretDefinitionList) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretDefinitionList.ProtoReflect.Descriptor instead.
func (*SecretDefinitionList) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{1}
}

func (x *SecretDefinitionList) GetItems() []*SecretDefinition {
	if x != nil {
		return x.Items
	}
	return nil
}

type Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TypeName string            `protobuf:"bytes,2,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	Secret   map[string]string `protobuf:"bytes,3,rep,name=secret,proto3" json:"secret,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Users    []string          `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Secret) Reset() {
	*x = Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{2}
}

func (x *Secret) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Secret) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Secret) GetSecret() map[string]string {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *Secret) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

type ModifySecretRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeName string `protobuf:"bytes,1,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// If set the secret will be deleted.
	Delete      bool     `protobuf:"varint,3,opt,name=delete,proto3" json:"delete,omitempty"`
	AddUsers    []string `protobuf:"bytes,4,rep,name=add_users,json=addUsers,proto3" json:"add_users,omitempty"`
	RemoveUsers []string `protobuf:"bytes,5,rep,name=remove_users,json=removeUsers,proto3" json:"remove_users,omitempty"`
}

func (x *ModifySecretRequest) Reset() {
	*x = ModifySecretRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secrets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySecretRequest) ProtoMessage() {}

func (x *ModifySecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secrets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySecretRequest.ProtoReflect.Descriptor instead.
func (*ModifySecretRequest) Descriptor() ([]byte, []int) {
	return file_secrets_proto_rawDescGZIP(), []int{3}
}

func (x *ModifySecretRequest) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *ModifySecretRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModifySecretRequest) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

func (x *ModifySecretRequest) GetAddUsers() []string {
	if x != nil {
		return x.AddUsers
	}
	return nil
}

func (x *ModifySecretRequest) GetRemoveUsers() []string {
	if x != nil {
		return x.RemoveUsers
	}
	return nil
}

var File_secrets_proto protoreflect.FileDescriptor

var file_secrets_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x10, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x14, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xbd, 0x01,
	0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9e, 0x01,
	0x0a, 0x13, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x31,
	0x5a, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x65, 0x6c, 0x6f, 0x63,
	0x69, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_secrets_proto_rawDescOnce sync.Once
	file_secrets_proto_rawDescData = file_secrets_proto_rawDesc
)

func file_secrets_proto_rawDescGZIP() []byte {
	file_secrets_proto_rawDescOnce.Do(func() {
		file_secrets_proto_rawDescData = protoimpl.X.CompressGZIP(file_secrets_proto_rawDescData)
	})
	return file_secrets_proto_rawDescData
}

var file_secrets_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_secrets_proto_goTypes = []interface{}{
	(*SecretDefinition)(nil),     // 0: proto.SecretDefinition
	(*SecretDefinitionList)(nil), // 1: proto.SecretDefinitionList
	(*Secret)(nil),               // 2: proto.Secret
	(*ModifySecretRequest)(nil),  // 3: proto.ModifySecretRequest
	nil,                          // 4: proto.Secret.SecretEntry
}
var file_secrets_proto_depIdxs = []int32{
	0, // 0: proto.SecretDefinitionList.items:type_name -> proto.SecretDefinition
	4, // 1: proto.Secret.secret:type_name -> proto.Secret.SecretEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_secrets_proto_init() }
func file_secrets_proto_init() {
	if File_secrets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_secrets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretDefinition); i {
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
		file_secrets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretDefinitionList); i {
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
		file_secrets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secret); i {
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
		file_secrets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifySecretRequest); i {
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
			RawDescriptor: file_secrets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_secrets_proto_goTypes,
		DependencyIndexes: file_secrets_proto_depIdxs,
		MessageInfos:      file_secrets_proto_msgTypes,
	}.Build()
	File_secrets_proto = out.File
	file_secrets_proto_rawDesc = nil
	file_secrets_proto_goTypes = nil
	file_secrets_proto_depIdxs = nil
}

// resourcedefinition contains the
// schema of the resource definition.
// regenerate with
//
// protoc ./aepc/schema/resourcedefinition.proto --go_opt paths=source_relative\
//   --go_out=.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: schema/resourcedefinition.proto

package schema

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

type Type int32

const (
	Type_UNSPECIFIED Type = 0
	Type_STRING      Type = 1
	Type_INT32       Type = 2
	Type_INT64       Type = 3
	Type_DOUBLE      Type = 4
	Type_FLOAT       Type = 5
	Type_BOOLEAN     Type = 6
	Type_OBJECT      Type = 7
	Type_ARRAY       Type = 8
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "STRING",
		2: "INT32",
		3: "INT64",
		4: "DOUBLE",
		5: "FLOAT",
		6: "BOOLEAN",
		7: "OBJECT",
		8: "ARRAY",
	}
	Type_value = map[string]int32{
		"UNSPECIFIED": 0,
		"STRING":      1,
		"INT32":       2,
		"INT64":       3,
		"DOUBLE":      4,
		"FLOAT":       5,
		"BOOLEAN":     6,
		"OBJECT":      7,
		"ARRAY":       8,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_resourcedefinition_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_schema_resourcedefinition_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{0}
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Resources []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	// These are additional objects (list of properties) that may be referenced from another property.
	Objects []*Object `protobuf:"bytes,3,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Service) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the resource. Used to programmatically
	// refer to and identify the resource.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// The plural of the resource. Used in documentation.
	Plural string `protobuf:"bytes,2,opt,name=plural,proto3" json:"plural,omitempty"`
	// The list of parent resources, referred to via the kind.
	Parents []string `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents,omitempty"`
	// Properties
	Properties map[string]*Property `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// method support
	Methods *Methods `protobuf:"bytes,5,opt,name=methods,proto3" json:"methods,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{1}
}

func (x *Resource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Resource) GetPlural() string {
	if x != nil {
		return x.Plural
	}
	return ""
}

func (x *Resource) GetParents() []string {
	if x != nil {
		return x.Parents
	}
	return nil
}

func (x *Resource) GetProperties() map[string]*Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Resource) GetMethods() *Methods {
	if x != nil {
		return x.Methods
	}
	return nil
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind       string               `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Properties map[string]*Property `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{2}
}

func (x *Object) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Object) GetProperties() map[string]*Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

type Methods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Create     *Methods_CreateMethod     `protobuf:"bytes,1,opt,name=create,proto3" json:"create,omitempty"`
	Read       *Methods_ReadMethod       `protobuf:"bytes,2,opt,name=read,proto3" json:"read,omitempty"`
	Update     *Methods_UpdateMethod     `protobuf:"bytes,3,opt,name=update,proto3" json:"update,omitempty"`
	Delete     *Methods_DeleteMethod     `protobuf:"bytes,4,opt,name=delete,proto3" json:"delete,omitempty"`
	List       *Methods_ListMethod       `protobuf:"bytes,5,opt,name=list,proto3" json:"list,omitempty"`
	GlobalList *Methods_GlobalListMethod `protobuf:"bytes,6,opt,name=global_list,json=globalList,proto3" json:"global_list,omitempty"`
	Apply      *Methods_ApplyMethod      `protobuf:"bytes,7,opt,name=apply,proto3" json:"apply,omitempty"`
}

func (x *Methods) Reset() {
	*x = Methods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Methods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Methods) ProtoMessage() {}

func (x *Methods) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Methods.ProtoReflect.Descriptor instead.
func (*Methods) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{3}
}

func (x *Methods) GetCreate() *Methods_CreateMethod {
	if x != nil {
		return x.Create
	}
	return nil
}

func (x *Methods) GetRead() *Methods_ReadMethod {
	if x != nil {
		return x.Read
	}
	return nil
}

func (x *Methods) GetUpdate() *Methods_UpdateMethod {
	if x != nil {
		return x.Update
	}
	return nil
}

func (x *Methods) GetDelete() *Methods_DeleteMethod {
	if x != nil {
		return x.Delete
	}
	return nil
}

func (x *Methods) GetList() *Methods_ListMethod {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *Methods) GetGlobalList() *Methods_GlobalListMethod {
	if x != nil {
		return x.GlobalList
	}
	return nil
}

func (x *Methods) GetApply() *Methods_ApplyMethod {
	if x != nil {
		return x.Apply
	}
	return nil
}

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=Type" json:"type,omitempty"`
	// field number used for protobuf or other systems where fields must
	// be explicitly enumerated.
	Number   int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	ReadOnly bool  `protobuf:"varint,3,opt,name=readOnly,proto3" json:"readOnly,omitempty"`
	Required bool  `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
	// If type is OBJECT, this is the name of the object in `messages`.
	ObjectType string `protobuf:"bytes,5,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	// Types that are assignable to ArrayType:
	//
	//	*Property_ArrayObjectType
	//	*Property_ArrayPrimitiveType
	ArrayType isProperty_ArrayType `protobuf_oneof:"array_type"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{4}
}

func (x *Property) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_UNSPECIFIED
}

func (x *Property) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Property) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

func (x *Property) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *Property) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (m *Property) GetArrayType() isProperty_ArrayType {
	if m != nil {
		return m.ArrayType
	}
	return nil
}

func (x *Property) GetArrayObjectType() string {
	if x, ok := x.GetArrayType().(*Property_ArrayObjectType); ok {
		return x.ArrayObjectType
	}
	return ""
}

func (x *Property) GetArrayPrimitiveType() Type {
	if x, ok := x.GetArrayType().(*Property_ArrayPrimitiveType); ok {
		return x.ArrayPrimitiveType
	}
	return Type_UNSPECIFIED
}

type isProperty_ArrayType interface {
	isProperty_ArrayType()
}

type Property_ArrayObjectType struct {
	ArrayObjectType string `protobuf:"bytes,6,opt,name=array_object_type,json=arrayObjectType,proto3,oneof"`
}

type Property_ArrayPrimitiveType struct {
	ArrayPrimitiveType Type `protobuf:"varint,7,opt,name=array_primitive_type,json=arrayPrimitiveType,proto3,enum=Type,oneof"`
}

func (*Property_ArrayObjectType) isProperty_ArrayType() {}

func (*Property_ArrayPrimitiveType) isProperty_ArrayType() {}

type Methods_CreateMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Methods_CreateMethod) Reset() {
	*x = Methods_CreateMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Methods_CreateMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Methods_CreateMethod) ProtoMessage() {}

func (x *Methods_CreateMethod) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Methods_CreateMethod.ProtoReflect.Descriptor instead.
func (*Methods_CreateMethod) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{3, 0}
}

type Methods_ReadMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Methods_ReadMethod) Reset() {
	*x = Methods_ReadMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Methods_ReadMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Methods_ReadMethod) ProtoMessage() {}

func (x *Methods_ReadMethod) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Methods_ReadMethod.ProtoReflect.Descriptor instead.
func (*Methods_ReadMethod) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{3, 1}
}

type Methods_UpdateMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Methods_UpdateMethod) Reset() {
	*x = Methods_UpdateMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Methods_UpdateMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Methods_UpdateMethod) ProtoMessage() {}

func (x *Methods_UpdateMethod) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Methods_UpdateMethod.ProtoReflect.Descriptor instead.
func (*Methods_UpdateMethod) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{3, 2}
}

type Methods_DeleteMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Methods_DeleteMethod) Reset() {
	*x = Methods_DeleteMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Methods_DeleteMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Methods_DeleteMethod) ProtoMessage() {}

func (x *Methods_DeleteMethod) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Methods_DeleteMethod.ProtoReflect.Descriptor instead.
func (*Methods_DeleteMethod) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{3, 3}
}

type Methods_ListMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Methods_ListMethod) Reset() {
	*x = Methods_ListMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Methods_ListMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Methods_ListMethod) ProtoMessage() {}

func (x *Methods_ListMethod) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Methods_ListMethod.ProtoReflect.Descriptor instead.
func (*Methods_ListMethod) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{3, 4}
}

type Methods_GlobalListMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Methods_GlobalListMethod) Reset() {
	*x = Methods_GlobalListMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Methods_GlobalListMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Methods_GlobalListMethod) ProtoMessage() {}

func (x *Methods_GlobalListMethod) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Methods_GlobalListMethod.ProtoReflect.Descriptor instead.
func (*Methods_GlobalListMethod) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{3, 5}
}

type Methods_ApplyMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Methods_ApplyMethod) Reset() {
	*x = Methods_ApplyMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_resourcedefinition_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Methods_ApplyMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Methods_ApplyMethod) ProtoMessage() {}

func (x *Methods_ApplyMethod) ProtoReflect() protoreflect.Message {
	mi := &file_schema_resourcedefinition_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Methods_ApplyMethod.ProtoReflect.Descriptor instead.
func (*Methods_ApplyMethod) Descriptor() ([]byte, []int) {
	return file_schema_resourcedefinition_proto_rawDescGZIP(), []int{3, 6}
}

var File_schema_resourcedefinition_proto protoreflect.FileDescriptor

var file_schema_resourcedefinition_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x69, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x27, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0xf9, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6c, 0x75, 0x72, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6c, 0x75, 0x72, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x07, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0x48,
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9f, 0x01, 0x0a, 0x06, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x1a, 0x48, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbf, 0x03, 0x0a, 0x07, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x2d,
	0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0a, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x05, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x1a, 0x0e, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x0c, 0x0a,
	0x0a, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x0e, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x0e, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x0c, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x12, 0x0a, 0x10, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x0d, 0x0a,
	0x0b, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x8d, 0x02, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x14, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x05, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x12, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x74, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c,
	0x45, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x05, 0x12, 0x0b,
	0x0a, 0x07, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x4f,
	0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x52, 0x52, 0x41, 0x59,
	0x10, 0x08, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_resourcedefinition_proto_rawDescOnce sync.Once
	file_schema_resourcedefinition_proto_rawDescData = file_schema_resourcedefinition_proto_rawDesc
)

func file_schema_resourcedefinition_proto_rawDescGZIP() []byte {
	file_schema_resourcedefinition_proto_rawDescOnce.Do(func() {
		file_schema_resourcedefinition_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_resourcedefinition_proto_rawDescData)
	})
	return file_schema_resourcedefinition_proto_rawDescData
}

var file_schema_resourcedefinition_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_schema_resourcedefinition_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_schema_resourcedefinition_proto_goTypes = []interface{}{
	(Type)(0),                        // 0: Type
	(*Service)(nil),                  // 1: Service
	(*Resource)(nil),                 // 2: Resource
	(*Object)(nil),                   // 3: Object
	(*Methods)(nil),                  // 4: Methods
	(*Property)(nil),                 // 5: Property
	nil,                              // 6: Resource.PropertiesEntry
	nil,                              // 7: Object.PropertiesEntry
	(*Methods_CreateMethod)(nil),     // 8: Methods.CreateMethod
	(*Methods_ReadMethod)(nil),       // 9: Methods.ReadMethod
	(*Methods_UpdateMethod)(nil),     // 10: Methods.UpdateMethod
	(*Methods_DeleteMethod)(nil),     // 11: Methods.DeleteMethod
	(*Methods_ListMethod)(nil),       // 12: Methods.ListMethod
	(*Methods_GlobalListMethod)(nil), // 13: Methods.GlobalListMethod
	(*Methods_ApplyMethod)(nil),      // 14: Methods.ApplyMethod
}
var file_schema_resourcedefinition_proto_depIdxs = []int32{
	2,  // 0: Service.resources:type_name -> Resource
	3,  // 1: Service.objects:type_name -> Object
	6,  // 2: Resource.properties:type_name -> Resource.PropertiesEntry
	4,  // 3: Resource.methods:type_name -> Methods
	7,  // 4: Object.properties:type_name -> Object.PropertiesEntry
	8,  // 5: Methods.create:type_name -> Methods.CreateMethod
	9,  // 6: Methods.read:type_name -> Methods.ReadMethod
	10, // 7: Methods.update:type_name -> Methods.UpdateMethod
	11, // 8: Methods.delete:type_name -> Methods.DeleteMethod
	12, // 9: Methods.list:type_name -> Methods.ListMethod
	13, // 10: Methods.global_list:type_name -> Methods.GlobalListMethod
	14, // 11: Methods.apply:type_name -> Methods.ApplyMethod
	0,  // 12: Property.type:type_name -> Type
	0,  // 13: Property.array_primitive_type:type_name -> Type
	5,  // 14: Resource.PropertiesEntry.value:type_name -> Property
	5,  // 15: Object.PropertiesEntry.value:type_name -> Property
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_schema_resourcedefinition_proto_init() }
func file_schema_resourcedefinition_proto_init() {
	if File_schema_resourcedefinition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_resourcedefinition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_schema_resourcedefinition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_schema_resourcedefinition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_schema_resourcedefinition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Methods); i {
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
		file_schema_resourcedefinition_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_schema_resourcedefinition_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Methods_CreateMethod); i {
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
		file_schema_resourcedefinition_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Methods_ReadMethod); i {
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
		file_schema_resourcedefinition_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Methods_UpdateMethod); i {
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
		file_schema_resourcedefinition_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Methods_DeleteMethod); i {
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
		file_schema_resourcedefinition_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Methods_ListMethod); i {
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
		file_schema_resourcedefinition_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Methods_GlobalListMethod); i {
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
		file_schema_resourcedefinition_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Methods_ApplyMethod); i {
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
	file_schema_resourcedefinition_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Property_ArrayObjectType)(nil),
		(*Property_ArrayPrimitiveType)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_resourcedefinition_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_resourcedefinition_proto_goTypes,
		DependencyIndexes: file_schema_resourcedefinition_proto_depIdxs,
		EnumInfos:         file_schema_resourcedefinition_proto_enumTypes,
		MessageInfos:      file_schema_resourcedefinition_proto_msgTypes,
	}.Build()
	File_schema_resourcedefinition_proto = out.File
	file_schema_resourcedefinition_proto_rawDesc = nil
	file_schema_resourcedefinition_proto_goTypes = nil
	file_schema_resourcedefinition_proto_depIdxs = nil
}

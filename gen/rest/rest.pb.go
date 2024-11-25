// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/restaurants/rest.proto

package rest

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

type RestId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RestId) Reset() {
	*x = RestId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurants_rest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestId) ProtoMessage() {}

func (x *RestId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurants_rest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestId.ProtoReflect.Descriptor instead.
func (*RestId) Descriptor() ([]byte, []int) {
	return file_proto_restaurants_rest_proto_rawDescGZIP(), []int{0}
}

func (x *RestId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurants_rest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurants_rest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_proto_restaurants_rest_proto_rawDescGZIP(), []int{1}
}

func (x *Category) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurants_rest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurants_rest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_proto_restaurants_rest_proto_rawDescGZIP(), []int{2}
}

func (x *Id) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Limit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Limit) Reset() {
	*x = Limit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurants_rest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Limit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Limit) ProtoMessage() {}

func (x *Limit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurants_rest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Limit.ProtoReflect.Descriptor instead.
func (*Limit) Descriptor() ([]byte, []int) {
	return file_proto_restaurants_rest_proto_rawDescGZIP(), []int{3}
}

func (x *Limit) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurants_rest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurants_rest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_restaurants_rest_proto_rawDescGZIP(), []int{4}
}

type Rest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShortDescription string  `protobuf:"bytes,3,opt,name=short_description,json=shortDescription,proto3" json:"short_description,omitempty"`
	LongDescription  string  `protobuf:"bytes,4,opt,name=long_description,json=longDescription,proto3" json:"long_description,omitempty"`
	Address          string  `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	ImgUrl           string  `protobuf:"bytes,6,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
	Rating           float32 `protobuf:"fixed32,7,opt,name=rating,proto3" json:"rating,omitempty"`
	CommentCount     uint32  `protobuf:"varint,8,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
}

func (x *Rest) Reset() {
	*x = Rest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurants_rest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rest) ProtoMessage() {}

func (x *Rest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurants_rest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rest.ProtoReflect.Descriptor instead.
func (*Rest) Descriptor() ([]byte, []int) {
	return file_proto_restaurants_rest_proto_rawDescGZIP(), []int{5}
}

func (x *Rest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Rest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rest) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *Rest) GetLongDescription() string {
	if x != nil {
		return x.LongDescription
	}
	return ""
}

func (x *Rest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Rest) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

func (x *Rest) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Rest) GetCommentCount() uint32 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

type RestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rest []*Rest `protobuf:"bytes,1,rep,name=rest,proto3" json:"rest,omitempty"`
}

func (x *RestList) Reset() {
	*x = RestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurants_rest_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestList) ProtoMessage() {}

func (x *RestList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurants_rest_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestList.ProtoReflect.Descriptor instead.
func (*RestList) Descriptor() ([]byte, []int) {
	return file_proto_restaurants_rest_proto_rawDescGZIP(), []int{6}
}

func (x *RestList) GetRest() []*Rest {
	if x != nil {
		return x.Rest
	}
	return nil
}

type CategoryList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C []*Category `protobuf:"bytes,1,rep,name=c,proto3" json:"c,omitempty"`
}

func (x *CategoryList) Reset() {
	*x = CategoryList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurants_rest_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryList) ProtoMessage() {}

func (x *CategoryList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurants_rest_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryList.ProtoReflect.Descriptor instead.
func (*CategoryList) Descriptor() ([]byte, []int) {
	return file_proto_restaurants_rest_proto_rawDescGZIP(), []int{7}
}

func (x *CategoryList) GetC() []*Category {
	if x != nil {
		return x.C
	}
	return nil
}

type UserAndLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *UserAndLimit) Reset() {
	*x = UserAndLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurants_rest_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAndLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAndLimit) ProtoMessage() {}

func (x *UserAndLimit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurants_rest_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAndLimit.ProtoReflect.Descriptor instead.
func (*UserAndLimit) Descriptor() ([]byte, []int) {
	return file_proto_restaurants_rest_proto_rawDescGZIP(), []int{8}
}

func (x *UserAndLimit) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserAndLimit) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type RestIdList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I []*Id `protobuf:"bytes,1,rep,name=i,proto3" json:"i,omitempty"`
}

func (x *RestIdList) Reset() {
	*x = RestIdList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_restaurants_rest_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestIdList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestIdList) ProtoMessage() {}

func (x *RestIdList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurants_rest_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestIdList.ProtoReflect.Descriptor instead.
func (*RestIdList) Descriptor() ([]byte, []int) {
	return file_proto_restaurants_rest_proto_rawDescGZIP(), []int{9}
}

func (x *RestIdList) GetI() []*Id {
	if x != nil {
		return x.I
	}
	return nil
}

var File_proto_restaurants_rest_proto protoreflect.FileDescriptor

var file_proto_restaurants_rest_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x72, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14,
	0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xf1, 0x01, 0x0a,
	0x04, 0x52, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6d, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x67,
	0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x2a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04,
	0x72, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x65, 0x73,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x52, 0x04, 0x72, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0c,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x01,
	0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x01, 0x63, 0x22, 0x3d, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x24, 0x0a, 0x0a, 0x52, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x01, 0x69, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x52, 0x01, 0x69, 0x32,
	0xf7, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x27,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0b, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x0c, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x1a, 0x0a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x29,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x08, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x72,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_restaurants_rest_proto_rawDescOnce sync.Once
	file_proto_restaurants_rest_proto_rawDescData = file_proto_restaurants_rest_proto_rawDesc
)

func file_proto_restaurants_rest_proto_rawDescGZIP() []byte {
	file_proto_restaurants_rest_proto_rawDescOnce.Do(func() {
		file_proto_restaurants_rest_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_restaurants_rest_proto_rawDescData)
	})
	return file_proto_restaurants_rest_proto_rawDescData
}

var file_proto_restaurants_rest_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_restaurants_rest_proto_goTypes = []any{
	(*RestId)(nil),       // 0: rest.RestId
	(*Category)(nil),     // 1: rest.Category
	(*Id)(nil),           // 2: rest.Id
	(*Limit)(nil),        // 3: rest.Limit
	(*Empty)(nil),        // 4: rest.Empty
	(*Rest)(nil),         // 5: rest.Rest
	(*RestList)(nil),     // 6: rest.RestList
	(*CategoryList)(nil), // 7: rest.CategoryList
	(*UserAndLimit)(nil), // 8: rest.UserAndLimit
	(*RestIdList)(nil),   // 9: rest.RestIdList
}
var file_proto_restaurants_rest_proto_depIdxs = []int32{
	5, // 0: rest.RestList.rest:type_name -> rest.Rest
	1, // 1: rest.CategoryList.c:type_name -> rest.Category
	2, // 2: rest.RestIdList.i:type_name -> rest.Id
	4, // 3: rest.RestWorker.GetAll:input_type -> rest.Empty
	0, // 4: rest.RestWorker.GetById:input_type -> rest.RestId
	2, // 5: rest.RestWorker.GetByFilter:input_type -> rest.Id
	4, // 6: rest.RestWorker.GetCategoryList:input_type -> rest.Empty
	8, // 7: rest.RestWorker.GetRecomendation:input_type -> rest.UserAndLimit
	6, // 8: rest.RestWorker.GetAll:output_type -> rest.RestList
	5, // 9: rest.RestWorker.GetById:output_type -> rest.Rest
	6, // 10: rest.RestWorker.GetByFilter:output_type -> rest.RestList
	7, // 11: rest.RestWorker.GetCategoryList:output_type -> rest.CategoryList
	6, // 12: rest.RestWorker.GetRecomendation:output_type -> rest.RestList
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_restaurants_rest_proto_init() }
func file_proto_restaurants_rest_proto_init() {
	if File_proto_restaurants_rest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_restaurants_rest_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RestId); i {
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
		file_proto_restaurants_rest_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Category); i {
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
		file_proto_restaurants_rest_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Id); i {
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
		file_proto_restaurants_rest_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Limit); i {
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
		file_proto_restaurants_rest_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_proto_restaurants_rest_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Rest); i {
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
		file_proto_restaurants_rest_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RestList); i {
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
		file_proto_restaurants_rest_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryList); i {
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
		file_proto_restaurants_rest_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UserAndLimit); i {
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
		file_proto_restaurants_rest_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*RestIdList); i {
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
			RawDescriptor: file_proto_restaurants_rest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_restaurants_rest_proto_goTypes,
		DependencyIndexes: file_proto_restaurants_rest_proto_depIdxs,
		MessageInfos:      file_proto_restaurants_rest_proto_msgTypes,
	}.Build()
	File_proto_restaurants_rest_proto = out.File
	file_proto_restaurants_rest_proto_rawDesc = nil
	file_proto_restaurants_rest_proto_goTypes = nil
	file_proto_restaurants_rest_proto_depIdxs = nil
}

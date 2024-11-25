// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/comment/comment.proto

package comment

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

type CommentId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CommentId) Reset() {
	*x = CommentId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentId) ProtoMessage() {}

func (x *CommentId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentId.ProtoReflect.Descriptor instead.
func (*CommentId) Descriptor() ([]byte, []int) {
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RestId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RestId) Reset() {
	*x = RestId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestId) ProtoMessage() {}

func (x *RestId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[1]
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
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{1}
}

func (x *RestId) GetId() uint64 {
	if x != nil {
		return x.Id
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
		mi := &file_proto_comment_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[2]
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
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{2}
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Image    string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	RestId   uint64 `protobuf:"varint,5,opt,name=rest_id,json=restId,proto3" json:"rest_id,omitempty"`
	Text     string `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	Rating   uint32 `protobuf:"varint,7,opt,name=rating,proto3" json:"rating,omitempty"`
	OrderId  uint64 `protobuf:"varint,8,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{3}
}

func (x *Comment) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comment) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Comment) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Comment) GetRestId() uint64 {
	if x != nil {
		return x.RestId
	}
	return 0
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Comment) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type CommentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment []*Comment `protobuf:"bytes,1,rep,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentList) Reset() {
	*x = CommentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comment_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentList) ProtoMessage() {}

func (x *CommentList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comment_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentList.ProtoReflect.Descriptor instead.
func (*CommentList) Descriptor() ([]byte, []int) {
	return file_proto_comment_comment_proto_rawDescGZIP(), []int{4}
}

func (x *CommentList) GetComment() []*Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

var File_proto_comment_comment_proto protoreflect.FileDescriptor

var file_proto_comment_comment_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xc5, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xbb, 0x01, 0x0a, 0x0d, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x52, 0x65, 0x73, 0x74, 0x12, 0x0f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a,
	0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_comment_comment_proto_rawDescOnce sync.Once
	file_proto_comment_comment_proto_rawDescData = file_proto_comment_comment_proto_rawDesc
)

func file_proto_comment_comment_proto_rawDescGZIP() []byte {
	file_proto_comment_comment_proto_rawDescOnce.Do(func() {
		file_proto_comment_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_comment_comment_proto_rawDescData)
	})
	return file_proto_comment_comment_proto_rawDescData
}

var file_proto_comment_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_comment_comment_proto_goTypes = []any{
	(*CommentId)(nil),   // 0: comment.CommentId
	(*RestId)(nil),      // 1: comment.RestId
	(*Empty)(nil),       // 2: comment.Empty
	(*Comment)(nil),     // 3: comment.Comment
	(*CommentList)(nil), // 4: comment.CommentList
}
var file_proto_comment_comment_proto_depIdxs = []int32{
	3, // 0: comment.CommentList.comment:type_name -> comment.Comment
	3, // 1: comment.CommentWorker.CreateComment:input_type -> comment.Comment
	0, // 2: comment.CommentWorker.DeleteComment:input_type -> comment.CommentId
	1, // 3: comment.CommentWorker.GetCommentsByRest:input_type -> comment.RestId
	3, // 4: comment.CommentWorker.CreateComment:output_type -> comment.Comment
	2, // 5: comment.CommentWorker.DeleteComment:output_type -> comment.Empty
	4, // 6: comment.CommentWorker.GetCommentsByRest:output_type -> comment.CommentList
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_comment_comment_proto_init() }
func file_proto_comment_comment_proto_init() {
	if File_proto_comment_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_comment_comment_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommentId); i {
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
		file_proto_comment_comment_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_proto_comment_comment_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_proto_comment_comment_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Comment); i {
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
		file_proto_comment_comment_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CommentList); i {
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
			RawDescriptor: file_proto_comment_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_comment_comment_proto_goTypes,
		DependencyIndexes: file_proto_comment_comment_proto_depIdxs,
		MessageInfos:      file_proto_comment_comment_proto_msgTypes,
	}.Build()
	File_proto_comment_comment_proto = out.File
	file_proto_comment_comment_proto_rawDesc = nil
	file_proto_comment_comment_proto_goTypes = nil
	file_proto_comment_comment_proto_depIdxs = nil
}

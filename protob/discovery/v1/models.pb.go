// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: discovery/v1/models.proto

package pbdiscovery

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

type AddPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *AddPost `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *AddPostRequest) Reset() {
	*x = AddPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_v1_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostRequest) ProtoMessage() {}

func (x *AddPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_v1_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostRequest.ProtoReflect.Descriptor instead.
func (*AddPostRequest) Descriptor() ([]byte, []int) {
	return file_discovery_v1_models_proto_rawDescGZIP(), []int{0}
}

func (x *AddPostRequest) GetPost() *AddPost {
	if x != nil {
		return x.Post
	}
	return nil
}

type AddPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Post    *Post `protobuf:"bytes,2,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *AddPostResponse) Reset() {
	*x = AddPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_v1_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPostResponse) ProtoMessage() {}

func (x *AddPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_v1_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPostResponse.ProtoReflect.Descriptor instead.
func (*AddPostResponse) Descriptor() ([]byte, []int) {
	return file_discovery_v1_models_proto_rawDescGZIP(), []int{1}
}

func (x *AddPostResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type GetPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start from this postId (excluded) to retrieve a new page
	Cursor string `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *GetPostsRequest) Reset() {
	*x = GetPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_v1_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsRequest) ProtoMessage() {}

func (x *GetPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_v1_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsRequest.ProtoReflect.Descriptor instead.
func (*GetPostsRequest) Descriptor() ([]byte, []int) {
	return file_discovery_v1_models_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostsRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type GetPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// last postId in the page, used to retrieve next page
	// with GetPostsRequest
	Cursor string  `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Data   []*Post `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetPostsResponse) Reset() {
	*x = GetPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_v1_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsResponse) ProtoMessage() {}

func (x *GetPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_v1_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsResponse.ProtoReflect.Descriptor instead.
func (*GetPostsResponse) Descriptor() ([]byte, []int) {
	return file_discovery_v1_models_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetPostsResponse) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *GetPostsResponse) GetData() []*Post {
	if x != nil {
		return x.Data
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner       string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	FrontPicUrl string `protobuf:"bytes,3,opt,name=front_pic_url,json=frontPicUrl,proto3" json:"front_pic_url,omitempty"`
	BackPicUrl  string `protobuf:"bytes,4,opt,name=back_pic_url,json=backPicUrl,proto3" json:"back_pic_url,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_v1_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_v1_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_discovery_v1_models_proto_rawDescGZIP(), []int{4}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Post) GetFrontPicUrl() string {
	if x != nil {
		return x.FrontPicUrl
	}
	return ""
}

func (x *Post) GetBackPicUrl() string {
	if x != nil {
		return x.BackPicUrl
	}
	return ""
}

type AddPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner       string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	FrontPicUrl string `protobuf:"bytes,3,opt,name=front_pic_url,json=frontPicUrl,proto3" json:"front_pic_url,omitempty"`
	BackPicUrl  string `protobuf:"bytes,4,opt,name=back_pic_url,json=backPicUrl,proto3" json:"back_pic_url,omitempty"`
}

func (x *AddPost) Reset() {
	*x = AddPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_v1_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPost) ProtoMessage() {}

func (x *AddPost) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_v1_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPost.ProtoReflect.Descriptor instead.
func (*AddPost) Descriptor() ([]byte, []int) {
	return file_discovery_v1_models_proto_rawDescGZIP(), []int{5}
}

func (x *AddPost) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AddPost) GetFrontPicUrl() string {
	if x != nil {
		return x.FrontPicUrl
	}
	return ""
}

func (x *AddPost) GetBackPicUrl() string {
	if x != nil {
		return x.BackPicUrl
	}
	return ""
}

var File_discovery_v1_models_proto protoreflect.FileDescriptor

var file_discovery_v1_models_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x3b, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x6c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x72, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x70, 0x69, 0x63, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x70,
	0x69, 0x63, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61,
	0x63, 0x6b, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x22, 0x65, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x5f, 0x70, 0x69, 0x63, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a,
	0x0c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x69, 0x63, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x69, 0x63, 0x55, 0x72, 0x6c, 0x42,
	0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x65,
	0x52, 0x65, 0x61, 0x6c, 0x2d, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x67, 0x6f, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discovery_v1_models_proto_rawDescOnce sync.Once
	file_discovery_v1_models_proto_rawDescData = file_discovery_v1_models_proto_rawDesc
)

func file_discovery_v1_models_proto_rawDescGZIP() []byte {
	file_discovery_v1_models_proto_rawDescOnce.Do(func() {
		file_discovery_v1_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_discovery_v1_models_proto_rawDescData)
	})
	return file_discovery_v1_models_proto_rawDescData
}

var file_discovery_v1_models_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_discovery_v1_models_proto_goTypes = []interface{}{
	(*AddPostRequest)(nil),   // 0: discovery.v1.AddPostRequest
	(*AddPostResponse)(nil),  // 1: discovery.v1.AddPostResponse
	(*GetPostsRequest)(nil),  // 2: discovery.v1.GetPostsRequest
	(*GetPostsResponse)(nil), // 3: discovery.v1.GetPostsResponse
	(*Post)(nil),             // 4: discovery.v1.Post
	(*AddPost)(nil),          // 5: discovery.v1.AddPost
}
var file_discovery_v1_models_proto_depIdxs = []int32{
	5, // 0: discovery.v1.AddPostRequest.post:type_name -> discovery.v1.AddPost
	4, // 1: discovery.v1.AddPostResponse.post:type_name -> discovery.v1.Post
	4, // 2: discovery.v1.GetPostsResponse.data:type_name -> discovery.v1.Post
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_discovery_v1_models_proto_init() }
func file_discovery_v1_models_proto_init() {
	if File_discovery_v1_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discovery_v1_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostRequest); i {
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
		file_discovery_v1_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPostResponse); i {
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
		file_discovery_v1_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsRequest); i {
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
		file_discovery_v1_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsResponse); i {
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
		file_discovery_v1_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_discovery_v1_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPost); i {
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
			RawDescriptor: file_discovery_v1_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_discovery_v1_models_proto_goTypes,
		DependencyIndexes: file_discovery_v1_models_proto_depIdxs,
		MessageInfos:      file_discovery_v1_models_proto_msgTypes,
	}.Build()
	File_discovery_v1_models_proto = out.File
	file_discovery_v1_models_proto_rawDesc = nil
	file_discovery_v1_models_proto_goTypes = nil
	file_discovery_v1_models_proto_depIdxs = nil
}
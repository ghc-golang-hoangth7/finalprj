// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.4
// source: proto/planes.proto

package planes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Plane struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaneId     string `protobuf:"bytes,1,opt,name=plane_id,json=planeId,proto3" json:"plane_id,omitempty"`
	PlaneNumber string `protobuf:"bytes,2,opt,name=plane_number,json=planeNumber,proto3" json:"plane_number,omitempty"`
	TotalSeats  int32  `protobuf:"varint,3,opt,name=total_seats,json=totalSeats,proto3" json:"total_seats,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Plane) Reset() {
	*x = Plane{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_planes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plane) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plane) ProtoMessage() {}

func (x *Plane) ProtoReflect() protoreflect.Message {
	mi := &file_proto_planes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plane.ProtoReflect.Descriptor instead.
func (*Plane) Descriptor() ([]byte, []int) {
	return file_proto_planes_proto_rawDescGZIP(), []int{0}
}

func (x *Plane) GetPlaneId() string {
	if x != nil {
		return x.PlaneId
	}
	return ""
}

func (x *Plane) GetPlaneNumber() string {
	if x != nil {
		return x.PlaneNumber
	}
	return ""
}

func (x *Plane) GetTotalSeats() int32 {
	if x != nil {
		return x.TotalSeats
	}
	return 0
}

func (x *Plane) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PlaneId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PlaneId) Reset() {
	*x = PlaneId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_planes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaneId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaneId) ProtoMessage() {}

func (x *PlaneId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_planes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaneId.ProtoReflect.Descriptor instead.
func (*PlaneId) Descriptor() ([]byte, []int) {
	return file_proto_planes_proto_rawDescGZIP(), []int{1}
}

func (x *PlaneId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PlaneList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planes []*Plane `protobuf:"bytes,1,rep,name=planes,proto3" json:"planes,omitempty"`
}

func (x *PlaneList) Reset() {
	*x = PlaneList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_planes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaneList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaneList) ProtoMessage() {}

func (x *PlaneList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_planes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaneList.ProtoReflect.Descriptor instead.
func (*PlaneList) Descriptor() ([]byte, []int) {
	return file_proto_planes_proto_rawDescGZIP(), []int{2}
}

func (x *PlaneList) GetPlanes() []*Plane {
	if x != nil {
		return x.Planes
	}
	return nil
}

type PlaneStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaneId string `protobuf:"bytes,1,opt,name=plane_id,json=planeId,proto3" json:"plane_id,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PlaneStatusRequest) Reset() {
	*x = PlaneStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_planes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaneStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaneStatusRequest) ProtoMessage() {}

func (x *PlaneStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_planes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaneStatusRequest.ProtoReflect.Descriptor instead.
func (*PlaneStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_planes_proto_rawDescGZIP(), []int{3}
}

func (x *PlaneStatusRequest) GetPlaneId() string {
	if x != nil {
		return x.PlaneId
	}
	return ""
}

func (x *PlaneStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_planes_proto protoreflect.FileDescriptor

var file_proto_planes_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x05, 0x50, 0x6c, 0x61,
	0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x61, 0x74,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x50, 0x6c, 0x61,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0xea, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x1a, 0x0f, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x31, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x65, 0x12, 0x47, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0b,
	0x5a, 0x09, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_planes_proto_rawDescOnce sync.Once
	file_proto_planes_proto_rawDescData = file_proto_planes_proto_rawDesc
)

func file_proto_planes_proto_rawDescGZIP() []byte {
	file_proto_planes_proto_rawDescOnce.Do(func() {
		file_proto_planes_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_planes_proto_rawDescData)
	})
	return file_proto_planes_proto_rawDescData
}

var file_proto_planes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_planes_proto_goTypes = []interface{}{
	(*Plane)(nil),              // 0: planes.Plane
	(*PlaneId)(nil),            // 1: planes.PlaneId
	(*PlaneList)(nil),          // 2: planes.PlaneList
	(*PlaneStatusRequest)(nil), // 3: planes.PlaneStatusRequest
	(*emptypb.Empty)(nil),      // 4: google.protobuf.Empty
}
var file_proto_planes_proto_depIdxs = []int32{
	0, // 0: planes.PlaneList.planes:type_name -> planes.Plane
	0, // 1: planes.PlanesService.UpsertPlane:input_type -> planes.Plane
	0, // 2: planes.PlanesService.GetPlanesList:input_type -> planes.Plane
	1, // 3: planes.PlanesService.GetPlaneById:input_type -> planes.PlaneId
	3, // 4: planes.PlanesService.ChangePlaneStatus:input_type -> planes.PlaneStatusRequest
	1, // 5: planes.PlanesService.UpsertPlane:output_type -> planes.PlaneId
	2, // 6: planes.PlanesService.GetPlanesList:output_type -> planes.PlaneList
	0, // 7: planes.PlanesService.GetPlaneById:output_type -> planes.Plane
	4, // 8: planes.PlanesService.ChangePlaneStatus:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_planes_proto_init() }
func file_proto_planes_proto_init() {
	if File_proto_planes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_planes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plane); i {
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
		file_proto_planes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaneId); i {
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
		file_proto_planes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaneList); i {
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
		file_proto_planes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaneStatusRequest); i {
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
			RawDescriptor: file_proto_planes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_planes_proto_goTypes,
		DependencyIndexes: file_proto_planes_proto_depIdxs,
		MessageInfos:      file_proto_planes_proto_msgTypes,
	}.Build()
	File_proto_planes_proto = out.File
	file_proto_planes_proto_rawDesc = nil
	file_proto_planes_proto_goTypes = nil
	file_proto_planes_proto_depIdxs = nil
}

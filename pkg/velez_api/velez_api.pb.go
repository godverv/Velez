// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: grpc/velez_api.proto

package velez_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PortBindingsProtocol int32

const (
	PortBindings_tcp PortBindingsProtocol = 0
	PortBindings_udp PortBindingsProtocol = 1
)

// Enum value maps for PortBindingsProtocol.
var (
	PortBindingsProtocol_name = map[int32]string{
		0: "tcp",
		1: "udp",
	}
	PortBindingsProtocol_value = map[string]int32{
		"tcp": 0,
		"udp": 1,
	}
)

func (x PortBindingsProtocol) Enum() *PortBindingsProtocol {
	p := new(PortBindingsProtocol)
	*p = x
	return p
}

func (x PortBindingsProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PortBindingsProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_velez_api_proto_enumTypes[0].Descriptor()
}

func (PortBindingsProtocol) Type() protoreflect.EnumType {
	return &file_grpc_velez_api_proto_enumTypes[0]
}

func (x PortBindingsProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PortBindingsProtocol.Descriptor instead.
func (PortBindingsProtocol) EnumDescriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{1, 0}
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{0}
}

type PortBindings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host      uint32               `protobuf:"varint,1,opt,name=host,proto3" json:"host,omitempty"`
	Container uint32               `protobuf:"varint,2,opt,name=container,proto3" json:"container,omitempty"`
	Protoc    PortBindingsProtocol `protobuf:"varint,3,opt,name=protoc,proto3,enum=velez_api.PortBindingsProtocol" json:"protoc,omitempty"`
}

func (x *PortBindings) Reset() {
	*x = PortBindings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortBindings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortBindings) ProtoMessage() {}

func (x *PortBindings) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortBindings.ProtoReflect.Descriptor instead.
func (*PortBindings) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{1}
}

func (x *PortBindings) GetHost() uint32 {
	if x != nil {
		return x.Host
	}
	return 0
}

func (x *PortBindings) GetContainer() uint32 {
	if x != nil {
		return x.Container
	}
	return 0
}

func (x *PortBindings) GetProtoc() PortBindingsProtocol {
	if x != nil {
		return x.Protoc
	}
	return PortBindings_tcp
}

type VolumeBindings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host      string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
}

func (x *VolumeBindings) Reset() {
	*x = VolumeBindings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VolumeBindings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VolumeBindings) ProtoMessage() {}

func (x *VolumeBindings) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VolumeBindings.ProtoReflect.Descriptor instead.
func (*VolumeBindings) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{2}
}

func (x *VolumeBindings) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *VolumeBindings) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tags []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{3}
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Smerd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string            `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name      string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageName string            `protobuf:"bytes,3,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	Ports     []*PortBindings   `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	Volumes   []*VolumeBindings `protobuf:"bytes,5,rep,name=volumes,proto3" json:"volumes,omitempty"`
}

func (x *Smerd) Reset() {
	*x = Smerd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Smerd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Smerd) ProtoMessage() {}

func (x *Smerd) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Smerd.ProtoReflect.Descriptor instead.
func (*Smerd) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{4}
}

func (x *Smerd) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Smerd) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Smerd) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *Smerd) GetPorts() []*PortBindings {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Smerd) GetVolumes() []*VolumeBindings {
	if x != nil {
		return x.Volumes
	}
	return nil
}

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{5}
}

type CreateSmerd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSmerd) Reset() {
	*x = CreateSmerd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSmerd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSmerd) ProtoMessage() {}

func (x *CreateSmerd) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSmerd.ProtoReflect.Descriptor instead.
func (*CreateSmerd) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{6}
}

type ListSmerds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSmerds) Reset() {
	*x = ListSmerds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSmerds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSmerds) ProtoMessage() {}

func (x *ListSmerds) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSmerds.ProtoReflect.Descriptor instead.
func (*ListSmerds) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{7}
}

type Version_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Version_Request) Reset() {
	*x = Version_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version_Request) ProtoMessage() {}

func (x *Version_Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version_Request.ProtoReflect.Descriptor instead.
func (*Version_Request) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{0, 0}
}

type Version_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Version_Response) Reset() {
	*x = Version_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version_Response) ProtoMessage() {}

func (x *Version_Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version_Response.ProtoReflect.Descriptor instead.
func (*Version_Response) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Version_Response) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Container_Hardware struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuAmount    *float32 `protobuf:"fixed32,1,opt,name=cpu_amount,json=cpuAmount,proto3,oneof" json:"cpu_amount,omitempty"`
	RamMb        *int32   `protobuf:"varint,2,opt,name=ram_mb,json=ramMb,proto3,oneof" json:"ram_mb,omitempty"`
	MemorySwapMb *int32   `protobuf:"varint,3,opt,name=memory_swap_mb,json=memorySwapMb,proto3,oneof" json:"memory_swap_mb,omitempty"`
}

func (x *Container_Hardware) Reset() {
	*x = Container_Hardware{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container_Hardware) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container_Hardware) ProtoMessage() {}

func (x *Container_Hardware) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container_Hardware.ProtoReflect.Descriptor instead.
func (*Container_Hardware) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{5, 0}
}

func (x *Container_Hardware) GetCpuAmount() float32 {
	if x != nil && x.CpuAmount != nil {
		return *x.CpuAmount
	}
	return 0
}

func (x *Container_Hardware) GetRamMb() int32 {
	if x != nil && x.RamMb != nil {
		return *x.RamMb
	}
	return 0
}

func (x *Container_Hardware) GetMemorySwapMb() int32 {
	if x != nil && x.MemorySwapMb != nil {
		return *x.MemorySwapMb
	}
	return 0
}

type Container_Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports   []*PortBindings   `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
	Volumes []*VolumeBindings `protobuf:"bytes,2,rep,name=volumes,proto3" json:"volumes,omitempty"`
}

func (x *Container_Settings) Reset() {
	*x = Container_Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container_Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container_Settings) ProtoMessage() {}

func (x *Container_Settings) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container_Settings.ProtoReflect.Descriptor instead.
func (*Container_Settings) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{5, 1}
}

func (x *Container_Settings) GetPorts() []*PortBindings {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Container_Settings) GetVolumes() []*VolumeBindings {
	if x != nil {
		return x.Volumes
	}
	return nil
}

type CreateSmerd_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ImageName       string              `protobuf:"bytes,2,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	Hardware        *Container_Hardware `protobuf:"bytes,3,opt,name=hardware,proto3,oneof" json:"hardware,omitempty"`
	Settings        *Container_Settings `protobuf:"bytes,4,opt,name=settings,proto3,oneof" json:"settings,omitempty"`
	AllowDuplicates bool                `protobuf:"varint,5,opt,name=allow_duplicates,json=allowDuplicates,proto3" json:"allow_duplicates,omitempty"`
}

func (x *CreateSmerd_Request) Reset() {
	*x = CreateSmerd_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSmerd_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSmerd_Request) ProtoMessage() {}

func (x *CreateSmerd_Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSmerd_Request.ProtoReflect.Descriptor instead.
func (*CreateSmerd_Request) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{6, 0}
}

func (x *CreateSmerd_Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSmerd_Request) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *CreateSmerd_Request) GetHardware() *Container_Hardware {
	if x != nil {
		return x.Hardware
	}
	return nil
}

func (x *CreateSmerd_Request) GetSettings() *Container_Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *CreateSmerd_Request) GetAllowDuplicates() bool {
	if x != nil {
		return x.AllowDuplicates
	}
	return false
}

type ListSmerds_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageName string `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
}

func (x *ListSmerds_Request) Reset() {
	*x = ListSmerds_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSmerds_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSmerds_Request) ProtoMessage() {}

func (x *ListSmerds_Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSmerds_Request.ProtoReflect.Descriptor instead.
func (*ListSmerds_Request) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ListSmerds_Request) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

type ListSmerds_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSmerds_Response) Reset() {
	*x = ListSmerds_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_velez_api_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSmerds_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSmerds_Response) ProtoMessage() {}

func (x *ListSmerds_Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_velez_api_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSmerds_Response.ProtoReflect.Descriptor instead.
func (*ListSmerds_Response) Descriptor() ([]byte, []int) {
	return file_grpc_velez_api_proto_rawDescGZIP(), []int{7, 1}
}

var File_grpc_velez_api_proto protoreflect.FileDescriptor

var file_grpc_velez_api_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70,
	0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x22, 0x1c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x07, 0x0a,
	0x03, 0x74, 0x63, 0x70, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x75, 0x64, 0x70, 0x10, 0x01, 0x22,
	0x42, 0x0a, 0x0e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x05, 0x53, 0x6d, 0x65, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x05, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x22, 0xa0, 0x02, 0x0a, 0x09, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x1a, 0xa2, 0x01, 0x0a, 0x08, 0x48, 0x61, 0x72, 0x64,
	0x77, 0x61, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x70, 0x75, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x09, 0x63, 0x70, 0x75, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06, 0x72, 0x61, 0x6d, 0x5f,
	0x6d, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x05, 0x72, 0x61, 0x6d, 0x4d,
	0x62, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73,
	0x77, 0x61, 0x70, 0x5f, 0x6d, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x0c,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x77, 0x61, 0x70, 0x4d, 0x62, 0x88, 0x01, 0x01, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x72, 0x61, 0x6d, 0x5f, 0x6d, 0x62, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x77, 0x61, 0x70, 0x5f, 0x6d, 0x62, 0x1a, 0x6e, 0x0a, 0x08,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x22, 0xba, 0x02, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x65, 0x72, 0x64, 0x1a, 0xaa, 0x02, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0a,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x27, 0xfa, 0x42, 0x24, 0x72, 0x22, 0x28, 0x32, 0x32, 0x1e, 0x28, 0x5b, 0x61, 0x2d, 0x7a,
	0x5d, 0x2b, 0x29, 0x2f, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x2b, 0x29, 0x3a, 0x28, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2e, 0x5d, 0x2b, 0x29, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x72,
	0x64, 0x77, 0x61, 0x72, 0x65, 0x48, 0x00, 0x52, 0x08, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x48, 0x01, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x64, 0x75,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x42, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x6d, 0x65, 0x72, 0x64, 0x73, 0x1a, 0x28, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x1a, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa3, 0x02,
	0x0a, 0x08, 0x56, 0x65, 0x6c, 0x65, 0x7a, 0x41, 0x50, 0x49, 0x12, 0x57, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x65,
	0x72, 0x64, 0x12, 0x1e, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x65, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x6d, 0x65, 0x72, 0x64, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22,
	0x0d, 0x2f, 0x73, 0x6d, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x63,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x65, 0x72, 0x64, 0x73, 0x12, 0x1d, 0x2e, 0x76,
	0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x65,
	0x72, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x65,
	0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6d, 0x65, 0x72,
	0x64, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x73, 0x6d, 0x65, 0x72, 0x64, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x76, 0x65, 0x6c, 0x65, 0x7a, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_velez_api_proto_rawDescOnce sync.Once
	file_grpc_velez_api_proto_rawDescData = file_grpc_velez_api_proto_rawDesc
)

func file_grpc_velez_api_proto_rawDescGZIP() []byte {
	file_grpc_velez_api_proto_rawDescOnce.Do(func() {
		file_grpc_velez_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_velez_api_proto_rawDescData)
	})
	return file_grpc_velez_api_proto_rawDescData
}

var file_grpc_velez_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_velez_api_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_grpc_velez_api_proto_goTypes = []interface{}{
	(PortBindingsProtocol)(0),   // 0: velez_api.PortBindings.protocol
	(*Version)(nil),             // 1: velez_api.Version
	(*PortBindings)(nil),        // 2: velez_api.PortBindings
	(*VolumeBindings)(nil),      // 3: velez_api.VolumeBindings
	(*Image)(nil),               // 4: velez_api.Image
	(*Smerd)(nil),               // 5: velez_api.Smerd
	(*Container)(nil),           // 6: velez_api.Container
	(*CreateSmerd)(nil),         // 7: velez_api.CreateSmerd
	(*ListSmerds)(nil),          // 8: velez_api.ListSmerds
	(*Version_Request)(nil),     // 9: velez_api.Version.Request
	(*Version_Response)(nil),    // 10: velez_api.Version.Response
	(*Container_Hardware)(nil),  // 11: velez_api.Container.Hardware
	(*Container_Settings)(nil),  // 12: velez_api.Container.Settings
	(*CreateSmerd_Request)(nil), // 13: velez_api.CreateSmerd.Request
	(*ListSmerds_Request)(nil),  // 14: velez_api.ListSmerds.Request
	(*ListSmerds_Response)(nil), // 15: velez_api.ListSmerds.Response
}
var file_grpc_velez_api_proto_depIdxs = []int32{
	0,  // 0: velez_api.PortBindings.protoc:type_name -> velez_api.PortBindings.protocol
	2,  // 1: velez_api.Smerd.ports:type_name -> velez_api.PortBindings
	3,  // 2: velez_api.Smerd.volumes:type_name -> velez_api.VolumeBindings
	2,  // 3: velez_api.Container.Settings.ports:type_name -> velez_api.PortBindings
	3,  // 4: velez_api.Container.Settings.volumes:type_name -> velez_api.VolumeBindings
	11, // 5: velez_api.CreateSmerd.Request.hardware:type_name -> velez_api.Container.Hardware
	12, // 6: velez_api.CreateSmerd.Request.settings:type_name -> velez_api.Container.Settings
	9,  // 7: velez_api.VelezAPI.Version:input_type -> velez_api.Version.Request
	13, // 8: velez_api.VelezAPI.CreateSmerd:input_type -> velez_api.CreateSmerd.Request
	14, // 9: velez_api.VelezAPI.ListSmerds:input_type -> velez_api.ListSmerds.Request
	10, // 10: velez_api.VelezAPI.Version:output_type -> velez_api.Version.Response
	5,  // 11: velez_api.VelezAPI.CreateSmerd:output_type -> velez_api.Smerd
	15, // 12: velez_api.VelezAPI.ListSmerds:output_type -> velez_api.ListSmerds.Response
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_grpc_velez_api_proto_init() }
func file_grpc_velez_api_proto_init() {
	if File_grpc_velez_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_velez_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_grpc_velez_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortBindings); i {
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
		file_grpc_velez_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VolumeBindings); i {
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
		file_grpc_velez_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_grpc_velez_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Smerd); i {
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
		file_grpc_velez_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
		file_grpc_velez_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSmerd); i {
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
		file_grpc_velez_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSmerds); i {
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
		file_grpc_velez_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version_Request); i {
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
		file_grpc_velez_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version_Response); i {
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
		file_grpc_velez_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container_Hardware); i {
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
		file_grpc_velez_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container_Settings); i {
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
		file_grpc_velez_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSmerd_Request); i {
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
		file_grpc_velez_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSmerds_Request); i {
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
		file_grpc_velez_api_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSmerds_Response); i {
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
	file_grpc_velez_api_proto_msgTypes[10].OneofWrappers = []interface{}{}
	file_grpc_velez_api_proto_msgTypes[12].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_velez_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_velez_api_proto_goTypes,
		DependencyIndexes: file_grpc_velez_api_proto_depIdxs,
		EnumInfos:         file_grpc_velez_api_proto_enumTypes,
		MessageInfos:      file_grpc_velez_api_proto_msgTypes,
	}.Build()
	File_grpc_velez_api_proto = out.File
	file_grpc_velez_api_proto_rawDesc = nil
	file_grpc_velez_api_proto_goTypes = nil
	file_grpc_velez_api_proto_depIdxs = nil
}

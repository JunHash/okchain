// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/admin.proto

// Key Transparency Administration
//
// The Key Transparency API consists of a map of user names to public
// keys. Each user name also has a history of public keys that have been
// associated with it.

package keytransparency_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	trillian "github.com/google/trillian"
	keyspb "github.com/google/trillian/crypto/keyspb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Domain contains information on a single domain
type Domain struct {
	// DomainId can be any URL safe string.
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// Log contains the Log-Tree's info.
	Log *trillian.Tree `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	// Map contains the Map-Tree's info.
	Map *trillian.Tree `protobuf:"bytes,3,opt,name=map,proto3" json:"map,omitempty"`
	// Vrf contains the VRF public key.
	Vrf *keyspb.PublicKey `protobuf:"bytes,4,opt,name=vrf,proto3" json:"vrf,omitempty"`
	// min_interval is the minimum time between epochs.
	MinInterval *duration.Duration `protobuf:"bytes,5,opt,name=min_interval,json=minInterval,proto3" json:"min_interval,omitempty"`
	// max_interval is the maximum time between epochs.
	MaxInterval *duration.Duration `protobuf:"bytes,6,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
	// Deleted indicates whether the domain has been marked as deleted.
	// By its presence in a response, this domain has not been garbage collected.
	Deleted              bool     `protobuf:"varint,7,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Domain) Reset()         { *m = Domain{} }
func (m *Domain) String() string { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()    {}
func (*Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{0}
}

func (m *Domain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Domain.Unmarshal(m, b)
}
func (m *Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Domain.Marshal(b, m, deterministic)
}
func (m *Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain.Merge(m, src)
}
func (m *Domain) XXX_Size() int {
	return xxx_messageInfo_Domain.Size(m)
}
func (m *Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Domain proto.InternalMessageInfo

func (m *Domain) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *Domain) GetLog() *trillian.Tree {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Domain) GetMap() *trillian.Tree {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *Domain) GetVrf() *keyspb.PublicKey {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *Domain) GetMinInterval() *duration.Duration {
	if m != nil {
		return m.MinInterval
	}
	return nil
}

func (m *Domain) GetMaxInterval() *duration.Duration {
	if m != nil {
		return m.MaxInterval
	}
	return nil
}

func (m *Domain) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

// ListDomains request.
// No pagination options are provided.
type ListDomainsRequest struct {
	// showDeleted requests domains that have been marked for deletion
	// but have not been garbage collected.
	ShowDeleted          bool     `protobuf:"varint,1,opt,name=show_deleted,json=showDeleted,proto3" json:"show_deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDomainsRequest) Reset()         { *m = ListDomainsRequest{} }
func (m *ListDomainsRequest) String() string { return proto.CompactTextString(m) }
func (*ListDomainsRequest) ProtoMessage()    {}
func (*ListDomainsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{1}
}

func (m *ListDomainsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDomainsRequest.Unmarshal(m, b)
}
func (m *ListDomainsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDomainsRequest.Marshal(b, m, deterministic)
}
func (m *ListDomainsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDomainsRequest.Merge(m, src)
}
func (m *ListDomainsRequest) XXX_Size() int {
	return xxx_messageInfo_ListDomainsRequest.Size(m)
}
func (m *ListDomainsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDomainsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDomainsRequest proto.InternalMessageInfo

func (m *ListDomainsRequest) GetShowDeleted() bool {
	if m != nil {
		return m.ShowDeleted
	}
	return false
}

// ListDomains response contains domains.
type ListDomainsResponse struct {
	Domains              []*Domain `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListDomainsResponse) Reset()         { *m = ListDomainsResponse{} }
func (m *ListDomainsResponse) String() string { return proto.CompactTextString(m) }
func (*ListDomainsResponse) ProtoMessage()    {}
func (*ListDomainsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{2}
}

func (m *ListDomainsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDomainsResponse.Unmarshal(m, b)
}
func (m *ListDomainsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDomainsResponse.Marshal(b, m, deterministic)
}
func (m *ListDomainsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDomainsResponse.Merge(m, src)
}
func (m *ListDomainsResponse) XXX_Size() int {
	return xxx_messageInfo_ListDomainsResponse.Size(m)
}
func (m *ListDomainsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDomainsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDomainsResponse proto.InternalMessageInfo

func (m *ListDomainsResponse) GetDomains() []*Domain {
	if m != nil {
		return m.Domains
	}
	return nil
}

// GetDomainRequest specifies the domain to retrieve information for.
type GetDomainRequest struct {
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// showDeleted requests domains that have been marked for deletion
	// but have not been garbage collected.
	ShowDeleted          bool     `protobuf:"varint,2,opt,name=show_deleted,json=showDeleted,proto3" json:"show_deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDomainRequest) Reset()         { *m = GetDomainRequest{} }
func (m *GetDomainRequest) String() string { return proto.CompactTextString(m) }
func (*GetDomainRequest) ProtoMessage()    {}
func (*GetDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{3}
}

func (m *GetDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDomainRequest.Unmarshal(m, b)
}
func (m *GetDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDomainRequest.Marshal(b, m, deterministic)
}
func (m *GetDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDomainRequest.Merge(m, src)
}
func (m *GetDomainRequest) XXX_Size() int {
	return xxx_messageInfo_GetDomainRequest.Size(m)
}
func (m *GetDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDomainRequest proto.InternalMessageInfo

func (m *GetDomainRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *GetDomainRequest) GetShowDeleted() bool {
	if m != nil {
		return m.ShowDeleted
	}
	return false
}

// CreateDomainRequest creates a new domain
type CreateDomainRequest struct {
	DomainId    string             `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	MinInterval *duration.Duration `protobuf:"bytes,2,opt,name=min_interval,json=minInterval,proto3" json:"min_interval,omitempty"`
	MaxInterval *duration.Duration `protobuf:"bytes,3,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
	// The private_key fields allows callers to set the private key.
	VrfPrivateKey        *any.Any `protobuf:"bytes,4,opt,name=vrf_private_key,json=vrfPrivateKey,proto3" json:"vrf_private_key,omitempty"`
	LogPrivateKey        *any.Any `protobuf:"bytes,5,opt,name=log_private_key,json=logPrivateKey,proto3" json:"log_private_key,omitempty"`
	MapPrivateKey        *any.Any `protobuf:"bytes,6,opt,name=map_private_key,json=mapPrivateKey,proto3" json:"map_private_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDomainRequest) Reset()         { *m = CreateDomainRequest{} }
func (m *CreateDomainRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDomainRequest) ProtoMessage()    {}
func (*CreateDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{4}
}

func (m *CreateDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDomainRequest.Unmarshal(m, b)
}
func (m *CreateDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDomainRequest.Marshal(b, m, deterministic)
}
func (m *CreateDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDomainRequest.Merge(m, src)
}
func (m *CreateDomainRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDomainRequest.Size(m)
}
func (m *CreateDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDomainRequest proto.InternalMessageInfo

func (m *CreateDomainRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *CreateDomainRequest) GetMinInterval() *duration.Duration {
	if m != nil {
		return m.MinInterval
	}
	return nil
}

func (m *CreateDomainRequest) GetMaxInterval() *duration.Duration {
	if m != nil {
		return m.MaxInterval
	}
	return nil
}

func (m *CreateDomainRequest) GetVrfPrivateKey() *any.Any {
	if m != nil {
		return m.VrfPrivateKey
	}
	return nil
}

func (m *CreateDomainRequest) GetLogPrivateKey() *any.Any {
	if m != nil {
		return m.LogPrivateKey
	}
	return nil
}

func (m *CreateDomainRequest) GetMapPrivateKey() *any.Any {
	if m != nil {
		return m.MapPrivateKey
	}
	return nil
}

// DeleteDomainRequest deletes a domain
type DeleteDomainRequest struct {
	DomainId             string   `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDomainRequest) Reset()         { *m = DeleteDomainRequest{} }
func (m *DeleteDomainRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDomainRequest) ProtoMessage()    {}
func (*DeleteDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{5}
}

func (m *DeleteDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDomainRequest.Unmarshal(m, b)
}
func (m *DeleteDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDomainRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDomainRequest.Merge(m, src)
}
func (m *DeleteDomainRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDomainRequest.Size(m)
}
func (m *DeleteDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDomainRequest proto.InternalMessageInfo

func (m *DeleteDomainRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

// UndeleteDomainRequest deletes a domain
type UndeleteDomainRequest struct {
	DomainId             string   `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UndeleteDomainRequest) Reset()         { *m = UndeleteDomainRequest{} }
func (m *UndeleteDomainRequest) String() string { return proto.CompactTextString(m) }
func (*UndeleteDomainRequest) ProtoMessage()    {}
func (*UndeleteDomainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{6}
}

func (m *UndeleteDomainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UndeleteDomainRequest.Unmarshal(m, b)
}
func (m *UndeleteDomainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UndeleteDomainRequest.Marshal(b, m, deterministic)
}
func (m *UndeleteDomainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UndeleteDomainRequest.Merge(m, src)
}
func (m *UndeleteDomainRequest) XXX_Size() int {
	return xxx_messageInfo_UndeleteDomainRequest.Size(m)
}
func (m *UndeleteDomainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UndeleteDomainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UndeleteDomainRequest proto.InternalMessageInfo

func (m *UndeleteDomainRequest) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

// GarbageCollect request.
type GarbageCollectRequest struct {
	// Soft-deleted domains with a deleted timestamp before this will be fully
	// deleted.
	Before               *timestamp.Timestamp `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GarbageCollectRequest) Reset()         { *m = GarbageCollectRequest{} }
func (m *GarbageCollectRequest) String() string { return proto.CompactTextString(m) }
func (*GarbageCollectRequest) ProtoMessage()    {}
func (*GarbageCollectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{7}
}

func (m *GarbageCollectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarbageCollectRequest.Unmarshal(m, b)
}
func (m *GarbageCollectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarbageCollectRequest.Marshal(b, m, deterministic)
}
func (m *GarbageCollectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarbageCollectRequest.Merge(m, src)
}
func (m *GarbageCollectRequest) XXX_Size() int {
	return xxx_messageInfo_GarbageCollectRequest.Size(m)
}
func (m *GarbageCollectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GarbageCollectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GarbageCollectRequest proto.InternalMessageInfo

func (m *GarbageCollectRequest) GetBefore() *timestamp.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

type GarbageCollectResponse struct {
	Domains              []*Domain `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GarbageCollectResponse) Reset()         { *m = GarbageCollectResponse{} }
func (m *GarbageCollectResponse) String() string { return proto.CompactTextString(m) }
func (*GarbageCollectResponse) ProtoMessage()    {}
func (*GarbageCollectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_599f1e5eaea78ae3, []int{8}
}

func (m *GarbageCollectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GarbageCollectResponse.Unmarshal(m, b)
}
func (m *GarbageCollectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GarbageCollectResponse.Marshal(b, m, deterministic)
}
func (m *GarbageCollectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GarbageCollectResponse.Merge(m, src)
}
func (m *GarbageCollectResponse) XXX_Size() int {
	return xxx_messageInfo_GarbageCollectResponse.Size(m)
}
func (m *GarbageCollectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GarbageCollectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GarbageCollectResponse proto.InternalMessageInfo

func (m *GarbageCollectResponse) GetDomains() []*Domain {
	if m != nil {
		return m.Domains
	}
	return nil
}

func init() {
	proto.RegisterType((*Domain)(nil), "google.keytransparency.v1.Domain")
	proto.RegisterType((*ListDomainsRequest)(nil), "google.keytransparency.v1.ListDomainsRequest")
	proto.RegisterType((*ListDomainsResponse)(nil), "google.keytransparency.v1.ListDomainsResponse")
	proto.RegisterType((*GetDomainRequest)(nil), "google.keytransparency.v1.GetDomainRequest")
	proto.RegisterType((*CreateDomainRequest)(nil), "google.keytransparency.v1.CreateDomainRequest")
	proto.RegisterType((*DeleteDomainRequest)(nil), "google.keytransparency.v1.DeleteDomainRequest")
	proto.RegisterType((*UndeleteDomainRequest)(nil), "google.keytransparency.v1.UndeleteDomainRequest")
	proto.RegisterType((*GarbageCollectRequest)(nil), "google.keytransparency.v1.GarbageCollectRequest")
	proto.RegisterType((*GarbageCollectResponse)(nil), "google.keytransparency.v1.GarbageCollectResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyTransparencyAdminClient is the client API for KeyTransparencyAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyTransparencyAdminClient interface {
	// ListDomains returns a list of all domains this Key Transparency server
	// operates on.
	ListDomains(ctx context.Context, in *ListDomainsRequest, opts ...grpc.CallOption) (*ListDomainsResponse, error)
	// GetDomain returns the confiuration information for a given domain.
	GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*Domain, error)
	// CreateDomain creates a new Trillian log/map pair.  A unique domainId must
	// be provided.  To create a new domain with the same name as a previously
	// deleted domain, a user must wait X days until the domain is garbage
	// collected.
	CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*Domain, error)
	// DeleteDomain marks a domain as deleted.  Domains will be garbage collected
	// after X days.
	DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// UndeleteDomain marks a previously deleted domain as active if it has not
	// already been garbage collected.
	UndeleteDomain(ctx context.Context, in *UndeleteDomainRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Fully delete soft-deleted domains that have been soft-deleted before the specified timestamp.
	GarbageCollect(ctx context.Context, in *GarbageCollectRequest, opts ...grpc.CallOption) (*GarbageCollectResponse, error)
}

type keyTransparencyAdminClient struct {
	cc *grpc.ClientConn
}

func NewKeyTransparencyAdminClient(cc *grpc.ClientConn) KeyTransparencyAdminClient {
	return &keyTransparencyAdminClient{cc}
}

func (c *keyTransparencyAdminClient) ListDomains(ctx context.Context, in *ListDomainsRequest, opts ...grpc.CallOption) (*ListDomainsResponse, error) {
	out := new(ListDomainsResponse)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/ListDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminClient) GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/GetDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminClient) CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/CreateDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminClient) DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/DeleteDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminClient) UndeleteDomain(ctx context.Context, in *UndeleteDomainRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/UndeleteDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencyAdminClient) GarbageCollect(ctx context.Context, in *GarbageCollectRequest, opts ...grpc.CallOption) (*GarbageCollectResponse, error) {
	out := new(GarbageCollectResponse)
	err := c.cc.Invoke(ctx, "/google.keytransparency.v1.KeyTransparencyAdmin/GarbageCollect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyTransparencyAdminServer is the server API for KeyTransparencyAdmin service.
type KeyTransparencyAdminServer interface {
	// ListDomains returns a list of all domains this Key Transparency server
	// operates on.
	ListDomains(context.Context, *ListDomainsRequest) (*ListDomainsResponse, error)
	// GetDomain returns the confiuration information for a given domain.
	GetDomain(context.Context, *GetDomainRequest) (*Domain, error)
	// CreateDomain creates a new Trillian log/map pair.  A unique domainId must
	// be provided.  To create a new domain with the same name as a previously
	// deleted domain, a user must wait X days until the domain is garbage
	// collected.
	CreateDomain(context.Context, *CreateDomainRequest) (*Domain, error)
	// DeleteDomain marks a domain as deleted.  Domains will be garbage collected
	// after X days.
	DeleteDomain(context.Context, *DeleteDomainRequest) (*empty.Empty, error)
	// UndeleteDomain marks a previously deleted domain as active if it has not
	// already been garbage collected.
	UndeleteDomain(context.Context, *UndeleteDomainRequest) (*empty.Empty, error)
	// Fully delete soft-deleted domains that have been soft-deleted before the specified timestamp.
	GarbageCollect(context.Context, *GarbageCollectRequest) (*GarbageCollectResponse, error)
}

func RegisterKeyTransparencyAdminServer(s *grpc.Server, srv KeyTransparencyAdminServer) {
	s.RegisterService(&_KeyTransparencyAdmin_serviceDesc, srv)
}

func _KeyTransparencyAdmin_ListDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).ListDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/ListDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).ListDomains(ctx, req.(*ListDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdmin_GetDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).GetDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/GetDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).GetDomain(ctx, req.(*GetDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdmin_CreateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).CreateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/CreateDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).CreateDomain(ctx, req.(*CreateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdmin_DeleteDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).DeleteDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/DeleteDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).DeleteDomain(ctx, req.(*DeleteDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdmin_UndeleteDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndeleteDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).UndeleteDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/UndeleteDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).UndeleteDomain(ctx, req.(*UndeleteDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencyAdmin_GarbageCollect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GarbageCollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencyAdminServer).GarbageCollect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.v1.KeyTransparencyAdmin/GarbageCollect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencyAdminServer).GarbageCollect(ctx, req.(*GarbageCollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyTransparencyAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.v1.KeyTransparencyAdmin",
	HandlerType: (*KeyTransparencyAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDomains",
			Handler:    _KeyTransparencyAdmin_ListDomains_Handler,
		},
		{
			MethodName: "GetDomain",
			Handler:    _KeyTransparencyAdmin_GetDomain_Handler,
		},
		{
			MethodName: "CreateDomain",
			Handler:    _KeyTransparencyAdmin_CreateDomain_Handler,
		},
		{
			MethodName: "DeleteDomain",
			Handler:    _KeyTransparencyAdmin_DeleteDomain_Handler,
		},
		{
			MethodName: "UndeleteDomain",
			Handler:    _KeyTransparencyAdmin_UndeleteDomain_Handler,
		},
		{
			MethodName: "GarbageCollect",
			Handler:    _KeyTransparencyAdmin_GarbageCollect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/admin.proto",
}

func init() { proto.RegisterFile("v1/admin.proto", fileDescriptor_599f1e5eaea78ae3) }

var fileDescriptor_599f1e5eaea78ae3 = []byte{
	// 753 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x4f, 0x53, 0xd3, 0x4e,
	0x18, 0xc7, 0x7f, 0x69, 0x7f, 0x14, 0xd8, 0x94, 0x8a, 0x5b, 0xc0, 0x10, 0x1c, 0x29, 0xf1, 0xd2,
	0xc1, 0x31, 0xb1, 0xd5, 0x19, 0x67, 0x90, 0x0b, 0x82, 0x22, 0x53, 0x0f, 0x4c, 0x06, 0x2e, 0x5e,
	0x3a, 0xdb, 0x66, 0x1b, 0x32, 0x24, 0xd9, 0xb8, 0xd9, 0x06, 0x32, 0xea, 0x41, 0xc7, 0xa3, 0x37,
	0xdf, 0x82, 0xef, 0xc8, 0xb7, 0xe0, 0xc1, 0x97, 0xe1, 0x24, 0xd9, 0x94, 0x36, 0x6d, 0x43, 0x3b,
	0x7a, 0x4a, 0x93, 0xef, 0xf3, 0x79, 0xfe, 0xec, 0xf3, 0x3c, 0x49, 0x41, 0x25, 0x68, 0x68, 0xc8,
	0x70, 0x2c, 0x57, 0xf5, 0x28, 0x61, 0x04, 0x6e, 0x9a, 0x84, 0x98, 0x36, 0x56, 0x2f, 0x71, 0xc8,
	0x28, 0x72, 0x7d, 0x0f, 0x51, 0xec, 0x76, 0x43, 0x35, 0x68, 0xc8, 0xf7, 0x13, 0x49, 0x43, 0x9e,
	0xa5, 0x21, 0xd7, 0x25, 0x0c, 0x31, 0x8b, 0xb8, 0x7e, 0x02, 0xca, 0x1c, 0xd4, 0xe2, 0xbb, 0x4e,
	0xbf, 0xa7, 0x21, 0x37, 0xe4, 0xd2, 0x56, 0x56, 0xc2, 0x8e, 0xc7, 0x52, 0xf1, 0x41, 0x56, 0x34,
	0xfa, 0x34, 0x76, 0xcc, 0xf5, 0xed, 0xac, 0xce, 0x2c, 0x07, 0xfb, 0x0c, 0x39, 0x1e, 0x37, 0xa8,
	0x30, 0x6a, 0xd9, 0xb6, 0x85, 0x52, 0x40, 0xee, 0xd2, 0xd0, 0x63, 0x44, 0xbb, 0xc4, 0xa1, 0xef,
	0x75, 0xf8, 0x25, 0xd1, 0x94, 0x1f, 0x05, 0x50, 0x3a, 0x22, 0x0e, 0xb2, 0x5c, 0xb8, 0x05, 0x96,
	0x8d, 0xf8, 0x57, 0xdb, 0x32, 0x24, 0xa1, 0x26, 0xd4, 0x97, 0xf5, 0xa5, 0xe4, 0xc1, 0x89, 0x01,
	0x6b, 0xa0, 0x68, 0x13, 0x53, 0x2a, 0xd4, 0x84, 0xba, 0xd8, 0xac, 0xa8, 0x83, 0x08, 0x67, 0x14,
	0x63, 0x3d, 0x92, 0x22, 0x0b, 0x07, 0x79, 0x52, 0x71, 0xb2, 0x85, 0x83, 0x3c, 0xf8, 0x10, 0x14,
	0x03, 0xda, 0x93, 0xfe, 0x8f, 0x2d, 0xee, 0xaa, 0x3c, 0x8f, 0xd3, 0x7e, 0xc7, 0xb6, 0xba, 0x2d,
	0x1c, 0xea, 0x91, 0x0a, 0xf7, 0x41, 0xd9, 0x89, 0x52, 0x70, 0x19, 0xa6, 0x01, 0xb2, 0xa5, 0x85,
	0xd8, 0x7a, 0x53, 0xe5, 0x5d, 0x48, 0x8b, 0x56, 0x8f, 0xf8, 0xa1, 0xe8, 0xa2, 0x63, 0xb9, 0x27,
	0xdc, 0x3a, 0xa6, 0xd1, 0xf5, 0x0d, 0x5d, 0xba, 0x9d, 0x46, 0xd7, 0x03, 0x5a, 0x02, 0x8b, 0x06,
	0xb6, 0x31, 0xc3, 0x86, 0xb4, 0x58, 0x13, 0xea, 0x4b, 0x7a, 0x7a, 0xab, 0x3c, 0x07, 0xf0, 0xad,
	0xe5, 0xb3, 0xe4, 0xa4, 0x7c, 0x1d, 0xbf, 0xef, 0x63, 0x9f, 0xc1, 0x1d, 0x50, 0xf6, 0x2f, 0xc8,
	0x55, 0x3b, 0x85, 0x84, 0x18, 0x12, 0xa3, 0x67, 0x47, 0x1c, 0xd4, 0x41, 0x75, 0x04, 0xf4, 0x3d,
	0xe2, 0xfa, 0x18, 0xbe, 0x00, 0x8b, 0xc9, 0xd1, 0xfa, 0x92, 0x50, 0x2b, 0xd6, 0xc5, 0xe6, 0x8e,
	0x3a, 0x75, 0xcc, 0xd4, 0x04, 0xd6, 0x53, 0x42, 0xd1, 0xc1, 0xea, 0x31, 0xe6, 0x2e, 0xd3, 0x54,
	0x72, 0x9b, 0x97, 0xcd, 0xb3, 0x30, 0x9e, 0xe7, 0xef, 0x02, 0xa8, 0x1e, 0x52, 0x8c, 0x18, 0x9e,
	0xc3, 0x6f, 0xb6, 0x57, 0x85, 0xbf, 0xea, 0x55, 0x71, 0xae, 0x5e, 0xed, 0x83, 0x3b, 0x01, 0xed,
	0xb5, 0x3d, 0x6a, 0x05, 0x88, 0xe1, 0xf6, 0x25, 0x0e, 0xf9, 0x60, 0xad, 0x8d, 0x39, 0x38, 0x70,
	0x43, 0x7d, 0x25, 0xa0, 0xbd, 0xd3, 0xc4, 0xb6, 0x85, 0xc3, 0x88, 0xb6, 0x89, 0x39, 0x42, 0x2f,
	0xe4, 0xd1, 0x36, 0x31, 0x47, 0x69, 0x07, 0x79, 0x23, 0x74, 0x29, 0x8f, 0x76, 0x90, 0x77, 0x43,
	0x2b, 0x4d, 0x50, 0x4d, 0x4e, 0x7d, 0xf6, 0x93, 0x56, 0x9e, 0x81, 0xf5, 0x73, 0xd7, 0x98, 0x97,
	0x6a, 0x81, 0xf5, 0x63, 0x44, 0x3b, 0xc8, 0xc4, 0x87, 0xc4, 0xb6, 0x71, 0x97, 0xa5, 0x54, 0x13,
	0x94, 0x3a, 0xb8, 0x47, 0x28, 0x8e, 0x11, 0xb1, 0x29, 0x8f, 0xe5, 0x7d, 0x96, 0xbe, 0x53, 0x74,
	0x6e, 0xa9, 0x9c, 0x83, 0x8d, 0xac, 0xb3, 0x7f, 0x30, 0xcc, 0xcd, 0x6f, 0x25, 0xb0, 0xd6, 0xc2,
	0xe1, 0xd9, 0x90, 0xd9, 0x41, 0xf4, 0xf6, 0x85, 0x9f, 0x05, 0x20, 0x0e, 0xad, 0x0e, 0x7c, 0x9c,
	0xe3, 0x74, 0x7c, 0x37, 0x65, 0x75, 0x56, 0xf3, 0xa4, 0x08, 0xa5, 0xfa, 0xe5, 0xe7, 0xaf, 0xef,
	0x85, 0x15, 0x28, 0x6a, 0x41, 0x43, 0xe3, 0xc9, 0xc1, 0x8f, 0x60, 0x79, 0xb0, 0x69, 0xf0, 0x51,
	0x8e, 0xc7, 0xec, 0x3e, 0xca, 0xb7, 0x1f, 0x81, 0xb2, 0x1d, 0x47, 0xdc, 0x84, 0xf7, 0x86, 0x22,
	0x6a, 0x1f, 0x06, 0xdd, 0xfc, 0x04, 0x43, 0x50, 0x1e, 0x5e, 0x49, 0x98, 0x57, 0xd2, 0x84, 0xdd,
	0x9d, 0x25, 0x87, 0x8d, 0x38, 0x87, 0x55, 0x65, 0xb8, 0xea, 0x3d, 0x61, 0x17, 0x5e, 0x81, 0xf2,
	0xf0, 0x8c, 0xe6, 0x86, 0x9e, 0x30, 0xcc, 0xf2, 0xc6, 0xd8, 0x40, 0xbd, 0x8a, 0xbe, 0x70, 0x69,
	0xcd, 0xbb, 0x53, 0x6b, 0xfe, 0x2a, 0x80, 0xca, 0xe8, 0xa4, 0xc3, 0x27, 0x39, 0xb1, 0x27, 0x2e,
	0xc5, 0xd4, 0xe8, 0xf5, 0x38, 0xba, 0xb2, 0x5b, 0x9b, 0x12, 0x7d, 0xaf, 0xcf, 0xdd, 0xc1, 0x2b,
	0x50, 0x19, 0x1d, 0xf6, 0xdc, 0x2c, 0x26, 0x2e, 0x99, 0xdc, 0x98, 0x83, 0xe0, 0x43, 0xf8, 0xdf,
	0xcb, 0x37, 0xef, 0x5e, 0x9b, 0x16, 0xbb, 0xe8, 0x77, 0xd4, 0x2e, 0x71, 0x34, 0xfe, 0xa5, 0xcf,
	0x38, 0xd0, 0xba, 0x84, 0x26, 0x7f, 0x3a, 0x82, 0x46, 0x56, 0x6b, 0x9b, 0xa4, 0x9d, 0x54, 0x5d,
	0x8a, 0x2f, 0x4f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x8d, 0x7d, 0xb9, 0xd0, 0x08, 0x00,
	0x00,
}
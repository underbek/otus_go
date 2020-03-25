// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

package events

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrorCode int32

const (
	ErrorCode_OK    ErrorCode = 0
	ErrorCode_ERROR ErrorCode = 1
)

var ErrorCode_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}

var ErrorCode_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}

type Event struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Heading              string               `protobuf:"bytes,2,opt,name=heading,proto3" json:"heading,omitempty"`
	DateTime             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date_time,json=dateTime,proto3" json:"date_time,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Owner                string               `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Event) GetHeading() string {
	if m != nil {
		return m.Heading
	}
	return ""
}

func (m *Event) GetDateTime() *timestamp.Timestamp {
	if m != nil {
		return m.DateTime
	}
	return nil
}

func (m *Event) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type Response struct {
	Error                ErrorCode `protobuf:"varint,1,opt,name=error,proto3,enum=events.ErrorCode" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetError() ErrorCode {
	if m != nil {
		return m.Error
	}
	return ErrorCode_OK
}

type EventsResponse struct {
	Error                ErrorCode `protobuf:"varint,1,opt,name=error,proto3,enum=events.ErrorCode" json:"error,omitempty"`
	Events               []*Event  `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EventsResponse) Reset()         { *m = EventsResponse{} }
func (m *EventsResponse) String() string { return proto.CompactTextString(m) }
func (*EventsResponse) ProtoMessage()    {}
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{2}
}

func (m *EventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsResponse.Unmarshal(m, b)
}
func (m *EventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsResponse.Marshal(b, m, deterministic)
}
func (m *EventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsResponse.Merge(m, src)
}
func (m *EventsResponse) XXX_Size() int {
	return xxx_messageInfo_EventsResponse.Size(m)
}
func (m *EventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventsResponse proto.InternalMessageInfo

func (m *EventsResponse) GetError() ErrorCode {
	if m != nil {
		return m.Error
	}
	return ErrorCode_OK
}

func (m *EventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type CreateRequest struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{3}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateRequest struct {
	//string uuid = 1; I don't understand why
	Event                *Event   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{4}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type RemoveRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{5}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type DataRequest struct {
	DateTime             *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date_time,json=dateTime,proto3" json:"date_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DataRequest) Reset()         { *m = DataRequest{} }
func (m *DataRequest) String() string { return proto.CompactTextString(m) }
func (*DataRequest) ProtoMessage()    {}
func (*DataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{6}
}

func (m *DataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRequest.Unmarshal(m, b)
}
func (m *DataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRequest.Marshal(b, m, deterministic)
}
func (m *DataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRequest.Merge(m, src)
}
func (m *DataRequest) XXX_Size() int {
	return xxx_messageInfo_DataRequest.Size(m)
}
func (m *DataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataRequest proto.InternalMessageInfo

func (m *DataRequest) GetDateTime() *timestamp.Timestamp {
	if m != nil {
		return m.DateTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("events.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterType((*Event)(nil), "events.Event")
	proto.RegisterType((*Response)(nil), "events.Response")
	proto.RegisterType((*EventsResponse)(nil), "events.EventsResponse")
	proto.RegisterType((*CreateRequest)(nil), "events.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "events.UpdateRequest")
	proto.RegisterType((*RemoveRequest)(nil), "events.RemoveRequest")
	proto.RegisterType((*DataRequest)(nil), "events.DataRequest")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9) }

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0xb1, 0x1b, 0xbb, 0xce, 0x84, 0x04, 0x33, 0x7c, 0xc8, 0xf8, 0x50, 0x22, 0x57, 0x88,
	0x88, 0x83, 0xab, 0xa6, 0x20, 0x4e, 0xc0, 0x21, 0x09, 0x48, 0x7c, 0xa8, 0x68, 0x05, 0xe2, 0x08,
	0x2e, 0x1e, 0x52, 0x8b, 0xc6, 0x6b, 0xbc, 0xeb, 0x22, 0x7e, 0x2d, 0x07, 0xfe, 0x08, 0xf2, 0x6e,
	0xd6, 0xa9, 0xa9, 0x85, 0x20, 0xa7, 0x64, 0x66, 0xde, 0x67, 0x67, 0xe6, 0x1d, 0xc3, 0x55, 0x3a,
	0xa7, 0x5c, 0x8a, 0xb8, 0x28, 0xb9, 0xe4, 0xe8, 0xea, 0x28, 0xbc, 0xbb, 0xe4, 0x7c, 0x79, 0x46,
	0x07, 0x2a, 0x7b, 0x52, 0x7d, 0x39, 0x90, 0xd9, 0x8a, 0x84, 0x4c, 0x56, 0x85, 0x16, 0x86, 0x7b,
	0x7f, 0x0a, 0xd2, 0xaa, 0x4c, 0x64, 0xc6, 0x73, 0x5d, 0x8f, 0x7e, 0x59, 0xe0, 0x2c, 0xea, 0xb7,
	0x10, 0xa1, 0x57, 0x55, 0x59, 0x1a, 0x58, 0x63, 0x6b, 0xd2, 0x67, 0xea, 0x3f, 0x06, 0xb0, 0x7b,
	0x4a, 0x49, 0x9a, 0xe5, 0xcb, 0xc0, 0x56, 0x69, 0x13, 0xe2, 0x63, 0xe8, 0xa7, 0x89, 0xa4, 0x8f,
	0x75, 0xbf, 0x60, 0x67, 0x6c, 0x4d, 0x06, 0xd3, 0x30, 0xd6, 0xbd, 0x62, 0xd3, 0x2b, 0x7e, 0x67,
	0x86, 0x61, 0x5e, 0x2d, 0xae, 0x43, 0x7c, 0x04, 0x9e, 0x19, 0x21, 0xe8, 0x29, 0xee, 0xce, 0x25,
	0x6e, 0xbe, 0x16, 0xb0, 0x46, 0x8a, 0x63, 0x18, 0xa4, 0x24, 0x3e, 0x97, 0x59, 0xa1, 0x48, 0x47,
	0x4d, 0x73, 0x31, 0x85, 0x37, 0xc1, 0xe1, 0xdf, 0x73, 0x2a, 0x03, 0x57, 0xd5, 0x74, 0xf0, 0xb2,
	0xe7, 0xed, 0xfa, 0x5e, 0x74, 0x04, 0x1e, 0x23, 0x51, 0xf0, 0x5c, 0x10, 0xde, 0x07, 0x87, 0xca,
	0x92, 0x97, 0x6a, 0xd1, 0xd1, 0xf4, 0x7a, 0xbc, 0x36, 0x76, 0x51, 0x27, 0x67, 0x3c, 0x25, 0xa6,
	0xeb, 0xd1, 0x27, 0x18, 0x29, 0x67, 0xc4, 0x7f, 0xa3, 0x78, 0x0f, 0xd6, 0x07, 0x0a, 0xec, 0xf1,
	0xce, 0x64, 0x30, 0x1d, 0x36, 0xca, 0xfa, 0x87, 0xad, 0x8b, 0xd1, 0x43, 0x18, 0xce, 0x4a, 0x4a,
	0x24, 0x31, 0xfa, 0x56, 0x91, 0x90, 0xb8, 0x0f, 0x8e, 0x2a, 0xa9, 0x06, 0x97, 0x30, 0x5d, 0xab,
	0xa9, 0xf7, 0x45, 0xda, 0x45, 0xd9, 0x7f, 0xa1, 0xf6, 0x61, 0xc8, 0x68, 0xc5, 0xcf, 0x1b, 0xaa,
	0xe3, 0xde, 0xd1, 0x73, 0x18, 0xcc, 0x13, 0x99, 0x18, 0x49, 0xeb, 0xc8, 0xd6, 0xbf, 0x1f, 0xf9,
	0xc1, 0x1e, 0xf4, 0x1b, 0x4f, 0xd0, 0x05, 0xfb, 0xf8, 0x95, 0x7f, 0x05, 0xfb, 0xe0, 0x2c, 0x18,
	0x3b, 0x66, 0xbe, 0x35, 0xfd, 0x69, 0x43, 0xef, 0x05, 0x7b, 0x3b, 0xc3, 0x43, 0x70, 0xb5, 0x03,
	0x78, 0xcb, 0x4c, 0xdd, 0x72, 0x24, 0xf4, 0x4d, 0xba, 0x39, 0xc2, 0x21, 0xb8, 0x7a, 0xfd, 0x0d,
	0xd2, 0xb2, 0xa3, 0x1b, 0xd1, 0xbb, 0x6f, 0x90, 0x96, 0x17, 0x1d, 0xc8, 0x13, 0x18, 0xcd, 0x93,
	0xec, 0xec, 0x87, 0xf2, 0xf0, 0x75, 0x26, 0x24, 0xde, 0x30, 0x9a, 0x0b, 0x0e, 0x85, 0xb7, 0x5b,
	0x5e, 0x6f, 0xbe, 0x94, 0xa7, 0x70, 0xed, 0x03, 0xd1, 0xd7, 0xad, 0xf9, 0x67, 0xe0, 0xbf, 0xe1,
	0xb9, 0x3c, 0xdd, 0xf6, 0x81, 0x13, 0x57, 0xdd, 0xe7, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x32, 0x66, 0x20, 0xd4, 0x37, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GRPCClient is the client API for GRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPCClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*Response, error)
	DailyEventList(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	WeeklyEventList(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*EventsResponse, error)
	MonthlyEventList(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*EventsResponse, error)
}

type gRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCClient(cc grpc.ClientConnInterface) GRPCClient {
	return &gRPCClient{cc}
}

func (c *gRPCClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/events.GRPC/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/events.GRPC/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/events.GRPC/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCClient) DailyEventList(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/events.GRPC/DailyEventList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCClient) WeeklyEventList(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/events.GRPC/WeeklyEventList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCClient) MonthlyEventList(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/events.GRPC/MonthlyEventList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCServer is the server API for GRPC service.
type GRPCServer interface {
	Create(context.Context, *CreateRequest) (*Response, error)
	Update(context.Context, *UpdateRequest) (*Response, error)
	Remove(context.Context, *RemoveRequest) (*Response, error)
	DailyEventList(context.Context, *DataRequest) (*EventsResponse, error)
	WeeklyEventList(context.Context, *DataRequest) (*EventsResponse, error)
	MonthlyEventList(context.Context, *DataRequest) (*EventsResponse, error)
}

// UnimplementedGRPCServer can be embedded to have forward compatible implementations.
type UnimplementedGRPCServer struct {
}

func (*UnimplementedGRPCServer) Create(ctx context.Context, req *CreateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedGRPCServer) Update(ctx context.Context, req *UpdateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedGRPCServer) Remove(ctx context.Context, req *RemoveRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedGRPCServer) DailyEventList(ctx context.Context, req *DataRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DailyEventList not implemented")
}
func (*UnimplementedGRPCServer) WeeklyEventList(ctx context.Context, req *DataRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WeeklyEventList not implemented")
}
func (*UnimplementedGRPCServer) MonthlyEventList(ctx context.Context, req *DataRequest) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MonthlyEventList not implemented")
}

func RegisterGRPCServer(s *grpc.Server, srv GRPCServer) {
	s.RegisterService(&_GRPC_serviceDesc, srv)
}

func _GRPC_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.GRPC/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPC_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.GRPC/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPC_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.GRPC/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPC_DailyEventList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServer).DailyEventList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.GRPC/DailyEventList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServer).DailyEventList(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPC_WeeklyEventList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServer).WeeklyEventList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.GRPC/WeeklyEventList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServer).WeeklyEventList(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPC_MonthlyEventList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServer).MonthlyEventList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.GRPC/MonthlyEventList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServer).MonthlyEventList(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.GRPC",
	HandlerType: (*GRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GRPC_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GRPC_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _GRPC_Remove_Handler,
		},
		{
			MethodName: "DailyEventList",
			Handler:    _GRPC_DailyEventList_Handler,
		},
		{
			MethodName: "WeeklyEventList",
			Handler:    _GRPC_WeeklyEventList_Handler,
		},
		{
			MethodName: "MonthlyEventList",
			Handler:    _GRPC_MonthlyEventList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events.proto",
}

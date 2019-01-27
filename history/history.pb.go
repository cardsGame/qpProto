// Code generated by protoc-gen-go. DO NOT EDIT.
// source: history/history.proto

package history

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// 获取用户战绩请求
type GetUserHistoryRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserHistoryRequest) Reset()         { *m = GetUserHistoryRequest{} }
func (m *GetUserHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserHistoryRequest) ProtoMessage()    {}
func (*GetUserHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_132197fc40cc0614, []int{0}
}

func (m *GetUserHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserHistoryRequest.Unmarshal(m, b)
}
func (m *GetUserHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetUserHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserHistoryRequest.Merge(m, src)
}
func (m *GetUserHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserHistoryRequest.Size(m)
}
func (m *GetUserHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserHistoryRequest proto.InternalMessageInfo

func (m *GetUserHistoryRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetUserHistoryRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// 获取用户战绩响应
type GetUserHistoryResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserHistoryResponse) Reset()         { *m = GetUserHistoryResponse{} }
func (m *GetUserHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserHistoryResponse) ProtoMessage()    {}
func (*GetUserHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_132197fc40cc0614, []int{1}
}

func (m *GetUserHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserHistoryResponse.Unmarshal(m, b)
}
func (m *GetUserHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetUserHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserHistoryResponse.Merge(m, src)
}
func (m *GetUserHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserHistoryResponse.Size(m)
}
func (m *GetUserHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserHistoryResponse proto.InternalMessageInfo

// 游戏记录
type GameRecord struct {
	State                int32    `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	CreateAt             string   `protobuf:"bytes,3,opt,name=createAt,proto3" json:"createAt,omitempty"`
	OverAt               string   `protobuf:"bytes,4,opt,name=overAt,proto3" json:"overAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameRecord) Reset()         { *m = GameRecord{} }
func (m *GameRecord) String() string { return proto.CompactTextString(m) }
func (*GameRecord) ProtoMessage()    {}
func (*GameRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_132197fc40cc0614, []int{2}
}

func (m *GameRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameRecord.Unmarshal(m, b)
}
func (m *GameRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameRecord.Marshal(b, m, deterministic)
}
func (m *GameRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameRecord.Merge(m, src)
}
func (m *GameRecord) XXX_Size() int {
	return xxx_messageInfo_GameRecord.Size(m)
}
func (m *GameRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_GameRecord.DiscardUnknown(m)
}

var xxx_messageInfo_GameRecord proto.InternalMessageInfo

func (m *GameRecord) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *GameRecord) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *GameRecord) GetOverAt() string {
	if m != nil {
		return m.OverAt
	}
	return ""
}

// 游戏玩家
type GamePlayerData struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Score                int32    `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Index                int32    `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	TotalScore           int32    `protobuf:"varint,5,opt,name=totalScore,proto3" json:"totalScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GamePlayerData) Reset()         { *m = GamePlayerData{} }
func (m *GamePlayerData) String() string { return proto.CompactTextString(m) }
func (*GamePlayerData) ProtoMessage()    {}
func (*GamePlayerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_132197fc40cc0614, []int{3}
}

func (m *GamePlayerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GamePlayerData.Unmarshal(m, b)
}
func (m *GamePlayerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GamePlayerData.Marshal(b, m, deterministic)
}
func (m *GamePlayerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GamePlayerData.Merge(m, src)
}
func (m *GamePlayerData) XXX_Size() int {
	return xxx_messageInfo_GamePlayerData.Size(m)
}
func (m *GamePlayerData) XXX_DiscardUnknown() {
	xxx_messageInfo_GamePlayerData.DiscardUnknown(m)
}

var xxx_messageInfo_GamePlayerData proto.InternalMessageInfo

func (m *GamePlayerData) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GamePlayerData) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *GamePlayerData) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *GamePlayerData) GetTotalScore() int32 {
	if m != nil {
		return m.TotalScore
	}
	return 0
}

// 房间记录
type RoomRecordData struct {
	RoomNo               string            `protobuf:"bytes,1,opt,name=roomNo,proto3" json:"roomNo,omitempty"`
	RealRoomNo           string            `protobuf:"bytes,2,opt,name=realRoomNo,proto3" json:"realRoomNo,omitempty"`
	GameType             int32             `protobuf:"varint,3,opt,name=gameType,proto3" json:"gameType,omitempty"`
	RoomType             int32             `protobuf:"varint,4,opt,name=roomType,proto3" json:"roomType,omitempty"`
	State                int32             `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`
	CreateAt             string            `protobuf:"bytes,6,opt,name=createAt,proto3" json:"createAt,omitempty"`
	OverAt               string            `protobuf:"bytes,7,opt,name=overAt,proto3" json:"overAt,omitempty"`
	FamilyCircleNo       string            `protobuf:"bytes,8,opt,name=familyCircleNo,proto3" json:"familyCircleNo,omitempty"`
	OverType             int32             `protobuf:"varint,9,opt,name=overType,proto3" json:"overType,omitempty"`
	Players              []*RoomPlayerData `protobuf:"bytes,10,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RoomRecordData) Reset()         { *m = RoomRecordData{} }
func (m *RoomRecordData) String() string { return proto.CompactTextString(m) }
func (*RoomRecordData) ProtoMessage()    {}
func (*RoomRecordData) Descriptor() ([]byte, []int) {
	return fileDescriptor_132197fc40cc0614, []int{4}
}

func (m *RoomRecordData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomRecordData.Unmarshal(m, b)
}
func (m *RoomRecordData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomRecordData.Marshal(b, m, deterministic)
}
func (m *RoomRecordData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomRecordData.Merge(m, src)
}
func (m *RoomRecordData) XXX_Size() int {
	return xxx_messageInfo_RoomRecordData.Size(m)
}
func (m *RoomRecordData) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomRecordData.DiscardUnknown(m)
}

var xxx_messageInfo_RoomRecordData proto.InternalMessageInfo

func (m *RoomRecordData) GetRoomNo() string {
	if m != nil {
		return m.RoomNo
	}
	return ""
}

func (m *RoomRecordData) GetRealRoomNo() string {
	if m != nil {
		return m.RealRoomNo
	}
	return ""
}

func (m *RoomRecordData) GetGameType() int32 {
	if m != nil {
		return m.GameType
	}
	return 0
}

func (m *RoomRecordData) GetRoomType() int32 {
	if m != nil {
		return m.RoomType
	}
	return 0
}

func (m *RoomRecordData) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *RoomRecordData) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *RoomRecordData) GetOverAt() string {
	if m != nil {
		return m.OverAt
	}
	return ""
}

func (m *RoomRecordData) GetFamilyCircleNo() string {
	if m != nil {
		return m.FamilyCircleNo
	}
	return ""
}

func (m *RoomRecordData) GetOverType() int32 {
	if m != nil {
		return m.OverType
	}
	return 0
}

func (m *RoomRecordData) GetPlayers() []*RoomPlayerData {
	if m != nil {
		return m.Players
	}
	return nil
}

// 房间玩家
type RoomPlayerData struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	HeadimgUrl           string   `protobuf:"bytes,2,opt,name=headimgUrl,proto3" json:"headimgUrl,omitempty"`
	Gender               int32    `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	NickName             string   `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Role                 int32    `protobuf:"varint,5,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomPlayerData) Reset()         { *m = RoomPlayerData{} }
func (m *RoomPlayerData) String() string { return proto.CompactTextString(m) }
func (*RoomPlayerData) ProtoMessage()    {}
func (*RoomPlayerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_132197fc40cc0614, []int{5}
}

func (m *RoomPlayerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomPlayerData.Unmarshal(m, b)
}
func (m *RoomPlayerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomPlayerData.Marshal(b, m, deterministic)
}
func (m *RoomPlayerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomPlayerData.Merge(m, src)
}
func (m *RoomPlayerData) XXX_Size() int {
	return xxx_messageInfo_RoomPlayerData.Size(m)
}
func (m *RoomPlayerData) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomPlayerData.DiscardUnknown(m)
}

var xxx_messageInfo_RoomPlayerData proto.InternalMessageInfo

func (m *RoomPlayerData) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RoomPlayerData) GetHeadimgUrl() string {
	if m != nil {
		return m.HeadimgUrl
	}
	return ""
}

func (m *RoomPlayerData) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *RoomPlayerData) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *RoomPlayerData) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUserHistoryRequest)(nil), "GetUserHistoryRequest")
	proto.RegisterType((*GetUserHistoryResponse)(nil), "GetUserHistoryResponse")
	proto.RegisterType((*GameRecord)(nil), "GameRecord")
	proto.RegisterType((*GamePlayerData)(nil), "GamePlayerData")
	proto.RegisterType((*RoomRecordData)(nil), "RoomRecordData")
	proto.RegisterType((*RoomPlayerData)(nil), "RoomPlayerData")
}

func init() { proto.RegisterFile("history/history.proto", fileDescriptor_132197fc40cc0614) }

var fileDescriptor_132197fc40cc0614 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x25, 0x69, 0x37, 0x69, 0x06, 0x29, 0x48, 0x16, 0x0d, 0xab, 0x1e, 0xaa, 0x68, 0x0f, 0x28,
	0x5c, 0x82, 0x54, 0xbe, 0x20, 0x2a, 0x52, 0xe1, 0x12, 0x21, 0x43, 0xb9, 0x9b, 0xdd, 0x21, 0x5d,
	0xb1, 0xce, 0x2c, 0x63, 0x17, 0xb1, 0x7f, 0xc1, 0xef, 0xf0, 0x77, 0xc8, 0x63, 0x6f, 0x12, 0x45,
	0x44, 0xea, 0x29, 0xfb, 0xde, 0xcb, 0x3c, 0xcf, 0x3c, 0x8f, 0xe1, 0xf2, 0xa1, 0x76, 0x9e, 0xb8,
	0x7b, 0x9b, 0x7e, 0x97, 0x2d, 0x93, 0xa7, 0x62, 0x05, 0x97, 0x77, 0xe8, 0xef, 0x1d, 0xf2, 0x87,
	0xc8, 0x6b, 0xfc, 0xf9, 0x88, 0xce, 0x2b, 0x05, 0xe7, 0xad, 0xd9, 0x60, 0x3e, 0x98, 0x0f, 0x16,
	0x99, 0x96, 0x6f, 0xf5, 0x12, 0xb2, 0xa6, 0xb6, 0xb5, 0xcf, 0x87, 0x42, 0x46, 0x50, 0xe4, 0x30,
	0x3b, 0xb6, 0x70, 0x2d, 0x6d, 0x1d, 0x16, 0x5f, 0x01, 0xee, 0x8c, 0x45, 0x8d, 0x25, 0x71, 0x15,
	0xaa, 0x9d, 0x37, 0xbe, 0xb7, 0x8c, 0x40, 0x5d, 0xc1, 0x45, 0xc9, 0x68, 0x3c, 0xae, 0x7c, 0x7e,
	0x36, 0x1f, 0x2c, 0x26, 0x7a, 0x87, 0xd5, 0x0c, 0x46, 0xf4, 0x0b, 0x79, 0xe5, 0xf3, 0x73, 0x51,
	0x12, 0x2a, 0x3c, 0x4c, 0x83, 0xef, 0xa7, 0xc6, 0x74, 0xc8, 0xef, 0x8d, 0x37, 0xe1, 0x9f, 0x8f,
	0x0e, 0xf9, 0x63, 0x95, 0xcc, 0x13, 0x92, 0x33, 0x4b, 0x62, 0xec, 0x3b, 0x16, 0x10, 0xd8, 0x7a,
	0x5b, 0xe1, 0x6f, 0xb1, 0xcd, 0x74, 0x04, 0xea, 0x1a, 0xc0, 0x93, 0x37, 0xcd, 0x67, 0x29, 0xc8,
	0x44, 0x3a, 0x60, 0x8a, 0xbf, 0x43, 0x98, 0x6a, 0x22, 0x1b, 0xc7, 0xe9, 0x8f, 0x65, 0x22, 0xbb,
	0x26, 0x39, 0x76, 0xa2, 0x13, 0x0a, 0x56, 0x8c, 0xa6, 0xd1, 0x51, 0x1b, 0x8a, 0x76, 0xc0, 0x84,
	0xa1, 0x37, 0xc6, 0xe2, 0x97, 0xae, 0x45, 0x19, 0x3a, 0xd3, 0x3b, 0x1c, 0xb4, 0xe0, 0x22, 0x5a,
	0xec, 0x6f, 0x87, 0xf7, 0x11, 0x66, 0xa7, 0x22, 0x1c, 0x9d, 0x8c, 0x70, 0x7c, 0x18, 0xa1, 0x7a,
	0x0d, 0xd3, 0xef, 0xc6, 0xd6, 0x4d, 0x77, 0x5b, 0x73, 0xd9, 0xe0, 0x9a, 0xf2, 0x0b, 0xd1, 0x8f,
	0xd8, 0xe0, 0x1d, 0x2a, 0xa4, 0x9b, 0x49, 0xec, 0xa6, 0xc7, 0xea, 0x0d, 0x8c, 0x5b, 0xb9, 0x02,
	0x97, 0xc3, 0xfc, 0x6c, 0xf1, 0xfc, 0xe6, 0xc5, 0x32, 0xcc, 0xb7, 0xbf, 0x16, 0xdd, 0xeb, 0xc5,
	0x9f, 0x41, 0xcc, 0xee, 0x09, 0x57, 0x76, 0x0d, 0xf0, 0x80, 0xa6, 0xaa, 0xed, 0xe6, 0x9e, 0x9b,
	0x3e, 0xbb, 0x3d, 0x13, 0xea, 0x36, 0xb8, 0xad, 0x90, 0x53, 0x72, 0x09, 0x85, 0x4e, 0xb7, 0x75,
	0xf9, 0x63, 0x6d, 0x2c, 0xa6, 0x75, 0xd9, 0xe1, 0xb0, 0xcc, 0x4c, 0x4d, 0x1f, 0x9b, 0x7c, 0xdf,
	0xac, 0x61, 0x9c, 0xf6, 0x55, 0xdd, 0xc2, 0xf4, 0x60, 0x83, 0x99, 0x3a, 0x35, 0x5b, 0xfe, 0xf7,
	0x55, 0x5c, 0xbd, 0x5a, 0x9e, 0x58, 0xf5, 0x67, 0xdf, 0x46, 0xf2, 0xa0, 0xde, 0xfd, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xa0, 0xaf, 0xbe, 0xae, 0x69, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HistoryClient is the client API for History service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HistoryClient interface {
	GetUserHistroy(ctx context.Context, in *GetUserHistoryRequest, opts ...grpc.CallOption) (*GetUserHistoryResponse, error)
}

type historyClient struct {
	cc *grpc.ClientConn
}

func NewHistoryClient(cc *grpc.ClientConn) HistoryClient {
	return &historyClient{cc}
}

func (c *historyClient) GetUserHistroy(ctx context.Context, in *GetUserHistoryRequest, opts ...grpc.CallOption) (*GetUserHistoryResponse, error) {
	out := new(GetUserHistoryResponse)
	err := c.cc.Invoke(ctx, "/History/GetUserHistroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryServer is the server API for History service.
type HistoryServer interface {
	GetUserHistroy(context.Context, *GetUserHistoryRequest) (*GetUserHistoryResponse, error)
}

func RegisterHistoryServer(s *grpc.Server, srv HistoryServer) {
	s.RegisterService(&_History_serviceDesc, srv)
}

func _History_GetUserHistroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServer).GetUserHistroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/History/GetUserHistroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServer).GetUserHistroy(ctx, req.(*GetUserHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _History_serviceDesc = grpc.ServiceDesc{
	ServiceName: "History",
	HandlerType: (*HistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserHistroy",
			Handler:    _History_GetUserHistroy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "history/history.proto",
}
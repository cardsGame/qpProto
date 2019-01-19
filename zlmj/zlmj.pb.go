// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zlmj/zlmj.proto

package zlmj

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 状态
type State int32

const (
	State_initing     State = 0
	State_can_discard State = 1
	State_over        State = 2
)

var State_name = map[int32]string{
	0: "initing",
	1: "can_discard",
	2: "over",
}

var State_value = map[string]int32{
	"initing":     0,
	"can_discard": 1,
	"over":        2,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{0}
}

// 游戏选项
type Option struct {
	Cheat                bool     `protobuf:"varint,1,opt,name=cheat,proto3" json:"cheat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Option) Reset()         { *m = Option{} }
func (m *Option) String() string { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()    {}
func (*Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{0}
}

func (m *Option) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Option.Unmarshal(m, b)
}
func (m *Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Option.Marshal(b, m, deterministic)
}
func (m *Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Option.Merge(m, src)
}
func (m *Option) XXX_Size() int {
	return xxx_messageInfo_Option.Size(m)
}
func (m *Option) XXX_DiscardUnknown() {
	xxx_messageInfo_Option.DiscardUnknown(m)
}

var xxx_messageInfo_Option proto.InternalMessageInfo

func (m *Option) GetCheat() bool {
	if m != nil {
		return m.Cheat
	}
	return false
}

// 游戏数据结构
type GameData struct {
	State                State         `protobuf:"varint,1,opt,name=state,proto3,enum=zlmj.State" json:"state,omitempty"`
	Option               *Option       `protobuf:"bytes,2,opt,name=option,proto3" json:"option,omitempty"`
	Players              []*PlayerData `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GameData) Reset()         { *m = GameData{} }
func (m *GameData) String() string { return proto.CompactTextString(m) }
func (*GameData) ProtoMessage()    {}
func (*GameData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{1}
}

func (m *GameData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameData.Unmarshal(m, b)
}
func (m *GameData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameData.Marshal(b, m, deterministic)
}
func (m *GameData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameData.Merge(m, src)
}
func (m *GameData) XXX_Size() int {
	return xxx_messageInfo_GameData.Size(m)
}
func (m *GameData) XXX_DiscardUnknown() {
	xxx_messageInfo_GameData.DiscardUnknown(m)
}

var xxx_messageInfo_GameData proto.InternalMessageInfo

func (m *GameData) GetState() State {
	if m != nil {
		return m.State
	}
	return State_initing
}

func (m *GameData) GetOption() *Option {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *GameData) GetPlayers() []*PlayerData {
	if m != nil {
		return m.Players
	}
	return nil
}

// 玩家数据结构
type PlayerData struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Index                int32    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerData) Reset()         { *m = PlayerData{} }
func (m *PlayerData) String() string { return proto.CompactTextString(m) }
func (*PlayerData) ProtoMessage()    {}
func (*PlayerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{2}
}

func (m *PlayerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerData.Unmarshal(m, b)
}
func (m *PlayerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerData.Marshal(b, m, deterministic)
}
func (m *PlayerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerData.Merge(m, src)
}
func (m *PlayerData) XXX_Size() int {
	return xxx_messageInfo_PlayerData.Size(m)
}
func (m *PlayerData) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerData.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerData proto.InternalMessageInfo

func (m *PlayerData) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PlayerData) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type Req struct {
	// Types that are valid to be assigned to Req:
	//	*Req_DiscardsReq
	//	*Req_WantHoldsReq
	Req                  isReq_Req `protobuf_oneof:"req"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}
func (*Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{3}
}

func (m *Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Req.Unmarshal(m, b)
}
func (m *Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Req.Marshal(b, m, deterministic)
}
func (m *Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Req.Merge(m, src)
}
func (m *Req) XXX_Size() int {
	return xxx_messageInfo_Req.Size(m)
}
func (m *Req) XXX_DiscardUnknown() {
	xxx_messageInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_Req proto.InternalMessageInfo

type isReq_Req interface {
	isReq_Req()
}

type Req_DiscardsReq struct {
	DiscardsReq *DiscardsReq `protobuf:"bytes,1,opt,name=discardsReq,proto3,oneof"`
}

type Req_WantHoldsReq struct {
	WantHoldsReq *WantHoldsReq `protobuf:"bytes,2,opt,name=wantHoldsReq,proto3,oneof"`
}

func (*Req_DiscardsReq) isReq_Req() {}

func (*Req_WantHoldsReq) isReq_Req() {}

func (m *Req) GetReq() isReq_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *Req) GetDiscardsReq() *DiscardsReq {
	if x, ok := m.GetReq().(*Req_DiscardsReq); ok {
		return x.DiscardsReq
	}
	return nil
}

func (m *Req) GetWantHoldsReq() *WantHoldsReq {
	if x, ok := m.GetReq().(*Req_WantHoldsReq); ok {
		return x.WantHoldsReq
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Req) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Req_DiscardsReq)(nil),
		(*Req_WantHoldsReq)(nil),
	}
}

type Rsp struct {
	// Types that are valid to be assigned to Rsp:
	//	*Rsp_ErrorRsp
	//	*Rsp_DiscardsRsp
	Rsp                  isRsp_Rsp `protobuf_oneof:"rsp"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Rsp) Reset()         { *m = Rsp{} }
func (m *Rsp) String() string { return proto.CompactTextString(m) }
func (*Rsp) ProtoMessage()    {}
func (*Rsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{4}
}

func (m *Rsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rsp.Unmarshal(m, b)
}
func (m *Rsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rsp.Marshal(b, m, deterministic)
}
func (m *Rsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rsp.Merge(m, src)
}
func (m *Rsp) XXX_Size() int {
	return xxx_messageInfo_Rsp.Size(m)
}
func (m *Rsp) XXX_DiscardUnknown() {
	xxx_messageInfo_Rsp.DiscardUnknown(m)
}

var xxx_messageInfo_Rsp proto.InternalMessageInfo

type isRsp_Rsp interface {
	isRsp_Rsp()
}

type Rsp_ErrorRsp struct {
	ErrorRsp *ErrorRsp `protobuf:"bytes,1,opt,name=errorRsp,proto3,oneof"`
}

type Rsp_DiscardsRsp struct {
	DiscardsRsp *DiscardsRsp `protobuf:"bytes,2,opt,name=discardsRsp,proto3,oneof"`
}

func (*Rsp_ErrorRsp) isRsp_Rsp() {}

func (*Rsp_DiscardsRsp) isRsp_Rsp() {}

func (m *Rsp) GetRsp() isRsp_Rsp {
	if m != nil {
		return m.Rsp
	}
	return nil
}

func (m *Rsp) GetErrorRsp() *ErrorRsp {
	if x, ok := m.GetRsp().(*Rsp_ErrorRsp); ok {
		return x.ErrorRsp
	}
	return nil
}

func (m *Rsp) GetDiscardsRsp() *DiscardsRsp {
	if x, ok := m.GetRsp().(*Rsp_DiscardsRsp); ok {
		return x.DiscardsRsp
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Rsp) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Rsp_ErrorRsp)(nil),
		(*Rsp_DiscardsRsp)(nil),
	}
}

type Event struct {
	// Types that are valid to be assigned to Event:
	//	*Event_DiscardsEvent
	//	*Event_GameStartEvent
	Event                isEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{5}
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

type isEvent_Event interface {
	isEvent_Event()
}

type Event_DiscardsEvent struct {
	DiscardsEvent *DiscardsEvent `protobuf:"bytes,1,opt,name=discardsEvent,proto3,oneof"`
}

type Event_GameStartEvent struct {
	GameStartEvent *GameStartEvent `protobuf:"bytes,2,opt,name=gameStartEvent,proto3,oneof"`
}

func (*Event_DiscardsEvent) isEvent_Event() {}

func (*Event_GameStartEvent) isEvent_Event() {}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetDiscardsEvent() *DiscardsEvent {
	if x, ok := m.GetEvent().(*Event_DiscardsEvent); ok {
		return x.DiscardsEvent
	}
	return nil
}

func (m *Event) GetGameStartEvent() *GameStartEvent {
	if x, ok := m.GetEvent().(*Event_GameStartEvent); ok {
		return x.GameStartEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_DiscardsEvent)(nil),
		(*Event_GameStartEvent)(nil),
	}
}

// region 请求
// 出牌请求
type WantHoldsReq struct {
	Cards                []int32  `protobuf:"varint,1,rep,packed,name=cards,proto3" json:"cards,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WantHoldsReq) Reset()         { *m = WantHoldsReq{} }
func (m *WantHoldsReq) String() string { return proto.CompactTextString(m) }
func (*WantHoldsReq) ProtoMessage()    {}
func (*WantHoldsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{6}
}

func (m *WantHoldsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WantHoldsReq.Unmarshal(m, b)
}
func (m *WantHoldsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WantHoldsReq.Marshal(b, m, deterministic)
}
func (m *WantHoldsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WantHoldsReq.Merge(m, src)
}
func (m *WantHoldsReq) XXX_Size() int {
	return xxx_messageInfo_WantHoldsReq.Size(m)
}
func (m *WantHoldsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_WantHoldsReq.DiscardUnknown(m)
}

var xxx_messageInfo_WantHoldsReq proto.InternalMessageInfo

func (m *WantHoldsReq) GetCards() []int32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

type DiscardsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscardsReq) Reset()         { *m = DiscardsReq{} }
func (m *DiscardsReq) String() string { return proto.CompactTextString(m) }
func (*DiscardsReq) ProtoMessage()    {}
func (*DiscardsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{7}
}

func (m *DiscardsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscardsReq.Unmarshal(m, b)
}
func (m *DiscardsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscardsReq.Marshal(b, m, deterministic)
}
func (m *DiscardsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscardsReq.Merge(m, src)
}
func (m *DiscardsReq) XXX_Size() int {
	return xxx_messageInfo_DiscardsReq.Size(m)
}
func (m *DiscardsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscardsReq.DiscardUnknown(m)
}

var xxx_messageInfo_DiscardsReq proto.InternalMessageInfo

// region 响应
// 错误响应
type ErrorRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorRsp) Reset()         { *m = ErrorRsp{} }
func (m *ErrorRsp) String() string { return proto.CompactTextString(m) }
func (*ErrorRsp) ProtoMessage()    {}
func (*ErrorRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{8}
}

func (m *ErrorRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorRsp.Unmarshal(m, b)
}
func (m *ErrorRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorRsp.Marshal(b, m, deterministic)
}
func (m *ErrorRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorRsp.Merge(m, src)
}
func (m *ErrorRsp) XXX_Size() int {
	return xxx_messageInfo_ErrorRsp.Size(m)
}
func (m *ErrorRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorRsp proto.InternalMessageInfo

// 出牌响应
type DiscardsRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscardsRsp) Reset()         { *m = DiscardsRsp{} }
func (m *DiscardsRsp) String() string { return proto.CompactTextString(m) }
func (*DiscardsRsp) ProtoMessage()    {}
func (*DiscardsRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{9}
}

func (m *DiscardsRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscardsRsp.Unmarshal(m, b)
}
func (m *DiscardsRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscardsRsp.Marshal(b, m, deterministic)
}
func (m *DiscardsRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscardsRsp.Merge(m, src)
}
func (m *DiscardsRsp) XXX_Size() int {
	return xxx_messageInfo_DiscardsRsp.Size(m)
}
func (m *DiscardsRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscardsRsp.DiscardUnknown(m)
}

var xxx_messageInfo_DiscardsRsp proto.InternalMessageInfo

// region 事件
type GameStartEvent struct {
	State                State    `protobuf:"varint,1,opt,name=state,proto3,enum=zlmj.State" json:"state,omitempty"`
	Cards                []int32  `protobuf:"varint,2,rep,packed,name=cards,proto3" json:"cards,omitempty"`
	DrawCount            int32    `protobuf:"varint,3,opt,name=drawCount,proto3" json:"drawCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameStartEvent) Reset()         { *m = GameStartEvent{} }
func (m *GameStartEvent) String() string { return proto.CompactTextString(m) }
func (*GameStartEvent) ProtoMessage()    {}
func (*GameStartEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{10}
}

func (m *GameStartEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStartEvent.Unmarshal(m, b)
}
func (m *GameStartEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStartEvent.Marshal(b, m, deterministic)
}
func (m *GameStartEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStartEvent.Merge(m, src)
}
func (m *GameStartEvent) XXX_Size() int {
	return xxx_messageInfo_GameStartEvent.Size(m)
}
func (m *GameStartEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStartEvent.DiscardUnknown(m)
}

var xxx_messageInfo_GameStartEvent proto.InternalMessageInfo

func (m *GameStartEvent) GetState() State {
	if m != nil {
		return m.State
	}
	return State_initing
}

func (m *GameStartEvent) GetCards() []int32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *GameStartEvent) GetDrawCount() int32 {
	if m != nil {
		return m.DrawCount
	}
	return 0
}

// 玩家打牌事件
type DiscardsEvent struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscardsEvent) Reset()         { *m = DiscardsEvent{} }
func (m *DiscardsEvent) String() string { return proto.CompactTextString(m) }
func (*DiscardsEvent) ProtoMessage()    {}
func (*DiscardsEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{11}
}

func (m *DiscardsEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscardsEvent.Unmarshal(m, b)
}
func (m *DiscardsEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscardsEvent.Marshal(b, m, deterministic)
}
func (m *DiscardsEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscardsEvent.Merge(m, src)
}
func (m *DiscardsEvent) XXX_Size() int {
	return xxx_messageInfo_DiscardsEvent.Size(m)
}
func (m *DiscardsEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscardsEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DiscardsEvent proto.InternalMessageInfo

func (m *DiscardsEvent) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterEnum("zlmj.State", State_name, State_value)
	proto.RegisterType((*Option)(nil), "zlmj.Option")
	proto.RegisterType((*GameData)(nil), "zlmj.GameData")
	proto.RegisterType((*PlayerData)(nil), "zlmj.PlayerData")
	proto.RegisterType((*Req)(nil), "zlmj.Req")
	proto.RegisterType((*Rsp)(nil), "zlmj.Rsp")
	proto.RegisterType((*Event)(nil), "zlmj.Event")
	proto.RegisterType((*WantHoldsReq)(nil), "zlmj.WantHoldsReq")
	proto.RegisterType((*DiscardsReq)(nil), "zlmj.DiscardsReq")
	proto.RegisterType((*ErrorRsp)(nil), "zlmj.ErrorRsp")
	proto.RegisterType((*DiscardsRsp)(nil), "zlmj.DiscardsRsp")
	proto.RegisterType((*GameStartEvent)(nil), "zlmj.GameStartEvent")
	proto.RegisterType((*DiscardsEvent)(nil), "zlmj.DiscardsEvent")
}

func init() { proto.RegisterFile("zlmj/zlmj.proto", fileDescriptor_ba28b2cf51ceff1b) }

var fileDescriptor_ba28b2cf51ceff1b = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x93, 0xe6, 0xd2, 0xc6, 0x49, 0xdb, 0xab, 0xe3, 0x21, 0x79, 0x10, 0xa9, 0xcb, 0x09,
	0xe5, 0x90, 0x3b, 0xa8, 0x08, 0xa2, 0xe0, 0xc3, 0x79, 0xc7, 0xc5, 0x27, 0x65, 0xef, 0xc1, 0x47,
	0x59, 0x9b, 0xa5, 0x46, 0xda, 0x24, 0xdd, 0xdd, 0xbb, 0x53, 0x41, 0xfc, 0x04, 0x7e, 0x67, 0xd9,
	0xd9, 0xf4, 0x9a, 0xed, 0x93, 0x2f, 0x25, 0x33, 0xff, 0xff, 0xcc, 0x6f, 0x3a, 0x99, 0xc0, 0xe1,
	0xaf, 0xd5, 0xfa, 0xfb, 0x99, 0xfd, 0x39, 0x6d, 0x54, 0x6d, 0x6a, 0x3c, 0xb0, 0xcf, 0xec, 0x29,
	0xf4, 0x3f, 0x36, 0xa6, 0xac, 0x2b, 0x3c, 0x82, 0x78, 0xf1, 0x4d, 0x0a, 0x93, 0x85, 0xd3, 0x70,
	0x96, 0x70, 0x17, 0xb0, 0x3f, 0x90, 0x5c, 0x89, 0xb5, 0xbc, 0x10, 0x46, 0xe0, 0x33, 0x88, 0xb5,
	0x11, 0x46, 0x92, 0x63, 0x3c, 0x4f, 0x4f, 0xa9, 0xdb, 0xb5, 0x4d, 0x71, 0xa7, 0xe0, 0x31, 0xf4,
	0x6b, 0x6a, 0x97, 0xf5, 0xa6, 0xe1, 0x2c, 0x9d, 0x0f, 0x9d, 0xc7, 0x21, 0x78, 0xab, 0xe1, 0x09,
	0x0c, 0x9a, 0x95, 0xf8, 0x29, 0x95, 0xce, 0xa2, 0x69, 0x34, 0x4b, 0xe7, 0x13, 0x67, 0xfb, 0x44,
	0x49, 0xcb, 0xe2, 0x5b, 0x03, 0x7b, 0x03, 0xb0, 0x4b, 0xe3, 0x63, 0xe8, 0xdf, 0x68, 0xa9, 0x3e,
	0x14, 0x34, 0x43, 0xcc, 0xdb, 0xc8, 0x0e, 0x5f, 0x56, 0x85, 0xfc, 0x41, 0xd8, 0x98, 0xbb, 0x80,
	0xfd, 0x86, 0x88, 0xcb, 0x0d, 0xbe, 0x82, 0xb4, 0x28, 0xf5, 0x42, 0xa8, 0x42, 0x73, 0xb9, 0xa1,
	0xca, 0x74, 0xfe, 0xd0, 0x21, 0x2f, 0x76, 0x42, 0x1e, 0xf0, 0xae, 0x0f, 0x5f, 0xc3, 0xf0, 0x4e,
	0x54, 0x26, 0xaf, 0x57, 0xae, 0xce, 0xfd, 0x23, 0x74, 0x75, 0x9f, 0x3b, 0x4a, 0x1e, 0x70, 0xcf,
	0x79, 0x1e, 0x43, 0xa4, 0xe4, 0x86, 0x6d, 0x20, 0xe2, 0xba, 0xc1, 0x17, 0x90, 0x48, 0xa5, 0x6a,
	0xc5, 0x75, 0xd3, 0xb2, 0xc7, 0xae, 0xc7, 0x65, 0x9b, 0xcd, 0x03, 0x7e, 0xef, 0xf0, 0x86, 0xd5,
	0x4d, 0x0b, 0xdd, 0x1f, 0x96, 0x6a, 0xba, 0x3e, 0x42, 0xea, 0x86, 0xfd, 0x0d, 0x21, 0xbe, 0xbc,
	0x95, 0x95, 0xc1, 0xb7, 0x30, 0xda, 0xea, 0x94, 0x68, 0xd1, 0x8f, 0xfc, 0x4e, 0x24, 0xe5, 0x01,
	0xf7, 0xbd, 0xf8, 0x0e, 0xc6, 0x4b, 0xb1, 0x96, 0xd7, 0x46, 0x28, 0xe3, 0xaa, 0xdd, 0x1c, 0x47,
	0xae, 0xfa, 0xca, 0xd3, 0xf2, 0x80, 0xef, 0xb9, 0xcf, 0x07, 0x10, 0x4b, 0xfb, 0xc0, 0x8e, 0x61,
	0xd8, 0xdd, 0x14, 0x1d, 0x99, 0xc5, 0x64, 0xe1, 0x34, 0xb2, 0xef, 0x89, 0x02, 0x36, 0x82, 0xb4,
	0xf3, 0x1e, 0x18, 0x40, 0xb2, 0x5d, 0x8d, 0x27, 0xe9, 0x86, 0x2d, 0x61, 0xec, 0xc3, 0xff, 0xe7,
	0x28, 0xef, 0xa1, 0xbd, 0x0e, 0x14, 0x9f, 0xc0, 0x83, 0x42, 0x89, 0xbb, 0xf7, 0xf5, 0x4d, 0x65,
	0xb2, 0x88, 0xce, 0x66, 0x97, 0x60, 0xcf, 0x61, 0xe4, 0xed, 0x68, 0x77, 0x61, 0x61, 0xe7, 0xc2,
	0x4e, 0xce, 0x20, 0x26, 0x14, 0xa6, 0x30, 0x28, 0xab, 0xd2, 0x94, 0xd5, 0x72, 0x12, 0xe0, 0x21,
	0xa4, 0x0b, 0x51, 0x7d, 0x69, 0x77, 0x3a, 0x09, 0x31, 0x81, 0x83, 0xfa, 0x56, 0xaa, 0x49, 0xef,
	0x6b, 0x9f, 0x3e, 0xbe, 0x97, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x23, 0x98, 0x19, 0xbe, 0x8f,
	0x03, 0x00, 0x00,
}

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
	State_cheating    State = 1
	State_can_discard State = 2
	State_over        State = 3
)

var State_name = map[int32]string{
	0: "initing",
	1: "cheating",
	2: "can_discard",
	3: "over",
}

var State_value = map[string]int32{
	"initing":     0,
	"cheating":    1,
	"can_discard": 2,
	"over":        3,
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
	BackCardCount        int32         `protobuf:"varint,4,opt,name=backCardCount,proto3" json:"backCardCount,omitempty"`
	Cards                []int32       `protobuf:"varint,5,rep,packed,name=cards,proto3" json:"cards,omitempty"`
	LastDisCard          int32         `protobuf:"varint,6,opt,name=lastDisCard,proto3" json:"lastDisCard,omitempty"`
	CurDiscardIndex      int32         `protobuf:"varint,7,opt,name=curDiscardIndex,proto3" json:"curDiscardIndex,omitempty"`
	LastDiscardIndex     int32         `protobuf:"varint,8,opt,name=lastDiscardIndex,proto3" json:"lastDiscardIndex,omitempty"`
	CardCount            int32         `protobuf:"varint,9,opt,name=cardCount,proto3" json:"cardCount,omitempty"`
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

func (m *GameData) GetBackCardCount() int32 {
	if m != nil {
		return m.BackCardCount
	}
	return 0
}

func (m *GameData) GetCards() []int32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *GameData) GetLastDisCard() int32 {
	if m != nil {
		return m.LastDisCard
	}
	return 0
}

func (m *GameData) GetCurDiscardIndex() int32 {
	if m != nil {
		return m.CurDiscardIndex
	}
	return 0
}

func (m *GameData) GetLastDiscardIndex() int32 {
	if m != nil {
		return m.LastDiscardIndex
	}
	return 0
}

func (m *GameData) GetCardCount() int32 {
	if m != nil {
		return m.CardCount
	}
	return 0
}

// 玩家数据结构
type PlayerData struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Index                int32    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	HoldsCount           int32    `protobuf:"varint,3,opt,name=holdsCount,proto3" json:"holdsCount,omitempty"`
	Holds                []int32  `protobuf:"varint,4,rep,packed,name=holds,proto3" json:"holds,omitempty"`
	Folds                []int32  `protobuf:"varint,5,rep,packed,name=folds,proto3" json:"folds,omitempty"`
	CanChupai            bool     `protobuf:"varint,6,opt,name=canChupai,proto3" json:"canChupai,omitempty"`
	Score                int32    `protobuf:"varint,7,opt,name=score,proto3" json:"score,omitempty"`
	TotalScore           int32    `protobuf:"varint,8,opt,name=totalScore,proto3" json:"totalScore,omitempty"`
	IsMaster             bool     `protobuf:"varint,9,opt,name=isMaster,proto3" json:"isMaster,omitempty"`
	CurCard              int32    `protobuf:"varint,10,opt,name=curCard,proto3" json:"curCard,omitempty"`
	DisableCards         []int32  `protobuf:"varint,11,rep,packed,name=disableCards,proto3" json:"disableCards,omitempty"`
	IsReportTing         []bool   `protobuf:"varint,12,rep,packed,name=isReportTing,proto3" json:"isReportTing,omitempty"`
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

func (m *PlayerData) GetHoldsCount() int32 {
	if m != nil {
		return m.HoldsCount
	}
	return 0
}

func (m *PlayerData) GetHolds() []int32 {
	if m != nil {
		return m.Holds
	}
	return nil
}

func (m *PlayerData) GetFolds() []int32 {
	if m != nil {
		return m.Folds
	}
	return nil
}

func (m *PlayerData) GetCanChupai() bool {
	if m != nil {
		return m.CanChupai
	}
	return false
}

func (m *PlayerData) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *PlayerData) GetTotalScore() int32 {
	if m != nil {
		return m.TotalScore
	}
	return 0
}

func (m *PlayerData) GetIsMaster() bool {
	if m != nil {
		return m.IsMaster
	}
	return false
}

func (m *PlayerData) GetCurCard() int32 {
	if m != nil {
		return m.CurCard
	}
	return 0
}

func (m *PlayerData) GetDisableCards() []int32 {
	if m != nil {
		return m.DisableCards
	}
	return nil
}

func (m *PlayerData) GetIsReportTing() []bool {
	if m != nil {
		return m.IsReportTing
	}
	return nil
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
	//	*Event_DealCardsEvent
	//	*Event_ResumeEvent
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

type Event_DealCardsEvent struct {
	DealCardsEvent *DealCardsEvent `protobuf:"bytes,3,opt,name=dealCardsEvent,proto3,oneof"`
}

type Event_ResumeEvent struct {
	ResumeEvent *ResumeEvent `protobuf:"bytes,4,opt,name=resumeEvent,proto3,oneof"`
}

func (*Event_DiscardsEvent) isEvent_Event() {}

func (*Event_GameStartEvent) isEvent_Event() {}

func (*Event_DealCardsEvent) isEvent_Event() {}

func (*Event_ResumeEvent) isEvent_Event() {}

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

func (m *Event) GetDealCardsEvent() *DealCardsEvent {
	if x, ok := m.GetEvent().(*Event_DealCardsEvent); ok {
		return x.DealCardsEvent
	}
	return nil
}

func (m *Event) GetResumeEvent() *ResumeEvent {
	if x, ok := m.GetEvent().(*Event_ResumeEvent); ok {
		return x.ResumeEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_DiscardsEvent)(nil),
		(*Event_GameStartEvent)(nil),
		(*Event_DealCardsEvent)(nil),
		(*Event_ResumeEvent)(nil),
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
	Card                 int32    `protobuf:"varint,1,opt,name=card,proto3" json:"card,omitempty"`
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

func (m *DiscardsReq) GetCard() int32 {
	if m != nil {
		return m.Card
	}
	return 0
}

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
	Card                 int32    `protobuf:"varint,2,opt,name=card,proto3" json:"card,omitempty"`
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

func (m *DiscardsEvent) GetCard() int32 {
	if m != nil {
		return m.Card
	}
	return 0
}

// 发牌事件
type DealCardsEvent struct {
	Players              []*PlayerData `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DealCardsEvent) Reset()         { *m = DealCardsEvent{} }
func (m *DealCardsEvent) String() string { return proto.CompactTextString(m) }
func (*DealCardsEvent) ProtoMessage()    {}
func (*DealCardsEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{12}
}

func (m *DealCardsEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DealCardsEvent.Unmarshal(m, b)
}
func (m *DealCardsEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DealCardsEvent.Marshal(b, m, deterministic)
}
func (m *DealCardsEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DealCardsEvent.Merge(m, src)
}
func (m *DealCardsEvent) XXX_Size() int {
	return xxx_messageInfo_DealCardsEvent.Size(m)
}
func (m *DealCardsEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DealCardsEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DealCardsEvent proto.InternalMessageInfo

func (m *DealCardsEvent) GetPlayers() []*PlayerData {
	if m != nil {
		return m.Players
	}
	return nil
}

// 轮到某人打牌事件
type ResumeEvent struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	DiscardLeft          int32    `protobuf:"varint,2,opt,name=discardLeft,proto3" json:"discardLeft,omitempty"`
	IsNewRound           bool     `protobuf:"varint,3,opt,name=isNewRound,proto3" json:"isNewRound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeEvent) Reset()         { *m = ResumeEvent{} }
func (m *ResumeEvent) String() string { return proto.CompactTextString(m) }
func (*ResumeEvent) ProtoMessage()    {}
func (*ResumeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba28b2cf51ceff1b, []int{13}
}

func (m *ResumeEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeEvent.Unmarshal(m, b)
}
func (m *ResumeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeEvent.Marshal(b, m, deterministic)
}
func (m *ResumeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeEvent.Merge(m, src)
}
func (m *ResumeEvent) XXX_Size() int {
	return xxx_messageInfo_ResumeEvent.Size(m)
}
func (m *ResumeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeEvent proto.InternalMessageInfo

func (m *ResumeEvent) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ResumeEvent) GetDiscardLeft() int32 {
	if m != nil {
		return m.DiscardLeft
	}
	return 0
}

func (m *ResumeEvent) GetIsNewRound() bool {
	if m != nil {
		return m.IsNewRound
	}
	return false
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
	proto.RegisterType((*DealCardsEvent)(nil), "zlmj.DealCardsEvent")
	proto.RegisterType((*ResumeEvent)(nil), "zlmj.ResumeEvent")
}

func init() { proto.RegisterFile("zlmj/zlmj.proto", fileDescriptor_ba28b2cf51ceff1b) }

var fileDescriptor_ba28b2cf51ceff1b = []byte{
	// 766 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdf, 0x6b, 0xdb, 0x48,
	0x10, 0xb6, 0x24, 0xcb, 0x56, 0x46, 0xb6, 0xe3, 0xdb, 0x0b, 0x87, 0x38, 0x42, 0x50, 0x44, 0x1e,
	0x4c, 0x38, 0x72, 0xe0, 0xe3, 0xe0, 0x8e, 0xfb, 0xf1, 0x10, 0x3b, 0xc4, 0x81, 0xbb, 0xb6, 0x6c,
	0x0a, 0x7d, 0x2c, 0x1b, 0x6b, 0xe3, 0xa8, 0xb5, 0x25, 0x79, 0x77, 0x9d, 0xb4, 0x85, 0xbe, 0xf5,
	0x1f, 0xec, 0x5b, 0xff, 0x9c, 0xb2, 0xb3, 0xb2, 0xb5, 0x4a, 0x29, 0xed, 0x8b, 0xd1, 0x7c, 0xf3,
	0x7d, 0x3b, 0xb3, 0xdf, 0xc8, 0x23, 0xd8, 0x7f, 0xb7, 0x5c, 0xbd, 0xfa, 0x55, 0xff, 0x9c, 0x95,
	0xa2, 0x50, 0x05, 0x69, 0xeb, 0xe7, 0xe4, 0x08, 0x3a, 0x4f, 0x4b, 0x95, 0x15, 0x39, 0x39, 0x00,
	0x7f, 0x7e, 0xc7, 0x99, 0x8a, 0x9c, 0xd8, 0x19, 0x05, 0xd4, 0x04, 0xc9, 0x47, 0x17, 0x82, 0x4b,
	0xb6, 0xe2, 0x53, 0xa6, 0x18, 0x39, 0x06, 0x5f, 0x2a, 0xa6, 0x38, 0x52, 0x06, 0xe3, 0xf0, 0x0c,
	0x8f, 0xbb, 0xd6, 0x10, 0x35, 0x19, 0x72, 0x02, 0x9d, 0x02, 0xcf, 0x8b, 0xdc, 0xd8, 0x19, 0x85,
	0xe3, 0x9e, 0xe1, 0x98, 0x1a, 0xb4, 0xca, 0x91, 0x53, 0xe8, 0x96, 0x4b, 0xf6, 0x96, 0x0b, 0x19,
	0x79, 0xb1, 0x37, 0x0a, 0xc7, 0x43, 0x43, 0x7b, 0x86, 0xa0, 0xae, 0x45, 0xb7, 0x04, 0x72, 0x02,
	0xfd, 0x1b, 0x36, 0x7f, 0x3d, 0x61, 0x22, 0x9d, 0x14, 0x9b, 0x5c, 0x45, 0xed, 0xd8, 0x19, 0xf9,
	0xb4, 0x09, 0x62, 0xf7, 0x4c, 0xa4, 0x32, 0xf2, 0x63, 0x6f, 0xe4, 0x53, 0x13, 0x90, 0x18, 0xc2,
	0x25, 0x93, 0x6a, 0x9a, 0x49, 0xcd, 0x8c, 0x3a, 0xa8, 0xb4, 0x21, 0x32, 0x82, 0xfd, 0xf9, 0x46,
	0x4c, 0x33, 0xa9, 0x05, 0x57, 0x79, 0xca, 0xdf, 0x44, 0x5d, 0x64, 0x3d, 0x86, 0xc9, 0x29, 0x0c,
	0x2b, 0x61, 0x4d, 0x0d, 0x90, 0xfa, 0x05, 0x4e, 0x0e, 0x61, 0x6f, 0xbe, 0xeb, 0x77, 0x0f, 0x49,
	0x35, 0x90, 0x7c, 0x72, 0x01, 0xea, 0x9b, 0x92, 0x9f, 0xa0, 0xb3, 0x91, 0x5c, 0x5c, 0xa5, 0x68,
	0xab, 0x4f, 0xab, 0x48, 0x5f, 0x29, 0xc3, 0x2a, 0x2e, 0xc2, 0x26, 0x20, 0x47, 0x00, 0x77, 0xc5,
	0x32, 0x95, 0xe6, 0x6c, 0x0f, 0x53, 0x16, 0xa2, 0x55, 0x18, 0x45, 0x6d, 0x63, 0x04, 0x06, 0x1a,
	0xbd, 0x45, 0xb4, 0xb2, 0x07, 0x03, 0xd3, 0x66, 0x3e, 0xb9, 0xdb, 0x94, 0x2c, 0x43, 0x73, 0x02,
	0x5a, 0x03, 0x5a, 0x23, 0xe7, 0x85, 0xe0, 0x95, 0x21, 0x26, 0xd0, 0xf5, 0x55, 0xa1, 0xd8, 0xf2,
	0x1a, 0x53, 0xc6, 0x00, 0x0b, 0x21, 0x3f, 0x43, 0x90, 0xc9, 0xff, 0x99, 0x54, 0x5c, 0xe0, 0xcd,
	0x03, 0xba, 0x8b, 0x49, 0x04, 0xdd, 0xf9, 0x46, 0xe0, 0x28, 0x00, 0x85, 0xdb, 0x90, 0x24, 0xd0,
	0x4b, 0x33, 0xc9, 0x6e, 0x96, 0x7c, 0x82, 0x53, 0x0c, 0xb1, 0xcd, 0x06, 0xa6, 0x39, 0x99, 0xa4,
	0xbc, 0x2c, 0x84, 0x7a, 0x9e, 0xe5, 0x8b, 0xa8, 0x17, 0x7b, 0xa3, 0x80, 0x36, 0xb0, 0xe4, 0x3d,
	0x78, 0x94, 0xaf, 0xc9, 0xef, 0x10, 0xa6, 0x66, 0x1e, 0x92, 0xf2, 0x35, 0xfa, 0x1a, 0x8e, 0x7f,
	0x30, 0xef, 0xd8, 0xb4, 0x4e, 0xcc, 0x5a, 0xd4, 0xe6, 0x91, 0x3f, 0xa0, 0xf7, 0xc0, 0x72, 0x35,
	0xd3, 0xe6, 0x68, 0x9d, 0x79, 0x85, 0x89, 0xd1, 0xbd, 0xb0, 0x32, 0xb3, 0x16, 0x6d, 0x30, 0xcf,
	0x7d, 0xf0, 0x04, 0x5f, 0x27, 0x6b, 0xf0, 0xa8, 0x2c, 0xc9, 0x2f, 0x10, 0x70, 0x21, 0x0a, 0x41,
	0x65, 0x59, 0xd5, 0x1e, 0x98, 0x33, 0x2e, 0x2a, 0x74, 0xd6, 0xa2, 0x3b, 0x46, 0xa3, 0x59, 0x59,
	0x56, 0x45, 0x1f, 0x37, 0x8b, 0x1a, 0x9b, 0x87, 0x25, 0x65, 0x99, 0x7c, 0x70, 0xc1, 0xbf, 0xb8,
	0xe7, 0xb9, 0x22, 0x7f, 0x41, 0x7f, 0x9b, 0x47, 0xa0, 0x2a, 0xfd, 0x63, 0xf3, 0x24, 0x4c, 0xcd,
	0x5a, 0xb4, 0xc9, 0x25, 0xff, 0xc2, 0x60, 0xc1, 0x56, 0xfc, 0x5a, 0x31, 0xa1, 0x8c, 0xda, 0xf4,
	0x71, 0x60, 0xd4, 0x97, 0x8d, 0xdc, 0xac, 0x45, 0x1f, 0xb1, 0xb5, 0x3e, 0xe5, 0x6c, 0x39, 0xa9,
	0xab, 0x7b, 0xb6, 0x7e, 0xda, 0xc8, 0x69, 0x7d, 0x93, 0xad, 0x4d, 0x10, 0x5c, 0x6e, 0x56, 0xdc,
	0x88, 0xdb, 0xb6, 0x09, 0xb4, 0x4e, 0x68, 0x13, 0x2c, 0xde, 0x79, 0x17, 0x7c, 0xae, 0x1f, 0x92,
	0x13, 0xe8, 0xd9, 0x03, 0xaa, 0xf7, 0x81, 0x63, 0xed, 0x83, 0xe4, 0x18, 0x42, 0x6b, 0xfc, 0x84,
	0x40, 0x5b, 0x3f, 0x57, 0xff, 0x3b, 0x7c, 0x4e, 0x00, 0x82, 0xed, 0x94, 0x92, 0xbe, 0x45, 0x97,
	0x65, 0xb2, 0x80, 0x41, 0xd3, 0x87, 0xef, 0x59, 0x88, 0xbb, 0x46, 0x5c, 0x7b, 0x31, 0x1d, 0xc2,
	0x5e, 0x2a, 0xd8, 0x83, 0xfd, 0x27, 0xae, 0x81, 0xe4, 0x4f, 0xe8, 0x37, 0xc6, 0x55, 0xaf, 0x02,
	0xc7, 0x5e, 0x05, 0xdb, 0xf6, 0x5d, 0xab, 0xfd, 0xbf, 0x61, 0xd0, 0xf4, 0xda, 0xde, 0xb5, 0xce,
	0x37, 0x76, 0x6d, 0xc2, 0x21, 0xb4, 0xcc, 0xfe, 0x4a, 0xd9, 0x78, 0xf7, 0xbe, 0xfe, 0xc7, 0x6f,
	0x55, 0x55, 0xdd, 0x86, 0xf4, 0x8e, 0xc8, 0xe4, 0x13, 0xfe, 0x40, 0x8b, 0x4d, 0x9e, 0xe2, 0xf5,
	0x02, 0x6a, 0x21, 0xa7, 0xff, 0x80, 0x8f, 0x1e, 0x91, 0x10, 0xba, 0x59, 0x9e, 0xa9, 0x2c, 0x5f,
	0x0c, 0x5b, 0xa4, 0x07, 0x01, 0x7e, 0x73, 0x74, 0xe4, 0x90, 0x7d, 0x08, 0xe7, 0x2c, 0x7f, 0x59,
	0x1d, 0x3b, 0x74, 0x49, 0x00, 0xed, 0xe2, 0x9e, 0x8b, 0xa1, 0x77, 0xd3, 0xc1, 0x0f, 0xd8, 0x6f,
	0x9f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x24, 0x06, 0x97, 0xb0, 0xd3, 0x06, 0x00, 0x00,
}

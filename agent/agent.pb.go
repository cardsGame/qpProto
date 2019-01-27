// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/agent/agent.proto

package agent

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

type ChangeAvaliableRequest_Types int32

const (
	ChangeAvaliableRequest_Unknonwn  ChangeAvaliableRequest_Types = 0
	ChangeAvaliableRequest_Blocked   ChangeAvaliableRequest_Types = 1
	ChangeAvaliableRequest_Avaliable ChangeAvaliableRequest_Types = 2
)

var ChangeAvaliableRequest_Types_name = map[int32]string{
	0: "Unknonwn",
	1: "Blocked",
	2: "Avaliable",
}

var ChangeAvaliableRequest_Types_value = map[string]int32{
	"Unknonwn":  0,
	"Blocked":   1,
	"Avaliable": 2,
}

func (x ChangeAvaliableRequest_Types) String() string {
	return proto.EnumName(ChangeAvaliableRequest_Types_name, int32(x))
}

func (ChangeAvaliableRequest_Types) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_656b6c96a18ce683, []int{5, 0}
}

// Request of Creating an agent.
// Prerequisite: User Must be already Created in Game Server
type CreateAgentRequest struct {
	NickName             string   `protobuf:"bytes,1,opt,name=NickName,proto3" json:"NickName,omitempty"`
	WxAccount            string   `protobuf:"bytes,2,opt,name=WxAccount,proto3" json:"WxAccount,omitempty"`
	UserID               int32    `protobuf:"varint,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	HeadImg              string   `protobuf:"bytes,5,opt,name=HeadImg,proto3" json:"HeadImg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAgentRequest) Reset()         { *m = CreateAgentRequest{} }
func (m *CreateAgentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAgentRequest) ProtoMessage()    {}
func (*CreateAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_656b6c96a18ce683, []int{0}
}

func (m *CreateAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAgentRequest.Unmarshal(m, b)
}
func (m *CreateAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAgentRequest.Marshal(b, m, deterministic)
}
func (m *CreateAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAgentRequest.Merge(m, src)
}
func (m *CreateAgentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAgentRequest.Size(m)
}
func (m *CreateAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAgentRequest proto.InternalMessageInfo

func (m *CreateAgentRequest) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *CreateAgentRequest) GetWxAccount() string {
	if m != nil {
		return m.WxAccount
	}
	return ""
}

func (m *CreateAgentRequest) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CreateAgentRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CreateAgentRequest) GetHeadImg() string {
	if m != nil {
		return m.HeadImg
	}
	return ""
}

// Request of getting an angent's personal info.
type GetInfoByUserIDRequest struct {
	UserID               int32    `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoByUserIDRequest) Reset()         { *m = GetInfoByUserIDRequest{} }
func (m *GetInfoByUserIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoByUserIDRequest) ProtoMessage()    {}
func (*GetInfoByUserIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_656b6c96a18ce683, []int{1}
}

func (m *GetInfoByUserIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoByUserIDRequest.Unmarshal(m, b)
}
func (m *GetInfoByUserIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoByUserIDRequest.Marshal(b, m, deterministic)
}
func (m *GetInfoByUserIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoByUserIDRequest.Merge(m, src)
}
func (m *GetInfoByUserIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoByUserIDRequest.Size(m)
}
func (m *GetInfoByUserIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoByUserIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoByUserIDRequest proto.InternalMessageInfo

func (m *GetInfoByUserIDRequest) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

// Request of getting an angent's personal info.
type GetInfoByIDRequest struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoByIDRequest) Reset()         { *m = GetInfoByIDRequest{} }
func (m *GetInfoByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoByIDRequest) ProtoMessage()    {}
func (*GetInfoByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_656b6c96a18ce683, []int{2}
}

func (m *GetInfoByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoByIDRequest.Unmarshal(m, b)
}
func (m *GetInfoByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetInfoByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoByIDRequest.Merge(m, src)
}
func (m *GetInfoByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoByIDRequest.Size(m)
}
func (m *GetInfoByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoByIDRequest proto.InternalMessageInfo

func (m *GetInfoByIDRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetAgentListRequest struct {
	Order                string   `protobuf:"bytes,1,opt,name=Order,proto3" json:"Order,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset               int32    `protobuf:"varint,3,opt,name=Offset,proto3" json:"Offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgentListRequest) Reset()         { *m = GetAgentListRequest{} }
func (m *GetAgentListRequest) String() string { return proto.CompactTextString(m) }
func (*GetAgentListRequest) ProtoMessage()    {}
func (*GetAgentListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_656b6c96a18ce683, []int{3}
}

func (m *GetAgentListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentListRequest.Unmarshal(m, b)
}
func (m *GetAgentListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentListRequest.Marshal(b, m, deterministic)
}
func (m *GetAgentListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentListRequest.Merge(m, src)
}
func (m *GetAgentListRequest) XXX_Size() int {
	return xxx_messageInfo_GetAgentListRequest.Size(m)
}
func (m *GetAgentListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentListRequest proto.InternalMessageInfo

func (m *GetAgentListRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *GetAgentListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetAgentListRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type GetAgentListResponse struct {
	List                 []*Agent `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgentListResponse) Reset()         { *m = GetAgentListResponse{} }
func (m *GetAgentListResponse) String() string { return proto.CompactTextString(m) }
func (*GetAgentListResponse) ProtoMessage()    {}
func (*GetAgentListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_656b6c96a18ce683, []int{4}
}

func (m *GetAgentListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentListResponse.Unmarshal(m, b)
}
func (m *GetAgentListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentListResponse.Marshal(b, m, deterministic)
}
func (m *GetAgentListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentListResponse.Merge(m, src)
}
func (m *GetAgentListResponse) XXX_Size() int {
	return xxx_messageInfo_GetAgentListResponse.Size(m)
}
func (m *GetAgentListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentListResponse proto.InternalMessageInfo

func (m *GetAgentListResponse) GetList() []*Agent {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *GetAgentListResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ChangeAvaliableRequest struct {
	AgentID              int32                        `protobuf:"varint,1,opt,name=AgentID,proto3" json:"AgentID,omitempty"`
	Type                 ChangeAvaliableRequest_Types `protobuf:"varint,2,opt,name=Type,proto3,enum=agent.ChangeAvaliableRequest_Types" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ChangeAvaliableRequest) Reset()         { *m = ChangeAvaliableRequest{} }
func (m *ChangeAvaliableRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeAvaliableRequest) ProtoMessage()    {}
func (*ChangeAvaliableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_656b6c96a18ce683, []int{5}
}

func (m *ChangeAvaliableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeAvaliableRequest.Unmarshal(m, b)
}
func (m *ChangeAvaliableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeAvaliableRequest.Marshal(b, m, deterministic)
}
func (m *ChangeAvaliableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAvaliableRequest.Merge(m, src)
}
func (m *ChangeAvaliableRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeAvaliableRequest.Size(m)
}
func (m *ChangeAvaliableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAvaliableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAvaliableRequest proto.InternalMessageInfo

func (m *ChangeAvaliableRequest) GetAgentID() int32 {
	if m != nil {
		return m.AgentID
	}
	return 0
}

func (m *ChangeAvaliableRequest) GetType() ChangeAvaliableRequest_Types {
	if m != nil {
		return m.Type
	}
	return ChangeAvaliableRequest_Unknonwn
}

type ChangeAvaliableResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeAvaliableResponse) Reset()         { *m = ChangeAvaliableResponse{} }
func (m *ChangeAvaliableResponse) String() string { return proto.CompactTextString(m) }
func (*ChangeAvaliableResponse) ProtoMessage()    {}
func (*ChangeAvaliableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_656b6c96a18ce683, []int{6}
}

func (m *ChangeAvaliableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeAvaliableResponse.Unmarshal(m, b)
}
func (m *ChangeAvaliableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeAvaliableResponse.Marshal(b, m, deterministic)
}
func (m *ChangeAvaliableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAvaliableResponse.Merge(m, src)
}
func (m *ChangeAvaliableResponse) XXX_Size() int {
	return xxx_messageInfo_ChangeAvaliableResponse.Size(m)
}
func (m *ChangeAvaliableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAvaliableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAvaliableResponse proto.InternalMessageInfo

func (m *ChangeAvaliableResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

// Information of single agent
type Agent struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	NickName             string   `protobuf:"bytes,3,opt,name=NickName,proto3" json:"NickName,omitempty"`
	HeadImg              string   `protobuf:"bytes,4,opt,name=HeadImg,proto3" json:"HeadImg,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Avaliable            bool     `protobuf:"varint,6,opt,name=Avaliable,proto3" json:"Avaliable,omitempty"`
	Balance              float64  `protobuf:"fixed64,7,opt,name=Balance,proto3" json:"Balance,omitempty"`
	WxAccount            string   `protobuf:"bytes,8,opt,name=WxAccount,proto3" json:"WxAccount,omitempty"`
	Scode                string   `protobuf:"bytes,9,opt,name=Scode,proto3" json:"Scode,omitempty"`
	AgentLevel           int32    `protobuf:"varint,10,opt,name=AgentLevel,proto3" json:"AgentLevel,omitempty"`
	WithdrawPassword     string   `protobuf:"bytes,11,opt,name=WithdrawPassword,proto3" json:"WithdrawPassword,omitempty"`
	PayPassword          string   `protobuf:"bytes,12,opt,name=PayPassword,proto3" json:"PayPassword,omitempty"`
	UserID               int32    `protobuf:"varint,13,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Father               *Agent   `protobuf:"bytes,14,opt,name=Father,proto3" json:"Father,omitempty"`
	CreatedAt            string   `protobuf:"bytes,15,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,16,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DelatedAt            string   `protobuf:"bytes,17,opt,name=DelatedAt,proto3" json:"DelatedAt,omitempty"`
	Link                 *Link    `protobuf:"bytes,18,opt,name=Link,proto3" json:"Link,omitempty"`
	TotalSon             int32    `protobuf:"varint,19,opt,name=TotalSon,proto3" json:"TotalSon,omitempty"`
	TotalIncome          float64  `protobuf:"fixed64,20,opt,name=TotalIncome,proto3" json:"TotalIncome,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Agent) Reset()         { *m = Agent{} }
func (m *Agent) String() string { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()    {}
func (*Agent) Descriptor() ([]byte, []int) {
	return fileDescriptor_656b6c96a18ce683, []int{7}
}

func (m *Agent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Agent.Unmarshal(m, b)
}
func (m *Agent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Agent.Marshal(b, m, deterministic)
}
func (m *Agent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Agent.Merge(m, src)
}
func (m *Agent) XXX_Size() int {
	return xxx_messageInfo_Agent.Size(m)
}
func (m *Agent) XXX_DiscardUnknown() {
	xxx_messageInfo_Agent.DiscardUnknown(m)
}

var xxx_messageInfo_Agent proto.InternalMessageInfo

func (m *Agent) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Agent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Agent) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *Agent) GetHeadImg() string {
	if m != nil {
		return m.HeadImg
	}
	return ""
}

func (m *Agent) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Agent) GetAvaliable() bool {
	if m != nil {
		return m.Avaliable
	}
	return false
}

func (m *Agent) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Agent) GetWxAccount() string {
	if m != nil {
		return m.WxAccount
	}
	return ""
}

func (m *Agent) GetScode() string {
	if m != nil {
		return m.Scode
	}
	return ""
}

func (m *Agent) GetAgentLevel() int32 {
	if m != nil {
		return m.AgentLevel
	}
	return 0
}

func (m *Agent) GetWithdrawPassword() string {
	if m != nil {
		return m.WithdrawPassword
	}
	return ""
}

func (m *Agent) GetPayPassword() string {
	if m != nil {
		return m.PayPassword
	}
	return ""
}

func (m *Agent) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *Agent) GetFather() *Agent {
	if m != nil {
		return m.Father
	}
	return nil
}

func (m *Agent) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Agent) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Agent) GetDelatedAt() string {
	if m != nil {
		return m.DelatedAt
	}
	return ""
}

func (m *Agent) GetLink() *Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *Agent) GetTotalSon() int32 {
	if m != nil {
		return m.TotalSon
	}
	return 0
}

func (m *Agent) GetTotalIncome() float64 {
	if m != nil {
		return m.TotalIncome
	}
	return 0
}

type Link struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FatherID             int32    `protobuf:"varint,2,opt,name=FatherID,proto3" json:"FatherID,omitempty"`
	GrandFatherID        int32    `protobuf:"varint,3,opt,name=GrandFatherID,proto3" json:"GrandFatherID,omitempty"`
	FatherBindAt         string   `protobuf:"bytes,4,opt,name=FatherBindAt,proto3" json:"FatherBindAt,omitempty"`
	GrandFatherBindAt    string   `protobuf:"bytes,5,opt,name=GrandFatherBindAt,proto3" json:"GrandFatherBindAt,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeleatedAt           string   `protobuf:"bytes,8,opt,name=DeleatedAt,proto3" json:"DeleatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_656b6c96a18ce683, []int{8}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Link) GetFatherID() int32 {
	if m != nil {
		return m.FatherID
	}
	return 0
}

func (m *Link) GetGrandFatherID() int32 {
	if m != nil {
		return m.GrandFatherID
	}
	return 0
}

func (m *Link) GetFatherBindAt() string {
	if m != nil {
		return m.FatherBindAt
	}
	return ""
}

func (m *Link) GetGrandFatherBindAt() string {
	if m != nil {
		return m.GrandFatherBindAt
	}
	return ""
}

func (m *Link) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Link) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Link) GetDeleatedAt() string {
	if m != nil {
		return m.DeleatedAt
	}
	return ""
}

func init() {
	proto.RegisterEnum("agent.ChangeAvaliableRequest_Types", ChangeAvaliableRequest_Types_name, ChangeAvaliableRequest_Types_value)
	proto.RegisterType((*CreateAgentRequest)(nil), "agent.CreateAgentRequest")
	proto.RegisterType((*GetInfoByUserIDRequest)(nil), "agent.GetInfoByUserIDRequest")
	proto.RegisterType((*GetInfoByIDRequest)(nil), "agent.GetInfoByIDRequest")
	proto.RegisterType((*GetAgentListRequest)(nil), "agent.GetAgentListRequest")
	proto.RegisterType((*GetAgentListResponse)(nil), "agent.GetAgentListResponse")
	proto.RegisterType((*ChangeAvaliableRequest)(nil), "agent.ChangeAvaliableRequest")
	proto.RegisterType((*ChangeAvaliableResponse)(nil), "agent.ChangeAvaliableResponse")
	proto.RegisterType((*Agent)(nil), "agent.Agent")
	proto.RegisterType((*Link)(nil), "agent.Link")
}

func init() { proto.RegisterFile("proto/agent/agent.proto", fileDescriptor_656b6c96a18ce683) }

var fileDescriptor_656b6c96a18ce683 = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0xc5, 0x21, 0xce, 0xcf, 0x24, 0x40, 0x58, 0xa2, 0xb0, 0x5f, 0x3e, 0x4a, 0x2d, 0x97, 0x8b,
	0xa8, 0xaa, 0x68, 0x81, 0x8b, 0xf6, 0xb2, 0x09, 0x51, 0x69, 0x24, 0x04, 0xc8, 0x80, 0x50, 0x2f,
	0x8d, 0x3d, 0x10, 0x2b, 0xce, 0x3a, 0xb5, 0x0d, 0x94, 0x17, 0xe8, 0x23, 0xf4, 0xb2, 0xcf, 0xd3,
	0xa7, 0xe8, 0xb3, 0x54, 0xbb, 0x6b, 0x6f, 0xec, 0x38, 0xed, 0x4d, 0xe4, 0x73, 0x66, 0x67, 0x66,
	0xcf, 0xee, 0xd9, 0x09, 0x6c, 0xcf, 0xc2, 0x20, 0x0e, 0xde, 0xda, 0xf7, 0xc8, 0x62, 0xf9, 0xbb,
	0x2f, 0x18, 0xa2, 0x0b, 0x60, 0xfe, 0xd0, 0x80, 0x1c, 0x87, 0x68, 0xc7, 0xd8, 0xe7, 0xd8, 0xc2,
	0xaf, 0x0f, 0x18, 0xc5, 0xa4, 0x0b, 0xb5, 0x33, 0xcf, 0x99, 0x9c, 0xd9, 0x53, 0xa4, 0x9a, 0xa1,
	0xf5, 0xea, 0x96, 0xc2, 0x64, 0x07, 0xea, 0x37, 0xdf, 0xfa, 0x8e, 0x13, 0x3c, 0xb0, 0x98, 0x96,
	0x44, 0x70, 0x4e, 0x90, 0x0e, 0x54, 0xae, 0x23, 0x0c, 0x47, 0x43, 0xba, 0x6a, 0x68, 0x3d, 0xdd,
	0x4a, 0x10, 0x69, 0x83, 0x7e, 0x31, 0x0e, 0x18, 0xd2, 0xb2, 0xc8, 0x90, 0x80, 0x50, 0xa8, 0x7e,
	0x46, 0xdb, 0x1d, 0x4d, 0xef, 0xa9, 0x2e, 0xf8, 0x14, 0x9a, 0xef, 0xa0, 0x73, 0x82, 0xf1, 0x88,
	0xdd, 0x05, 0x83, 0x67, 0x59, 0x22, 0xdd, 0xdb, 0xbc, 0x83, 0x96, 0xed, 0x60, 0xee, 0x01, 0x51,
	0x19, 0xf3, 0xd5, 0xeb, 0x50, 0x52, 0x2b, 0x4b, 0xa3, 0xa1, 0xf9, 0x05, 0xb6, 0x4e, 0x30, 0x16,
	0x62, 0x4f, 0xbd, 0x48, 0x09, 0x6e, 0x83, 0x7e, 0x1e, 0xba, 0x18, 0x26, 0x6a, 0x25, 0xe0, 0xec,
	0xa9, 0x37, 0xf5, 0xa4, 0x4c, 0xdd, 0x92, 0x80, 0x6f, 0xe0, 0xfc, 0xee, 0x2e, 0xc2, 0x38, 0x95,
	0x28, 0x91, 0x79, 0x06, 0xed, 0x7c, 0xe9, 0x68, 0x16, 0xb0, 0x08, 0x89, 0x01, 0x65, 0x8e, 0xa9,
	0x66, 0xac, 0xf6, 0x1a, 0x87, 0xcd, 0x7d, 0x79, 0x0d, 0xf2, 0xbc, 0x45, 0x84, 0xf7, 0x39, 0x56,
	0xc7, 0xa9, 0x5b, 0x12, 0x98, 0x3f, 0x35, 0xe8, 0x1c, 0x8f, 0x6d, 0x76, 0x8f, 0xfd, 0x47, 0xdb,
	0xf7, 0xec, 0x5b, 0x1f, 0xd3, 0xed, 0x52, 0xa8, 0x8a, 0x7c, 0x25, 0x2d, 0x85, 0xe4, 0x3d, 0x94,
	0xaf, 0x9e, 0x67, 0x28, 0x2a, 0xad, 0x1f, 0xbe, 0x4a, 0x9a, 0x2d, 0x2f, 0xb3, 0xcf, 0x57, 0x46,
	0x96, 0x48, 0x30, 0x0f, 0x40, 0x17, 0x90, 0x34, 0xa1, 0x76, 0xcd, 0x26, 0x2c, 0x60, 0x4f, 0xac,
	0xb5, 0x42, 0x1a, 0x50, 0x1d, 0xf8, 0x81, 0x33, 0x41, 0xb7, 0xa5, 0x91, 0x35, 0xa8, 0xab, 0x1a,
	0xad, 0x92, 0x79, 0x00, 0xdb, 0x85, 0xc2, 0x89, 0xe6, 0x0e, 0x54, 0x42, 0x8c, 0x1e, 0xfc, 0x38,
	0x39, 0xd0, 0x04, 0x99, 0xbf, 0xca, 0xa0, 0x8b, 0xad, 0x2e, 0x5e, 0x0c, 0x21, 0x50, 0x16, 0x76,
	0x93, 0x8e, 0x12, 0xdf, 0x39, 0x1b, 0xae, 0x2e, 0xd8, 0x30, 0x63, 0x9d, 0x72, 0xce, 0x3a, 0x73,
	0xab, 0xe9, 0x59, 0xab, 0xed, 0x64, 0xf6, 0x4e, 0x2b, 0x86, 0xd6, 0xab, 0x59, 0x73, 0x82, 0x57,
	0x1b, 0xd8, 0xbe, 0xcd, 0x1c, 0xa4, 0x55, 0x43, 0xeb, 0x69, 0x56, 0x0a, 0xf3, 0x76, 0xaf, 0x2d,
	0xda, 0xbd, 0x0d, 0xfa, 0xa5, 0x13, 0xb8, 0x48, 0xeb, 0xb2, 0x97, 0x00, 0x64, 0x17, 0x40, 0xda,
	0x00, 0x1f, 0xd1, 0xa7, 0x20, 0x34, 0x66, 0x18, 0xf2, 0x1a, 0x5a, 0x37, 0x5e, 0x3c, 0x76, 0x43,
	0xfb, 0xe9, 0xc2, 0x8e, 0xa2, 0xa7, 0x20, 0x74, 0x69, 0x43, 0x14, 0x28, 0xf0, 0xc4, 0x80, 0xc6,
	0x85, 0xfd, 0xac, 0x96, 0x35, 0xc5, 0xb2, 0x2c, 0x95, 0x79, 0x10, 0x6b, 0xb9, 0x27, 0xb7, 0x07,
	0x95, 0x4f, 0x76, 0x3c, 0xc6, 0x90, 0xae, 0x1b, 0x5a, 0xc1, 0x79, 0x49, 0x8c, 0xeb, 0x93, 0x03,
	0xc0, 0xed, 0xc7, 0x74, 0x43, 0xea, 0x53, 0x04, 0x8f, 0x5e, 0xcf, 0xdc, 0x24, 0xda, 0x92, 0x51,
	0x45, 0xf0, 0xe8, 0x10, 0xfd, 0x24, 0xba, 0x29, 0xa3, 0x8a, 0x20, 0x2f, 0xb9, 0xef, 0xd9, 0x84,
	0x12, 0xd1, 0xbd, 0x91, 0x74, 0xe7, 0x94, 0x25, 0x02, 0xfc, 0x7a, 0xaf, 0x82, 0xd8, 0xf6, 0x2f,
	0x03, 0x46, 0xb7, 0xc4, 0xd6, 0x15, 0xe6, 0xb2, 0xc5, 0xf7, 0x88, 0x39, 0xc1, 0x14, 0x69, 0x5b,
	0x5c, 0x4a, 0x96, 0x32, 0xbf, 0x97, 0x64, 0xfd, 0x82, 0x93, 0xba, 0x50, 0x93, 0xda, 0x46, 0xc3,
	0xe4, 0x41, 0x29, 0x4c, 0xf6, 0x60, 0xed, 0x24, 0xb4, 0x99, 0xab, 0x16, 0xc8, 0x27, 0x9c, 0x27,
	0x89, 0x09, 0x4d, 0xf9, 0x3d, 0xf0, 0x18, 0x97, 0x26, 0x0d, 0x96, 0xe3, 0xc8, 0x1b, 0xd8, 0xcc,
	0x24, 0x25, 0x0b, 0xa5, 0xe3, 0x8a, 0x81, 0xfc, 0x29, 0x57, 0xfe, 0x79, 0xca, 0xd5, 0xc5, 0x53,
	0xde, 0x05, 0x18, 0xa2, 0x9f, 0x26, 0x4b, 0x0b, 0x66, 0x98, 0xc3, 0xdf, 0x25, 0x68, 0x8a, 0x3b,
	0xbd, 0xc4, 0xf0, 0xd1, 0x73, 0x90, 0x1c, 0x41, 0x45, 0xd6, 0x26, 0xff, 0xa5, 0xef, 0xbf, 0x30,
	0xe2, 0xbb, 0x39, 0x37, 0x98, 0x2b, 0xe4, 0x23, 0x6c, 0x2c, 0x0c, 0x5c, 0xf2, 0x22, 0x59, 0xb2,
	0x7c, 0x10, 0x17, 0x2a, 0x7c, 0x80, 0x46, 0x66, 0x00, 0xab, 0xde, 0xc5, 0xa1, 0x5c, 0xc8, 0x1c,
	0x41, 0x33, 0x3b, 0x39, 0x49, 0x77, 0x9e, 0xba, 0x38, 0xa9, 0xbb, 0xff, 0x2f, 0x8d, 0xc9, 0xb1,
	0x63, 0xae, 0x10, 0x0b, 0x36, 0x16, 0x66, 0x92, 0x92, 0xb1, 0x7c, 0x08, 0x76, 0x77, 0xff, 0x16,
	0x4e, 0x6b, 0xde, 0x56, 0xc4, 0x5f, 0xe6, 0xd1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x74,
	0x24, 0x15, 0x4d, 0x07, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room/room.proto

package room

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

// 游戏类型
type GameType int32

const (
	GameType_none_game GameType = 0
	GameType_zlmj      GameType = 1
	GameType_ddz       GameType = 2
)

var GameType_name = map[int32]string{
	0: "none_game",
	1: "zlmj",
	2: "ddz",
}

var GameType_value = map[string]int32{
	"none_game": 0,
	"zlmj":      1,
	"ddz":       2,
}

func (x GameType) String() string {
	return proto.EnumName(GameType_name, int32(x))
}

func (GameType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b6c5d20c74c34643, []int{0}
}

// 房间类型
type RoomType int32

const (
	RoomType_none_room RoomType = 0
	RoomType_friend    RoomType = 1
	RoomType_gold      RoomType = 2
	RoomType_division  RoomType = 3
	RoomType_match     RoomType = 4
)

var RoomType_name = map[int32]string{
	0: "none_room",
	1: "friend",
	2: "gold",
	3: "division",
	4: "match",
}

var RoomType_value = map[string]int32{
	"none_room": 0,
	"friend":    1,
	"gold":      2,
	"division":  3,
	"match":     4,
}

func (x RoomType) String() string {
	return proto.EnumName(RoomType_name, int32(x))
}

func (RoomType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b6c5d20c74c34643, []int{1}
}

// 房间支付类型
type RoomPay int32

const (
	RoomPay_none_pay   RoomPay = 0
	RoomPay_aa_pay     RoomPay = 1
	RoomPay_master_pay RoomPay = 2
)

var RoomPay_name = map[int32]string{
	0: "none_pay",
	1: "aa_pay",
	2: "master_pay",
}

var RoomPay_value = map[string]int32{
	"none_pay":   0,
	"aa_pay":     1,
	"master_pay": 2,
}

func (x RoomPay) String() string {
	return proto.EnumName(RoomPay_name, int32(x))
}

func (RoomPay) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b6c5d20c74c34643, []int{2}
}

// 房间结束类型
type RoomOverType int32

const (
	RoomOverType_none_over  RoomOverType = 0
	RoomOverType_not_over   RoomOverType = 1
	RoomOverType_normal     RoomOverType = 2
	RoomOverType_force_over RoomOverType = 3
)

var RoomOverType_name = map[int32]string{
	0: "none_over",
	1: "not_over",
	2: "normal",
	3: "force_over",
}

var RoomOverType_value = map[string]int32{
	"none_over":  0,
	"not_over":   1,
	"normal":     2,
	"force_over": 3,
}

func (x RoomOverType) String() string {
	return proto.EnumName(RoomOverType_name, int32(x))
}

func (RoomOverType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b6c5d20c74c34643, []int{3}
}

// 房间玩家角色
type RoomPlayerRole int32

const (
	RoomPlayerRole_none_role RoomPlayerRole = 0
	RoomPlayerRole_player    RoomPlayerRole = 1
	RoomPlayerRole_guest     RoomPlayerRole = 2
)

var RoomPlayerRole_name = map[int32]string{
	0: "none_role",
	1: "player",
	2: "guest",
}

var RoomPlayerRole_value = map[string]int32{
	"none_role": 0,
	"player":    1,
	"guest":     2,
}

func (x RoomPlayerRole) String() string {
	return proto.EnumName(RoomPlayerRole_name, int32(x))
}

func (RoomPlayerRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b6c5d20c74c34643, []int{4}
}

func init() {
	proto.RegisterEnum("room.GameType", GameType_name, GameType_value)
	proto.RegisterEnum("room.RoomType", RoomType_name, RoomType_value)
	proto.RegisterEnum("room.RoomPay", RoomPay_name, RoomPay_value)
	proto.RegisterEnum("room.RoomOverType", RoomOverType_name, RoomOverType_value)
	proto.RegisterEnum("room.RoomPlayerRole", RoomPlayerRole_name, RoomPlayerRole_value)
}

func init() { proto.RegisterFile("room/room.proto", fileDescriptor_b6c5d20c74c34643) }

var fileDescriptor_b6c5d20c74c34643 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4e, 0xc3, 0x30,
	0x10, 0x84, 0xf3, 0x47, 0x9b, 0xae, 0x4a, 0x59, 0xf9, 0x31, 0x22, 0x04, 0x87, 0x4a, 0xbc, 0x42,
	0xb9, 0x81, 0x2a, 0xee, 0x68, 0x69, 0xb6, 0x21, 0xc8, 0xf6, 0x46, 0x8e, 0x89, 0x94, 0x3e, 0x3d,
	0x5a, 0x87, 0x43, 0x7b, 0xb1, 0x76, 0xe4, 0x99, 0xcf, 0xe3, 0x85, 0x87, 0x20, 0xe2, 0x9e, 0xf5,
	0x78, 0x1a, 0x82, 0x44, 0x31, 0x95, 0xce, 0xcd, 0x23, 0xd4, 0x07, 0x72, 0xfc, 0x31, 0x0f, 0x6c,
	0xee, 0x61, 0xe3, 0xc5, 0xf3, 0x67, 0x47, 0x8e, 0x31, 0x33, 0x35, 0x54, 0x17, 0xeb, 0x7e, 0x30,
	0x37, 0x6b, 0x28, 0xdb, 0xf6, 0x82, 0x45, 0xf3, 0x0a, 0xf5, 0x51, 0xc4, 0xdd, 0xb8, 0x15, 0x83,
	0x99, 0x01, 0x58, 0x9d, 0x43, 0xcf, 0xbe, 0xc5, 0x5c, 0x93, 0x9d, 0xd8, 0x16, 0x0b, 0xb3, 0x85,
	0xba, 0xed, 0xa7, 0x7e, 0xec, 0xc5, 0x63, 0x69, 0x36, 0x70, 0xe7, 0x28, 0x9e, 0xbe, 0xb1, 0x6a,
	0xf6, 0xb0, 0x56, 0xd2, 0x3b, 0xcd, 0xea, 0x49, 0xa0, 0x81, 0xe6, 0x85, 0x43, 0x94, 0xe6, 0xdc,
	0xec, 0x00, 0x1c, 0x8d, 0x91, 0x43, 0xd2, 0x45, 0x73, 0x80, 0xad, 0x86, 0xde, 0x26, 0x0e, 0x37,
	0x15, 0x64, 0xe2, 0x80, 0xd9, 0x02, 0x8a, 0x8b, 0xca, 0x15, 0xe4, 0x25, 0x38, 0xb2, 0x58, 0x28,
	0xe8, 0x2c, 0xe1, 0xf4, 0xef, 0x2c, 0x9b, 0x17, 0xd8, 0xa5, 0xd7, 0x2d, 0xcd, 0x1c, 0x8e, 0x62,
	0xaf, 0x7f, 0x63, 0x79, 0x69, 0x31, 0xa4, 0x4b, 0xcc, 0xb5, 0x75, 0xf7, 0xcb, 0x63, 0xc4, 0xe2,
	0x6b, 0x95, 0x56, 0xb7, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xe1, 0xc3, 0xf5, 0x4d, 0x01,
	0x00, 0x00,
}

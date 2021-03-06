// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room/room.proto

package room

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	return fileDescriptor_room_8a3676b2ed9df2a3, []int{0}
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
	return fileDescriptor_room_8a3676b2ed9df2a3, []int{1}
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
	return fileDescriptor_room_8a3676b2ed9df2a3, []int{2}
}

// 房间结束类型
type RoomOverType int32

const (
	RoomOverType_not_over      RoomOverType = 0
	RoomOverType_normal        RoomOverType = 2
	RoomOverType_force_over    RoomOverType = 3
	RoomOverType_dissolve_over RoomOverType = 4
)

var RoomOverType_name = map[int32]string{
	0: "not_over",
	2: "normal",
	3: "force_over",
	4: "dissolve_over",
}
var RoomOverType_value = map[string]int32{
	"not_over":      0,
	"normal":        2,
	"force_over":    3,
	"dissolve_over": 4,
}

func (x RoomOverType) String() string {
	return proto.EnumName(RoomOverType_name, int32(x))
}
func (RoomOverType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_room_8a3676b2ed9df2a3, []int{3}
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
	return fileDescriptor_room_8a3676b2ed9df2a3, []int{4}
}

func init() {
	proto.RegisterEnum("room.GameType", GameType_name, GameType_value)
	proto.RegisterEnum("room.RoomType", RoomType_name, RoomType_value)
	proto.RegisterEnum("room.RoomPay", RoomPay_name, RoomPay_value)
	proto.RegisterEnum("room.RoomOverType", RoomOverType_name, RoomOverType_value)
	proto.RegisterEnum("room.RoomPlayerRole", RoomPlayerRole_name, RoomPlayerRole_value)
}

func init() { proto.RegisterFile("room/room.proto", fileDescriptor_room_8a3676b2ed9df2a3) }

var fileDescriptor_room_8a3676b2ed9df2a3 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x4e, 0xc3, 0x30,
	0x0c, 0x87, 0xd7, 0x3f, 0x6c, 0x9d, 0xb5, 0x0d, 0x93, 0xc7, 0xa8, 0x10, 0x1c, 0x26, 0xf1, 0x0a,
	0x20, 0x71, 0x00, 0x4d, 0xdc, 0x27, 0xb3, 0x78, 0xa5, 0x28, 0x89, 0xab, 0x24, 0x54, 0xea, 0x9e,
	0x1e, 0x25, 0xe5, 0xc0, 0x2e, 0x91, 0x13, 0xfb, 0xfb, 0xf4, 0x8b, 0xe1, 0xd6, 0x8b, 0xd8, 0xc7,
	0x74, 0x3c, 0x0c, 0x5e, 0xa2, 0xa8, 0x3a, 0xd5, 0xed, 0x3d, 0x34, 0xcf, 0x64, 0xf9, 0x63, 0x1a,
	0x58, 0x6d, 0x61, 0xed, 0xc4, 0xf1, 0xb1, 0x23, 0xcb, 0xb8, 0x50, 0x0d, 0xd4, 0x17, 0x63, 0xbf,
	0xb1, 0x50, 0x2b, 0xa8, 0xb4, 0xbe, 0x60, 0xd9, 0xbe, 0x40, 0x73, 0x10, 0xb1, 0x57, 0xd3, 0x49,
	0x83, 0x0b, 0x05, 0xb0, 0x3c, 0xfb, 0x9e, 0x9d, 0xc6, 0x22, 0x91, 0x9d, 0x18, 0x8d, 0xa5, 0xda,
	0x40, 0xa3, 0xfb, 0xb1, 0x0f, 0xbd, 0x38, 0xac, 0xd4, 0x1a, 0x6e, 0x2c, 0xc5, 0xd3, 0x17, 0xd6,
	0xed, 0x1e, 0x56, 0xc9, 0xf4, 0x4e, 0x53, 0x9a, 0xc9, 0xa2, 0x81, 0xa6, 0xd9, 0x43, 0x94, 0xeb,
	0x42, 0xed, 0x00, 0x2c, 0x85, 0xc8, 0x3e, 0xdf, 0xcb, 0xf6, 0x15, 0x36, 0x09, 0x7a, 0x1b, 0xd9,
	0xe7, 0x08, 0x99, 0x8c, 0x47, 0x19, 0xd9, 0xcf, 0xa4, 0x13, 0x6f, 0xc9, 0x60, 0x99, 0xc8, 0xb3,
	0xf8, 0x13, 0xcf, 0xbd, 0x4a, 0xdd, 0xc1, 0x56, 0xf7, 0x21, 0x88, 0x19, 0xff, 0x9e, 0xea, 0xf6,
	0x09, 0x76, 0x39, 0x81, 0xa1, 0x89, 0xfd, 0x41, 0xcc, 0xff, 0x1f, 0x19, 0x9e, 0x7d, 0x43, 0x6e,
	0x62, 0x91, 0x92, 0x77, 0x3f, 0x1c, 0x22, 0x96, 0x9f, 0xcb, 0xbc, 0xbe, 0xfd, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x00, 0xf5, 0x3d, 0xf6, 0x51, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: zlmj/zlmj.proto

/*
Package zlmj is a generated protocol buffer package.

It is generated from these files:
	zlmj/zlmj.proto

It has these top-level messages:
	Option
	GameData
	PlayerData
	Action
	Combination
	Req
	Rsp
	Event
	WantHoldsReq
	DiscardsReq
	DoActionReq
	PassReq
	ErrorRsp
	DiscardsRsp
	GameStartEvent
	DiscardsEvent
	DealCardsEvent
	ResumeEvent
	DrawCardEvent
	NewActionEvent
	RemoveActionEvent
	DispatchActionEvent
	ChangeBaoCardEvent
*/
package zlmj

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

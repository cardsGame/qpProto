// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: circle/circle.proto

/*
Package circle is a generated protocol buffer package.

It is generated from these files:
	circle/circle.proto

It has these top-level messages:
	BuildCircleRequest
	BuildCircleResponse
	UpdateCircleRequest
	UpdateCircleResponse
	GetCircleListRequest
	GetCircleListResponse
	DelCircleRequest
	DelCircleResponse
	ApplyJoinCircleRequest
	ApplyJoinCircleResponse
	GetCircleLogRequest
	GetCircleLogResponse
	DealApplyRequest
	DealApplyResponse
	DealMessageRequest
	DealMessageResponse
	GetApplyListRequest
	GetApplyListResponse
	GetQuiteListRequest
	GetQuiteListResponse
	DealQuitRequest
	DealQuitResponse
	ApplyQuitCircleRequest
	ApplyQuitCircleResponse
	GerCircleDetailRequest
	GerCircleDetailResponse
	GetCircleMembersRequest
	GetCircleMembersResponse
	GetCircleInfoRequest
	GetCircleInfoResponse
	CircleUser
	User
	CircleMsg
	Circle
*/
package circle

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CircleMasterService service

type CircleMasterService interface {
	// 创建
	BuildCircle(ctx context.Context, in *BuildCircleRequest, opts ...client.CallOption) (*BuildCircleResponse, error)
	// 更新
	UpdateCircle(ctx context.Context, in *UpdateCircleRequest, opts ...client.CallOption) (*UpdateCircleResponse, error)
	// 获取圈子列表
	GetCircleList(ctx context.Context, in *GetCircleListRequest, opts ...client.CallOption) (*GetCircleListResponse, error)
	// 获取成员列表
	GetCircleMembers(ctx context.Context, in *GetCircleMembersRequest, opts ...client.CallOption) (*GetCircleMembersResponse, error)
	// 删除圈子
	DelCircle(ctx context.Context, in *DelCircleRequest, opts ...client.CallOption) (*DelCircleResponse, error)
	// 获取圈子的日志
	GetCircleLog(ctx context.Context, in *GetCircleLogRequest, opts ...client.CallOption) (*GetCircleLogResponse, error)
	// 处理圈子消息
	DealMessage(ctx context.Context, in *DealMessageRequest, opts ...client.CallOption) (*DealMessageResponse, error)
	// 处理申请
	DealApply(ctx context.Context, in *DealApplyRequest, opts ...client.CallOption) (*DealApplyResponse, error)
	// 获取加入圈子申请
	GetApplyList(ctx context.Context, in *GetApplyListRequest, opts ...client.CallOption) (*GetApplyListResponse, error)
	// 获取退出圈子申请
	GetQuiteList(ctx context.Context, in *GetQuiteListRequest, opts ...client.CallOption) (*GetQuiteListResponse, error)
	// 处理用户退出申请
	DealQuit(ctx context.Context, in *DealQuitRequest, opts ...client.CallOption) (*DealQuitResponse, error)
	// 详细信息
	GetCircleInfo(ctx context.Context, in *GetCircleInfoRequest, opts ...client.CallOption) (*GetCircleInfoResponse, error)
}

type circleMasterService struct {
	c    client.Client
	name string
}

func NewCircleMasterService(name string, c client.Client) CircleMasterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "circlemasterservice"
	}
	return &circleMasterService{
		c:    c,
		name: name,
	}
}

func (c *circleMasterService) BuildCircle(ctx context.Context, in *BuildCircleRequest, opts ...client.CallOption) (*BuildCircleResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.BuildCircle", in)
	out := new(BuildCircleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) UpdateCircle(ctx context.Context, in *UpdateCircleRequest, opts ...client.CallOption) (*UpdateCircleResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.UpdateCircle", in)
	out := new(UpdateCircleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) GetCircleList(ctx context.Context, in *GetCircleListRequest, opts ...client.CallOption) (*GetCircleListResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.GetCircleList", in)
	out := new(GetCircleListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) GetCircleMembers(ctx context.Context, in *GetCircleMembersRequest, opts ...client.CallOption) (*GetCircleMembersResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.GetCircleMembers", in)
	out := new(GetCircleMembersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) DelCircle(ctx context.Context, in *DelCircleRequest, opts ...client.CallOption) (*DelCircleResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.DelCircle", in)
	out := new(DelCircleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) GetCircleLog(ctx context.Context, in *GetCircleLogRequest, opts ...client.CallOption) (*GetCircleLogResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.GetCircleLog", in)
	out := new(GetCircleLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) DealMessage(ctx context.Context, in *DealMessageRequest, opts ...client.CallOption) (*DealMessageResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.DealMessage", in)
	out := new(DealMessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) DealApply(ctx context.Context, in *DealApplyRequest, opts ...client.CallOption) (*DealApplyResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.DealApply", in)
	out := new(DealApplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) GetApplyList(ctx context.Context, in *GetApplyListRequest, opts ...client.CallOption) (*GetApplyListResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.GetApplyList", in)
	out := new(GetApplyListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) GetQuiteList(ctx context.Context, in *GetQuiteListRequest, opts ...client.CallOption) (*GetQuiteListResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.GetQuiteList", in)
	out := new(GetQuiteListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) DealQuit(ctx context.Context, in *DealQuitRequest, opts ...client.CallOption) (*DealQuitResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.DealQuit", in)
	out := new(DealQuitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleMasterService) GetCircleInfo(ctx context.Context, in *GetCircleInfoRequest, opts ...client.CallOption) (*GetCircleInfoResponse, error) {
	req := c.c.NewRequest(c.name, "CircleMasterService.GetCircleInfo", in)
	out := new(GetCircleInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CircleMasterService service

type CircleMasterServiceHandler interface {
	// 创建
	BuildCircle(context.Context, *BuildCircleRequest, *BuildCircleResponse) error
	// 更新
	UpdateCircle(context.Context, *UpdateCircleRequest, *UpdateCircleResponse) error
	// 获取圈子列表
	GetCircleList(context.Context, *GetCircleListRequest, *GetCircleListResponse) error
	// 获取成员列表
	GetCircleMembers(context.Context, *GetCircleMembersRequest, *GetCircleMembersResponse) error
	// 删除圈子
	DelCircle(context.Context, *DelCircleRequest, *DelCircleResponse) error
	// 获取圈子的日志
	GetCircleLog(context.Context, *GetCircleLogRequest, *GetCircleLogResponse) error
	// 处理圈子消息
	DealMessage(context.Context, *DealMessageRequest, *DealMessageResponse) error
	// 处理申请
	DealApply(context.Context, *DealApplyRequest, *DealApplyResponse) error
	// 获取加入圈子申请
	GetApplyList(context.Context, *GetApplyListRequest, *GetApplyListResponse) error
	// 获取退出圈子申请
	GetQuiteList(context.Context, *GetQuiteListRequest, *GetQuiteListResponse) error
	// 处理用户退出申请
	DealQuit(context.Context, *DealQuitRequest, *DealQuitResponse) error
	// 详细信息
	GetCircleInfo(context.Context, *GetCircleInfoRequest, *GetCircleInfoResponse) error
}

func RegisterCircleMasterServiceHandler(s server.Server, hdlr CircleMasterServiceHandler, opts ...server.HandlerOption) error {
	type circleMasterService interface {
		BuildCircle(ctx context.Context, in *BuildCircleRequest, out *BuildCircleResponse) error
		UpdateCircle(ctx context.Context, in *UpdateCircleRequest, out *UpdateCircleResponse) error
		GetCircleList(ctx context.Context, in *GetCircleListRequest, out *GetCircleListResponse) error
		GetCircleMembers(ctx context.Context, in *GetCircleMembersRequest, out *GetCircleMembersResponse) error
		DelCircle(ctx context.Context, in *DelCircleRequest, out *DelCircleResponse) error
		GetCircleLog(ctx context.Context, in *GetCircleLogRequest, out *GetCircleLogResponse) error
		DealMessage(ctx context.Context, in *DealMessageRequest, out *DealMessageResponse) error
		DealApply(ctx context.Context, in *DealApplyRequest, out *DealApplyResponse) error
		GetApplyList(ctx context.Context, in *GetApplyListRequest, out *GetApplyListResponse) error
		GetQuiteList(ctx context.Context, in *GetQuiteListRequest, out *GetQuiteListResponse) error
		DealQuit(ctx context.Context, in *DealQuitRequest, out *DealQuitResponse) error
		GetCircleInfo(ctx context.Context, in *GetCircleInfoRequest, out *GetCircleInfoResponse) error
	}
	type CircleMasterService struct {
		circleMasterService
	}
	h := &circleMasterServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CircleMasterService{h}, opts...))
}

type circleMasterServiceHandler struct {
	CircleMasterServiceHandler
}

func (h *circleMasterServiceHandler) BuildCircle(ctx context.Context, in *BuildCircleRequest, out *BuildCircleResponse) error {
	return h.CircleMasterServiceHandler.BuildCircle(ctx, in, out)
}

func (h *circleMasterServiceHandler) UpdateCircle(ctx context.Context, in *UpdateCircleRequest, out *UpdateCircleResponse) error {
	return h.CircleMasterServiceHandler.UpdateCircle(ctx, in, out)
}

func (h *circleMasterServiceHandler) GetCircleList(ctx context.Context, in *GetCircleListRequest, out *GetCircleListResponse) error {
	return h.CircleMasterServiceHandler.GetCircleList(ctx, in, out)
}

func (h *circleMasterServiceHandler) GetCircleMembers(ctx context.Context, in *GetCircleMembersRequest, out *GetCircleMembersResponse) error {
	return h.CircleMasterServiceHandler.GetCircleMembers(ctx, in, out)
}

func (h *circleMasterServiceHandler) DelCircle(ctx context.Context, in *DelCircleRequest, out *DelCircleResponse) error {
	return h.CircleMasterServiceHandler.DelCircle(ctx, in, out)
}

func (h *circleMasterServiceHandler) GetCircleLog(ctx context.Context, in *GetCircleLogRequest, out *GetCircleLogResponse) error {
	return h.CircleMasterServiceHandler.GetCircleLog(ctx, in, out)
}

func (h *circleMasterServiceHandler) DealMessage(ctx context.Context, in *DealMessageRequest, out *DealMessageResponse) error {
	return h.CircleMasterServiceHandler.DealMessage(ctx, in, out)
}

func (h *circleMasterServiceHandler) DealApply(ctx context.Context, in *DealApplyRequest, out *DealApplyResponse) error {
	return h.CircleMasterServiceHandler.DealApply(ctx, in, out)
}

func (h *circleMasterServiceHandler) GetApplyList(ctx context.Context, in *GetApplyListRequest, out *GetApplyListResponse) error {
	return h.CircleMasterServiceHandler.GetApplyList(ctx, in, out)
}

func (h *circleMasterServiceHandler) GetQuiteList(ctx context.Context, in *GetQuiteListRequest, out *GetQuiteListResponse) error {
	return h.CircleMasterServiceHandler.GetQuiteList(ctx, in, out)
}

func (h *circleMasterServiceHandler) DealQuit(ctx context.Context, in *DealQuitRequest, out *DealQuitResponse) error {
	return h.CircleMasterServiceHandler.DealQuit(ctx, in, out)
}

func (h *circleMasterServiceHandler) GetCircleInfo(ctx context.Context, in *GetCircleInfoRequest, out *GetCircleInfoResponse) error {
	return h.CircleMasterServiceHandler.GetCircleInfo(ctx, in, out)
}

// Client API for CircleUserService service

type CircleUserService interface {
	// 获取圈子列表
	GetCircleList(ctx context.Context, in *GetCircleListRequest, opts ...client.CallOption) (*GetCircleListResponse, error)
	// 申请加入圈子
	ApplyJoinCircle(ctx context.Context, in *ApplyJoinCircleRequest, opts ...client.CallOption) (*ApplyJoinCircleResponse, error)
	// 获取成员列表
	GetCircleMembers(ctx context.Context, in *GetCircleMembersRequest, opts ...client.CallOption) (*GetCircleMembersResponse, error)
	// 获取圈子的日志
	GetCircleLog(ctx context.Context, in *GetCircleLogRequest, opts ...client.CallOption) (*GetCircleLogResponse, error)
	// 玩家申请离开圈子
	ApplyQuitCircle(ctx context.Context, in *ApplyQuitCircleRequest, opts ...client.CallOption) (*ApplyQuitCircleResponse, error)
	// 获取圈子详细信息
	GerCircleDetail(ctx context.Context, in *GerCircleDetailRequest, opts ...client.CallOption) (*GerCircleDetailResponse, error)
}

type circleUserService struct {
	c    client.Client
	name string
}

func NewCircleUserService(name string, c client.Client) CircleUserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "circleuserservice"
	}
	return &circleUserService{
		c:    c,
		name: name,
	}
}

func (c *circleUserService) GetCircleList(ctx context.Context, in *GetCircleListRequest, opts ...client.CallOption) (*GetCircleListResponse, error) {
	req := c.c.NewRequest(c.name, "CircleUserService.GetCircleList", in)
	out := new(GetCircleListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleUserService) ApplyJoinCircle(ctx context.Context, in *ApplyJoinCircleRequest, opts ...client.CallOption) (*ApplyJoinCircleResponse, error) {
	req := c.c.NewRequest(c.name, "CircleUserService.ApplyJoinCircle", in)
	out := new(ApplyJoinCircleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleUserService) GetCircleMembers(ctx context.Context, in *GetCircleMembersRequest, opts ...client.CallOption) (*GetCircleMembersResponse, error) {
	req := c.c.NewRequest(c.name, "CircleUserService.GetCircleMembers", in)
	out := new(GetCircleMembersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleUserService) GetCircleLog(ctx context.Context, in *GetCircleLogRequest, opts ...client.CallOption) (*GetCircleLogResponse, error) {
	req := c.c.NewRequest(c.name, "CircleUserService.GetCircleLog", in)
	out := new(GetCircleLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleUserService) ApplyQuitCircle(ctx context.Context, in *ApplyQuitCircleRequest, opts ...client.CallOption) (*ApplyQuitCircleResponse, error) {
	req := c.c.NewRequest(c.name, "CircleUserService.ApplyQuitCircle", in)
	out := new(ApplyQuitCircleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *circleUserService) GerCircleDetail(ctx context.Context, in *GerCircleDetailRequest, opts ...client.CallOption) (*GerCircleDetailResponse, error) {
	req := c.c.NewRequest(c.name, "CircleUserService.GerCircleDetail", in)
	out := new(GerCircleDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CircleUserService service

type CircleUserServiceHandler interface {
	// 获取圈子列表
	GetCircleList(context.Context, *GetCircleListRequest, *GetCircleListResponse) error
	// 申请加入圈子
	ApplyJoinCircle(context.Context, *ApplyJoinCircleRequest, *ApplyJoinCircleResponse) error
	// 获取成员列表
	GetCircleMembers(context.Context, *GetCircleMembersRequest, *GetCircleMembersResponse) error
	// 获取圈子的日志
	GetCircleLog(context.Context, *GetCircleLogRequest, *GetCircleLogResponse) error
	// 玩家申请离开圈子
	ApplyQuitCircle(context.Context, *ApplyQuitCircleRequest, *ApplyQuitCircleResponse) error
	// 获取圈子详细信息
	GerCircleDetail(context.Context, *GerCircleDetailRequest, *GerCircleDetailResponse) error
}

func RegisterCircleUserServiceHandler(s server.Server, hdlr CircleUserServiceHandler, opts ...server.HandlerOption) error {
	type circleUserService interface {
		GetCircleList(ctx context.Context, in *GetCircleListRequest, out *GetCircleListResponse) error
		ApplyJoinCircle(ctx context.Context, in *ApplyJoinCircleRequest, out *ApplyJoinCircleResponse) error
		GetCircleMembers(ctx context.Context, in *GetCircleMembersRequest, out *GetCircleMembersResponse) error
		GetCircleLog(ctx context.Context, in *GetCircleLogRequest, out *GetCircleLogResponse) error
		ApplyQuitCircle(ctx context.Context, in *ApplyQuitCircleRequest, out *ApplyQuitCircleResponse) error
		GerCircleDetail(ctx context.Context, in *GerCircleDetailRequest, out *GerCircleDetailResponse) error
	}
	type CircleUserService struct {
		circleUserService
	}
	h := &circleUserServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CircleUserService{h}, opts...))
}

type circleUserServiceHandler struct {
	CircleUserServiceHandler
}

func (h *circleUserServiceHandler) GetCircleList(ctx context.Context, in *GetCircleListRequest, out *GetCircleListResponse) error {
	return h.CircleUserServiceHandler.GetCircleList(ctx, in, out)
}

func (h *circleUserServiceHandler) ApplyJoinCircle(ctx context.Context, in *ApplyJoinCircleRequest, out *ApplyJoinCircleResponse) error {
	return h.CircleUserServiceHandler.ApplyJoinCircle(ctx, in, out)
}

func (h *circleUserServiceHandler) GetCircleMembers(ctx context.Context, in *GetCircleMembersRequest, out *GetCircleMembersResponse) error {
	return h.CircleUserServiceHandler.GetCircleMembers(ctx, in, out)
}

func (h *circleUserServiceHandler) GetCircleLog(ctx context.Context, in *GetCircleLogRequest, out *GetCircleLogResponse) error {
	return h.CircleUserServiceHandler.GetCircleLog(ctx, in, out)
}

func (h *circleUserServiceHandler) ApplyQuitCircle(ctx context.Context, in *ApplyQuitCircleRequest, out *ApplyQuitCircleResponse) error {
	return h.CircleUserServiceHandler.ApplyQuitCircle(ctx, in, out)
}

func (h *circleUserServiceHandler) GerCircleDetail(ctx context.Context, in *GerCircleDetailRequest, out *GerCircleDetailResponse) error {
	return h.CircleUserServiceHandler.GerCircleDetail(ctx, in, out)
}

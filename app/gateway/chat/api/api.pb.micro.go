// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api.proto

package api

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Chat service

func NewChatEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Chat service

type ChatService interface {
	PushMessage(ctx context.Context, in *PushMessageReq, opts ...client.CallOption) (*Empty, error)
}

type chatService struct {
	c    client.Client
	name string
}

func NewChatService(name string, c client.Client) ChatService {
	return &chatService{
		c:    c,
		name: name,
	}
}

func (c *chatService) PushMessage(ctx context.Context, in *PushMessageReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "Chat.PushMessage", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chat service

type ChatHandler interface {
	PushMessage(context.Context, *PushMessageReq, *Empty) error
}

func RegisterChatHandler(s server.Server, hdlr ChatHandler, opts ...server.HandlerOption) error {
	type chat interface {
		PushMessage(ctx context.Context, in *PushMessageReq, out *Empty) error
	}
	type Chat struct {
		chat
	}
	h := &chatHandler{hdlr}
	return s.Handle(s.NewHandler(&Chat{h}, opts...))
}

type chatHandler struct {
	ChatHandler
}

func (h *chatHandler) PushMessage(ctx context.Context, in *PushMessageReq, out *Empty) error {
	return h.ChatHandler.PushMessage(ctx, in, out)
}

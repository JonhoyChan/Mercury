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

// Api Endpoints for ChatAdmin service

func NewChatAdminEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ChatAdmin service

type ChatAdminService interface {
	// Get client
	GetClient(ctx context.Context, in *GetClientReq, opts ...client.CallOption) (*GetClientResp, error)
	// Create new client
	CreateClient(ctx context.Context, in *CreateClientReq, opts ...client.CallOption) (*CreateClientResp, error)
	// Update client
	UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...client.CallOption) (*Empty, error)
	// Delete client
	DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...client.CallOption) (*Empty, error)
}

type chatAdminService struct {
	c    client.Client
	name string
}

func NewChatAdminService(name string, c client.Client) ChatAdminService {
	return &chatAdminService{
		c:    c,
		name: name,
	}
}

func (c *chatAdminService) GetClient(ctx context.Context, in *GetClientReq, opts ...client.CallOption) (*GetClientResp, error) {
	req := c.c.NewRequest(c.name, "ChatAdmin.GetClient", in)
	out := new(GetClientResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminService) CreateClient(ctx context.Context, in *CreateClientReq, opts ...client.CallOption) (*CreateClientResp, error) {
	req := c.c.NewRequest(c.name, "ChatAdmin.CreateClient", in)
	out := new(CreateClientResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminService) UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "ChatAdmin.UpdateClient", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatAdminService) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "ChatAdmin.DeleteClient", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChatAdmin service

type ChatAdminHandler interface {
	// Get client
	GetClient(context.Context, *GetClientReq, *GetClientResp) error
	// Create new client
	CreateClient(context.Context, *CreateClientReq, *CreateClientResp) error
	// Update client
	UpdateClient(context.Context, *UpdateClientReq, *Empty) error
	// Delete client
	DeleteClient(context.Context, *DeleteClientReq, *Empty) error
}

func RegisterChatAdminHandler(s server.Server, hdlr ChatAdminHandler, opts ...server.HandlerOption) error {
	type chatAdmin interface {
		GetClient(ctx context.Context, in *GetClientReq, out *GetClientResp) error
		CreateClient(ctx context.Context, in *CreateClientReq, out *CreateClientResp) error
		UpdateClient(ctx context.Context, in *UpdateClientReq, out *Empty) error
		DeleteClient(ctx context.Context, in *DeleteClientReq, out *Empty) error
	}
	type ChatAdmin struct {
		chatAdmin
	}
	h := &chatAdminHandler{hdlr}
	return s.Handle(s.NewHandler(&ChatAdmin{h}, opts...))
}

type chatAdminHandler struct {
	ChatAdminHandler
}

func (h *chatAdminHandler) GetClient(ctx context.Context, in *GetClientReq, out *GetClientResp) error {
	return h.ChatAdminHandler.GetClient(ctx, in, out)
}

func (h *chatAdminHandler) CreateClient(ctx context.Context, in *CreateClientReq, out *CreateClientResp) error {
	return h.ChatAdminHandler.CreateClient(ctx, in, out)
}

func (h *chatAdminHandler) UpdateClient(ctx context.Context, in *UpdateClientReq, out *Empty) error {
	return h.ChatAdminHandler.UpdateClient(ctx, in, out)
}

func (h *chatAdminHandler) DeleteClient(ctx context.Context, in *DeleteClientReq, out *Empty) error {
	return h.ChatAdminHandler.DeleteClient(ctx, in, out)
}

// Api Endpoints for ChatClientAdmin service

func NewChatClientAdminEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ChatClientAdmin service

type ChatClientAdminService interface {
	// Get client
	GetClient(ctx context.Context, in *GetClientReq, opts ...client.CallOption) (*GetClientResp, error)
	// Create new client
	CreateClient(ctx context.Context, in *CreateClientReq, opts ...client.CallOption) (*CreateClientResp, error)
	// Update client
	UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...client.CallOption) (*Empty, error)
	// Delete client
	DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...client.CallOption) (*Empty, error)
	// Generate a new token for client
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...client.CallOption) (*TokenResp, error)
	// Create new user
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...client.CallOption) (*CreateUserResp, error)
	// Update user activated
	UpdateActivated(ctx context.Context, in *UpdateActivatedReq, opts ...client.CallOption) (*Empty, error)
	// Delete user
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...client.CallOption) (*Empty, error)
	// Generate a new token for user
	GenerateUserToken(ctx context.Context, in *GenerateUserTokenReq, opts ...client.CallOption) (*TokenResp, error)
	// Add friend
	AddFriend(ctx context.Context, in *AddFriendReq, opts ...client.CallOption) (*Empty, error)
	// Get friends
	GetFriends(ctx context.Context, in *GetFriendsReq, opts ...client.CallOption) (*GetFriendsResp, error)
	// Delete friend
	DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...client.CallOption) (*Empty, error)
	// Create new group
	CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...client.CallOption) (*CreateGroupResp, error)
	// Get groups
	GetGroups(ctx context.Context, in *GetGroupsReq, opts ...client.CallOption) (*GetGroupsResp, error)
	// Add user to group
	AddMember(ctx context.Context, in *AddMemberReq, opts ...client.CallOption) (*Empty, error)
	// Get users of the group
	GetMembers(ctx context.Context, in *GetMembersReq, opts ...client.CallOption) (*GetMembersResp, error)
	// Listening all real-time messages under the client to which the current token belongs
	Listen(ctx context.Context, in *ListenReq, opts ...client.CallOption) (ChatClientAdmin_ListenService, error)
}

type chatClientAdminService struct {
	c    client.Client
	name string
}

func NewChatClientAdminService(name string, c client.Client) ChatClientAdminService {
	return &chatClientAdminService{
		c:    c,
		name: name,
	}
}

func (c *chatClientAdminService) GetClient(ctx context.Context, in *GetClientReq, opts ...client.CallOption) (*GetClientResp, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.GetClient", in)
	out := new(GetClientResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) CreateClient(ctx context.Context, in *CreateClientReq, opts ...client.CallOption) (*CreateClientResp, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.CreateClient", in)
	out := new(CreateClientResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.UpdateClient", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.DeleteClient", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...client.CallOption) (*TokenResp, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.GenerateToken", in)
	out := new(TokenResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) CreateUser(ctx context.Context, in *CreateUserReq, opts ...client.CallOption) (*CreateUserResp, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.CreateUser", in)
	out := new(CreateUserResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) UpdateActivated(ctx context.Context, in *UpdateActivatedReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.UpdateActivated", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.DeleteUser", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) GenerateUserToken(ctx context.Context, in *GenerateUserTokenReq, opts ...client.CallOption) (*TokenResp, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.GenerateUserToken", in)
	out := new(TokenResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) AddFriend(ctx context.Context, in *AddFriendReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.AddFriend", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) GetFriends(ctx context.Context, in *GetFriendsReq, opts ...client.CallOption) (*GetFriendsResp, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.GetFriends", in)
	out := new(GetFriendsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.DeleteFriend", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...client.CallOption) (*CreateGroupResp, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.CreateGroup", in)
	out := new(CreateGroupResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) GetGroups(ctx context.Context, in *GetGroupsReq, opts ...client.CallOption) (*GetGroupsResp, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.GetGroups", in)
	out := new(GetGroupsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) AddMember(ctx context.Context, in *AddMemberReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.AddMember", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) GetMembers(ctx context.Context, in *GetMembersReq, opts ...client.CallOption) (*GetMembersResp, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.GetMembers", in)
	out := new(GetMembersResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClientAdminService) Listen(ctx context.Context, in *ListenReq, opts ...client.CallOption) (ChatClientAdmin_ListenService, error) {
	req := c.c.NewRequest(c.name, "ChatClientAdmin.Listen", &ListenReq{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &chatClientAdminServiceListen{stream}, nil
}

type ChatClientAdmin_ListenService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Message, error)
}

type chatClientAdminServiceListen struct {
	stream client.Stream
}

func (x *chatClientAdminServiceListen) Close() error {
	return x.stream.Close()
}

func (x *chatClientAdminServiceListen) Context() context.Context {
	return x.stream.Context()
}

func (x *chatClientAdminServiceListen) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *chatClientAdminServiceListen) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *chatClientAdminServiceListen) Recv() (*Message, error) {
	m := new(Message)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChatClientAdmin service

type ChatClientAdminHandler interface {
	// Get client
	GetClient(context.Context, *GetClientReq, *GetClientResp) error
	// Create new client
	CreateClient(context.Context, *CreateClientReq, *CreateClientResp) error
	// Update client
	UpdateClient(context.Context, *UpdateClientReq, *Empty) error
	// Delete client
	DeleteClient(context.Context, *DeleteClientReq, *Empty) error
	// Generate a new token for client
	GenerateToken(context.Context, *GenerateTokenReq, *TokenResp) error
	// Create new user
	CreateUser(context.Context, *CreateUserReq, *CreateUserResp) error
	// Update user activated
	UpdateActivated(context.Context, *UpdateActivatedReq, *Empty) error
	// Delete user
	DeleteUser(context.Context, *DeleteUserReq, *Empty) error
	// Generate a new token for user
	GenerateUserToken(context.Context, *GenerateUserTokenReq, *TokenResp) error
	// Add friend
	AddFriend(context.Context, *AddFriendReq, *Empty) error
	// Get friends
	GetFriends(context.Context, *GetFriendsReq, *GetFriendsResp) error
	// Delete friend
	DeleteFriend(context.Context, *DeleteFriendReq, *Empty) error
	// Create new group
	CreateGroup(context.Context, *CreateGroupReq, *CreateGroupResp) error
	// Get groups
	GetGroups(context.Context, *GetGroupsReq, *GetGroupsResp) error
	// Add user to group
	AddMember(context.Context, *AddMemberReq, *Empty) error
	// Get users of the group
	GetMembers(context.Context, *GetMembersReq, *GetMembersResp) error
	// Listening all real-time messages under the client to which the current token belongs
	Listen(context.Context, *ListenReq, ChatClientAdmin_ListenStream) error
}

func RegisterChatClientAdminHandler(s server.Server, hdlr ChatClientAdminHandler, opts ...server.HandlerOption) error {
	type chatClientAdmin interface {
		GetClient(ctx context.Context, in *GetClientReq, out *GetClientResp) error
		CreateClient(ctx context.Context, in *CreateClientReq, out *CreateClientResp) error
		UpdateClient(ctx context.Context, in *UpdateClientReq, out *Empty) error
		DeleteClient(ctx context.Context, in *DeleteClientReq, out *Empty) error
		GenerateToken(ctx context.Context, in *GenerateTokenReq, out *TokenResp) error
		CreateUser(ctx context.Context, in *CreateUserReq, out *CreateUserResp) error
		UpdateActivated(ctx context.Context, in *UpdateActivatedReq, out *Empty) error
		DeleteUser(ctx context.Context, in *DeleteUserReq, out *Empty) error
		GenerateUserToken(ctx context.Context, in *GenerateUserTokenReq, out *TokenResp) error
		AddFriend(ctx context.Context, in *AddFriendReq, out *Empty) error
		GetFriends(ctx context.Context, in *GetFriendsReq, out *GetFriendsResp) error
		DeleteFriend(ctx context.Context, in *DeleteFriendReq, out *Empty) error
		CreateGroup(ctx context.Context, in *CreateGroupReq, out *CreateGroupResp) error
		GetGroups(ctx context.Context, in *GetGroupsReq, out *GetGroupsResp) error
		AddMember(ctx context.Context, in *AddMemberReq, out *Empty) error
		GetMembers(ctx context.Context, in *GetMembersReq, out *GetMembersResp) error
		Listen(ctx context.Context, stream server.Stream) error
	}
	type ChatClientAdmin struct {
		chatClientAdmin
	}
	h := &chatClientAdminHandler{hdlr}
	return s.Handle(s.NewHandler(&ChatClientAdmin{h}, opts...))
}

type chatClientAdminHandler struct {
	ChatClientAdminHandler
}

func (h *chatClientAdminHandler) GetClient(ctx context.Context, in *GetClientReq, out *GetClientResp) error {
	return h.ChatClientAdminHandler.GetClient(ctx, in, out)
}

func (h *chatClientAdminHandler) CreateClient(ctx context.Context, in *CreateClientReq, out *CreateClientResp) error {
	return h.ChatClientAdminHandler.CreateClient(ctx, in, out)
}

func (h *chatClientAdminHandler) UpdateClient(ctx context.Context, in *UpdateClientReq, out *Empty) error {
	return h.ChatClientAdminHandler.UpdateClient(ctx, in, out)
}

func (h *chatClientAdminHandler) DeleteClient(ctx context.Context, in *DeleteClientReq, out *Empty) error {
	return h.ChatClientAdminHandler.DeleteClient(ctx, in, out)
}

func (h *chatClientAdminHandler) GenerateToken(ctx context.Context, in *GenerateTokenReq, out *TokenResp) error {
	return h.ChatClientAdminHandler.GenerateToken(ctx, in, out)
}

func (h *chatClientAdminHandler) CreateUser(ctx context.Context, in *CreateUserReq, out *CreateUserResp) error {
	return h.ChatClientAdminHandler.CreateUser(ctx, in, out)
}

func (h *chatClientAdminHandler) UpdateActivated(ctx context.Context, in *UpdateActivatedReq, out *Empty) error {
	return h.ChatClientAdminHandler.UpdateActivated(ctx, in, out)
}

func (h *chatClientAdminHandler) DeleteUser(ctx context.Context, in *DeleteUserReq, out *Empty) error {
	return h.ChatClientAdminHandler.DeleteUser(ctx, in, out)
}

func (h *chatClientAdminHandler) GenerateUserToken(ctx context.Context, in *GenerateUserTokenReq, out *TokenResp) error {
	return h.ChatClientAdminHandler.GenerateUserToken(ctx, in, out)
}

func (h *chatClientAdminHandler) AddFriend(ctx context.Context, in *AddFriendReq, out *Empty) error {
	return h.ChatClientAdminHandler.AddFriend(ctx, in, out)
}

func (h *chatClientAdminHandler) GetFriends(ctx context.Context, in *GetFriendsReq, out *GetFriendsResp) error {
	return h.ChatClientAdminHandler.GetFriends(ctx, in, out)
}

func (h *chatClientAdminHandler) DeleteFriend(ctx context.Context, in *DeleteFriendReq, out *Empty) error {
	return h.ChatClientAdminHandler.DeleteFriend(ctx, in, out)
}

func (h *chatClientAdminHandler) CreateGroup(ctx context.Context, in *CreateGroupReq, out *CreateGroupResp) error {
	return h.ChatClientAdminHandler.CreateGroup(ctx, in, out)
}

func (h *chatClientAdminHandler) GetGroups(ctx context.Context, in *GetGroupsReq, out *GetGroupsResp) error {
	return h.ChatClientAdminHandler.GetGroups(ctx, in, out)
}

func (h *chatClientAdminHandler) AddMember(ctx context.Context, in *AddMemberReq, out *Empty) error {
	return h.ChatClientAdminHandler.AddMember(ctx, in, out)
}

func (h *chatClientAdminHandler) GetMembers(ctx context.Context, in *GetMembersReq, out *GetMembersResp) error {
	return h.ChatClientAdminHandler.GetMembers(ctx, in, out)
}

func (h *chatClientAdminHandler) Listen(ctx context.Context, stream server.Stream) error {
	m := new(ListenReq)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ChatClientAdminHandler.Listen(ctx, m, &chatClientAdminListenStream{stream})
}

type ChatClientAdmin_ListenStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Message) error
}

type chatClientAdminListenStream struct {
	stream server.Stream
}

func (x *chatClientAdminListenStream) Close() error {
	return x.stream.Close()
}

func (x *chatClientAdminListenStream) Context() context.Context {
	return x.stream.Context()
}

func (x *chatClientAdminListenStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *chatClientAdminListenStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *chatClientAdminListenStream) Send(m *Message) error {
	return x.stream.Send(m)
}

// Api Endpoints for Chat service

func NewChatEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Chat service

type ChatService interface {
	// Connect a connection
	Connect(ctx context.Context, in *ConnectReq, opts ...client.CallOption) (*ConnectResp, error)
	// Disconnect a connection
	Disconnect(ctx context.Context, in *DisconnectReq, opts ...client.CallOption) (*Empty, error)
	// Heartbeat a connection
	Heartbeat(ctx context.Context, in *HeartbeatReq, opts ...client.CallOption) (*Empty, error)
	// Push message
	PushMessage(ctx context.Context, in *PushMessageReq, opts ...client.CallOption) (*PushMessageResp, error)
	// Pull message
	PullMessage(ctx context.Context, in *PullMessageReq, opts ...client.CallOption) (*PullMessageResp, error)
	// Read message
	ReadMessage(ctx context.Context, in *ReadMessageReq, opts ...client.CallOption) (*Empty, error)
	// Keypress
	Keypress(ctx context.Context, in *KeypressReq, opts ...client.CallOption) (*Empty, error)
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

func (c *chatService) Connect(ctx context.Context, in *ConnectReq, opts ...client.CallOption) (*ConnectResp, error) {
	req := c.c.NewRequest(c.name, "Chat.Connect", in)
	out := new(ConnectResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatService) Disconnect(ctx context.Context, in *DisconnectReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "Chat.Disconnect", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatService) Heartbeat(ctx context.Context, in *HeartbeatReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "Chat.Heartbeat", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatService) PushMessage(ctx context.Context, in *PushMessageReq, opts ...client.CallOption) (*PushMessageResp, error) {
	req := c.c.NewRequest(c.name, "Chat.PushMessage", in)
	out := new(PushMessageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatService) PullMessage(ctx context.Context, in *PullMessageReq, opts ...client.CallOption) (*PullMessageResp, error) {
	req := c.c.NewRequest(c.name, "Chat.PullMessage", in)
	out := new(PullMessageResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatService) ReadMessage(ctx context.Context, in *ReadMessageReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "Chat.ReadMessage", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatService) Keypress(ctx context.Context, in *KeypressReq, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "Chat.Keypress", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chat service

type ChatHandler interface {
	// Connect a connection
	Connect(context.Context, *ConnectReq, *ConnectResp) error
	// Disconnect a connection
	Disconnect(context.Context, *DisconnectReq, *Empty) error
	// Heartbeat a connection
	Heartbeat(context.Context, *HeartbeatReq, *Empty) error
	// Push message
	PushMessage(context.Context, *PushMessageReq, *PushMessageResp) error
	// Pull message
	PullMessage(context.Context, *PullMessageReq, *PullMessageResp) error
	// Read message
	ReadMessage(context.Context, *ReadMessageReq, *Empty) error
	// Keypress
	Keypress(context.Context, *KeypressReq, *Empty) error
}

func RegisterChatHandler(s server.Server, hdlr ChatHandler, opts ...server.HandlerOption) error {
	type chat interface {
		Connect(ctx context.Context, in *ConnectReq, out *ConnectResp) error
		Disconnect(ctx context.Context, in *DisconnectReq, out *Empty) error
		Heartbeat(ctx context.Context, in *HeartbeatReq, out *Empty) error
		PushMessage(ctx context.Context, in *PushMessageReq, out *PushMessageResp) error
		PullMessage(ctx context.Context, in *PullMessageReq, out *PullMessageResp) error
		ReadMessage(ctx context.Context, in *ReadMessageReq, out *Empty) error
		Keypress(ctx context.Context, in *KeypressReq, out *Empty) error
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

func (h *chatHandler) Connect(ctx context.Context, in *ConnectReq, out *ConnectResp) error {
	return h.ChatHandler.Connect(ctx, in, out)
}

func (h *chatHandler) Disconnect(ctx context.Context, in *DisconnectReq, out *Empty) error {
	return h.ChatHandler.Disconnect(ctx, in, out)
}

func (h *chatHandler) Heartbeat(ctx context.Context, in *HeartbeatReq, out *Empty) error {
	return h.ChatHandler.Heartbeat(ctx, in, out)
}

func (h *chatHandler) PushMessage(ctx context.Context, in *PushMessageReq, out *PushMessageResp) error {
	return h.ChatHandler.PushMessage(ctx, in, out)
}

func (h *chatHandler) PullMessage(ctx context.Context, in *PullMessageReq, out *PullMessageResp) error {
	return h.ChatHandler.PullMessage(ctx, in, out)
}

func (h *chatHandler) ReadMessage(ctx context.Context, in *ReadMessageReq, out *Empty) error {
	return h.ChatHandler.ReadMessage(ctx, in, out)
}

func (h *chatHandler) Keypress(ctx context.Context, in *KeypressReq, out *Empty) error {
	return h.ChatHandler.Keypress(ctx, in, out)
}
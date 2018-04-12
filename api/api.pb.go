// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/api.proto

It has these top-level messages:
	Return
	AccountList
	WitnessList
	AssetIssueList
	NodeList
	Node
	Address
	EmptyMessage
	NumberMessage
	BytesMessage
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protocol1 "github.com/tronprotocol/grpc-gateway/core"
import protocol2 "github.com/tronprotocol/grpc-gateway/core"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Return struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *Return) Reset()                    { *m = Return{} }
func (m *Return) String() string            { return proto.CompactTextString(m) }
func (*Return) ProtoMessage()               {}
func (*Return) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Return) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type AccountList struct {
	Accounts []*protocol1.Account `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *AccountList) Reset()                    { *m = AccountList{} }
func (m *AccountList) String() string            { return proto.CompactTextString(m) }
func (*AccountList) ProtoMessage()               {}
func (*AccountList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AccountList) GetAccounts() []*protocol1.Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type WitnessList struct {
	Witnesses []*protocol1.Witness `protobuf:"bytes,1,rep,name=witnesses" json:"witnesses,omitempty"`
}

func (m *WitnessList) Reset()                    { *m = WitnessList{} }
func (m *WitnessList) String() string            { return proto.CompactTextString(m) }
func (*WitnessList) ProtoMessage()               {}
func (*WitnessList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WitnessList) GetWitnesses() []*protocol1.Witness {
	if m != nil {
		return m.Witnesses
	}
	return nil
}

type AssetIssueList struct {
	AssetIssue []*protocol2.AssetIssueContract `protobuf:"bytes,1,rep,name=assetIssue" json:"assetIssue,omitempty"`
}

func (m *AssetIssueList) Reset()                    { *m = AssetIssueList{} }
func (m *AssetIssueList) String() string            { return proto.CompactTextString(m) }
func (*AssetIssueList) ProtoMessage()               {}
func (*AssetIssueList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AssetIssueList) GetAssetIssue() []*protocol2.AssetIssueContract {
	if m != nil {
		return m.AssetIssue
	}
	return nil
}

// Gossip node list
type NodeList struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *NodeList) Reset()                    { *m = NodeList{} }
func (m *NodeList) String() string            { return proto.CompactTextString(m) }
func (*NodeList) ProtoMessage()               {}
func (*NodeList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NodeList) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// Gossip node
type Node struct {
	Address *Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Node) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

// Gossip node address
type Address struct {
	Host []byte `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Address) GetHost() []byte {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Address) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type NumberMessage struct {
	Num int64 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
}

func (m *NumberMessage) Reset()                    { *m = NumberMessage{} }
func (m *NumberMessage) String() string            { return proto.CompactTextString(m) }
func (*NumberMessage) ProtoMessage()               {}
func (*NumberMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *NumberMessage) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type BytesMessage struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BytesMessage) Reset()                    { *m = BytesMessage{} }
func (m *BytesMessage) String() string            { return proto.CompactTextString(m) }
func (*BytesMessage) ProtoMessage()               {}
func (*BytesMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *BytesMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Return)(nil), "protocol.Return")
	proto.RegisterType((*AccountList)(nil), "protocol.AccountList")
	proto.RegisterType((*WitnessList)(nil), "protocol.WitnessList")
	proto.RegisterType((*AssetIssueList)(nil), "protocol.AssetIssueList")
	proto.RegisterType((*NodeList)(nil), "protocol.NodeList")
	proto.RegisterType((*Node)(nil), "protocol.Node")
	proto.RegisterType((*Address)(nil), "protocol.Address")
	proto.RegisterType((*EmptyMessage)(nil), "protocol.EmptyMessage")
	proto.RegisterType((*NumberMessage)(nil), "protocol.NumberMessage")
	proto.RegisterType((*BytesMessage)(nil), "protocol.BytesMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Wallet service

type WalletClient interface {
	GetAccount(ctx context.Context, in *protocol1.Account, opts ...grpc.CallOption) (*protocol1.Account, error)
	CreateTransaction(ctx context.Context, in *protocol2.TransferContract, opts ...grpc.CallOption) (*protocol1.Transaction, error)
	BroadcastTransaction(ctx context.Context, in *protocol1.Transaction, opts ...grpc.CallOption) (*Return, error)
	ListAccounts(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*AccountList, error)
	UpdateAccount(ctx context.Context, in *protocol2.AccountUpdateContract, opts ...grpc.CallOption) (*protocol1.Transaction, error)
	CreateAccount(ctx context.Context, in *protocol2.AccountCreateContract, opts ...grpc.CallOption) (*protocol1.Transaction, error)
	VoteWitnessAccount(ctx context.Context, in *protocol2.VoteWitnessContract, opts ...grpc.CallOption) (*protocol1.Transaction, error)
	CreateAssetIssue(ctx context.Context, in *protocol2.AssetIssueContract, opts ...grpc.CallOption) (*protocol1.Transaction, error)
	ListWitnesses(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*WitnessList, error)
	UpdateWitness(ctx context.Context, in *protocol2.WitnessUpdateContract, opts ...grpc.CallOption) (*protocol1.Transaction, error)
	CreateWitness(ctx context.Context, in *protocol2.WitnessCreateContract, opts ...grpc.CallOption) (*protocol1.Transaction, error)
	TransferAsset(ctx context.Context, in *protocol2.TransferAssetContract, opts ...grpc.CallOption) (*protocol1.Transaction, error)
	ParticipateAssetIssue(ctx context.Context, in *protocol2.ParticipateAssetIssueContract, opts ...grpc.CallOption) (*protocol1.Transaction, error)
	ListNodes(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*NodeList, error)
	GetAssetIssueList(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*AssetIssueList, error)
	GetAssetIssueByAccount(ctx context.Context, in *protocol1.Account, opts ...grpc.CallOption) (*AssetIssueList, error)
	GetAssetIssueByName(ctx context.Context, in *BytesMessage, opts ...grpc.CallOption) (*protocol2.AssetIssueContract, error)
	GetNowBlock(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*protocol1.Block, error)
	GetBlockByNum(ctx context.Context, in *NumberMessage, opts ...grpc.CallOption) (*protocol1.Block, error)
	TotalTransaction(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*NumberMessage, error)
}

type walletClient struct {
	cc *grpc.ClientConn
}

func NewWalletClient(cc *grpc.ClientConn) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) GetAccount(ctx context.Context, in *protocol1.Account, opts ...grpc.CallOption) (*protocol1.Account, error) {
	out := new(protocol1.Account)
	err := grpc.Invoke(ctx, "/protocol.Wallet/GetAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) CreateTransaction(ctx context.Context, in *protocol2.TransferContract, opts ...grpc.CallOption) (*protocol1.Transaction, error) {
	out := new(protocol1.Transaction)
	err := grpc.Invoke(ctx, "/protocol.Wallet/CreateTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) BroadcastTransaction(ctx context.Context, in *protocol1.Transaction, opts ...grpc.CallOption) (*Return, error) {
	out := new(Return)
	err := grpc.Invoke(ctx, "/protocol.Wallet/BroadcastTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ListAccounts(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*AccountList, error) {
	out := new(AccountList)
	err := grpc.Invoke(ctx, "/protocol.Wallet/ListAccounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) UpdateAccount(ctx context.Context, in *protocol2.AccountUpdateContract, opts ...grpc.CallOption) (*protocol1.Transaction, error) {
	out := new(protocol1.Transaction)
	err := grpc.Invoke(ctx, "/protocol.Wallet/UpdateAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) CreateAccount(ctx context.Context, in *protocol2.AccountCreateContract, opts ...grpc.CallOption) (*protocol1.Transaction, error) {
	out := new(protocol1.Transaction)
	err := grpc.Invoke(ctx, "/protocol.Wallet/CreateAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) VoteWitnessAccount(ctx context.Context, in *protocol2.VoteWitnessContract, opts ...grpc.CallOption) (*protocol1.Transaction, error) {
	out := new(protocol1.Transaction)
	err := grpc.Invoke(ctx, "/protocol.Wallet/VoteWitnessAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) CreateAssetIssue(ctx context.Context, in *protocol2.AssetIssueContract, opts ...grpc.CallOption) (*protocol1.Transaction, error) {
	out := new(protocol1.Transaction)
	err := grpc.Invoke(ctx, "/protocol.Wallet/CreateAssetIssue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ListWitnesses(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*WitnessList, error) {
	out := new(WitnessList)
	err := grpc.Invoke(ctx, "/protocol.Wallet/ListWitnesses", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) UpdateWitness(ctx context.Context, in *protocol2.WitnessUpdateContract, opts ...grpc.CallOption) (*protocol1.Transaction, error) {
	out := new(protocol1.Transaction)
	err := grpc.Invoke(ctx, "/protocol.Wallet/UpdateWitness", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) CreateWitness(ctx context.Context, in *protocol2.WitnessCreateContract, opts ...grpc.CallOption) (*protocol1.Transaction, error) {
	out := new(protocol1.Transaction)
	err := grpc.Invoke(ctx, "/protocol.Wallet/CreateWitness", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) TransferAsset(ctx context.Context, in *protocol2.TransferAssetContract, opts ...grpc.CallOption) (*protocol1.Transaction, error) {
	out := new(protocol1.Transaction)
	err := grpc.Invoke(ctx, "/protocol.Wallet/TransferAsset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ParticipateAssetIssue(ctx context.Context, in *protocol2.ParticipateAssetIssueContract, opts ...grpc.CallOption) (*protocol1.Transaction, error) {
	out := new(protocol1.Transaction)
	err := grpc.Invoke(ctx, "/protocol.Wallet/ParticipateAssetIssue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ListNodes(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*NodeList, error) {
	out := new(NodeList)
	err := grpc.Invoke(ctx, "/protocol.Wallet/ListNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetAssetIssueList(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*AssetIssueList, error) {
	out := new(AssetIssueList)
	err := grpc.Invoke(ctx, "/protocol.Wallet/GetAssetIssueList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetAssetIssueByAccount(ctx context.Context, in *protocol1.Account, opts ...grpc.CallOption) (*AssetIssueList, error) {
	out := new(AssetIssueList)
	err := grpc.Invoke(ctx, "/protocol.Wallet/GetAssetIssueByAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetAssetIssueByName(ctx context.Context, in *BytesMessage, opts ...grpc.CallOption) (*protocol2.AssetIssueContract, error) {
	out := new(protocol2.AssetIssueContract)
	err := grpc.Invoke(ctx, "/protocol.Wallet/GetAssetIssueByName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetNowBlock(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*protocol1.Block, error) {
	out := new(protocol1.Block)
	err := grpc.Invoke(ctx, "/protocol.Wallet/GetNowBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetBlockByNum(ctx context.Context, in *NumberMessage, opts ...grpc.CallOption) (*protocol1.Block, error) {
	out := new(protocol1.Block)
	err := grpc.Invoke(ctx, "/protocol.Wallet/GetBlockByNum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) TotalTransaction(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*NumberMessage, error) {
	out := new(NumberMessage)
	err := grpc.Invoke(ctx, "/protocol.Wallet/TotalTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Wallet service

type WalletServer interface {
	GetAccount(context.Context, *protocol1.Account) (*protocol1.Account, error)
	CreateTransaction(context.Context, *protocol2.TransferContract) (*protocol1.Transaction, error)
	BroadcastTransaction(context.Context, *protocol1.Transaction) (*Return, error)
	ListAccounts(context.Context, *EmptyMessage) (*AccountList, error)
	UpdateAccount(context.Context, *protocol2.AccountUpdateContract) (*protocol1.Transaction, error)
	CreateAccount(context.Context, *protocol2.AccountCreateContract) (*protocol1.Transaction, error)
	VoteWitnessAccount(context.Context, *protocol2.VoteWitnessContract) (*protocol1.Transaction, error)
	CreateAssetIssue(context.Context, *protocol2.AssetIssueContract) (*protocol1.Transaction, error)
	ListWitnesses(context.Context, *EmptyMessage) (*WitnessList, error)
	UpdateWitness(context.Context, *protocol2.WitnessUpdateContract) (*protocol1.Transaction, error)
	CreateWitness(context.Context, *protocol2.WitnessCreateContract) (*protocol1.Transaction, error)
	TransferAsset(context.Context, *protocol2.TransferAssetContract) (*protocol1.Transaction, error)
	ParticipateAssetIssue(context.Context, *protocol2.ParticipateAssetIssueContract) (*protocol1.Transaction, error)
	ListNodes(context.Context, *EmptyMessage) (*NodeList, error)
	GetAssetIssueList(context.Context, *EmptyMessage) (*AssetIssueList, error)
	GetAssetIssueByAccount(context.Context, *protocol1.Account) (*AssetIssueList, error)
	GetAssetIssueByName(context.Context, *BytesMessage) (*protocol2.AssetIssueContract, error)
	GetNowBlock(context.Context, *EmptyMessage) (*protocol1.Block, error)
	GetBlockByNum(context.Context, *NumberMessage) (*protocol1.Block, error)
	TotalTransaction(context.Context, *EmptyMessage) (*NumberMessage, error)
}

func RegisterWalletServer(s *grpc.Server, srv WalletServer) {
	s.RegisterService(&_Wallet_serviceDesc, srv)
}

func _Wallet_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol1.Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAccount(ctx, req.(*protocol1.Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol2.TransferContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CreateTransaction(ctx, req.(*protocol2.TransferContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_BroadcastTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol1.Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).BroadcastTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/BroadcastTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).BroadcastTransaction(ctx, req.(*protocol1.Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ListAccounts(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol2.AccountUpdateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).UpdateAccount(ctx, req.(*protocol2.AccountUpdateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol2.AccountCreateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CreateAccount(ctx, req.(*protocol2.AccountCreateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_VoteWitnessAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol2.VoteWitnessContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).VoteWitnessAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/VoteWitnessAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).VoteWitnessAccount(ctx, req.(*protocol2.VoteWitnessContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_CreateAssetIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol2.AssetIssueContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CreateAssetIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/CreateAssetIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CreateAssetIssue(ctx, req.(*protocol2.AssetIssueContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ListWitnesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ListWitnesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/ListWitnesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ListWitnesses(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_UpdateWitness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol2.WitnessUpdateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).UpdateWitness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/UpdateWitness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).UpdateWitness(ctx, req.(*protocol2.WitnessUpdateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_CreateWitness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol2.WitnessCreateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CreateWitness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/CreateWitness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CreateWitness(ctx, req.(*protocol2.WitnessCreateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_TransferAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol2.TransferAssetContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).TransferAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/TransferAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).TransferAsset(ctx, req.(*protocol2.TransferAssetContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ParticipateAssetIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol2.ParticipateAssetIssueContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ParticipateAssetIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/ParticipateAssetIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ParticipateAssetIssue(ctx, req.(*protocol2.ParticipateAssetIssueContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ListNodes(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetAssetIssueList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAssetIssueList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/GetAssetIssueList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAssetIssueList(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetAssetIssueByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol1.Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAssetIssueByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/GetAssetIssueByAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAssetIssueByAccount(ctx, req.(*protocol1.Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetAssetIssueByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BytesMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAssetIssueByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/GetAssetIssueByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAssetIssueByName(ctx, req.(*BytesMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetNowBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetNowBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/GetNowBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetNowBlock(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetBlockByNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetBlockByNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/GetBlockByNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetBlockByNum(ctx, req.(*NumberMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_TotalTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).TotalTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Wallet/TotalTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).TotalTransaction(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wallet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _Wallet_GetAccount_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _Wallet_CreateTransaction_Handler,
		},
		{
			MethodName: "BroadcastTransaction",
			Handler:    _Wallet_BroadcastTransaction_Handler,
		},
		{
			MethodName: "ListAccounts",
			Handler:    _Wallet_ListAccounts_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _Wallet_UpdateAccount_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Wallet_CreateAccount_Handler,
		},
		{
			MethodName: "VoteWitnessAccount",
			Handler:    _Wallet_VoteWitnessAccount_Handler,
		},
		{
			MethodName: "CreateAssetIssue",
			Handler:    _Wallet_CreateAssetIssue_Handler,
		},
		{
			MethodName: "ListWitnesses",
			Handler:    _Wallet_ListWitnesses_Handler,
		},
		{
			MethodName: "UpdateWitness",
			Handler:    _Wallet_UpdateWitness_Handler,
		},
		{
			MethodName: "CreateWitness",
			Handler:    _Wallet_CreateWitness_Handler,
		},
		{
			MethodName: "TransferAsset",
			Handler:    _Wallet_TransferAsset_Handler,
		},
		{
			MethodName: "ParticipateAssetIssue",
			Handler:    _Wallet_ParticipateAssetIssue_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _Wallet_ListNodes_Handler,
		},
		{
			MethodName: "GetAssetIssueList",
			Handler:    _Wallet_GetAssetIssueList_Handler,
		},
		{
			MethodName: "GetAssetIssueByAccount",
			Handler:    _Wallet_GetAssetIssueByAccount_Handler,
		},
		{
			MethodName: "GetAssetIssueByName",
			Handler:    _Wallet_GetAssetIssueByName_Handler,
		},
		{
			MethodName: "GetNowBlock",
			Handler:    _Wallet_GetNowBlock_Handler,
		},
		{
			MethodName: "GetBlockByNum",
			Handler:    _Wallet_GetBlockByNum_Handler,
		},
		{
			MethodName: "TotalTransaction",
			Handler:    _Wallet_TotalTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

// Client API for Database service

type DatabaseClient interface {
}

type databaseClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseClient(cc *grpc.ClientConn) DatabaseClient {
	return &databaseClient{cc}
}

// Server API for Database service

type DatabaseServer interface {
}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
}

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api/api.proto",
}

// Client API for Network service

type NetworkClient interface {
}

type networkClient struct {
	cc *grpc.ClientConn
}

func NewNetworkClient(cc *grpc.ClientConn) NetworkClient {
	return &networkClient{cc}
}

// Server API for Network service

type NetworkServer interface {
}

func RegisterNetworkServer(s *grpc.Server, srv NetworkServer) {
	s.RegisterService(&_Network_serviceDesc, srv)
}

var _Network_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Network",
	HandlerType: (*NetworkServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api/api.proto",
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xed, 0x6e, 0xf2, 0x36,
	0x14, 0xc7, 0xe1, 0xe9, 0x53, 0xa0, 0x07, 0xd2, 0x52, 0xf7, 0x65, 0x88, 0x75, 0x1a, 0xb3, 0x2a,
	0x0d, 0x6d, 0x2a, 0x59, 0xe9, 0x97, 0xbd, 0x54, 0xd3, 0xa0, 0xab, 0x18, 0x52, 0x8b, 0xaa, 0xa8,
	0x5b, 0xb5, 0x7e, 0x33, 0xc1, 0xa3, 0x51, 0x43, 0x1c, 0xd9, 0x4e, 0x11, 0x5f, 0x77, 0x0b, 0xbb,
	0xa2, 0x5d, 0xc3, 0x6e, 0x61, 0x17, 0x32, 0xc5, 0x4e, 0xc0, 0xb4, 0x04, 0xba, 0x4f, 0xd8, 0xe7,
	0xe5, 0xc7, 0xf1, 0xdf, 0xc7, 0x47, 0x01, 0x8b, 0x84, 0x9e, 0x4d, 0x42, 0xaf, 0x15, 0x72, 0x26,
	0x19, 0x2a, 0xa9, 0x1f, 0x97, 0xf9, 0xf5, 0x3d, 0x97, 0x71, 0x6a, 0xdf, 0x73, 0x16, 0x68, 0x57,
	0xfd, 0x40, 0x19, 0xae, 0x58, 0x20, 0x39, 0x71, 0x65, 0x62, 0x3c, 0x19, 0x33, 0x36, 0xf6, 0xa9,
	0xad, 0x28, 0x41, 0xc0, 0x24, 0x91, 0x1e, 0x0b, 0x84, 0xf6, 0xe2, 0x06, 0x14, 0x1c, 0x2a, 0x23,
	0x1e, 0xa0, 0x63, 0x28, 0x70, 0x2a, 0x22, 0x5f, 0xd6, 0xf2, 0x8d, 0x7c, 0xb3, 0xe4, 0x24, 0x3b,
	0x7c, 0x09, 0xe5, 0x8e, 0xeb, 0xb2, 0x28, 0x90, 0x37, 0x9e, 0x90, 0xe8, 0x0c, 0x4a, 0x44, 0x6f,
	0x45, 0x2d, 0xdf, 0xd8, 0x6a, 0x96, 0xdb, 0xfb, 0xad, 0xb4, 0xa2, 0x56, 0x12, 0xe8, 0xcc, 0x43,
	0xf0, 0x8f, 0x50, 0x7e, 0xf0, 0x64, 0x40, 0x85, 0x50, 0xd9, 0x36, 0xec, 0x4c, 0xf5, 0x96, 0xae,
	0x48, 0x4f, 0x22, 0x9d, 0x45, 0x0c, 0x1e, 0xc0, 0x6e, 0x47, 0x08, 0x2a, 0xfb, 0x42, 0x44, 0x54,
	0x21, 0x2e, 0x01, 0xc8, 0xdc, 0x92, 0x30, 0x4e, 0x8c, 0x12, 0xe6, 0xbe, 0x54, 0x07, 0xc7, 0x88,
	0xc7, 0xdf, 0x40, 0x69, 0xc0, 0x46, 0x9a, 0x74, 0x0a, 0xdb, 0x01, 0x1b, 0xcd, 0x0b, 0xd9, 0x5d,
	0x40, 0xe2, 0x10, 0x47, 0x3b, 0xf1, 0x05, 0x7c, 0x8c, 0xb7, 0xe8, 0x6b, 0x28, 0x92, 0xd1, 0x88,
	0x53, 0x21, 0x94, 0x40, 0xcb, 0xe7, 0xd6, 0x0e, 0x27, 0x8d, 0xc0, 0xe7, 0x50, 0x4c, 0x6c, 0x08,
	0xc1, 0xc7, 0x27, 0x26, 0xb4, 0xaa, 0x15, 0x47, 0xad, 0x63, 0x5b, 0xc8, 0xb8, 0xac, 0x7d, 0x68,
	0xe4, 0x9b, 0xdb, 0x8e, 0x5a, 0xe3, 0x5d, 0xa8, 0x5c, 0x4f, 0x42, 0x39, 0xbb, 0xa5, 0x42, 0x90,
	0x31, 0xc5, 0x5f, 0x80, 0x35, 0x88, 0x26, 0x43, 0xca, 0x13, 0x03, 0xaa, 0xc2, 0x56, 0x10, 0x4d,
	0x14, 0x67, 0xcb, 0x89, 0x97, 0xf8, 0x14, 0x2a, 0xdd, 0x99, 0xa4, 0x22, 0x8d, 0x38, 0x84, 0xed,
	0x17, 0xe2, 0x2b, 0x55, 0xe2, 0xff, 0xd2, 0x9b, 0xf6, 0xdf, 0x65, 0x28, 0x3c, 0x10, 0xdf, 0xa7,
	0x12, 0xf5, 0x00, 0x7a, 0x54, 0x26, 0xb7, 0x84, 0xde, 0x5e, 0x5c, 0xfd, 0xad, 0x09, 0x1f, 0xfc,
	0xf9, 0xcf, 0xbf, 0x7f, 0x7d, 0xb0, 0x70, 0xc9, 0x4e, 0x2e, 0xf5, 0xfb, 0xfc, 0x57, 0xe8, 0x17,
	0xd8, 0xbf, 0xe2, 0x94, 0x48, 0x7a, 0xcf, 0x49, 0x20, 0x88, 0x1b, 0xb7, 0x14, 0xaa, 0x2f, 0x92,
	0x95, 0xf9, 0x0f, 0xca, 0xd3, 0x3b, 0xa8, 0x1f, 0xbd, 0xf2, 0xe9, 0x14, 0x9c, 0x43, 0x1d, 0x38,
	0xec, 0x72, 0x46, 0x46, 0x2e, 0x11, 0xd2, 0x84, 0xad, 0x4e, 0xa8, 0x57, 0x17, 0x66, 0xdd, 0xb7,
	0x38, 0x87, 0x1e, 0xa1, 0x12, 0xdf, 0x67, 0x52, 0xb0, 0x40, 0xc7, 0x8b, 0x18, 0x53, 0x51, 0xb3,
	0x06, 0xa3, 0xa3, 0xf1, 0xa7, 0xea, 0x80, 0x47, 0xb8, 0x6a, 0xbf, 0x9c, 0xa7, 0x67, 0xb4, 0x7d,
	0x4f, 0xa8, 0x83, 0xf6, 0xc1, 0xfa, 0x35, 0x1c, 0x11, 0x49, 0x53, 0xd1, 0x3e, 0x7f, 0x03, 0xd1,
	0xfe, 0xcd, 0x27, 0xed, 0x83, 0xa5, 0x35, 0xcb, 0x46, 0x69, 0xff, 0x66, 0xd4, 0x0d, 0xa0, 0xdf,
	0x98, 0xa4, 0xc9, 0x7b, 0x49, 0x79, 0x9f, 0x2d, 0xc2, 0x0d, 0xef, 0x7b, 0x0a, 0xab, 0x26, 0x85,
	0xcd, 0xdf, 0x09, 0x5a, 0xfb, 0xa2, 0xb2, 0x51, 0x3f, 0x81, 0x15, 0x6b, 0xfa, 0x90, 0xbe, 0xdf,
	0xf7, 0xdc, 0x85, 0x31, 0x1f, 0xb4, 0x4a, 0x5a, 0xd0, 0xc4, 0x6c, 0xaa, 0x94, 0x98, 0xfe, 0xb7,
	0xe0, 0xd9, 0xa8, 0xf7, 0x0a, 0xde, 0x07, 0x2b, 0x6d, 0x69, 0x25, 0x87, 0x89, 0x5a, 0x72, 0x6c,
	0x46, 0xfd, 0x0e, 0x47, 0x77, 0x84, 0x4b, 0xcf, 0xf5, 0xc2, 0x65, 0xc9, 0xbf, 0x5c, 0x64, 0xac,
	0x0c, 0xd8, 0x8c, 0xfe, 0x0e, 0x76, 0x62, 0x15, 0xe3, 0x71, 0x95, 0xad, 0x3c, 0x5a, 0x1e, 0x73,
	0x89, 0xec, 0x3d, 0xd8, 0x8f, 0x27, 0xc3, 0xf2, 0xa8, 0xcd, 0x42, 0xd4, 0x56, 0x35, 0xc7, 0x1c,
	0x74, 0xbc, 0x04, 0xea, 0xce, 0xd6, 0x8c, 0x9b, 0x75, 0xa0, 0x5b, 0x38, 0x78, 0x05, 0x1a, 0x90,
	0x09, 0x35, 0x6b, 0x32, 0x67, 0x5f, 0x7d, 0x6d, 0xc3, 0xe2, 0x1c, 0xfa, 0x16, 0xca, 0x3d, 0x2a,
	0x07, 0x6c, 0xda, 0xf5, 0x99, 0xfb, 0x9c, 0x79, 0xb4, 0x3d, 0x03, 0x1f, 0x07, 0xe2, 0x1c, 0xfa,
	0x01, 0xac, 0x1e, 0x95, 0x6a, 0xd7, 0x9d, 0x0d, 0xa2, 0x09, 0xfa, 0xc4, 0x50, 0xd0, 0x9c, 0xd0,
	0xab, 0x92, 0xaf, 0xa1, 0x7a, 0xcf, 0x24, 0xf1, 0xcd, 0xd1, 0x96, 0xf5, 0xdf, 0x59, 0x5c, 0x9c,
	0x6b, 0x03, 0x94, 0x7e, 0x26, 0x92, 0x0c, 0x89, 0xa0, 0xed, 0x1d, 0x28, 0x0e, 0xa8, 0x9c, 0x32,
	0xfe, 0xdc, 0xed, 0x40, 0x85, 0xf1, 0x71, 0x4b, 0xc6, 0x9f, 0x00, 0x24, 0xf4, 0xba, 0xc5, 0x1e,
	0x0f, 0xdd, 0xce, 0x5d, 0xff, 0xb1, 0x39, 0xf6, 0xe4, 0x53, 0x34, 0x6c, 0xb9, 0x6c, 0x62, 0xc7,
	0xde, 0x14, 0x6c, 0x8f, 0x79, 0xe8, 0x9e, 0x8d, 0x89, 0xa4, 0x53, 0x32, 0x8b, 0x3f, 0x09, 0x86,
	0x05, 0xe5, 0xba, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xee, 0xb5, 0x33, 0x66, 0x08, 0x00,
	0x00,
}

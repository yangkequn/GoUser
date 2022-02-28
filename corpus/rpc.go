// Code generated by goctl. DO NOT EDIT!
// Source: corpus.proto

package corpus

import (
	"context"


	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (

	Rpc interface {
		MergeUserTheme(ctx context.Context, in *MergeUserThemeRequest, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultRpc struct {
		cli zrpc.Client
	}
)

func NewRpc(cli zrpc.Client) Rpc {
	return &defaultRpc{
		cli: cli,
	}
}

func (m *defaultRpc) MergeUserTheme(ctx context.Context, in *MergeUserThemeRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := NewRpcClient(m.cli.Conn())
	return client.MergeUserTheme(ctx, in, opts...)
}
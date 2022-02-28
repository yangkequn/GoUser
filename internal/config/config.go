package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	DBPostgresql struct {
		Connection string
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
	UserRpc zrpc.RpcClientConf
}

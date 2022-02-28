package svc

import (
	"user/internal/config"
	"user/model"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UsersModel  model.UsersModel
	RedisClient *redis.Client
	UserRpc     zrpc.RpcClientConf
}

func NewServiceContext(c config.Config) *ServiceContext {
	connPostGres := sqlx.NewSqlConn("postgres", c.DBPostgresql.Connection)
	redisOptions := &redis.Options{
		Addr:     c.Redis.Addr,     // use default Addr
		Password: c.Redis.Password, // no password set
		DB:       c.Redis.DB,       // use default DB
	}
	return &ServiceContext{
		Config:      c,
		UsersModel:  model.NewUsersModel(connPostGres),
		RedisClient: redis.NewClient(redisOptions),
	}
}

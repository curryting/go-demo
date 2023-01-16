package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"wx-chat/service/user/cmd/api/internal/config"
	"wx-chat/service/user/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: usercenter.NewUserCenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}

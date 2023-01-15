package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"wx-chat/service/chat/cmd/api/internal/config"
	"wx-chat/service/chat/cmd/rpc/chatcenter"
)

type ServiceContext struct {
	Config        config.Config
	ChatRpcClient chatcenter.ChatCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ChatRpcClient: chatcenter.NewChatCenter(zrpc.MustNewClient(c.ChatRpcConf)),
	}
}

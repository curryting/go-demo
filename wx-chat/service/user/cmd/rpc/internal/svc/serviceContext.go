package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wx-chat/service/user/cmd/rpc/internal/config"
	"wx-chat/service/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.WxUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewWxUserModel(conn),
	}
}

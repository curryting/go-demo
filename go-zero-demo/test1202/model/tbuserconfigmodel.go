package model

import (
	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbUserConfigModel = (*customTbUserConfigModel)(nil)

type (
	// TbUserConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserConfigModel.
	TbUserConfigModel interface {
		tbUserConfigModel
	}

	customTbUserConfigModel struct {
		*defaultTbUserConfigModel
	}
)

// NewTbUserConfigModel returns a model for the database table.
func NewTbUserConfigModel(conn sqlx.SqlConn, c cache.CacheConf) TbUserConfigModel {
	return &customTbUserConfigModel{
		defaultTbUserConfigModel: newTbUserConfigModel(conn, c),
	}
}

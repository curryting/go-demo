package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WxUserModel = (*customWxUserModel)(nil)

type (
	// WxUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWxUserModel.
	WxUserModel interface {
		wxUserModel
	}

	customWxUserModel struct {
		*defaultWxUserModel
	}
)

// NewWxUserModel returns a model for the database table.
func NewWxUserModel(conn sqlx.SqlConn) WxUserModel {
	return &customWxUserModel{
		defaultWxUserModel: newWxUserModel(conn),
	}
}

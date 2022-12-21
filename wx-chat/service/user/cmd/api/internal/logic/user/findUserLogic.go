package user

import (
	"context"

	"wx-chat/service/user/cmd/api/internal/svc"
	"wx-chat/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserLogic) FindUser(req *types.FindUserReq) (resp *types.FindUserRes, err error) {
	// todo: add your logic here and delete this line

	return
}

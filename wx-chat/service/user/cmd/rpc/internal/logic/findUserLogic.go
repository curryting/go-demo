package logic

import (
	"context"

	"wx-chat/service/user/cmd/rpc/internal/svc"
	"wx-chat/service/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *__.UserReq) (*__.UserRes, error) {
	// todo: add your logic here and delete this line

	return &__.UserRes{}, nil
}

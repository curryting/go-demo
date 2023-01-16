package logic

import (
	"context"

	"wx-chat/service/user/cmd/rpc/internal/svc"
	"wx-chat/service/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *__.UserReq) (*__.CommonRes, error) {
	// todo: add your logic here and delete this line

	return &__.CommonRes{}, nil
}

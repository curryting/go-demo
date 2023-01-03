package logic

import (
	"context"

	"user-rpc/internal/svc"
	"user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line

	m := map[int64]string{
		1: "张三",
		2: "赵六",
	}
	nickname := "unknown"
	if name, ok := m[in.Id]; ok {
		nickname = name
	}

	return &pb.GetUserInfoResp{
		Id:       in.Id,
		Nickname: nickname,
	}, nil
}

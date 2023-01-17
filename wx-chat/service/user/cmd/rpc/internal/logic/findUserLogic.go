package logic

import (
	"context"
	"errors"
	"fmt"

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
	if in.Uid < 0 {
		return nil, errors.New("非法的uid")
	}
	wxUser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Uid)
	if err != nil {
		fmt.Println("rpcFindUser err返回为: ", err)
		return nil, err
	}

	return &__.UserRes{
		Uid:      wxUser.Id,
		Username: wxUser.Username,
		Password: wxUser.Password,
		Gender:   wxUser.Gender,
	}, nil
}

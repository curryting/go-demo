package user

import (
	"context"
	"errors"
	"log"
	"strings"
	"wx-chat/service/user/model"

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
	if len(strings.TrimSpace(req.Username)) == 0 {
		return nil, errors.New("参数错误")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	log.Println("查询用户信息结果：", userInfo)
	if err == model.ErrNotFound {
		return nil, errors.New("用户名不存在")
	}
	if err != nil {
		return nil, err
	}

	return &types.FindUserRes{
		Id:       userInfo.Id,
		Username: userInfo.Username,
		Password: userInfo.Password,
		Gender:   userInfo.Gender,
	}, nil
}

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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	log.Println("查询用户信息结果：", userInfo)
	if err != nil {
		return nil, err
	}
	if err != model.ErrNotFound {
		return nil, errors.New("用户名已存在")
	}

	insertRes, err := l.svcCtx.UserModel.Insert(l.ctx, &model.WxUser{
		Username: req.Username,
		Password: req.Password,
		Gender:   req.Gender,
	})
	if err != nil {
		return nil, err
	}
	log.Println("注册结果返回：", insertRes)
	return resp, nil
}

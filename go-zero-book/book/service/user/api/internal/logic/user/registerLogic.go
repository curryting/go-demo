package user

import (
	"book/service/user/model"
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"

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
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Name)) == 0 && len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}
	res, err1 := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Number:     req.Number,
		Name:       req.Name,
		Password:   req.Password,
		Gender:     req.Gender,
		CreateTime: time.Time{},
		UpdateTime: time.Time{},
	})
	log.Println("res is ", res)
	if err1 != nil {
		return nil, errors.New("插入失败")
	}
	return &types.RegisterRes{
		Msg: "success",
	}, nil
}

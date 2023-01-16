package logic

import (
	"context"
	"strings"
	"time"
	"wx-chat/service/user/model"

	"wx-chat/service/user/cmd/rpc/internal/svc"
	"wx-chat/service/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *__.AddUserReq) (*__.CommonRes, error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(in.Username)) == 0 {
		return CommonErrorRes(100011, "Username不能为空"), nil
	}
	if len(strings.TrimSpace(in.Username)) > 10 {
		return CommonErrorRes(100012, "Username过长"), nil
	}
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err == model.ErrNotFound {
		l.svcCtx.UserModel.Insert(l.ctx, &model.WxUser{
			Username:   in.Username,
			Password:   "",
			Gender:     in.Gender,
			CreateTime: time.Time{},
			UpdateTime: time.Time{},
		})
	} else if err == nil {
		return CommonErrorRes(100014, "username已存在"), nil
	} else {
		return nil, err
	}

	return CommonErrorRes(0, "新增成功"), nil
}

func CommonErrorRes(iRet int64, sMsg string) *__.CommonRes {
	return &__.CommonRes{
		IRet: iRet,
		SMsg: sMsg,
	}
}

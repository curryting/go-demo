package user

import (
	"context"
	"fmt"
	"wx-chat/service/user/cmd/api/internal/svc"
	"wx-chat/service/user/cmd/api/internal/types"
	__ "wx-chat/service/user/cmd/rpc/pb"

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
	if req.UserId < 0 {
		return commonResponse(100001, "userId异常", types.FindUser{}), nil
	}
	rpcFindUser, err1 := l.svcCtx.UserRpcClient.FindUser(l.ctx, &__.UserReq{Uid: req.UserId})
	if err1 != nil {
		fmt.Println("rpcFindUser err1返回为: ", err1)
		return commonResponse(100001, err1.Error(), types.FindUser{}), nil
	}
	return commonResponse(0, "success", types.FindUser{
		Id:       rpcFindUser.Uid,
		Username: rpcFindUser.Username,
		Password: "",
		Gender:   rpcFindUser.Gender,
	}), nil
}

func commonResponse(IRet int64, SMsg string, SData types.FindUser) *types.FindUserRes {
	return &types.FindUserRes{
		IRet:  IRet,
		SMsg:  SMsg,
		SData: SData,
	}
}

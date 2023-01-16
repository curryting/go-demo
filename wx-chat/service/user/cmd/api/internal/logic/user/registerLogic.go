package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"wx-chat/common"
	"wx-chat/service/user/cmd/api/internal/svc"
	"wx-chat/service/user/cmd/api/internal/types"
	__ "wx-chat/service/user/cmd/rpc/pb"
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

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *common.HttpResponse, err error) {
	//if len(strings.TrimSpace(req.Username)) == 0 {
	//	return common.JsonResponse(100011, "Username不能为空"), nil
	//}
	//if len(strings.TrimSpace(req.Username)) > 10 {
	//	return common.JsonResponse(100012, "Username过长"), nil
	//}
	userRpcRes, error := l.svcCtx.UserRpcClient.AddUser(l.ctx, &__.AddUserReq{
		Username: req.Username,
		Gender:   req.Gender,
	})
	if error != nil {
		return nil, err
	}
	return common.JsonResponse(userRpcRes.IRet, userRpcRes.SMsg), nil
}

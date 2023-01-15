package logic

import (
	"context"

	"wx-chat/service/chat/cmd/rpc/internal/svc"
	"wx-chat/service/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatListLogic {
	return &GetChatListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChatListLogic) GetChatList(in *pb.ChatListReq) (*pb.ChatListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ChatListResp{List: []*pb.ChatContent{
		{
			Uid:      in.Uid,
			Username: "老婆大人",
			Content:  "大猪蹄子",
		},
	}}, nil
}

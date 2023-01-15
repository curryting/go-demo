package chat

import (
	"context"
	"wx-chat/service/chat/cmd/rpc/pb"

	"wx-chat/service/chat/cmd/api/internal/svc"
	"wx-chat/service/chat/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ChatListReq) (resp *types.ChatListRes, err error) {
	// todo: add your logic here and delete this line
	_chatList, err := l.svcCtx.ChatRpcClient.GetChatList(l.ctx, &pb.ChatListReq{
		Uid: int64(req.Uid),
	})
	if err != nil {
		return nil, err
	}
	chatTemp := make([]types.ChatList, len(_chatList.List))
	for key, value := range _chatList.List {
		chatTemp[key].Username = value.Username
		chatTemp[key].Uid = int(value.Uid)
		chatTemp[key].Content = value.Content
	}

	chatList := types.ChatListRes{List: chatTemp}

	//chatList := types.ChatListRes{List: []types.ChatList{
	//	//{
	//	//	Uid:      111,
	//	//	Username: "curryting",
	//	//	Content:  "小伙子有点东西啊",
	//	//},
	//
	//}}
	return &chatList, nil
}

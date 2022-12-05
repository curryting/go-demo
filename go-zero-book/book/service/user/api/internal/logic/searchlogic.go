package logic

import (
	"book/service/user/model"
	"context"
	"errors"
	"log"
	"strings"

	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchReply, err error) {
	// todo: add your logic here and delete this line
	log.Println("name is ", req.Name)
	// 验证参数
	if len(strings.TrimSpace(req.Name)) == 0 {
		return nil, errors.New("参数错误")
	}
	userInfo, err := s.svcCtx.UserModel.FindOneByNumber(s.ctx, req.Name)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户名不存在")
	default:
		return nil, err
	}
	return &types.SearchReply{
		Id:       userInfo.Id,
		Name:     userInfo.Name,
		Gender:   userInfo.Gender,
		Number:   userInfo.Number,
		Password: userInfo.Password,
	}, nil
	return
}

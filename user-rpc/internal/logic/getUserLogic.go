package logic

import (
	"context"

	"go-zore-demo/user-rpc/internal/svc"
	"go-zore-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.IdRequest) (*pb.UserResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.UserResponse{
		Id:     "1",
		Name:   "xs,i am from rpc",
		Gender: "男",
	}, nil
}

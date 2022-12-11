// Code generated by goctl. DO NOT EDIT!
// Source: user-api.proto

package server

import (
	"context"

	"go-zore-demo/user-rpc/internal/logic"
	"go-zore-demo/user-rpc/internal/svc"
	"go-zore-demo/user-rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *pb.IdRequest) (*pb.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

package logic

import (
	"context"

	"go-zore-demo/user-rpc/internal/svc"
	"go-zore-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderLogic) GetOrder(in *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetOrderResponse{Id: "订单id,from rpc"}, nil
}

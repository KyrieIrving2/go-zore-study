package logic

import (
	"context"
	"go-zore-demo/user-api/internal/svc"
	"go-zore-demo/user-api/internal/types"
	"go-zore-demo/user-rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserRsp, err error) {
	// todo: add your logic here and delete this line

	//if err := l.Test1(); err != nil {
	//	logx.Errorf("err:%+v",err)
	//	return nil,err
	//}

	lawyer, err := l.svcCtx.LawyerModel.FindOne(l.ctx, int64(req.Id))
	if err != nil {
		return nil, errors.New("数据库查询错误")
	}

	if lawyer == nil {
		return nil, errors.New("数据不存在")
	}

	userRsp,err:= l.svcCtx.UserRpcClient.GetUser(l.ctx, &user.IdRequest{Id: "1"} )
	if err!=nil{
		return nil, err
	}

	return &types.GetUserRsp{
		Name: userRsp.Name,
	}, nil

}

func (l *GetUserLogic) Test1() error {
	return l.Test2()
}

func (l *GetUserLogic) Test2() error {
	return errors.Wrap(errors.New("error"),"sdfsd")
}

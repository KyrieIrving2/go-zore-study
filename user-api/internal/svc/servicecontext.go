package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zore-demo/user-api/internal/config"
	"go-zore-demo/user-api/internal/middleware"
	"go-zore-demo/user-api/model"
	"go-zore-demo/user-rpc/user"
)

type ServiceContext struct {
	Config config.Config
	LawyerModel model.LawyerModel
	TestMiddleWare rest.Middleware
	UserRpcClient user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		TestMiddleWare: middleware.NewTestMiddleWareMiddleware().Handle,
		LawyerModel: model.NewLawyerModel(sqlx.NewMysql(c.DB.DataSource)),
		UserRpcClient: user.NewUser(zrpc.MustNewClient(c.UserRpcConfig)),
	}
}

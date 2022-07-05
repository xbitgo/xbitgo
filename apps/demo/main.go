package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/xbitgo/components/tracing"
	"github.com/xbitgo/core/di"

	"github.com/xbitgo/components/dtx"
	"github.com/xbitgo/components/sequence"

	"xbitgo/common/middleware"
	"xbitgo/common/server"

	"xbitgo/apps/demo/conf"
	"xbitgo/apps/demo/domain/event"
	"xbitgo/apps/demo/domain/extend"
	"xbitgo/apps/demo/domain/service"
	"xbitgo/apps/demo/handler/entry"
	"xbitgo/apps/demo/repo_impl"
	pb "xbitgo/proto/apps/demo"
)

func main() {
	// 初始化
	Init()
	// handler层代理
	handlerProxy := entry.NewDemoHandler()
	// 创建服务
	app := server.NewApp(conf.App.Server)
	// grpc server
	app.InitGRPC(func(grpcService grpc.ServiceRegistrar) {
		pb.RegisterDemoServer(grpcService, handlerProxy)
	}, tracing.GrpcServerTrace())
	// http server
	app.InitHTTP(func(r *gin.Engine) {
		// 可选 全局中间件
		// r.Use(middleware.HTTPCors())
		// 可选 全局路由
		// r.OPTIONS("/*wild", func(c *gin.Context) {
		//	return
		// })
		// 必须；解析参数可以定制
		r.Use(middleware.HTTPParams())
		{
			entry.DemoHttpInit(r, handlerProxy)
		}
	})
	// 服务停止前处理
	app.OnClose(func() {
		// something close ...
	})
	err := app.Start(conf.App.GetEtcd())
	if err != nil {
		panic(err)
	}
}

func Init() {
	// 配置初始化
	conf.Init()
	// 自定义初始化配置
	conf.CustomInit()
	// 分布式ID生成器初始化
	sequence.Init()
	// 分布式事务管理器初始化
	dtx.Init(nil, 0) // 需要配置mq开启跨服务
	// 注册rpc客户端
	conf.RegisterRPCClients()
	// 注册DI存储层实现
	repo_impl.DIRegister()
	// 注册DI业务拓展层
	extend.DIRegister()
	// 注册DI服务层
	service.DIRegister()
	// event注册
	event.Register()
	// DI注入
	di.MustBindALL()
}

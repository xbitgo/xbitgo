package conf

import (
	"xbitgo/proto/rpc_client"
)

// RegisterRPCClients 注册Rpc客户端
func RegisterRPCClients() {
	err := rpc_client.RegisterRPC(Namespace, App.GetEtcd(), rpc_client.User) // 在这里注册依赖的其他服务 注册后 代码中可以直接使用
	if err != nil {
		panic(err)
	}
}

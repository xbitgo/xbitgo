// Code generated by xbit. DO NOT EDIT.
package conf

import (
	"log"

	"github.com/go-redis/redis/v8"
	etcdv3 "go.etcd.io/etcd/client/v3"

	"github.com/xbitgo/core/di"
)

func _dIRegister() {
	instEtcd, err := App.Etcd.CreateInstance()
	if err != nil {
		log.Panicf("init DI conf.Etcd] err: %v", err)
	}
	di.Register("conf.Etcd", instEtcd)
	instDB, err := App.DB.CreateInstance()
	if err != nil {
		log.Panicf("init DI conf.DB] err: %v", err)
	}
	di.Register("conf.DB", instDB)
	instRedis, err := App.Redis.CreateInstance()
	if err != nil {
		log.Panicf("init DI conf.Redis] err: %v", err)
	}
	di.Register("conf.Redis", instRedis)
}

// GetEtcd .
func (c *Config) GetEtcd() *etcdv3.Client {
	inst := di.GetInst("conf.Etcd")
	if v, ok := inst.(*etcdv3.Client); ok {
		return v
	}
	return nil
}

// GetRedis .
func (c *Config) GetRedis() *redis.Client {
	inst := di.GetInst("conf.Redis")
	if v, ok := inst.(*redis.Client); ok {
		return v
	}
	return nil
}

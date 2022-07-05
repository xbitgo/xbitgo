package conf

import "xbitgo/common/cfg"

// UnitTestingInit 单元测试时的配置初始化
func UnitTestingInit() {
	//Init() 一般配置中心模式可以直接调用正常初始化
	App = &Config{
		Name: "demo_server",
		Env:  "local",
		Server: &cfg.Server{
			HTTPAddr: "0.0.0.0:9901",
			GRPCAddr: "0.0.0.0:9902",
		},
		Tracing: &cfg.Tracing{
			ServiceName:                "demo_server",
			SamplerType:                "const",
			SamplerParam:               1,
			ReporterLocalAgentHostPort: "127.0.0.1:6831",
			LogSpans:                   false,
		},
		Etcd: &cfg.Etcd{
			Endpoints: []string{"127.0.0.1:2379"},
			Timeout:   5,
		},
		DB: &cfg.DB{
			DSN:             "root:@tcp(localhost:3306)/test?charset=utf8mb4&interpolateParams=true&parseTime=true&loc=Local",
			MaxOpenConn:     10,
			MaxIdleConn:     10,
			ConnMaxLifetime: 300,
			ConnMaxIdleTime: 300,
		},
		Redis: &cfg.Redis{
			Addr:     "127.0.0.1:6379",
			Password: "",
			PoolSize: 4,
			DB:       0,
		},
	}
	_dIRegister()
}

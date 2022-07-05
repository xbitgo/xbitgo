package conf

import (
	"log"
	"xbitgo/common/cfg"
)

const Namespace = "demo"

var App = &Config{}

type Config struct {
	Name    string       `json:"name" yaml:"name"`
	Env     string       `json:"env" yaml:"env"`
	Server  *cfg.Server  `json:"server" yaml:"server"`
	Tracing *cfg.Tracing `json:"trace" yaml:"tracing"`
	Etcd    *cfg.Etcd    `json:"etcd" yaml:"etcd" sdi:"etcd""`
	DB      *cfg.DB      `json:"DB" yaml:"DB" sdi:"DB"`
	Redis   *cfg.Redis   `json:"redis" yaml:"redis" sdi:"redis"`
}

func Init() {
	c := cfg.NewCfg(Namespace)
	err := c.Apply(App)
	if err != nil {
		log.Panic(err)
	}
	_dIRegister()
}

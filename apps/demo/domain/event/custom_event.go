package event

import (
	"github.com/xbitgo/core/di"
)

type CustomEvent struct {
}

func NewCustomEvent() *CustomEvent {
	custom := &CustomEvent{}
	di.MustBind(custom)
	return custom
}

func (c *CustomEvent) RegisterFunc() map[string]func(args ...string) error {
	return map[string]func(args ...string) error{
		// "DemoFunc": c.DemoFunc,
	}
}

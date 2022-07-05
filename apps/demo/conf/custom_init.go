package conf

import (
	"os"

	"github.com/gin-gonic/gin/binding"
	"github.com/json-iterator/go/extra"
	"github.com/xbitgo/components/tracing"
	"github.com/xbitgo/core/log"
)

// CustomInit 业务自定义配置
func CustomInit() {
	// JSON配置
	binding.EnableDecoderUseNumber = true
	extra.RegisterFuzzyDecoders()

	// 链路追踪配置
	tracing.InitJaegerTracer(App.Tracing.ToJaegerCfg())

	// log 配置
	log.InitLogger(os.Stderr)
	log.SetLevel(log.DebugLevel)
	log.SetTraceIdFunc(tracing.TraceID)
}

package kernel

import (
	"github.com/JoeZhao1/go-start/framework/gin"
	"net/http"
)

// 引擎服务
type StartKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewStartKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &StartKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *StartKernelService) HttpEngine() http.Handler {
	return s.engine
}

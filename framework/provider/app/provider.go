package app

import (
	"github.com/JoeZhao1/go-start/framework"
	"github.com/JoeZhao1/go-start/framework/contract"
)

// StartAppProvider 提供App的具体实现方法
type StartAppProvider struct {
	BaseFolder string
}

// Register 注册StartApp方法
func (s *StartAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewStartApp
}

// Boot 启动调用
func (s *StartAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (s *StartAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (s *StartAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, s.BaseFolder}
}

// Name 获取字符串凭证
func (s *StartAppProvider) Name() string {
	return contract.AppKey
}

package app

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/JoeZhao1/go-start/framework"
	"github.com/JoeZhao1/go-start/framework/util"
)

// HadeApp 代表hade框架的App实现
type StartApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
}

// Version 实现版本
func (s StartApp) Version() string {
	return "0.0.3"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (s StartApp) BaseFolder() string {
	if s.baseFolder != "" {
		return s.baseFolder
	}

	// 如果没有设置，则使用参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (s StartApp) ConfigFolder() string {
	return filepath.Join(s.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (s StartApp) LogFolder() string {
	return filepath.Join(s.StorageFolder(), "log")
}

func (s StartApp) HttpFolder() string {
	return filepath.Join(s.BaseFolder(), "http")
}

func (s StartApp) ConsoleFolder() string {
	return filepath.Join(s.BaseFolder(), "console")
}

func (s StartApp) StorageFolder() string {
	return filepath.Join(s.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (s StartApp) ProviderFolder() string {
	return filepath.Join(s.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (s StartApp) MiddlewareFolder() string {
	return filepath.Join(s.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (s StartApp) CommandFolder() string {
	return filepath.Join(s.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (s StartApp) RuntimeFolder() string {
	return filepath.Join(s.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (s StartApp) TestFolder() string {
	return filepath.Join(s.BaseFolder(), "test")
}

// NewHadeApp 初始化HadeApp
func NewStartApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &StartApp{baseFolder: baseFolder, container: container}, nil
}

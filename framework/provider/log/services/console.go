package services

import (
	"github.com/JoeZhao1/go-start/framework"
	"github.com/JoeZhao1/go-start/framework/contract"
	"os"
)

// StartConsoleLog 代表控制台输出
type StartConsoleLog struct {
	StartLog
}

// NewStartConsoleLog 实例化StartConsoleLog
func NewStartConsoleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &StartConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
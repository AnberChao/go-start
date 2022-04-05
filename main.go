// Copyright 2021 JoeZhao1.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/JoeZhao1/go-start/app/console"
	"github.com/JoeZhao1/go-start/app/http"
	"github.com/JoeZhao1/go-start/framework"
	"github.com/JoeZhao1/go-start/framework/provider/app"
	"github.com/JoeZhao1/go-start/framework/provider/config"
	"github.com/JoeZhao1/go-start/framework/provider/distributed"
	"github.com/JoeZhao1/go-start/framework/provider/env"
	"github.com/JoeZhao1/go-start/framework/provider/id"
	"github.com/JoeZhao1/go-start/framework/provider/kernel"
	"github.com/JoeZhao1/go-start/framework/provider/log"
	"github.com/JoeZhao1/go-start/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewStartContainer()
	// 绑定App服务提供者
	container.Bind(&app.StartAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.StartEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.StartConfigProvider{})
	container.Bind(&id.StartIDProvider{})
	container.Bind(&trace.StartTraceProvider{})
	container.Bind(&log.StartLogServiceProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.StartKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}

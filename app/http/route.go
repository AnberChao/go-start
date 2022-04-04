package http

import (
	"github.com/JoeZhao1/go-start/app/http/module/demo"
	"github.com/JoeZhao1/go-start/framework/gin"
	"github.com/JoeZhao1/go-start/framework/middleware"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	r.Use(middleware.Trace())
	demo.Register(r)
}

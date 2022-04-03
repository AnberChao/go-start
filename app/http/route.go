package http

import (
	"github.com/JoeZhao1/go-start/app/http/module/demo"
	"github.com/JoeZhao1/go-start/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}

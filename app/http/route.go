package http

import (
	"github.com/fusjoke/fweb/app/http/module/demo"
	"github.com/fusjoke/fweb/framework/gin"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
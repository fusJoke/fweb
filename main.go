package main

import (
	"github.com/fusjoke/fweb/app/console"
	"github.com/fusjoke/fweb/app/http"
	"github.com/fusjoke/fweb/framework"
	"github.com/fusjoke/fweb/framework/provider/app"
	"github.com/fusjoke/fweb/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewHadeContainer()
	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{})
	// 后续初始化需要绑定的服务提供者...

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
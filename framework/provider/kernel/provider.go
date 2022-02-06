package kernel

import (
	"github.com/fusjoke/fweb/framework"
	"github.com/fusjoke/fweb/framework/contract"
	"github.com/fusjoke/fweb/framework/gin"
)

type HadeKernelProvider struct {
	HttpEngine *gin.Engine
}

func (h *HadeKernelProvider) Name() string {
	return contract.KernelKey
}

func (h *HadeKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeKernelService
}

// Boot 启动的时候判断是否由外界注入了Engine，如果注入的化，用注入的，如果没有，重新实例化
func (provider *HadeKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

// Params 参数就是一个HttpEngine
func (provider *HadeKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}

func (provider *HadeKernelProvider) IsDefer () bool {
	return true
}

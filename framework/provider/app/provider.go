package app

import (
	"fmt"
	"github.com/fusjoke/fweb/framework"
	"github.com/fusjoke/fweb/framework/contract"
)

type HadeAppProvider struct {
	BaseFolder string
}

func (h *HadeAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}


func (h *HadeAppProvider) Name () string {
	return contract.AppKey
}

func (h *HadeAppProvider) Register (c framework.Container) framework.NewInstance {
	return NewHadeApp
}

func (h *HadeAppProvider) IsDefer () bool {
	return false
}


func (h *HadeAppProvider) Boot (c framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}


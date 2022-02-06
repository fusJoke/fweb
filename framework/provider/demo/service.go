package demo

import (
	"fmt"
	"github.com/fusjoke/fweb/framework"
)

type DemoService struct {
	Service
	c framework.Container
}

func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	// 这里需要将参数展开
	c := params[0].(framework.Container)

	fmt.Println("new demo service")
	// 返回实例
	return &DemoService{c: c}, nil
}
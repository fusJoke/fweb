package main

import (
	"github.com/fusjoke/fweb/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context)  {
	foo, _ := c.DefaultQueryString("foo", "def")
	time.Sleep(5*time.Second)
	c.ISetOkStatus().IJson("ok, UserLoginController" + foo)
}


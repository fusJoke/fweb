package middleware

import (
	"github.com/fusjoke/fweb/framework/gin"
	"context"
	"fmt"
	"log"
	"time"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context)  {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancal := context.WithTimeout(c.BaseContext(), d)
		defer cancal()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			c.Next()
			finish <- struct{}{}
		}()
		select {
		case p:= <-panicChan:
			c.ISetStatus(500).IJson("time out")
			log.Println(p)
		case <- finish:
			fmt.Println("finish")
		case <- durationCtx.Done():
			c.ISetStatus(500).IJson("time out")
		}
	}
}
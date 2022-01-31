package middleware

import (
	"github.com/fusjoke/fweb/framework"
	"context"
	"fmt"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
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
			c.SetStatus(500).Json("panic")
			log.Println(p)
		case <- finish:
			fmt.Println("finish")
		case <- durationCtx.Done():
			c.SetHasTimeout()
			c.SetStatus(500).Json("time out")
		}
		return nil
	}
}
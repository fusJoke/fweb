package main

import (

	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fusjoke/fweb/framework/gin"
	"github.com/fusjoke/fweb/framework/middleware"
)

func main() {
	fmt.Println("start framework work")
	core := gin.New()

	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit

	timeoutCtx, cancal := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancal()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("server shutdown:", err)
	}
}
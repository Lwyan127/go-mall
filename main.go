package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-study-lab/go-mall/api/router"
	"github.com/go-study-lab/go-mall/common/enum"
	"github.com/go-study-lab/go-mall/common/logger"
	"github.com/go-study-lab/go-mall/config"
	"github.com/go-study-lab/go-mall/event"
)

func main() {
	if config.App.Env == enum.ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	router.RegisterRoutes(g)
	server := http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	log := logger.New(context.Background())

	// 创建带取消的上下文
	shutdownWaitCtx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	// 创建系统信号接收器
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done
		wg.Add(2)

		// 关闭事件分发器
		go func() {
			defer wg.Done()
			err := event.ShutdownEventDispatcher(10 * time.Second)
			if err != nil {
				log.Error("ShutdownEventDispatcherError", "err", err)
			}
		}()

		// 关闭 HTTP 服务
		go func() {
			defer wg.Done()
			// 设置一个独立的超时上下文
			shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer shutdownCancel()

			if err := server.Shutdown(shutdownCtx); err != nil {
				log.Error("ShutdownServerError", "err", err)
			}
		}()

		// 等待所有关闭操作完成或超时
		waitCh := make(chan struct{})
		go func() {
			wg.Wait()
			close(waitCh)
		}()

		// 设置一个总体超时时间
		select {
		case <-waitCh:
			log.Info("All shutdown operations completed successfully")
		case <-time.After(15 * time.Second):
			log.Error("Shutdown operations timed out")
		}
		// 所有清理工作完成后，再调用 cancel
		cancel() // 取消上下文
	}()

	log.Info("Starting GO MALL HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Info("Server closed under request")
		} else {
			log.Error("Server closed unexpected", "err", err)
			panic(err)
		}
	}
	// 等待所有清理工作完成
	<-shutdownWaitCtx.Done()
}

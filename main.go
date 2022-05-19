package main

import (
	"github.com/SkyAPM/go2sky"
	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-hello-world/pkg/shutdown"
	"go-hello-world/router"
	"log"
	"net/http"
	"syscall"
	"time"
)

var SKYWALKING_ENABLED = false

func main() {
	r := gin.New()

	// 配置skywalking
	if SKYWALKING_ENABLED {
		rp, err := reporter.NewGRPCReporter("skywalking-oap:11800", reporter.WithCheckInterval(time.Second))
		if err != nil {
			log.Printf("create gosky reporter failed. err: %s", err)
		}
		defer rp.Close()
		tracer, _ := go2sky.NewTracer("go-hello-world", go2sky.WithReporter(rp))
		r.Use(v3.Middleware(r, tracer))
	}

	// 注册路由
	router.SetupRouter(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// 启动metrics服务
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9527", nil); err != nil {
			log.Printf("metrics port listen failed. err: %s", err)
		}
	}()

	// 运行服务
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server.ListenAndServe err: %v", err)
		}
	}()

	// 优雅退出
	quit := shutdown.New(10)
	quit.Add(syscall.SIGINT, syscall.SIGTERM)
	quit.Start(server)
}

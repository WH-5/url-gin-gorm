package server

import (
	"errors"
	"github.com/WH-5/url-gin-gorm/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// NewHttpServer 创建http服务器
func NewHttpServer(urlBiz *service.UrlBiz) *gin.Engine {
	r := gin.Default()

	// 注册中间件
	r.Use(gin.Recovery()) // 捕获 panic
	r.Use(gin.Logger())   // 日志中间件

	// 注册自定义校验规则
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("customURL", service.ValidateURL)
		if err != nil {
			log.Fatalf("register custom validator failed: %v", err)
			return nil
		}
	}

	// 依赖注入 Handler
	urlHandler := service.NewUrlHandler(*urlBiz)

	// 注册路由
	RegisterRoutes(r, urlHandler)

	return r
}

// RegisterRoutes 注册路由
func RegisterRoutes(router *gin.Engine, handler *service.UrlHandler) {
	api := router.Group("/api")
	router.GET("/:code", handler.RedirectUrl)
	{
		api.POST("/url", handler.CreateUrl)
	}
}

// RunServer 启动http服务器，包括关闭
func RunServer(port string, urlBiz *service.UrlBiz) {
	router := NewHttpServer(urlBiz)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	// 监听系统信号（优雅关闭）
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %s\n", err)
		}
	}()
	log.Println("Server started on port", port)

	// 等待系统信号，优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// 关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}

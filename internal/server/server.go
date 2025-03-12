package server

import (
	"errors"
	"github.com/WH-5/url-gin-gorm/internal/service"
	"github.com/gin-contrib/cors"
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
func NewHttpServer(urlBiz service.UrlBiz) *gin.Engine {
	r := gin.Default()

	// 注册中间件
	r.Use(gin.Recovery()) // 捕获 panic
	//r.Use(gin.Logger())   // 日志中间件 Default会注册一次这个
	r.Use(HeaderInfoMiddleware(urlBiz))

	r.Use(cors.Default())
	// 注册自定义校验规则
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("customURL", service.ValidateURL)
		if err != nil {
			log.Fatalf("register custom validator failed: %v", err)
			return nil
		}
	}

	// 依赖注入 Handler
	urlHandler := service.NewUrlHandler(urlBiz)

	// 注册路由
	RegisterRoutes(r, urlHandler)

	return r
}

// HeaderInfoMiddleware 获取请求头信息的中间件
func HeaderInfoMiddleware(biz service.UrlBiz) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头信息
		ip := c.ClientIP()                 // 获取客户端 IP 地址
		userAgent := c.Request.UserAgent() // 获取 User-Agent
		url := c.Request.URL.String()
		method := c.Request.Method
		full := c.FullPath()
		// 打印日志或者存储信息
		log.Printf("IP: %s, User-Agent: %s, Url: %s, method: %s", ip, userAgent, url, method)

		err := biz.ApiLOG(ip, userAgent, url, method, full)
		if err != nil {
			return
		}
		// 将信息保存在上下文中，后续可以在其他地方访问
		//c.Set("ip", ip)
		//c.Set("userAgent", userAgent)
		//
		//c.Set("url", url)
		//c.Set("method", method)
		// 调用后续处理
		c.Next()
	}
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
func RunServer(port string, urlBiz service.UrlBiz) {
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

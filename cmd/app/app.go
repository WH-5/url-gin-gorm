package main

import (
	"github.com/WH-5/url-gin-gorm/configs"
	"github.com/WH-5/url-gin-gorm/internal/data/cache"
	"github.com/WH-5/url-gin-gorm/internal/data/database"
	"github.com/WH-5/url-gin-gorm/internal/server"
	"github.com/WH-5/url-gin-gorm/pkg/shortCode"
)

type Application struct {
	//e                  *echo.Echo
	db          *database.DBClient
	redisClient *cache.RedisClient
	//urlService         *service.URLService
	//urlHandler         *api.URLHandler
	config             *configs.Config
	shortCodeGenerator *shortCode.ShortCode
}

func main() {
	//加载配置
	config, err := configs.LoadConfig()
	if err != nil {
		return
	}
	//启动服务
	server.RunServer()
}

package app

import (
	"fmt"
	"github.com/WH-5/url-gin-gorm/configs"
	"github.com/WH-5/url-gin-gorm/internal/biz"
	"github.com/WH-5/url-gin-gorm/internal/data/cache"
	"github.com/WH-5/url-gin-gorm/internal/data/database"
	"github.com/WH-5/url-gin-gorm/internal/server"
	"github.com/WH-5/url-gin-gorm/internal/service"
	"github.com/WH-5/url-gin-gorm/pkg/shortCode"
	"log"
	"time"
)

type Application struct {
	db          *database.DBClient
	cacheClient biz.Cache
	urlBiz      service.UrlBiz
	//urlHandler         *service.UrlHandler
	config             *configs.Config
	shortCodeGenerator biz.ShortCodeGen
}

func (a *Application) Init(filepath string) error {
	config, err := configs.LoadConfig(filepath)
	if err != nil {
		return fmt.Errorf("load config error: %v", err)
	}
	a.config = config
	a.db, err = database.NewDB(config.DB)
	a.cacheClient, err = cache.NewRedisClient(config.RD)
	a.shortCodeGenerator = shortCode.NewShortCode(a.config.SCC.Length)
	baseUrl := a.config.AC.BaseHost + a.config.AC.BasePort
	log.Println("baseUrl:", baseUrl)
	a.urlBiz = biz.NewUrl(a.shortCodeGenerator, a.cacheClient, a.config.AC.DefaultDuration, baseUrl, a.db)
	//a.urlHandler = service.NewUrlHandler(a.urlBiz)
	//g:=server.NewHttpServer(a.urlBiz)
	return nil
}
func (a *Application) CleanExpired() {
	t := time.NewTicker(a.config.AC.CleanUpInterval)
	defer t.Stop()
	for range t.C {
		err := a.urlBiz.DeleteUrlByExpiredTime()
		if err != nil {
			log.Println(err)
		}
	}
}
func Run(filePath string) error {

	a := Application{}
	err := a.Init(filePath)
	if err != nil {
		return err
	}
	server.RunServer("8080", a.urlBiz)
	go a.CleanExpired()
	return nil
}

//func main() {
//	//加载配置
//	err := run("")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//启动服务
//	//server.RunServer()
//}

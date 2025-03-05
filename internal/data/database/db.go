package database

import (
	"fmt"
	"github.com/WH-5/url-gin-gorm/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	db *gorm.DB
}

func NewDB(config configs.DbConfig) (*DBClient, error) {
	var dsn string
	if config.Driver == "postgres" {
		dsn = config.PgDsn()
	} else {
		dsn = config.PgDsn()
		//现在只支持pg，防止配置文件写错导致错误
	}

	//连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	//自动迁移
	err = db.AutoMigrate(&UrlShortcode{})
	err = db.AutoMigrate(&IpAccess{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}

	return &DBClient{db: db}, nil
}

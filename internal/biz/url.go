package biz

import (
	"errors"
	"fmt"
	"github.com/WH-5/url-gin-gorm/internal/data/database"
	"github.com/WH-5/url-gin-gorm/internal/service"
	"strconv"
	"time"
)

// Url url的业务逻辑
type Url struct {
	codeGen  ShortCodeGen
	cache    Cache
	duration time.Duration
	baseurl  string
	dbClient *database.DBClient
}

func NewUrl(codeGen ShortCodeGen, cache Cache, duration time.Duration, baseurl string, dbClient *database.DBClient) *Url {
	return &Url{codeGen: codeGen, cache: cache, duration: duration, baseurl: baseurl, dbClient: dbClient}
}

type Cache interface {
	GetURL(shortcode string) (url string, err error)
	SetURL(shortcode, url string) error
}
type ShortCodeGen interface {
	GenerateShortCode() string
}

// CreateUrl 传入code和过期时间然后存到数据库和cache
func (u *Url) CreateUrl(request service.CreateUrlRequest) (string, error) {
	//数据库结构
	//ShortCode   string        `gorm:"size:255;not null;uniqueIndex"` //加个唯一索引加快查询速度
	//OriginalURL string        `gorm:"size:255;not null"`
	//ExpireTime  time.Time     `gorm:"not null"`
	//IsExpired   bool          `gorm:"default:false;not null"`
	//Duration    time.Duration `gorm:"type:int;not null"`

	//cache结构 <code，url>
	code := request.CustomCode
	count := 0
	var available bool
	var err error
	if code == "" {
		for ; count < 5; count++ {
			code = u.codeGen.GenerateShortCode()
			available, err = u.dbClient.IsShortCodeAvailable(code)
			if err != nil {
				return "", err
			}
			if available {
				break
			}
		}
		if !available {
			return "", errors.New("程序可能被入侵了，出现异常")
		}
	} else {
		available, err = u.dbClient.IsShortCodeAvailable(code)
		if !available {
			return "", errors.New(request.CustomCode + " already used")
		}
	}
	d, err := time.ParseDuration(strconv.Itoa(request.Duration) + "h")
	err = u.dbClient.CreateShortcode(code, request.OriginalUrl, d)
	if err != nil {
		return "", err
	}

	err = u.cache.SetURL(code, request.OriginalUrl)
	if err != nil {
		return "", err
	}
	return u.baseurl + code, nil

}

func (u *Url) GetUrl(s string) (string, error) {
	url, err := u.cache.GetURL(s)
	if err != nil {
		return "", err
	}
	if url != "" {
		return url, nil
	}
	code, err := u.dbClient.GetURLByShortCode(s)
	if err != nil {
		return "", err
	}
	return code, nil
}
func (u *Url) createShortCode() (string, error) {
	var code string
	for i := 0; i < 6; i++ {
		code = u.codeGen.GenerateShortCode()
		avail, err := u.dbClient.IsShortCodeAvailable(code)
		if err != nil {
			return "", fmt.Errorf("IsShortCodeAvailable() error: %v", err)
		}
		if avail {
			return code, nil
		}
	}
	return "", errors.New("can not generate short code")
}
func (u *Url) DeleteUrlByExpiredTime() error {
	err := u.dbClient.DeleteURLExpired()
	if err != nil {
		return err
	}
	return nil
}

var _ service.UrlBiz = (*Url)(nil)

package biz

import (
	"errors"
	"fmt"
	"github.com/WH-5/url-gin-gorm/internal/data/database"
	"time"
)

// Url url的业务逻辑
type Url struct {
	sc       ShortCodeGen
	cc       Cache
	duration time.Duration
	baseurl  string
	dbClient database.DBClient
}
type Cache interface {
	GetURL(shortcode string) (url string, err error)
	SetURL(shortcode, url string) error
}
type ShortCodeGen interface {
	GenerateShortCode() string
}

func (u *Url) CreateUrl(s string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *Url) GetUrl(s string) (string, error) {
	//TODO implement me
	panic("implement me")
}
func (u *Url) createShortCode() (string, error) {
	var code string
	for i := 0; i < 6; i++ {
		code = u.sc.GenerateShortCode()
		avail, err := u.dbClient.IsShortCodeAvailable(code)
		if err != nil {
			return "", fmt.Errorf("IsShortCodeAvailable() error: %v", err)
		}
		if avail {
			return code, nil
		}
	}
	return "", errors.New("can generate short code")
}
func (u *Url) DeleteUrlByExpiredTime() error {
	err := u.dbClient.DeleteURLExpired()
	if err != nil {
		return err
	}
	return nil
}

package database

import (
	"gorm.io/gorm"
	"time"
)

//数据库操作
//方法应该用接口类型的，设计的时候没考虑到

type UrlShortcode struct {
	gorm.Model
	ShortCode   string        `gorm:"size:255;not null;uniqueIndex"` //加个唯一索引加快查询速度
	OriginalURL string        `gorm:"size:255;not null"`
	ExpireTime  time.Time     `gorm:"not null"`
	IsExpired   bool          `gorm:"default:false;not null"`
	Duration    time.Duration `gorm:"type:int;not null"`
}

// CreateShortcode 1.创建一条记录
func (db *DBClient) CreateShortcode(shortcode string, original string, dua time.Duration) error {
	sc := UrlShortcode{
		ShortCode:   shortcode,
		OriginalURL: original,
		IsExpired:   false,
		Duration:    dua,
		ExpireTime:  time.Now().Add(dua),
	}
	result := db.db.Create(&sc)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// IsShortCodeAvailable 2.查询一个shortcode是否可以使用，如果可以返回true
func (db *DBClient) IsShortCodeAvailable(shortcode string) (bool, error) {
	//var sc UrlShortcode
	//var count int64
	var exists bool
	//result := db.db.First(&sc, "short_code=?", shortcode).Count(&count)
	err := db.db.Model(&UrlShortcode{}).
		Select("1").
		Where("short_code = ?", shortcode).
		Limit(1).
		Scan(&exists).Error
	//这样快一点
	if err == nil {
		return !exists, nil
	}
	return false, err
}

// GetURLByShortCode 3.通过shortcode查原始url
func (db *DBClient) GetURLByShortCode(shortcode string) (string, error) {
	var url UrlShortcode
	err := db.db.Where("short_code = ?", shortcode).Take(&url).Error
	//有索引，这样是最快的
	if err != nil {
		return "", err
	}
	return url.OriginalURL, nil
}

// DeleteURLExpired 4.删除过期的shortcode
func (db *DBClient) DeleteURLExpired() error {
	return db.db.Where("expire_time < ？", time.Now()).Delete(&UrlShortcode{}).Error
}

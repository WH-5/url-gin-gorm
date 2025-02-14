package database

import (
	"gorm.io/gorm"
	"time"
)

//数据库操作

type Shortcode struct {
	gorm.Model
	ShortCode   string        `gorm:"size:255;not null"`
	OriginalURL string        `gorm:"size:255;not null"`
	ExpireTime  time.Time     `gorm:"not null"`
	IsExpired   bool          `gorm:"default:false;not null"`
	Duration    time.Duration `gorm:"type:int;not null"`
}

// CreateShortcode 1.创建一条记录
func (db *DB) CreateShortcode() {

}

// IsShortCodeAvailable 2.查询数据库里是否存在一个shortcode
func (db *DB) IsShortCodeAvailable() {

}

// GetURLByShortCode 3.通过url查询出一条shortcode
func (db *DB) GetURLByShortCode() {

}

// DeleteURLExpired 4.删除过期的shortcode
func (db *DB) DeleteURLExpired() {

}

package database

import "gorm.io/gorm"

// IpAccess 通过ip记录访问次数
type IpAccess struct {
	gorm.Model
	IPAdr     string `gorm:"size:255;uniqueIndex;not null"`
	UserAgent string `gorm:"size:255"`
	//AccessNumber int    `gorm:"default:1;not null"`
	ApiAdr string `gorm:"size:255"`
	Url    string `gorm:"size:255"`
}

// AddAccess 加入一条访问记录
func (db *DBClient) AddAccess(ip, userAgent, api, reqUrl string) error {
	a := &IpAccess{
		IPAdr:     ip,
		UserAgent: userAgent,
		ApiAdr:    api,
		Url:       reqUrl,
	}
	result := db.db.Create(a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//func (db *DBClient) IpToNumber(ip string) (int, error) {
//	var access IpAccess
//	// 使用 GORM 的查询方法，减少 SQL 注入等安全问题
//	err := db.db.Where("ip_adr = ?", ip).First(&access).Error
//	if err != nil {
//		return 0, err
//	}
//	return access.AccessNumber, nil
//}

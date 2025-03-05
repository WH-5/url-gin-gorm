package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type UrlBiz interface {
	CreateUrl(CreateUrlRequest) (string, error)
	GetUrl(string) (string, error)
	DeleteUrlByExpiredTime() error
	ApiLOG(ip, userAgent, api, reqUrl, full string) error
}
type UrlHandler struct {
	UrlBiz UrlBiz
}

func NewUrlHandler(urlBiz UrlBiz) *UrlHandler {
	return &UrlHandler{
		UrlBiz: urlBiz,
	}
}

type CreateUrlRequest struct {
	OriginalUrl string `json:"original_url" binding:"required,customURL"`
	CustomCode  string `json:"custom_code" binding:"omitempty,min=1,max=10"`
	Duration    int    `json:"duration" binding:"omitempty,gt=0,lt=100"`
}

// 自定义正则表达式，匹配常见网站域名（支持 `http://`、`https://` 也可以不带协议）
var urlRegex = `^((https?://)?[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}(/\S*)?)$`

// ValidateURL 自定义 URL 校验函数
func ValidateURL(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(urlRegex)
	return re.MatchString(fl.Field().String())
}

// CreateUrl POST original_url,custom_code,duration -> short_url,expired_time
func (u *UrlHandler) CreateUrl(c *gin.Context) {
	var req CreateUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}
	if req.OriginalUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "original_url is required",
		})
		return
	}
	//if req.CustomCode=="" {}
	//空值放在在业务逻辑层处理
	var d string
	//没传或者不合法或传了0
	if req.Duration == 0 {
		d = "4h"
		req.Duration = 4
	} else {
		//duration传入的数字单位为小时
		d = strconv.Itoa(req.Duration) + "h"
	}

	//存入
	createUrl, err := u.UrlBiz.CreateUrl(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	//返回响应
	duration, _ := time.ParseDuration(d)
	location, _ := time.LoadLocation("Asia/Shanghai")
	c.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"short_url":    createUrl,
		"expired_time": time.Now().Add(duration).In(location).Format("2006-01-02 15:04:05"),
	})

}

// RedirectUrl GET DirectUrl /:code -> redirect
func (u *UrlHandler) RedirectUrl(c *gin.Context) {
	shortcode := c.Param("code")
	if shortcode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "code is required",
		})
		return
	}
	url, err := u.UrlBiz.GetUrl(shortcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": err.Error()})
		return
	}
	//c.JSON(http.StatusFound, gin.H{
	//	"code": http.StatusFound,
	//	"url":  url,
	//})

	url = AddHTTPPrefix(url)
	c.Redirect(http.StatusMovedPermanently, url)

	//c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf(`
	//    <html>
	//    <head>
	//        <meta http-equiv="refresh" content="0;url=%s">
	//    </head>
	//    <body>
	//        <p>If you are not redirected, <a href="%s">click here</a>.</p>
	//    </body>
	//    </html>
	//`, url, url)))
	return
}

// AddHTTPPrefix 检查字符串是否带协议，如果没有则添加 "http://"
func AddHTTPPrefix(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "http://" + url
}

package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
)

type UrlBiz interface {
	CreateUrl(string) (string, error)
	GetUrl(string) (string, error)
}
type UrlHandler struct {
	UrlBiz UrlBiz
}

func NewUrlHandler(urlBiz UrlBiz) *UrlHandler {
	return &UrlHandler{
		UrlBiz: urlBiz,
	}
}

type createUrlRequest struct {
	OriginalUrl string `json:"original_url" binding:"required,customURL"`
	CustomCode  string `json:"custom_code" binding:"omitempty min=1,max=10"`
	Duration    string `json:"duration"`
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
	panic("implement me")

}

// RedirectUrl GET DirectUrl /:code -> redirect
func (u *UrlHandler) RedirectUrl(c *gin.Context) {
	shortcode := c.Param("code")
	url, err := u.UrlBiz.GetUrl(shortcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, url)
}

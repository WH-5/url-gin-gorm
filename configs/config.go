package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	DbConfig     `mapstructure:"database"`
	RdConfig     `mapstructure:"redis"`
	ServerConfig `mapstructure:"server"`
	AppConfig    `mapstructure:"app"`
	ScConfig     `mapstructure:"shortcode"`
}
type DbConfig struct {
	Driver     string `mapstructure:"driver"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	DBName     string `mapstructure:"dbname"`
	SSLMode    string `mapstructure:"ssl_mode"`
	MaxIdleCon int    `mapstructure:"max_idle_con"`
	MaxOpenCon int    `mapstructure:"max_open_con"`
}
type RdConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
type ServerConfig struct {
	Address      string        `mapstructure:"address"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
}
type AppConfig struct {
	BaseHost        string        `mapstructure:"base_host"`
	BasePort        string        `mapstructure:"base_port"`
	DefaultDuration time.Duration `mapstructure:"default_duration"`
	CleanUpInterval time.Duration `mapstructure:"cleanup_interval"`
}
type ScConfig struct {
	Length int `mapstructure:"length"`
}

// 连接数据库的字符串 pgsql格式
func (d DbConfig) PgDsn() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s", d.Driver, d.User, d.Password, d.Host, d.Port, d.DBName, d.SSLMode)
}

// 加载配置
func LoadConfig(filePath string) (*Config, error) {

	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file failed, err:%v", err)
	}
	// 遍历所有键，展开环境变量并更新 viper
	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		expandedValue := os.ExpandEnv(value) // 替换 ${VAR} 格式的环境变量
		viper.Set(key, expandedValue)
	}
	// 将配置映射到结构体
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config to struct failed, err:%v", err)
	}
	return &cfg, nil
}

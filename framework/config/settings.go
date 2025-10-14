package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Setting 系统配置
//
// 从根目录的`application.yml`读取
type Setting struct {
	// 系统配置
	System struct {
		// 名称
		Name string `yaml:"name"`
		// 版本
		Version string `yaml:"version"`
		// 域名
		Host string `yaml:"host"`
		// 文件上传路径
		UploadPath string `yaml:"uploadPath"`
	} `yaml:"system"`

	// 服务器配置
	Server struct {
		// 端口
		Port int `yaml:"port"`
		// 模式，可选值：debug、test、release
		Mode string `yaml:"mode"`
	} `yaml:"server"`

	// 数据库配置
	DB struct {
		Host string `yaml:"host"`
		// 端口，默认为3306
		Port int `yaml:"port"`
		// 数据库名称
		Database string `yaml:"database"`
		// 用户名
		Username string `yaml:"username"`
		// 密码
		Password string `yaml:"password"`
		// 编码
		Charset string `yaml:"charset"`
		// 连接池最大连接数
		MaxIdleConn int `yaml:"maxIdleConn"`
		// 连接池最大打开连接数
		MaxOpenConn int `yaml:"maxOpenConn"`
	} `yaml:"db"`

	// Redis配置
	Cache struct {
		Host string `yaml:"host"`
		// 端口，默认为6379
		Port int `yaml:"port"`
		// 数据库索引
		Database int `yaml:"database"`
		// 密码
		Password string `yaml:"password"`
	} `yaml:"cache"`

	// 用户配置
	Auth struct {
		// Token配置
		Token struct {
			// 令牌自定义标识
			Header string `yaml:"header"`
			// 令牌密钥
			Secret string `yaml:"secret"`
			// 令牌有效期（默认30分钟）
			ExpireIn int `yaml:"expireIn"`
		} `yaml:"token"`
		// 密码配置
		Password struct {
			// 密码最大错误次数
			MaxRetryCount int `yaml:"maxRetryCount"`
			// 密码锁定时间（默认10分钟）
			LockTime int `yaml:"lockTime"`
		} `yaml:"password"`
	} `yaml:"auth"`
}

var conf *Setting

func init() {
	file, err := os.ReadFile("application.yml")
	if err != nil {
		panic(err)
	}

	conf = &Setting{}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		panic(err)
	}
}

func GetSetting() *Setting {
	return conf
}

package redis_key

import (
	"github.com/hugo8680/goat/framework/config"
)

var prefix = config.GetSetting().System.Name

var (
	CaptchaCodeKey        = prefix + "captcha:code:"          // 验证码
	LoginPasswordErrorKey = prefix + ":login:password:error:" // 登录账户密码错误次数
	UserTokenKey          = prefix + ":user:token:"           // 登录用户
	RepeatSubmitKey       = prefix + ":repeat:submit:"        // 防重提交
	SysConfigKey          = prefix + ":system:config"         // 配置表数据
	SysDictKey            = prefix + ":system:dict:data"      // 字典表数据
)

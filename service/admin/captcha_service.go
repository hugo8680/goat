package admin

import (
	"github.com/mojocn/base64Captcha"
)

// CaptchaService 图形验证码服务
type CaptchaService struct {
	captcha *base64Captcha.Captcha
}

// NewCaptchaService 构造函数
func NewCaptchaService() *CaptchaService {
	driver := base64Captcha.NewDriverDigit(40, 120, 4, 0, 8)
	return &CaptchaService{
		captcha: base64Captcha.NewCaptcha(driver, &CaptchaStoreService{}),
	}
}

// Generate 生成验证码
// uuid, base64, answer
func (c *CaptchaService) Generate() (string, string) {
	id, b64s, _, err := c.captcha.Generate()
	if err != nil {
		return "", ""
	}
	return id, b64s
}

// Verify 验证验证码
func (c *CaptchaService) Verify(id, answer string) bool {
	return c.captcha.Verify(id, answer, true)
}

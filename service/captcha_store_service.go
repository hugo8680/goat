package service

import (
	"context"
	"forum-service/common/constant/redis_key"
	"forum-service/framework/connector"
	"time"
)

// CaptchaStoreService 图形验证码存储服务
type CaptchaStoreService struct {
}

func (s *CaptchaStoreService) Set(id string, value string) error {
	return connector.GetCache().Set(context.Background(), redis_key.CaptchaCodeKey+id, value, time.Minute*5).Err()
}

func (s *CaptchaStoreService) Get(id string, clear bool) string {
	cache := connector.GetCache()
	captcha, err := cache.Get(context.Background(), redis_key.CaptchaCodeKey+id).Result()
	if err != nil {
		return ""
	}
	if clear {
		if err = cache.Del(context.Background(), redis_key.CaptchaCodeKey+id).Err(); err != nil {
			return ""
		}
	}
	return captcha
}

func (s *CaptchaStoreService) Verify(id, answer string, clear bool) bool {
	return s.Get(id, clear) == answer
}

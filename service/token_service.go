package service

import (
	"context"
	"errors"
	"forum-service/common/constant/redis_key"
	"forum-service/common/serializer/datetime"
	"forum-service/common/uuid"
	"forum-service/framework/config"
	"forum-service/framework/connector"
	"forum-service/model/dto"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// TokenService 授权声明
type TokenService struct {
	jwt.RegisteredClaims
	Key     string `json:"key"`
	Setting *config.Setting
}

// NewTokenService 获取授权声明
func NewTokenService() *TokenService {
	key, _ := uuid.CreateId()
	conf := config.GetSetting()
	return &TokenService{
		Key: key,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
			Issuer:    conf.System.Name,               // 签发人
		},
		Setting: conf,
	}
}

// Create 生成token
func (s *TokenService) Create(user *dto.UserTokenResponse) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, s).SignedString([]byte(s.Setting.Auth.Token.Secret))
	if err != nil {
		return "", err
	}
	expireTime := time.Minute * time.Duration(s.Setting.Auth.Token.ExpireIn)
	user.ExpireTime = datetime.Datetime{Time: time.Now().Add(expireTime)}
	err = connector.GetCache().Set(context.Background(), redis_key.UserTokenKey+s.Key, user, expireTime).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

// Refresh 刷新token
func (s *TokenService) Refresh(ctx *gin.Context, user *dto.UserTokenResponse) {
	tokenKey, err := s.getUserTokenKey(ctx)
	if err != nil {
		return
	}
	expireTime := time.Minute * time.Duration(s.Setting.Auth.Token.ExpireIn)
	user.ExpireTime = datetime.Datetime{Time: time.Now().Add(expireTime)}
	connector.GetCache().Set(ctx.Request.Context(), tokenKey, user, time.Minute*time.Duration(s.Setting.Auth.Token.ExpireIn))
}

// Parse 将token解析为用户信息
func (s *TokenService) Parse(ctx *gin.Context) (*dto.UserTokenResponse, error) {
	tokenKey, err := s.getUserTokenKey(ctx)
	if err != nil {
		return nil, err
	}
	var user *dto.UserTokenResponse
	if err = connector.GetCache().Get(ctx.Request.Context(), tokenKey).Scan(user); err != nil {
		return nil, err
	}
	return user, nil
}

// Delete 删除token
func (s *TokenService) Delete(ctx *gin.Context) error {
	tokenKey, err := s.getUserTokenKey(ctx)
	if err != nil {
		return err
	}
	return connector.GetCache().Del(ctx.Request.Context(), tokenKey).Err()
}

// 获取授权用户的token key
func (s *TokenService) getUserTokenKey(ctx *gin.Context) (string, error) {
	authorization := ctx.GetHeader(s.Setting.Auth.Token.Header)
	if authorization == "" {
		return "", errors.New("请先登录")
	}
	tokenSplit := strings.Split(authorization, " ")
	if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
		return "", errors.New("authorization format error")
	}
	token, err := jwt.ParseWithClaims(tokenSplit[1], s, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Setting.Auth.Token.Secret), nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", errors.New("token格式错误")
			}
			if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", errors.New("token未生效")
			}
			return "", errors.New("token校验失败")
		}
		return "", err
	}
	if claims, ok := token.Claims.(*TokenService); ok && token.Valid {
		return redis_key.UserTokenKey + claims.Key, nil
	}
	return "", errors.New("token校验失败")
}

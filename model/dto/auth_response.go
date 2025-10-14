package dto

// CaptchaResponse 图形验证码响应
type CaptchaResponse struct {
	Uuid           string `json:"uuid"`
	Img            string `json:"img"`
	CaptchaEnabled bool   `json:"captchaEnabled"`
}

type RegisterResponse struct {
}

type AuthInfoResponse struct {
	User        AuthUserInfoResponse `json:"user"`
	Roles       []string             `json:"roles"`
	Permissions []string             `json:"permissions"`
}

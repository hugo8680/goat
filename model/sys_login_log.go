package model

import (
	"github.com/hugo8680/goat/common/serializer/datetime"
)

type SysLoginLog struct {
	InfoId        int `gorm:"primaryKey;autoIncrement"`
	UserName      string
	Ipaddr        string
	LoginLocation string
	Browser       string
	Os            string
	Status        string `gorm:"default:0"`
	Msg           string
	LoginTime     datetime.Datetime
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}

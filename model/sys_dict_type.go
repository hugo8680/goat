package model

import (
	"github.com/hugo8680/goat/common/serializer/datetime"
)

type SysDictType struct {
	DictId     int `gorm:"primaryKey;autoIncrement"`
	DictName   string
	DictType   string
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	Remark     string
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}

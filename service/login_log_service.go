package service

import (
	"forum-service/common/constant/redis_key"
	"forum-service/framework/connector"
	"forum-service/model"
	"forum-service/model/dto"
	"log"

	"github.com/gin-gonic/gin"
)

type LoginLogService struct {
}

// Delete 删除登录日志
func (s *LoginLogService) Delete(infoIds []int) error {
	db := connector.GetDB()
	if len(infoIds) > 0 {
		return db.Model(model.SysLoginLog{}).Where("info_id IN ?", infoIds).Delete(&model.SysLoginLog{}).Error
	}
	// 为解决 WHERE conditions required 错误，添加 Where("info_id > ?", 0) 这个条件
	return db.Model(model.SysLoginLog{}).Where("info_id > ?", 0).Delete(&model.SysLoginLog{}).Error
}

// List 获取登录日志列表
func (s *LoginLogService) List(param dto.LoginLogListRequest, isPaging bool) ([]dto.LoginLogListResponse, int) {
	var count int64
	loginLogs := make([]dto.LoginLogListResponse, 0)
	query := connector.GetDB().Model(model.SysLoginLog{}).Order(param.OrderByColumn + " " + param.OrderRule)
	if param.Ipaddr != "" {
		query = query.Where("ipaddr LIKE ?", "%"+param.Ipaddr+"%")
	}
	if param.UserName != "" {
		query = query.Where("user_name LIKE ?", "%"+param.UserName+"%")
	}
	if param.Status != "" {
		query = query.Where("status = ?", param.Status)
	}
	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("login_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&loginLogs)
	return loginLogs, int(count)
}

// Create 记录登录信息
func (s *LoginLogService) Create(param dto.SaveLoginLogRequest) error {
	go func() {
		err := func() error {
			return connector.GetDB().Model(model.SysLoginLog{}).Create(&model.SysLoginLog{
				UserName:      param.UserName,
				Ipaddr:        param.Ipaddr,
				LoginLocation: param.LoginLocation,
				Browser:       param.Browser,
				Os:            param.Os,
				Status:        param.Status,
				Msg:           param.Msg,
				LoginTime:     param.LoginTime,
			}).Error
		}()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func (s *LoginLogService) UnLock(ctx *gin.Context) error {
	return connector.GetCache().Del(ctx.Request.Context(), redis_key.LoginPasswordErrorKey+ctx.Param("userName")).Err()
}

package service

import (
	"forum-service/framework/connector"
	"forum-service/model"
	"forum-service/model/dto"
	"log"
)

type OperLogService struct {
}

// Delete 删除操作日志
func (s *OperLogService) Delete(operIds []int) error {
	db := connector.GetDB()
	if len(operIds) > 0 {
		return db.Model(model.SysOperLog{}).Where("oper_id IN ?", operIds).Delete(&model.SysOperLog{}).Error
	}
	// 为解决 WHERE conditions required 错误，添加 Where("oper_id > ?", 0) 这个条件
	return db.Model(model.SysOperLog{}).Where("oper_id > ?", 0).Delete(&model.SysOperLog{}).Error
}

// List 操作日志列表
func (s *OperLogService) List(param dto.OperLogListRequest, isPaging bool) ([]dto.OperLogListResponse, int) {
	var count int64
	operLogs := make([]dto.OperLogListResponse, 0)
	query := connector.GetDB().Model(model.SysOperLog{}).Order(param.OrderByColumn + " " + param.OrderRule)
	if param.OperIp != "" {
		query = query.Where("oper_ip LIKE ?", "%"+param.OperIp+"%")
	}
	if param.Title != "" {
		query = query.Where("title LIKE ?", "%"+param.Title+"%")
	}
	if param.OperName != "" {
		query = query.Where("oper_name LIKE ?", "%"+param.OperName+"%")
	}
	if param.BusinessType != "" {
		query = query.Where("business_type = ?", param.BusinessType)
	}
	if param.Status != "" {
		query = query.Where("status = ?", param.Status)
	}
	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("oper_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&operLogs)
	return operLogs, int(count)
}

// Create 记录操作日志
func (s *OperLogService) Create(param dto.SaveOperLogRequest) error {
	go func() {
		err := func() error {
			return connector.GetDB().Model(model.SysOperLog{}).Create(&model.SysOperLog{
				Title:         param.Title,
				BusinessType:  param.BusinessType,
				Method:        param.Method,
				RequestMethod: param.RequestMethod,
				OperName:      param.OperName,
				DeptName:      param.DeptName,
				OperUrl:       param.OperUrl,
				OperIp:        param.OperIp,
				OperLocation:  param.OperLocation,
				OperParam:     param.OperParam,
				JsonResult:    param.JsonResult,
				Status:        param.Status,
				ErrorMsg:      param.ErrorMsg,
				OperTime:      param.OperTime,
				CostTime:      param.CostTime,
			}).Error
		}()
		if err != nil {
			log.Println(err)
		}
	}()
	return nil
}

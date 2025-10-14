package service

import (
	"context"
	"encoding/json"
	"errors"
	"forum-service/common/constant/redis_key"
	"forum-service/framework/connector"
	"forum-service/model"
	"forum-service/model/dto"

	"github.com/gin-gonic/gin"
)

type ConfigService struct {
}

// Create 创建参数
func (s *ConfigService) Create(param dto.SaveConfigRequest) error {
	if config := s.GetCacheByConfigKey(param.ConfigKey); config.ConfigId > 0 {
		return errors.New("新增参数" + param.ConfigName + "失败，参数键名已存在")
	}
	return connector.GetDB().Model(model.SysConfig{}).Create(&model.SysConfig{
		ConfigName:  param.ConfigName,
		ConfigKey:   param.ConfigKey,
		ConfigValue: param.ConfigValue,
		ConfigType:  param.ConfigType,
		CreateBy:    param.CreateBy,
		Remark:      param.Remark,
	}).Error
}

// Update 更新参数
func (s *ConfigService) Update(param dto.SaveConfigRequest) error {
	if config := s.GetCacheByConfigKey(param.ConfigKey); config.ConfigId > 0 && config.ConfigId != param.ConfigId {
		return errors.New("修改参数" + param.ConfigName + "失败，参数键名已存在")
	}
	return connector.GetDB().Model(model.SysConfig{}).Where("config_id = ?", param.ConfigId).Updates(&model.SysConfig{
		ConfigName:  param.ConfigName,
		ConfigKey:   param.ConfigKey,
		ConfigValue: param.ConfigValue,
		ConfigType:  param.ConfigType,
		UpdateBy:    param.UpdateBy,
		Remark:      param.Remark,
	}).Error
}

// Delete 删除参数
func (s *ConfigService) Delete(configIds []int) error {
	return connector.GetDB().Model(model.SysConfig{}).Where("config_id IN ?", configIds).Delete(&model.SysConfig{}).Error
}

// List 获取参数列表
func (s *ConfigService) List(param dto.ConfigListRequest, isPaging bool) ([]dto.ConfigListResponse, int) {
	var count int64
	configs := make([]dto.ConfigListResponse, 0)
	query := connector.GetDB().Model(model.SysConfig{}).Order("config_id")
	if param.ConfigName != "" {
		query = query.Where("config_name LIKE ?", "%"+param.ConfigName+"%")
	}
	if param.ConfigKey != "" {
		query = query.Where("config_key LIKE ?", "%"+param.ConfigKey+"%")
	}
	if param.ConfigType != "" {
		query = query.Where("config_type = ?", param.ConfigType)
	}
	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("create_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&configs)
	return configs, int(count)
}

// Get 获取参数详情
func (s *ConfigService) Get(configId int) dto.ConfigDetailResponse {
	var config dto.ConfigDetailResponse
	connector.GetDB().Model(model.SysConfig{}).Where("config_id = ?", configId).Last(&config)
	return config
}

// GetByConfigKey 根据参数key获取参数值
func (s *ConfigService) GetByConfigKey(configKey string) dto.ConfigDetailResponse {
	var config dto.ConfigDetailResponse
	connector.GetDB().Model(model.SysConfig{}).Where("config_key = ?", configKey).Last(&config)
	return config
}

// GetCacheByConfigKey 根据参数key获取参数配置
func (s *ConfigService) GetCacheByConfigKey(configKey string) dto.ConfigDetailResponse {
	cache := connector.GetCache()
	var config dto.ConfigDetailResponse
	// 缓存不为空不从数据库读取，减少数据库压力
	if configCache, _ := cache.HGet(context.Background(), redis_key.SysConfigKey, configKey).Result(); configCache != "" {
		if err := json.Unmarshal([]byte(configCache), &config); err == nil {
			return config
		}
	}
	// 从数据库读取配置并且记录到缓存
	config = s.GetByConfigKey(configKey)
	if config.ConfigId > 0 {
		configBytes, _ := json.Marshal(&config)
		cache.HSet(context.Background(), redis_key.SysConfigKey, configKey, string(configBytes)).Result()
	}
	return config
}

// RefreshCache 刷新缓存
func (s *ConfigService) RefreshCache(ctx *gin.Context) error {
	return connector.GetCache().Del(ctx.Request.Context(), redis_key.SysConfigKey).Err()
}

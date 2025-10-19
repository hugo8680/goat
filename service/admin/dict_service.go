package admin

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hugo8680/goat/common/constant/redis_key"
	"github.com/hugo8680/goat/framework/connector"
	"github.com/hugo8680/goat/model"
	"github.com/hugo8680/goat/model/dto"

	"github.com/gin-gonic/gin"
)

type DictTypeService struct {
}

// Create 创建字典类型
func (s *DictTypeService) Create(param dto.SaveDictTypeRequest) error {
	if dictType := s.GetByDictType(param.DictType); dictType.DictId > 0 {
		return errors.New("新增字典" + param.DictName + "失败，字典类型已存在")
	}
	return connector.GetDB().Model(model.SysDictType{}).Create(&model.SysDictType{
		DictName: param.DictName,
		DictType: param.DictType,
		Status:   param.Status,
		Remark:   param.Remark,
		CreateBy: param.CreateBy,
	}).Error
}

// Update 更新字典类型
func (s *DictTypeService) Update(param dto.SaveDictTypeRequest) error {
	if dictType := s.GetByDictType(param.DictType); dictType.DictId > 0 && dictType.DictId != param.DictId {
		return errors.New("修改字典" + param.DictName + "失败，字典类型已存在")
	}
	return connector.GetDB().Model(model.SysDictType{}).Where("dict_id = ?", param.DictId).Updates(&model.SysDictType{
		DictName: param.DictName,
		DictType: param.DictType,
		Status:   param.Status,
		Remark:   param.Remark,
		UpdateBy: param.UpdateBy,
	}).Error
}

// Delete 删除字典类型
func (s *DictTypeService) Delete(dictIds []int) error {
	return connector.GetDB().Model(model.SysDictType{}).Where("dict_id IN ?", dictIds).Delete(&model.SysDictType{}).Error
}

// List 字典类型列表
func (s *DictTypeService) List(param dto.DictTypeListRequest, isPaging bool) ([]dto.DictTypeListResponse, int) {
	var count int64
	dictTypes := make([]dto.DictTypeListResponse, 0)
	query := connector.GetDB().Model(model.SysDictType{}).Order("dict_id")
	if param.DictName != "" {
		query = query.Where("dict_name LIKE ?", "%"+param.DictName+"%")
	}
	if param.DictType != "" {
		query = query.Where("dict_type LIKE ?", "%"+param.DictType+"%")
	}
	if param.Status != "" {
		query = query.Where("status = ?", param.Status)
	}
	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("create_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&dictTypes)
	return dictTypes, int(count)
}

// Get 字典类型详情
func (s *DictTypeService) Get(dictId int) dto.DictTypeDetailResponse {
	var dictType dto.DictTypeDetailResponse
	connector.GetDB().Model(model.SysDictType{}).Where("dict_id = ?", dictId).Last(&dictType)
	return dictType
}

// GetByDictType 根据字典类型查询详情
func (s *DictTypeService) GetByDictType(dictType string) dto.DictTypeDetailResponse {
	var dictTypeResult dto.DictTypeDetailResponse
	connector.GetDB().Model(model.SysDictType{}).Where("dict_type = ?", dictType).Last(&dictTypeResult)
	return dictTypeResult
}

// RefreshCache 刷新缓存
func (s *DictTypeService) RefreshCache(ctx *gin.Context) error {
	return connector.GetCache().Del(ctx.Request.Context(), redis_key.SysDictKey).Err()
}

type DictDataService struct {
}

// Create 创建字典数据
func (s *DictDataService) Create(param dto.SaveDictDataRequest) error {
	return connector.GetDB().Model(model.SysDictData{}).Create(&model.SysDictData{
		DictSort:  param.DictSort,
		DictLabel: param.DictLabel,
		DictValue: param.DictValue,
		DictType:  param.DictType,
		CssClass:  param.CssClass,
		ListClass: param.ListClass,
		IsDefault: param.IsDefault,
		Status:    param.Status,
		Remark:    param.Remark,
		CreateBy:  param.CreateBy,
	}).Error
}

// 更新字典数据
func (s *DictDataService) Update(param dto.SaveDictDataRequest) error {
	return connector.GetDB().Model(model.SysDictData{}).Where("dict_code = ?", param.DictCode).Updates(&model.SysDictData{
		DictSort:  param.DictSort,
		DictLabel: param.DictLabel,
		DictValue: param.DictValue,
		DictType:  param.DictType,
		CssClass:  param.CssClass,
		ListClass: param.ListClass,
		IsDefault: param.IsDefault,
		Status:    param.Status,
		Remark:    param.Remark,
		UpdateBy:  param.UpdateBy,
	}).Error
}

// Delete 删除字典数据
func (s *DictDataService) Delete(dictCodes []int) error {
	return connector.GetDB().Model(model.SysDictData{}).Where("dict_code IN ?", dictCodes).Delete(&model.SysDictData{}).Error
}

// List 字典数据列表
func (s *DictDataService) List(param dto.DictDataListRequest, isPaging bool) ([]dto.DictDataListResponse, int) {
	var count int64
	dictDataList := make([]dto.DictDataListResponse, 0)
	query := connector.GetDB().Model(model.SysDictData{}).Order("dict_code")
	if param.DictLabel != "" {
		query = query.Where("dict_label LIKE ?", "%"+param.DictLabel+"%")
	}
	if param.DictType != "" {
		query = query.Where("dict_type LIKE ?", "%"+param.DictType+"%")
	}
	if param.Status != "" {
		query = query.Where("status = ?", param.Status)
	}
	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}
	query.Find(&dictDataList)
	return dictDataList, int(count)
}

// GetByDictCode 根据字典数据编码获取字典数据详情
func (s *DictDataService) GetByDictCode(dictCode int) dto.DictDataDetailResponse {
	var dictData dto.DictDataDetailResponse
	connector.GetDB().Model(model.SysDictData{}).Where("dict_code = ?", dictCode).Last(&dictData)
	return dictData
}

// GetByDictType 根据字典类型查询字典数据
func (s *DictDataService) GetByDictType(dictType string) []dto.DictDataListResponse {
	dictDataList := make([]dto.DictDataListResponse, 0)
	connector.GetDB().Model(model.SysDictData{}).Where("status = 0 AND dict_type = ?", dictType).Find(&dictDataList)
	return dictDataList
}

// GetCacheByDictType 根据字典类型查询字典数据
func (s *DictDataService) GetCacheByDictType(dictType string) []dto.DictDataListResponse {
	cache := connector.GetCache()
	dictDataList := make([]dto.DictDataListResponse, 0)
	// 缓存不为空不从数据库读取，减少数据库压力
	if dictDataListCache, _ := cache.HGet(context.Background(), redis_key.SysDictKey, dictType).Result(); dictDataListCache != "" {
		if err := json.Unmarshal([]byte(dictDataListCache), &dictDataList); err == nil {
			return dictDataList
		}
	}
	// 从数据库读取配置并且记录到缓存
	dictDataList = s.GetByDictType(dictType)
	if len(dictDataList) > 0 {
		dictDadasBytes, _ := json.Marshal(&dictDataList)
		cache.HSet(context.Background(), redis_key.SysDictKey, dictType, string(dictDadasBytes))
	}
	return dictDataList
}

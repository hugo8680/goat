package admin

import (
	"forum-service/api/validator/admin"
	"forum-service/common/constant/auth"
	"forum-service/common/serializer/datetime"
	"forum-service/common/utils"
	"forum-service/framework/response"
	"forum-service/model/dto"
	"forum-service/service"
	"strconv"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
)

type DictDataController struct {
	dictDataService *service.DictDataService
}

func NewDictDataController() *DictDataController {
	return &DictDataController{
		dictDataService: &service.DictDataService{},
	}
}

// List 获取字典数据列表
func (c *DictDataController) List(ctx *gin.Context) {
	var param dto.DictDataListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	dictDatas, total := c.dictDataService.List(param, true)
	response.Success(ctx).SetPageData(dictDatas, total).Json()
}

// Get 获取字典数据详情
func (c *DictDataController) Get(ctx *gin.Context) {
	dictCode, _ := strconv.Atoi(ctx.Param("dictCode"))
	dictData := c.dictDataService.GetByDictCode(dictCode)
	response.Success(ctx).SetData("data", dictData).Json()
}

// Create 新增字典数据
func (c *DictDataController) Create(ctx *gin.Context) {
	var param dto.CreateDictDataRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.CreateDictDataValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.dictDataService.Create(dto.SaveDictDataRequest{
		DictSort:  param.DictSort,
		DictLabel: param.DictLabel,
		DictValue: param.DictValue,
		DictType:  param.DictType,
		CssClass:  param.CssClass,
		ListClass: param.ListClass,
		IsDefault: param.IsDefault,
		Status:    param.Status,
		CreateBy:  user.(dto.UserTokenResponse).UserName,
		Remark:    param.Remark,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Update 更新字典数据
func (c *DictDataController) Update(ctx *gin.Context) {
	var param dto.UpdateDictDataRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.UpdateDictDataValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.dictDataService.Update(dto.SaveDictDataRequest{
		DictCode:  param.DictCode,
		DictSort:  param.DictSort,
		DictLabel: param.DictLabel,
		DictValue: param.DictValue,
		DictType:  param.DictType,
		CssClass:  param.CssClass,
		ListClass: param.ListClass,
		IsDefault: param.IsDefault,
		Status:    param.Status,
		UpdateBy:  user.(dto.UserTokenResponse).UserName,
		Remark:    param.Remark,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Delete 删除字典数据
func (c *DictDataController) Delete(ctx *gin.Context) {
	dictCodes, err := utils.StringToIntSlice(ctx.Param("dictCodes"), ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err = c.dictDataService.Delete(dictCodes); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// DictDataOptions 根据字典类型查询字典数据
func (c *DictDataController) DictDataOptions(ctx *gin.Context) {
	dictType := ctx.Param("dictType")
	dictDatas := c.dictDataService.GetCacheByDictType(dictType)
	for key, dictData := range dictDatas {
		dictDatas[key].Default = dictData.IsDefault == "Y"
	}
	response.Success(ctx).SetData("data", dictDatas).Json()
}

// Export 数据导出
func (c *DictDataController) Export(ctx *gin.Context) {
	var param dto.DictDataListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	list := make([]dto.DictDataExportResponse, 0)
	dictDatas, _ := c.dictDataService.List(param, false)
	for _, dictData := range dictDatas {
		list = append(list, dto.DictDataExportResponse{
			DictCode:  dictData.DictCode,
			DictSort:  dictData.DictSort,
			DictLabel: dictData.DictLabel,
			DictValue: dictData.DictValue,
			DictType:  dictData.DictType,
			IsDefault: dictData.IsDefault,
			Status:    dictData.Status,
		})
	}
	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	excel.DownLoadExcel("data_"+time.Now().Format(datetime.DATETIME_FORMAT2), ctx.Writer, file)
}

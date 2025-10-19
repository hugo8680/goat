package admin

import (
	"github.com/hugo8680/goat/api/validator/admin"
	"github.com/hugo8680/goat/common/constant/auth"
	"github.com/hugo8680/goat/common/serializer/datetime"
	"github.com/hugo8680/goat/common/utils"
	"github.com/hugo8680/goat/framework/response"
	"github.com/hugo8680/goat/model/dto"
	adminService "github.com/hugo8680/goat/service/admin"
	"strconv"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
)

type DictTypeController struct {
	dictTypeService *adminService.DictTypeService
}

func NewDictTypeController() *DictTypeController {
	return &DictTypeController{
		dictTypeService: &adminService.DictTypeService{},
	}
}

// List 字典类型列表
func (c *DictTypeController) List(ctx *gin.Context) {
	var param dto.DictTypeListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	dictTypes, total := c.dictTypeService.List(param, true)
	response.Success(ctx).SetPageData(dictTypes, total).Json()
}

// Get 字典类型详情
func (c *DictTypeController) Get(ctx *gin.Context) {
	dictId, _ := strconv.Atoi(ctx.Param("dictId"))
	dictType := c.dictTypeService.Get(dictId)
	response.Success(ctx).SetData("data", dictType).Json()
}

// Create 新增字典类型
func (c *DictTypeController) Create(ctx *gin.Context) {
	var param dto.CreateDictTypeRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.CreateDictTypeValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.dictTypeService.Create(dto.SaveDictTypeRequest{
		DictName: param.DictName,
		DictType: param.DictType,
		Status:   param.Status,
		CreateBy: user.(*dto.UserTokenResponse).UserName,
		Remark:   param.Remark,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Update 更新字典类型
func (c *DictTypeController) Update(ctx *gin.Context) {
	var param dto.UpdateDictTypeRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.UpdateDictTypeValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.dictTypeService.Update(dto.SaveDictTypeRequest{
		DictId:   param.DictId,
		DictName: param.DictName,
		DictType: param.DictType,
		Status:   param.Status,
		UpdateBy: user.(*dto.UserTokenResponse).UserName,
		Remark:   param.Remark,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Delete 删除字典类型
func (c *DictTypeController) Delete(ctx *gin.Context) {
	dictIds, err := utils.StringToIntSlice(ctx.Param("dictIds"), ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err = c.dictTypeService.Delete(dictIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// DictOptions 获取字典选择框列表
func (c *DictTypeController) DictOptions(ctx *gin.Context) {
	dictTypes, _ := c.dictTypeService.List(dto.DictTypeListRequest{
		Status: "0",
	}, false)
	response.Success(ctx).SetData("data", dictTypes).Json()
}

// Export 数据导出
func (c *DictTypeController) Export(ctx *gin.Context) {
	var param dto.DictTypeListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	list := make([]dto.DictTypeExportResponse, 0)
	dictTypes, _ := c.dictTypeService.List(param, false)
	for _, dictType := range dictTypes {
		list = append(list, dto.DictTypeExportResponse{
			DictId:   dictType.DictId,
			DictName: dictType.DictName,
			DictType: dictType.DictType,
			Status:   dictType.Status,
		})
	}
	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	excel.DownLoadExcel("type_"+time.Now().Format(datetime.DATETIME_FORMAT2), ctx.Writer, file)
}

// RefreshCache 刷新缓存
func (c *DictTypeController) RefreshCache(ctx *gin.Context) {
	err := c.dictTypeService.RefreshCache(ctx)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

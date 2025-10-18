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

type ConfigController struct {
	configService *service.ConfigService
}

func NewConfigController() *ConfigController {
	return &ConfigController{
		configService: &service.ConfigService{},
	}
}

// List 参数列表
func (c *ConfigController) List(ctx *gin.Context) {
	var param dto.ConfigListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	configs, total := c.configService.List(param, true)
	response.Success(ctx).SetPageData(configs, total).Json()
}

// Get 参数详情
func (c *ConfigController) Get(ctx *gin.Context) {
	configId, _ := strconv.Atoi(ctx.Param("configId"))
	config := c.configService.Get(configId)
	response.Success(ctx).SetData("data", config).Json()
}

// Create 新增参数
func (c *ConfigController) Create(ctx *gin.Context) {
	var param dto.CreateConfigRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.CreateConfigValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.configService.Create(dto.SaveConfigRequest{
		ConfigName:  param.ConfigName,
		ConfigKey:   param.ConfigKey,
		ConfigValue: param.ConfigValue,
		ConfigType:  param.ConfigType,
		Remark:      param.Remark,
		CreateBy:    user.(*dto.UserTokenResponse).UserName,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Update 更新参数
func (c *ConfigController) Update(ctx *gin.Context) {
	var param dto.UpdateConfigRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.UpdateConfigValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.configService.Update(dto.SaveConfigRequest{
		ConfigId:    param.ConfigId,
		ConfigName:  param.ConfigName,
		ConfigKey:   param.ConfigKey,
		ConfigValue: param.ConfigValue,
		ConfigType:  param.ConfigType,
		Remark:      param.Remark,
		UpdateBy:    user.(*dto.UserTokenResponse).UserName,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Delete 删除参数
func (c *ConfigController) Delete(ctx *gin.Context) {
	configIds, err := utils.StringToIntSlice(ctx.Param("configIds"), ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err = c.configService.Delete(configIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// ConfigKey 根据配置key获取配置值
func (c *ConfigController) ConfigKey(ctx *gin.Context) {
	configKey := ctx.Param("configKey")
	config := c.configService.GetCacheByConfigKey(configKey)
	response.Success(ctx).SetMsg(config.ConfigValue).Json()
}

// Export 数据导出
func (c *ConfigController) Export(ctx *gin.Context) {
	var param dto.ConfigListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	list := make([]dto.ConfigExportResponse, 0)
	configs, _ := c.configService.List(param, false)
	for _, config := range configs {
		list = append(list, dto.ConfigExportResponse{
			ConfigId:    config.ConfigId,
			ConfigName:  config.ConfigName,
			ConfigKey:   config.ConfigKey,
			ConfigValue: config.ConfigValue,
			ConfigType:  config.ConfigType,
		})
	}
	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	excel.DownLoadExcel("config_"+time.Now().Format(datetime.DATETIME_FORMAT2), ctx.Writer, file)
}

// RefreshCache 刷新缓存
func (c *ConfigController) RefreshCache(ctx *gin.Context) {
	if err := c.configService.RefreshCache(ctx); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

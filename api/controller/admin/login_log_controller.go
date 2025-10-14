package admin

import (
	"forum-service/common/serializer/datetime"
	"forum-service/common/utils"
	"forum-service/framework/response"
	"forum-service/model/dto"
	"forum-service/service"
	"regexp"
	"strings"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
)

type LoginLogController struct {
	loginLogService *service.LoginLogService
}

func NewLoginLogController() *LoginLogController {
	return &LoginLogController{
		loginLogService: &service.LoginLogService{},
	}
}

// List 登录日志列表
func (c *LoginLogController) List(ctx *gin.Context) {
	var param dto.LoginLogListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	// 排序规则默认为倒序（DESC）
	param.OrderRule = "DESC"
	if strings.HasPrefix(param.IsAsc, "asc") {
		param.OrderRule = "ASC"
	}
	// 排序字段小驼峰转蛇形
	if param.OrderByColumn == "" {
		param.OrderByColumn = "loginTime"
	}
	param.OrderByColumn = strings.ToLower(regexp.MustCompile("([A-Z])").ReplaceAllString(param.OrderByColumn, "_${1}"))
	logininfors, total := c.loginLogService.List(param, true)
	response.Success(ctx).SetPageData(logininfors, total).Json()
}

// Delete 删除登录日志
func (c *LoginLogController) Delete(ctx *gin.Context) {
	infoIds, err := utils.StringToIntSlice(ctx.Param("infoIds"), ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err = c.loginLogService.Delete(infoIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Clean 清空登录日志
func (c *LoginLogController) Clean(ctx *gin.Context) {
	if err := c.loginLogService.Delete(nil); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Unlock 账户解锁（删除登录错误次数限制10分钟缓存）
func (c *LoginLogController) Unlock(ctx *gin.Context) {
	err := c.loginLogService.UnLock(ctx)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Export 数据导出
func (c *LoginLogController) Export(ctx *gin.Context) {
	var param dto.LoginLogListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	// 排序规则默认为倒序（DESC）
	param.OrderRule = "DESC"
	if strings.HasPrefix(param.IsAsc, "asc") {
		param.OrderRule = "ASC"
	}
	// 排序字段小驼峰转蛇形
	if param.OrderByColumn == "" {
		param.OrderByColumn = "loginTime"
	}
	param.OrderByColumn = strings.ToLower(regexp.MustCompile("([A-Z])").ReplaceAllString(param.OrderByColumn, "_${1}"))
	list := make([]dto.LoginLogExportResponse, 0)
	logininfors, _ := c.loginLogService.List(param, false)
	for _, logininfor := range logininfors {
		list = append(list, dto.LoginLogExportResponse{
			InfoId:        logininfor.InfoId,
			UserName:      logininfor.UserName,
			Status:        logininfor.Status,
			Ipaddr:        logininfor.Ipaddr,
			LoginLocation: logininfor.LoginLocation,
			Browser:       logininfor.Browser,
			Os:            logininfor.Os,
			Msg:           logininfor.Msg,
			LoginTime:     logininfor.LoginTime.Format(datetime.DATETIME_FORMAT0),
		})
	}
	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	excel.DownLoadExcel("logininfor_"+time.Now().Format(datetime.DATETIME_FORMAT2), ctx.Writer, file)
}

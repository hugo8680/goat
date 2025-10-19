package admin

import (
	"github.com/hugo8680/goat/common/serializer/datetime"
	"github.com/hugo8680/goat/common/utils"
	"github.com/hugo8680/goat/framework/response"
	"github.com/hugo8680/goat/model/dto"
	"github.com/hugo8680/goat/service/admin"
	"regexp"
	"strings"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
)

type LoginLogController struct {
	loginLogService *admin.LoginLogService
}

func NewLoginLogController() *LoginLogController {
	return &LoginLogController{
		loginLogService: &admin.LoginLogService{},
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
	loginLogs, total := c.loginLogService.List(param, true)
	response.Success(ctx).SetPageData(loginLogs, total).Json()
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
	loginLogs, _ := c.loginLogService.List(param, false)
	for _, loginLog := range loginLogs {
		list = append(list, dto.LoginLogExportResponse{
			InfoId:        loginLog.InfoId,
			UserName:      loginLog.UserName,
			Status:        loginLog.Status,
			Ipaddr:        loginLog.Ipaddr,
			LoginLocation: loginLog.LoginLocation,
			Browser:       loginLog.Browser,
			Os:            loginLog.Os,
			Msg:           loginLog.Msg,
			LoginTime:     loginLog.LoginTime.Format(datetime.DATETIME_FORMAT0),
		})
	}
	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	excel.DownLoadExcel("login_log_"+time.Now().Format(datetime.DATETIME_FORMAT2), ctx.Writer, file)
}

package admin

import (
	"forum-service/common/serializer/datetime"
	"forum-service/common/utils"
	"forum-service/framework/response"
	"forum-service/model/dto"
	"forum-service/service"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
)

type OperLogController struct {
	operLogService *service.OperLogService
}

func NewOperLogController() *OperLogController {
	return &OperLogController{
		operLogService: &service.OperLogService{},
	}
}

// List 操作日志列表
func (c *OperLogController) List(ctx *gin.Context) {
	var param dto.OperLogListRequest
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
		param.OrderByColumn = "operTime"
	}
	param.OrderByColumn = strings.ToLower(regexp.MustCompile("([A-Z])").ReplaceAllString(param.OrderByColumn, "_${1}"))
	operLogs, total := c.operLogService.List(param, true)
	response.Success(ctx).SetPageData(operLogs, total).Json()
}

// Delete 删除操作日志
func (c *OperLogController) Delete(ctx *gin.Context) {
	operIds, err := utils.StringToIntSlice(ctx.Param("operIds"), ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err = c.operLogService.Delete(operIds); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Clean 清空操作日志
func (c *OperLogController) Clean(ctx *gin.Context) {
	if err := c.operLogService.Delete(nil); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Export 数据导出
func (c *OperLogController) Export(ctx *gin.Context) {
	var param dto.OperLogListRequest
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
		param.OrderByColumn = "operTime"
	}
	param.OrderByColumn = strings.ToLower(regexp.MustCompile("([A-Z])").ReplaceAllString(param.OrderByColumn, "_${1}"))
	list := make([]dto.OperLogExportResponse, 0)
	operLogs, _ := c.operLogService.List(param, false)
	for _, operLog := range operLogs {
		list = append(list, dto.OperLogExportResponse{
			OperId:        operLog.OperId,
			Title:         operLog.Title,
			BusinessType:  operLog.BusinessType,
			Method:        operLog.Method,
			RequestMethod: operLog.RequestMethod,
			OperName:      operLog.OperName,
			DeptName:      operLog.DeptName,
			OperUrl:       operLog.OperUrl,
			OperIp:        operLog.OperIp,
			OperLocation:  operLog.OperLocation,
			OperParam:     operLog.OperParam,
			JsonResult:    operLog.JsonResult,
			Status:        operLog.Status,
			ErrorMsg:      operLog.ErrorMsg,
			OperTime:      operLog.OperTime.Format(datetime.DATETIME_FORMAT0),
			CostTime:      strconv.Itoa(operLog.CostTime) + "毫秒",
		})
	}
	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	excel.DownLoadExcel("operlog_"+time.Now().Format(datetime.DATETIME_FORMAT2), ctx.Writer, file)
}

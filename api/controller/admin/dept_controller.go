package admin

import (
	"github.com/hugo8680/goat/api/validator/admin"
	"github.com/hugo8680/goat/common/constant/auth"
	"github.com/hugo8680/goat/common/utils"
	"github.com/hugo8680/goat/framework/response"
	"github.com/hugo8680/goat/model/dto"
	adminService "github.com/hugo8680/goat/service/admin"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type DeptController struct {
	deptService *adminService.DeptService
}

func NewDeptController() *DeptController {
	return &DeptController{
		deptService: &adminService.DeptService{},
	}
}

// List 部门列表
func (c *DeptController) List(ctx *gin.Context) {
	var param dto.DeptListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	deptList := c.deptService.List(param, user.(*dto.UserTokenResponse).UserId)
	response.Success(ctx).SetData("data", deptList).Json()
}

// ListExclude 查询部门列表（排除节点）
func (c *DeptController) ListExclude(ctx *gin.Context) {
	deptId, _ := strconv.Atoi(ctx.Param("deptId"))
	data := make([]dto.DeptListResponse, 0)
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	deptList := c.deptService.List(dto.DeptListRequest{}, user.(*dto.UserTokenResponse).UserId)
	for _, dept := range deptList {
		if dept.DeptId == deptId || utils.Contains(strings.Split(dept.Ancestors, ","), strconv.Itoa(deptId)) {
			continue
		}
		data = append(data, dept)
	}
	response.Success(ctx).SetData("data", data).Json()
}

// Get 获取部门详情
func (c *DeptController) Get(ctx *gin.Context) {
	deptId, _ := strconv.Atoi(ctx.Param("deptId"))
	dept := c.deptService.Get(deptId)
	response.Success(ctx).SetData("data", dept).Json()
}

// Create 新增部门
func (c *DeptController) Create(ctx *gin.Context) {
	var param dto.CreateDeptRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.CreateDeptValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.deptService.Create(dto.SaveDeptRequest{
		ParentId: param.ParentId,
		DeptName: param.DeptName,
		OrderNum: param.OrderNum,
		Leader:   param.Leader,
		Phone:    param.Phone,
		Email:    param.Email,
		Status:   param.Status,
		CreateBy: user.(*dto.UserTokenResponse).UserName,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Update 更新部门
func (c *DeptController) Update(ctx *gin.Context) {
	var param dto.UpdateDeptRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.UpdateDeptValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.deptService.Update(dto.SaveDeptRequest{
		DeptId:   param.DeptId,
		ParentId: param.ParentId,
		DeptName: param.DeptName,
		OrderNum: param.OrderNum,
		Leader:   param.Leader,
		Phone:    param.Phone,
		Email:    param.Email,
		Status:   param.Status,
		UpdateBy: user.(*dto.UserTokenResponse).UserName,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Delete 删除部门
func (c *DeptController) Delete(ctx *gin.Context) {
	deptId, _ := strconv.Atoi(ctx.Param("deptId"))
	if err := c.deptService.Delete(deptId); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

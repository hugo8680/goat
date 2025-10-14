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

type PostController struct {
	postService *service.PostService
}

func NewPostController() *PostController {
	return &PostController{
		postService: &service.PostService{},
	}
}

// List 岗位列表
func (c *PostController) List(ctx *gin.Context) {
	var param dto.PostListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	posts, total := c.postService.List(param, true)
	response.Success(ctx).SetPageData(posts, total).Json()
}

// Get 岗位详情
func (c *PostController) Get(ctx *gin.Context) {
	postId, _ := strconv.Atoi(ctx.Param("postId"))
	post := c.postService.Get(postId)
	response.Success(ctx).SetData("data", post).Json()
}

// Create 新增岗位
func (c *PostController) Create(ctx *gin.Context) {
	var param dto.CreatePostRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.CreatePostValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.postService.Create(dto.SavePostRequest{
		PostCode: param.PostCode,
		PostName: param.PostName,
		PostSort: param.PostSort,
		Status:   param.Status,
		CreateBy: user.(dto.UserTokenResponse).UserName,
		Remark:   param.Remark,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Update 更新岗位
func (c *PostController) Update(ctx *gin.Context) {
	var param dto.UpdatePostRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err := admin.UpdatePostValidator(param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	user, _ := ctx.Get(auth.CONTEXT_USER_KEY)
	if err := c.postService.Update(dto.SavePostRequest{
		PostId:   param.PostId,
		PostCode: param.PostCode,
		PostName: param.PostName,
		PostSort: param.PostSort,
		Status:   param.Status,
		UpdateBy: user.(dto.UserTokenResponse).UserName,
		Remark:   param.Remark,
	}); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	response.Success(ctx).Json()
}

// Delete 删除岗位
func (c *PostController) Delete(ctx *gin.Context) {
	postIds, err := utils.StringToIntSlice(ctx.Param("postIds"), ",")
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	if err = c.postService.Delete(postIds); err != nil {
		response.Error(ctx).SetMsg(err.Error())
		return
	}
	response.Success(ctx).Json()
}

// Export 数据导出
func (c *PostController) Export(ctx *gin.Context) {
	var param dto.PostListRequest
	if err := ctx.ShouldBind(&param); err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	list := make([]dto.PostExportResponse, 0)
	posts, _ := c.postService.List(param, false)
	for _, post := range posts {
		list = append(list, dto.PostExportResponse{
			PostId:   post.PostId,
			PostCode: post.PostCode,
			PostName: post.PostName,
			PostSort: post.PostSort,
			Status:   post.Status,
		})
	}
	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.Error(ctx).SetMsg(err.Error()).Json()
		return
	}
	excel.DownLoadExcel("post_"+time.Now().Format(datetime.DATETIME_FORMAT2), ctx.Writer, file)
}

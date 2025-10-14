package route

import (
	"forum-service/api/controller/admin"
	"forum-service/common/constant/log_request_type"
	"forum-service/framework"
	"forum-service/middleware"

	"github.com/gin-gonic/gin"
)

// DefaultRoutes 在此处编写路由
//
// 组Name为组的名称，用于方便维护
//
// 组RelativePath为路由前缀(可重复)
//
// 组Middlewares为中间件集合，按顺序传递一个HandlerFunc数组
//
// 组Routes为组内管理的路由集合
//
// 路由Method为HTTP方法
//
// 路由RelativePath为路由前缀
//
// 路由Middlewares为路由的中间件集合，按顺序传递一个HandlerFunc数组
//
// 组的Middlewares和路由的Middlewares会形成一个并集按顺序传递
//
// 路由Function为实际处理方法
//
// 若希望自行添加新路由文件，将新的路由组按照可变参数的形式传递给RunServer即可，如：RunServer(groups1, groups2, groups3)
func DefaultRoutes() []framework.RouteGroup {
	return []framework.RouteGroup{
		{
			Name:         "管理后台（无需授权登录）",
			RelativePath: "/",
			Routes: []framework.Route{
				{
					Method:       "GET",
					RelativePath: "/captchaImage",
					Function:     admin.NewAuthController().GetCaptchaImage,
				},
				{
					Method:       "POST",
					RelativePath: "/register",
					Function:     admin.NewAuthController().Register,
				},
				{
					Method:       "POST",
					RelativePath: "/login",
					Middlewares:  gin.HandlersChain{middleware.LoginLogMiddleware()},
					Function:     admin.NewAuthController().Login,
				},
				{
					Method:       "POST",
					RelativePath: "/logout",
					Function:     admin.NewAuthController().Logout,
				},
			},
		},
		{
			Name:         "管理后台（需要授权登录）",
			RelativePath: "/",
			Middlewares:  gin.HandlersChain{middleware.AdminAuthMiddleware()},
			Routes: []framework.Route{
				{
					Method:       "GET",
					RelativePath: "/getInfo",
					Function:     admin.NewAuthController().GetInfo,
				},
				{
					Method:       "GET",
					RelativePath: "/getRouters",
					Function:     admin.NewAuthController().GetRouters,
				},
				{
					Method:       "GET",
					RelativePath: "/system/user/profile",
					Function:     admin.NewUserController().GetProfile,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/user/profile",
					Function:     admin.NewUserController().UpdateProfile,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/user/profile/updatePwd",
					Function:     admin.NewUserController().UpdatePassword,
				},
				{
					Method:       "POST",
					RelativePath: "/system/user/profile/avatar",
					Function:     admin.NewUserController().UpdateAvatar,
				},
				{
					Method:       "GET",
					RelativePath: "/system/user/deptTree",
					Middlewares:  gin.HandlersChain{middleware.PermissionCheckMiddleware("system:user:list")},
					Function:     admin.NewUserController().DeptTree,
				},
				{
					Method:       "GET",
					RelativePath: "/system/user/list",
					Middlewares:  gin.HandlersChain{middleware.PermissionCheckMiddleware("system:user:list")},
					Function:     admin.NewUserController().List,
				},
				{
					Method:       "GET",
					RelativePath: "/system/user",
					Middlewares:  gin.HandlersChain{middleware.PermissionCheckMiddleware("system:user:query")},
					Function:     admin.NewUserController().Get,
				},
				{
					Method:       "GET",
					RelativePath: "/system/user/:userId",
					Middlewares:  gin.HandlersChain{middleware.PermissionCheckMiddleware("system:user:query")},
					Function:     admin.NewUserController().Get,
				},
				{
					Method:       "GET",
					RelativePath: "/system/user/authRole/:userId",
					Middlewares:  gin.HandlersChain{middleware.PermissionCheckMiddleware("system:user:query")},
					Function:     admin.NewUserController().ListRoleByUserId,
				},
				{
					Method:       "POST",
					RelativePath: "/system/user",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:user:add"),
						middleware.OperLogMiddleware("新增用户", log_request_type.REQUEST_BUSINESS_TYPE_INSERT),
					},
					Function: admin.NewUserController().Create,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/user",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:user:edit"),
						middleware.OperLogMiddleware("更新用户", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewUserController().Update,
				},
				{
					Method:       "DELETE",
					RelativePath: "/system/user/:userIds",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:user:remove"),
						middleware.OperLogMiddleware("删除用户", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewUserController().Delete,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/user/changeStatus",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:user:edit"),
						middleware.OperLogMiddleware("修改用户状态", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewUserController().ChangeStatus,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/user/resetPwd",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:user:edit"),
						middleware.OperLogMiddleware("修改用户密码", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewUserController().ResetPassword,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/user/authRole",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:user:edit"),
						middleware.OperLogMiddleware("用户授权角色", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewUserController().AuthRoles,
				},
				{
					Method:       "POST",
					RelativePath: "/system/user/export",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:user:export"),
						middleware.OperLogMiddleware("导出用户", log_request_type.REQUEST_BUSINESS_TYPE_EXPORT),
					},
					Function: admin.NewUserController().Export,
				},
				{
					Method:       "POST",
					RelativePath: "/system/user/importData",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:user:import"),
						middleware.OperLogMiddleware("导入用户", log_request_type.REQUEST_BUSINESS_TYPE_IMPORT),
					},
					Function: admin.NewUserController().Import,
				},
				{
					Method:       "POST",
					RelativePath: "/system/user/importTemplate",
					Middlewares: gin.HandlersChain{
						middleware.OperLogMiddleware("导入用户模板", log_request_type.REQUEST_BUSINESS_TYPE_EXPORT),
					},
					Function: admin.NewUserController().DownloadImportTemplate,
				},
				{
					Method:       "GET",
					RelativePath: "/system/role/list",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:list"),
					},
					Function: admin.NewRoleController().List,
				},
				{
					Method:       "GET",
					RelativePath: "/system/role/:roleId",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:query"),
					},
					Function: admin.NewRoleController().Get,
				},
				{
					Method:       "GET",
					RelativePath: "/system/role/deptTree/:roleId",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:query"),
					},
					Function: admin.NewRoleController().DeptTree,
				},
				{
					Method:       "GET",
					RelativePath: "/system/role/authUser/allocatedList",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:list"),
					},
					Function: admin.NewRoleController().RoleUsersAllocated,
				},
				{
					Method:       "GET",
					RelativePath: "/system/role/authUser/unallocatedList",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:list"),
					},
					Function: admin.NewRoleController().RoleUsersUnAllocated,
				},
				{
					Method:       "POST",
					RelativePath: "/system/role",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:add"),
						middleware.OperLogMiddleware("新增角色", log_request_type.REQUEST_BUSINESS_TYPE_INSERT),
					},
					Function: admin.NewRoleController().Create,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/role",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:edit"),
						middleware.OperLogMiddleware("更新角色", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewRoleController().Update,
				},
				{
					Method:       "DELETE",
					RelativePath: "/system/role/:roleIds",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:remove"),
						middleware.OperLogMiddleware("删除角色", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewRoleController().Delete,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/role/changeStatus",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:edit"),
						middleware.OperLogMiddleware("修改角色状态", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewRoleController().ChangeStatus,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/role/dataScope",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:edit"),
						middleware.OperLogMiddleware("分配数据权限", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewRoleController().AssignDataScope,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/role/authUser/selectAll",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:edit"),
						middleware.OperLogMiddleware("批量选择用户授权", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewRoleController().AuthUsers,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/role/authUser/cancel",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:edit"),
						middleware.OperLogMiddleware("取消用户授权", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewRoleController().UnAuthUser,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/role/authUser/cancelAll",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:edit"),
						middleware.OperLogMiddleware("批量取消用户授权", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewRoleController().UnAuthUsers,
				},
				{
					Method:       "POST",
					RelativePath: "/system/role/export",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:role:export"),
						middleware.OperLogMiddleware("导出角色", log_request_type.REQUEST_BUSINESS_TYPE_EXPORT),
					},
					Function: admin.NewRoleController().Export,
				},
				{
					Method:       "GET",
					RelativePath: "/system/menu/list",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:menu:list"),
					},
					Function: admin.NewMenuController().List,
				},
				{
					Method:       "GET",
					RelativePath: "/system/menu/treeselect",
					Function:     admin.NewMenuController().Tree,
				},
				{
					Method:       "GET",
					RelativePath: "/system/menu/roleMenuTreeselect/:roleId",
					Function:     admin.NewMenuController().RoleMenuTree,
				},
				{
					Method:       "GET",
					RelativePath: "/system/menu/:menuId",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:menu:query"),
					},
					Function: admin.NewMenuController().Get,
				},
				{
					Method:       "POST",
					RelativePath: "/system/menu",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:menu:add"),
						middleware.OperLogMiddleware("新增菜单", log_request_type.REQUEST_BUSINESS_TYPE_INSERT),
					},
					Function: admin.NewRoleController().Create,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/menu",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:menu:edit"),
						middleware.OperLogMiddleware("修改菜单", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewRoleController().Update,
				},
				{
					Method:       "DELETE",
					RelativePath: "/system/menu/:menuId",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:menu:remove"),
						middleware.OperLogMiddleware("删除菜单", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewRoleController().Delete,
				},
				{
					Method:       "GET",
					RelativePath: "/system/dept/list",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dept:list"),
					},
					Function: admin.NewDeptController().List,
				},
				{
					Method:       "GET",
					RelativePath: "/system/dept/list/exclude/:deptId",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dept:list"),
					},
					Function: admin.NewDeptController().ListExclude,
				},
				{
					Method:       "GET",
					RelativePath: "/system/dept/:deptId",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dept:query"),
					},
					Function: admin.NewDeptController().Get,
				},
				{
					Method:       "POST",
					RelativePath: "/system/dept",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dept:add"),
						middleware.OperLogMiddleware("新增部门", log_request_type.REQUEST_BUSINESS_TYPE_INSERT),
					},
					Function: admin.NewDeptController().Create,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/dept",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dept:edit"),
						middleware.OperLogMiddleware("修改部门", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewDeptController().Update,
				},
				{
					Method:       "DELETE",
					RelativePath: "/system/dept/:deptId",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dept:remove"),
						middleware.OperLogMiddleware("删除部门", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewDeptController().Delete,
				},
				{
					Method:       "GET",
					RelativePath: "/system/post/list",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:post:list"),
					},
					Function: admin.NewPostController().List,
				},
				{
					Method:       "GET",
					RelativePath: "/system/post/:postId",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:post:query"),
					},
					Function: admin.NewPostController().Get,
				},
				{
					Method:       "POST",
					RelativePath: "/system/post",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:post:add"),
						middleware.OperLogMiddleware("新增岗位", log_request_type.REQUEST_BUSINESS_TYPE_INSERT),
					},
					Function: admin.NewPostController().Create,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/post",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:post:edit"),
						middleware.OperLogMiddleware("更新岗位", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewPostController().Update,
				},
				{
					Method:       "DELETE",
					RelativePath: "/system/post/:postIds",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:post:remove"),
						middleware.OperLogMiddleware("删除岗位", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewPostController().Delete,
				},
				{
					Method:       "POST",
					RelativePath: "/system/post/export",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:post:export"),
						middleware.OperLogMiddleware("导出岗位", log_request_type.REQUEST_BUSINESS_TYPE_EXPORT),
					},
					Function: admin.NewPostController().Export,
				},
				{
					Method:       "GET",
					RelativePath: "/system/dict/list",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:list"),
					},
					Function: admin.NewDictTypeController().List,
				},
				{
					Method:       "GET",
					RelativePath: "/system/dict/type/:dictId",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:query"),
					},
					Function: admin.NewDictTypeController().Get,
				},
				{
					Method:       "GET",
					RelativePath: "/system/dict/type/optionselect",
					Function:     admin.NewDictTypeController().DictOptions,
				},
				{
					Method:       "POST",
					RelativePath: "/system/dict/type",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:add"),
						middleware.OperLogMiddleware("新增字典类型", log_request_type.REQUEST_BUSINESS_TYPE_INSERT),
					},
					Function: admin.NewDictTypeController().Create,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/dict/type",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:edit"),
						middleware.OperLogMiddleware("更新字典类型", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewDictTypeController().Update,
				},
				{
					Method:       "DELETE",
					RelativePath: "/system/dict/type/:dictIds",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:remove"),
						middleware.OperLogMiddleware("删除字典类型", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewDictTypeController().Delete,
				},
				{
					Method:       "POST",
					RelativePath: "/system/dict/type/export",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:export"),
						middleware.OperLogMiddleware("导出字典类型", log_request_type.REQUEST_BUSINESS_TYPE_EXPORT),
					},
					Function: admin.NewDictTypeController().Export,
				},
				{
					Method:       "DELETE",
					RelativePath: "/system/dict/type/refreshCache",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:remove"),
						middleware.OperLogMiddleware("刷新字典类型缓存", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewDictTypeController().RefreshCache,
				},
				{
					Method:       "GET",
					RelativePath: "/system/dict/data/list",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:list"),
					},
					Function: admin.NewDictDataController().List,
				},
				{
					Method:       "GET",
					RelativePath: "/system/dict/data/:dictCode",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:query"),
					},
					Function: admin.NewDictDataController().Get,
				},
				{
					Method:       "GET",
					RelativePath: "/system/dict/data/type/:dictType",
					Function:     admin.NewDictDataController().DictDataOptions,
				},
				{
					Method:       "POST",
					RelativePath: "/system/dict/data",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:add"),
						middleware.OperLogMiddleware("新增字典数据", log_request_type.REQUEST_BUSINESS_TYPE_INSERT),
					},
					Function: admin.NewDictDataController().Create,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/dict/data",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:edit"),
						middleware.OperLogMiddleware("更新字典数据", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewDictDataController().Update,
				},
				{
					Method:       "DELETE",
					RelativePath: "/system/dict/data/:dictCodes",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:remove"),
						middleware.OperLogMiddleware("删除字典数据", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewDictDataController().Delete,
				},
				{
					Method:       "POST",
					RelativePath: "/system/dict/data/export",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:dict:export"),
						middleware.OperLogMiddleware("导出字典数据", log_request_type.REQUEST_BUSINESS_TYPE_EXPORT),
					},
					Function: admin.NewDictDataController().Export,
				},
				{
					Method:       "GET",
					RelativePath: "/system/config/list",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:config:list"),
					},
					Function: admin.NewConfigController().List,
				},
				{
					Method:       "GET",
					RelativePath: "/system/config/:configId",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:config:query"),
					},
					Function: admin.NewConfigController().Get,
				},
				{
					Method:       "GET",
					RelativePath: "/system/config/configKey/:configKey",
					Function:     admin.NewConfigController().ConfigKey,
				},
				{
					Method:       "POST",
					RelativePath: "/system/config",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:config:add"),
						middleware.OperLogMiddleware("新增参数配置", log_request_type.REQUEST_BUSINESS_TYPE_INSERT),
					},
					Function: admin.NewConfigController().Create,
				},
				{
					Method:       "PUT",
					RelativePath: "/system/config",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:config:edit"),
						middleware.OperLogMiddleware("更新参数配置", log_request_type.REQUEST_BUSINESS_TYPE_UPDATE),
					},
					Function: admin.NewConfigController().Update,
				},
				{
					Method:       "DELETE",
					RelativePath: "/system/config/:configIds",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:config:remove"),
						middleware.OperLogMiddleware("删除参数配置", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewConfigController().Delete,
				},
				{
					Method:       "POST",
					RelativePath: "/system/config/export",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:config:export"),
						middleware.OperLogMiddleware("导出参数配置", log_request_type.REQUEST_BUSINESS_TYPE_EXPORT),
					},
					Function: admin.NewConfigController().Export,
				},
				{
					Method:       "DELETE",
					RelativePath: "/system/config/refreshCache",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("system:config:remove"),
						middleware.OperLogMiddleware("刷新参数配置缓存", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewConfigController().RefreshCache,
				},
				{
					Method:       "GET",
					RelativePath: "/monitor/logininfor/list",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("monitor:loginlog:list"),
					},
					Function: admin.NewLoginLogController().List,
				},
				{
					Method:       "DELETE",
					RelativePath: "/monitor/logininfor/:infoIds",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("monitor:loginlog:remove"),
						middleware.OperLogMiddleware("删除登录日志", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewLoginLogController().Delete,
				},
				{
					Method:       "DELETE",
					RelativePath: "/monitor/logininfor/clean",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("monitor:loginlog:remove"),
						middleware.OperLogMiddleware("清空登录日志", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewLoginLogController().Clean,
				},
				{
					Method:       "GET",
					RelativePath: "/monitor/logininfor/unlock/:userName",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("monitor:loginlog:unlock"),
						middleware.OperLogMiddleware("账户解锁", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewLoginLogController().Unlock,
				},
				{
					Method:       "POST",
					RelativePath: "/monitor/logininfor/export",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("monitor:loginlog:export"),
						middleware.OperLogMiddleware("导出登录日志", log_request_type.REQUEST_BUSINESS_TYPE_EXPORT),
					},
					Function: admin.NewLoginLogController().Export,
				},
				{
					Method:       "GET",
					RelativePath: "/monitor/operlog/list",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("monitor:operlog:list"),
					},
					Function: admin.NewOperLogController().List,
				},
				{
					Method:       "DELETE",
					RelativePath: "/monitor/operlog/:operIds",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("monitor:operlog:remove"),
						middleware.OperLogMiddleware("删除操作日志", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewOperLogController().Delete,
				},
				{
					Method:       "DELETE",
					RelativePath: "/monitor/operlog/clean",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("monitor:operlog:remove"),
						middleware.OperLogMiddleware("清空操作日志", log_request_type.REQUEST_BUSINESS_TYPE_DELETE),
					},
					Function: admin.NewOperLogController().Clean,
				},
				{
					Method:       "POST",
					RelativePath: "/monitor/operlog/export",
					Middlewares: gin.HandlersChain{
						middleware.PermissionCheckMiddleware("monitor:operlog:export"),
						middleware.OperLogMiddleware("导出操作日志", log_request_type.REQUEST_BUSINESS_TYPE_EXPORT),
					},
					Function: admin.NewOperLogController().Export,
				},
			},
		},
	}
}

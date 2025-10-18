package framework

import (
	"forum-service/middleware"

	"github.com/gin-gonic/gin"
)

// Route 路由配置
//
// Method Http方法 GET POST PUT DELETE OPTIONS
// RelativePath 路由前缀
// Middlewares 中间件组
// Function 控制器方法
type Route struct {
	Method       string
	RelativePath string
	Middlewares  gin.HandlersChain
	Function     gin.HandlerFunc
}

// RouteGroup 路由组
//
// Name 名称
// RelativePath 路由前缀
// Middlewares 中间件组
// Routes 路由组子项
type RouteGroup struct {
	Name         string
	RelativePath string
	Middlewares  gin.HandlersChain
	Routes       []Route
}

func registerCommonMiddlewares(server *gin.Engine) {
	server.Use(middleware.RecoveryMiddleware())
	server.Use(middleware.CorsMiddleware())
}

// registerRouteGroups 注册分组路由
func registerRouteGroups(groups []RouteGroup) {
	server := getServer()
	rootRoute := server.Group("/api")
	for _, group := range groups {
		g := rootRoute.Group(group.RelativePath)
		if group.Middlewares != nil {
			g.Use(group.Middlewares...)
		}
		for _, route := range group.Routes {
			if route.Middlewares != nil {
				handlers := route.Middlewares
				handlers = append(handlers, route.Function)
				g.Handle(route.Method, route.RelativePath, handlers...)
			} else {
				g.Handle(route.Method, route.RelativePath, route.Function)
			}
		}
	}
}

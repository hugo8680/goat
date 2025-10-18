package framework

import (
	"fmt"
	"forum-service/framework/config"
	"forum-service/framework/connector"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	Port string
}

var server *Server
var once sync.Once

func init() {
	conf := config.GetSetting()
	gin.SetMode(conf.Server.Mode)
	server = &Server{
		Engine: gin.New(),
		Port:   strconv.Itoa(conf.Server.Port),
	}
	once.Do(func() {
		registerCommonMiddlewares(server.Engine)
		connector.ConnectToMySQL()
		connector.ConnectToRedis()
		connector.InitializeLogger(server.Engine)
	})
	server.Static(conf.System.UploadPath, conf.System.UploadPath)
}

func getServer() *Server {
	return server
}

// RunServer 运行程序
//
// routeGroups 路由组，可传多个
func RunServer(routeGroups ...[]RouteGroup) {
	if routeGroups != nil && len(routeGroups) != 0 {
		for _, group := range routeGroups {
			registerRouteGroups(group)
		}
	}
	err := server.Run(fmt.Sprintf(":%s", server.Port))
	if err != nil {
		panic(err)
	}
}

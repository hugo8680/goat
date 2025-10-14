package main

import (
	"forum-service/framework"
	"forum-service/route"
)

func main() {
	framework.RunServer(route.DefaultRoutes(), route.APIRoutes())
}

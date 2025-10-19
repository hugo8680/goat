package main

import (
	"github.com/hugo8680/goat/framework"
	"github.com/hugo8680/goat/route"
)

func main() {
	framework.RunServer(route.DefaultRoutes(), route.APIRoutes())
}

package router

import (
	"cosmonaut_api/app/api"
	"cosmonaut_api/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.Ctx, service.Middleware.CORS)
		group.ALL("/hello", api.Hello)
		group.ALL("/user", api.User)
		group.Middleware(service.Middleware.Auth)
		group.GET("/profile", api.Profile.GetProfile)
	})
}

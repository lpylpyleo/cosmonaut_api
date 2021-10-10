package router

import (
	"cosmonaut_api/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.ALL("/hello", api.Hello)
		group.POST("/user/sign-up", api.User.SignUp)
	})
}

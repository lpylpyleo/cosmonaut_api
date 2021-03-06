package router

import (
	"cosmonaut_api/app/api"
	"cosmonaut_api/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.SetIndexFolder(false)
	s.AddSearchPath("./upload/image/avatar")
	s.AddStaticPath("/avatars", "./upload/image/avatar")

	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.Ctx, service.Middleware.CORS)
		group.GET("/hello", api.Hello)
		group.POST("/user/sign-in", api.User.SignIn)
		group.POST("/user/sign-up", api.User.SignUp)

		group.Middleware(service.Middleware.Auth)
		group.POST("/user/sign-out", api.User.SignOut)
		group.GET("/profile", api.Profile.GetProfile)
		group.POST("/profile/changeAvatar", api.Profile.ChangeAvatar)
		group.POST("/profile/edit", api.Profile.Edit)
		group.POST("/post", api.Post.Create)
		group.GET("/post", api.Post.GetAll)
		group.GET("/post/:id", api.Post.GetOne)
		group.POST("/post/like", api.Post.Like)
	})
}

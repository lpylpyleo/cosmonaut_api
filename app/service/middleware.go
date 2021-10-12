package service

import (
	"cosmonaut_api/app/model"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var Middleware = middlewareService{}

type middlewareService struct {
}

func (s middlewareService) Ctx(r *ghttp.Request) {
	customContext := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customContext)
	if user := Session.GetUser(r.Context()); user != nil {
		customContext.User = &model.ContextUser{
			Id:          user.Uid,
			Email:       user.Email,
			DisplayName: user.Email,
		}
	}
	r.Middleware.Next()
}

func (s middlewareService) Auth(r *ghttp.Request) {
	if User.IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
}

func (s middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

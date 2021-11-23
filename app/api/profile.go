package api

import (
	"cosmonaut_api/app/service"
	"cosmonaut_api/library/response"
	"net/http"

	"github.com/gogf/gf/net/ghttp"
)

var Profile = profileApi{}

type profileApi struct{}

// GetProfile 个人资料
func (a profileApi) GetProfile(r *ghttp.Request) {
	profile, err := service.Profile.GetProfile(r.Context())
	if err != nil {
		r.Response.Status = http.StatusNotFound
		response.JsonExit(r, response.CODE_USER_NOT_EXIST, "用户不存在", err.Error())
	}
	_ = r.Response.WriteJson(profile)
}

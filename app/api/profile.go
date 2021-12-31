package api

import (
	"cosmonaut_api/app/model"
	"cosmonaut_api/app/service"
	"cosmonaut_api/library/response"
	"net/http"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
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

func (a profileApi) ChangeAvatar(r *ghttp.Request) {
	file := r.GetUploadFile("avatar")
	if file == nil {
		response.JsonExit(r, response.CODE_PARSE_FORM_ERROR, "必须包含头像文件")
	}
	err := service.Profile.ChangeAvatar(r.Context(), file)
	if err != nil {
		r.Response.Status = http.StatusNotFound
		response.JsonExit(r, response.CODE_USER_NOT_EXIST, "用户不存在", err.Error())
	}
	r.Exit()
}

func (a profileApi) Edit(r *ghttp.Request) {
	var (
		apiReq     *model.ProfileApiEditAvatarReq
		serviceReq *model.ProfileServiceEditAvatarReq
	)

	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, response.CODE_PARSE_FORM_ERROR, err.Error())
	}

	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, response.CODE_PARSE_FORM_ERROR, err.Error())
	}

	if err := service.Profile.Edit(r.Context(), serviceReq); err != nil {
		response.JsonExit(r, response.CODE_UNKNOWN, err.Error())
	}

	r.Exit()
}

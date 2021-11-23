package api

import (
	"cosmonaut_api/app/model"
	"cosmonaut_api/app/service"
	"cosmonaut_api/library/response"
	"cosmonaut_api/library/util"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"strings"
)

var User = userApi{}

type userApi struct{}

func (a userApi) Index(r *ghttp.Request) {
	_ = r.Response.WriteTpl("user/index.html")
}

func (a userApi) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignUpReq
		serviceReq *model.UserServiceSignUpReq
	)

	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, response.CODE_PARSE_FORM_ERROR, err.Error())
	}
	serviceReq = &model.UserServiceSignUpReq{
		Id:       util.GenUid(),
		Email:    strings.TrimSpace(apiReq.Email),
		Password: util.GetSha256Digest(apiReq.Password),
	}
	if err := service.User.SignUp(serviceReq); err != nil {
		r.Response.Status = http.StatusConflict
		response.JsonExit(r, response.CODE_UNKNOWN, err.Error())
	}

	r.Exit()
}

func (a userApi) SignIn(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignInReq
		serviceReq *model.UserServiceSignUpReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		r.Response.Status = http.StatusBadRequest
		response.JsonExit(r, response.CODE_PARSE_FORM_ERROR, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		r.Response.Status = http.StatusInternalServerError
		response.JsonExit(r, response.CODE_UNKNOWN, err.Error())
	}
	if err := service.User.SignIn(r.Context(), serviceReq.Email, serviceReq.Password); err != nil {
		response.JsonExit(r, response.CODE_UNKNOWN, err.Error())
	}
	r.Exit()
}

func (a userApi) SignOut(r *ghttp.Request) {
	if err := service.User.SignOut(r.Context()); err != nil {
		response.JsonExit(r, response.CODE_UNKNOWN, err.Error())
	}
	r.Exit()
}

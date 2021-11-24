package api

import (
	"cosmonaut_api/app/model"
	"cosmonaut_api/app/service"
	"cosmonaut_api/library/response"

	"github.com/gogf/gf/net/ghttp"
)

var Post = postApi{}

type postApi struct{}

func (a postApi) Create(r *ghttp.Request) {
	var apiReq *model.PostApiCreateReq
	var serviceReq *model.PostServiceCreateReq

	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, response.CODE_PARSE_FORM_ERROR, err.Error())
	}

	user := service.Session.GetUser(r.Context())

	serviceReq = &model.PostServiceCreateReq{
		Creator:  user.Id,
		Content:  apiReq.Content,
		IsPublic: apiReq.IsPublic,
	}

	if err := service.Post.Create(serviceReq); err != nil {
		response.JsonExit(r, response.CODE_UNKNOWN, err.Error())
	}

	r.Exit()
}

func (a postApi) Get(r *ghttp.Request) {
	list, err := service.Post.GetAll()
	response.ExitIfErr(r, err)

	response.Json(r, &list)
}
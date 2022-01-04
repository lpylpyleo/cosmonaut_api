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
	user := service.Session.GetUser(r.Context())
	list, err := service.Post.GetAll(&model.PostServiceGetReq{Uid: user.Id})
	response.ExitIfErr(r, err)

	response.Json(r, &list)
}

func (a postApi) Like(r *ghttp.Request) {
	var (
		apiReq     *model.PostApiLikeReq
		serviceReq *model.PostServiceLikeReq
	)

	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, response.CODE_PARSE_FORM_ERROR, err.Error())
	}

	user := service.Session.GetUser(r.Context())
	serviceReq = &model.PostServiceLikeReq{
		Uid:    user.Id,
		Pid:    apiReq.Pid,
		IsLike: apiReq.IsLike,
	}

	err := service.Post.Like(*serviceReq)
	response.ExitIfErr(r, err)
}

package service

import (
	"cosmonaut_api/app/dao"
	"cosmonaut_api/app/model"
)

var Post = postService{}

type postService struct{}

func (s postService) Create(req *model.PostServiceCreateReq) error {
	_, err := dao.Post.Insert(req)
	return err
}

func (s postService) GetAll() (*[]model.PostResponse, error) {
	profile := dao.Profile.GetFieldsExStr(`id,created_at,updated_at,deleted_at`, `"profile".`)
	post := dao.Post.GetFieldsStr(`"post".`)
	r, err := dao.Post.Fields(post+","+profile).
		LeftJoin("profile", `"profile"."uid" = "post"."creator"`).
		OrderDesc(`"post"."id"`).
		All(`"post"."is_public"=true`)
	if err != nil {
		return nil, err
	}
	var posts []model.PostResponse
	err = r.Structs(&posts)
	return &posts, err
}

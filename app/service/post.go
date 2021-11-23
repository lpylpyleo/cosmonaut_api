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

func (s postService) Get() (*[]model.PostProfile, error) {
	r, err := dao.Post.LeftJoin("profile", `"profile"."uid" = "post"."creater"`).All()
	if err != nil {
		return nil, err
	}
	var posts []model.PostProfile
	err = r.Structs(&posts)
	return &posts, err
}

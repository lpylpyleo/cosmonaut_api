package model

type PostApiCreateReq struct {
	Content  string `v:"required"`
	IsPublic bool
}

type PostServiceCreateReq struct {
	Creater  string `v:"required"`
	Content  string `v:"required"`
	IsPublic bool
}

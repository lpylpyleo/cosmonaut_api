package model

type PostServiceLikeReq struct {
	Uid    string `v:"required"`
	Pid    int    `v:"required"`
	IsLike bool   `v:"required"`
}

type PostApiLikeReq struct {
	Pid    int  `v:"required"`
	IsLike bool `v:"required"`
}

package model

import "github.com/gogf/gf/os/gtime"

type PostApiCreateReq struct {
	Content  string `v:"required"`
	IsPublic bool
}

type PostServiceCreateReq struct {
	Creator  string `v:"required"`
	Content  string `v:"required"`
	IsPublic bool
}

type PostResponse struct {
	Id        int64       `orm:"id"         json:"id"`        //
	Creator   string      `orm:"creator"    json:"creator"`   //
	Title     string      `orm:"title"      json:"title"`     //
	Content   string      `orm:"content"    json:"content"`   //
	IsPublic  bool        `orm:"is_public"  json:"isPublic"`  //
	Deleted   bool        `orm:"deleted"    json:"deleted"`   //
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` //

	Uid      string `orm:"uid"        json:"uid"`      //
	Avatar   string `orm:"avatar"     json:"avatar"`   //
	Nickname string `orm:"nickname"   json:"nickname"` //
}

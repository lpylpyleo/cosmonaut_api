// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// PostDao is the manager for logic model data accessing and custom defined data operations functions management.
type PostDao struct {
	gmvc.M             // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      postColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB      // DB is the raw underlying database management object.
	Table  string      // Table is the underlying table name of the DAO.
}

// PostColumns defines and stores column names for table post.
type postColumns struct {
	Id           string //
	Creator      string //
	Title        string //
	Content      string //
	CommentCount string //
	LikeCount    string //
	IsPublic     string //
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
}

// NewPostDao creates and returns a new DAO object for table data access.
func NewPostDao() *PostDao {
	columns := postColumns{
		Id:           "id",
		Creator:      "creator",
		Title:        "title",
		Content:      "content",
		CommentCount: "comment_count",
		LikeCount:    "like_count",
		IsPublic:     "is_public",
		CreatedAt:    "created_at",
		UpdatedAt:    "updated_at",
		DeletedAt:    "deleted_at",
	}
	return &PostDao{
		C:     columns,
		M:     g.DB("default").Model("post").Safe(),
		DB:    g.DB("default"),
		Table: "post",
	}
}

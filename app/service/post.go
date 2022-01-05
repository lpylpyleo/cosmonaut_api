package service

import (
	"cosmonaut_api/app/dao"
	"cosmonaut_api/app/model"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var Post = postService{}

type postService struct{}

func (s postService) Create(req *model.PostServiceCreateReq) error {
	_, err := dao.Post.Insert(req)
	return err
}

func (s postService) GetAll(req *model.PostServiceGetReq) (gdb.Result, error) {
	sql := `SELECT DISTINCT ON("po"."id")
    "po"."id",
    "po"."creator",
    "po"."title",
    "po"."content",
    "po"."comment_count",
    "po"."like_count",
    "po"."is_public",
    "po"."created_at",
    "po"."updated_at",
    "po"."deleted_at",
    "pf"."uid",
    "pf"."avatar",
    "pf"."nickname",
    "like"."id" IS NOT NULL AS liked
FROM
    "post" as "po"
    LEFT JOIN "profile" as "pf" ON ("pf"."uid" = "po"."creator")
    LEFT JOIN "user_post_like" AS "like" ON ("like"."uid" = ?
            AND "like"."pid" = "po"."id"
			AND "like"."deleted_at" IS NULL
            )
WHERE
    "po"."is_public" = TRUE
    AND "po"."deleted_at" IS NULL
    AND "pf"."deleted_at" IS NULL
ORDER BY
    "po"."id" DESC;
`

	r, err := dao.Post.DB.GetAll(sql, g.Slice{req.Uid})
	if err != nil {
		return nil, err
	}
	// var posts []model.PostResponse
	// err = r.Structs(&posts)
	// err = gconv.Struct(r, &posts)
	return r, err
}

// TODO use redis to cache
func (s postService) Like(req model.PostServiceLikeReq) error {
	cnt, err := dao.UserPostLike.Count("uid = ? AND pid = ?", req.Uid, req.Pid)
	if err != nil {
		return err
	}
	if req.IsLike && cnt == 0 {
		_, err = dao.UserPostLike.Data(g.Map{"uid": req.Uid, "pid": req.Pid}).Insert()
		if err != nil {
			return err
		}
		_, err = dao.Post.Update(`"like_count" = "like_count" + 1`, `"id" = ?`, req.Pid)
		if err != nil {
			return err
		}

	} else if !req.IsLike && cnt > 0 {
		_, err = dao.UserPostLike.Delete("uid = ? AND pid = ?", req.Uid, req.Pid)
		if err != nil {
			return err
		}
		_, err = dao.Post.Update(`"like_count" = "like_count" - 1`, `"id" = ?`, req.Pid)
		if err != nil {
			return err
		}
	}
	return err
}

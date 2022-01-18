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

func (s postService) GetOne(req *model.PostServiceGetOneReq) (gdb.Record, error) {
	profile := dao.Profile.GetFieldsExStr(`id,created_at,updated_at,deleted_at`, `"pf".`)
	post := dao.Post.GetFieldsStr(`"po".`)

	sql := `SELECT DISTINCT ON("po"."id")
	` + post + "," + profile +
		`, "like"."id" IS NOT NULL AS liked
		FROM
		    "post" as "po"
		    LEFT JOIN "profile" as "pf" ON ("pf"."uid" = "po"."creator")
		    LEFT JOIN "user_post_like" AS "like" ON
			(
				    "like"."uid" = ?
		            AND "like"."pid" = "po"."id"
					AND "like"."deleted_at" IS NULL
		    )
		WHERE
		    "po"."is_public" = TRUE
		    AND "po"."id" = ?
		    AND "po"."deleted_at" IS NULL
		    AND "pf"."deleted_at" IS NULL
		ORDER BY
		    "po"."id" DESC;`

	r, err := dao.Post.DB.GetOne(sql, g.Slice{req.Uid, req.Id})
	if err != nil {
		return nil, err
	}

	return r, err
}

func (s postService) GetAll(req *model.PostServiceGetAllReq) (gdb.Result, error) {
	profile := dao.Profile.GetFieldsExStr(`id,created_at,updated_at,deleted_at`, `"pf".`)
	post := dao.Post.GetFieldsStr(`"po".`)

	sql := `SELECT DISTINCT ON("po"."id")
	` + post + "," + profile +
		`, "like"."id" IS NOT NULL AS liked
		FROM
		    "post" as "po"
		    LEFT JOIN "profile" as "pf" ON ("pf"."uid" = "po"."creator")
		    LEFT JOIN "user_post_like" AS "like" ON
			(
				    "like"."uid" = ?
		            AND "like"."pid" = "po"."id"
					AND "like"."deleted_at" IS NULL
		    )
		WHERE
		    "po"."is_public" = TRUE
		    AND "po"."deleted_at" IS NULL
		    AND "pf"."deleted_at" IS NULL
		ORDER BY
		    "po"."id" DESC;`

	r, err := dao.Post.DB.GetAll(sql, g.Slice{req.Uid})
	if err != nil {
		return nil, err
	}

	return r, err
}

// TODO use redis to cache
func (s postService) Like(req *model.PostServiceLikeReq) error {
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

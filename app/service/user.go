package service

import (
	"context"
	"cosmonaut_api/app/dao"
	"cosmonaut_api/app/model"
	"cosmonaut_api/library/util"
	"errors"
	"fmt"
)

var User = userService{}

type userService struct{}

func (s userService) SignUp(r *model.UserServiceSignUpReq) error {
	if !s.CheckEmail(r.Email) {
		return fmt.Errorf("账号 %s 已经存在", r.Email)
	}
	_, err := dao.User.Insert(r)
	return err
}

// SignIn 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (s *userService) SignIn(ctx context.Context, email, password string) error {
	var user *model.User
	err := dao.User.Where("email=? and password=?", email, util.GetSha256Digest(password)).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("账号或密码错误")
	}
	var profile *model.Profile
	err = dao.Profile.Where(dao.Profile.C.Uid, user.Id).Scan(&profile)
	if err != nil {
		return errors.New(err.Error())
	}

	if err := Session.SetUser(ctx, &model.UserProfile{User: user, Profile: profile}); err != nil {
		return err
	}
	Context.SetUser(ctx, &model.ContextUser{
		Id:          user.Id,
		Email:       user.Email,
		DisplayName: profile.Nickname,
	})
	return nil
}

func (s userService) IsSignedIn(ctx context.Context) bool {
	return Context.Get(ctx).User != nil
}

func (s userService) CheckEmail(email string) bool {
	if i, err := dao.User.FindCount(dao.User.C.Email, email); err != nil {
		return false
	} else {
		return i == 0
	}
}

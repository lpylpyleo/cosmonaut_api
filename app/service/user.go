package service

import (
	"cosmonaut_api/app/dao"
	"cosmonaut_api/app/model"
	"fmt"
)

var User = userService{}

type userService struct {
}

func (s userService) SignUp(r *model.UserServiceSignUpReq) error {
	if !s.CheckEmail(r.Email) {
		return fmt.Errorf("账号 %s 已经存在", r.Email)
	}
	_, err := dao.User.Insert(r)
	return err
}

func (s userService) CheckEmail(email string) bool {
	if i, err := dao.User.FindCount(dao.User.C.Email, email); err != nil {
		return false
	} else {
		return i == 0
	}
}

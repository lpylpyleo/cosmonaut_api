package service

import (
	"context"
	"cosmonaut_api/app/dao"
	"cosmonaut_api/app/model"
	"cosmonaut_api/config"
	"errors"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Profile = profileService{}

type profileService struct{}

func (s profileService) GetProfile(ctx context.Context) (*model.Profile, error) {
	if user := Session.GetUser(ctx); user != nil {
		var profile *model.Profile
		err := dao.Profile.Where(dao.Profile.C.Uid, user.Id).Scan(&profile)
		if err != nil {
			return nil, err
		}
		return profile, nil
	}

	return nil, errors.New("user not found")
}

func (s profileService) ChangeAvatar(ctx context.Context, avatar *ghttp.UploadFile) error {
	if user := Session.GetUser(ctx); user != nil {
		filename, err := avatar.Save("./upload/image/avatar", true)
		if err != nil {
			return err
		}
		_, err = dao.Profile.
			Data(g.Map{"avatar": config.ServerURL + "/avatars/" + filename}).
			Where(dao.Profile.C.Uid, user.Id).
			Update()
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("user not found")
}

func (s profileService) Edit(ctx context.Context, req *model.ProfileServiceEditAvatarReq) error {

	if user := Session.GetUser(ctx); user != nil {
		if _, err := dao.Profile.
			OmitEmpty().
			Data(req).
			Where(dao.Profile.C.Uid, user.Id).
			Update(); err != nil {
			return err
		}
		return nil
	}

	return errors.New("user not found")
}

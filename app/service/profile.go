package service

import (
	"context"
	"cosmonaut_api/app/dao"
	"cosmonaut_api/app/model"
	"errors"
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

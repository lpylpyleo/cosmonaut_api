package service

import (
	"context"
	"cosmonaut_api/app/model"
	"errors"
)

var Profile = profileService{}

type profileService struct{}

func (s profileService) GetProfile(ctx context.Context) (*model.UserProfile, error) {
	if user := Session.GetUser(ctx); user != nil {
		//var profile *model.Profile
		//err := dao.Profile.Where(dao.Profile.C.Uid, user).Scan(profile)
		//if err != nil {
		//	return nil, err
		//}
		//return &model.UserProfile{User: user, Profile: profile}, nil
		return user, nil
	}

	return nil, errors.New("user not found")
}

package service

import (
	"context"
	"cosmonaut_api/app/model"
)

const (
	SessionUserKey = "SessionUserKey"
)

var Session = sessionService{}

type sessionService struct{}

func (s sessionService) GetUser(ctx context.Context) *model.UserProfile {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.GetVar(SessionUserKey); !v.IsNil() {
			var user *model.UserProfile
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

func (s sessionService) SetUser(ctx context.Context, user *model.UserProfile) error {
	return Context.Get(ctx).Session.Set(SessionUserKey, user)
}

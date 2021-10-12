package service

import (
	"context"
	"cosmonaut_api/app/model"
	"github.com/gogf/gf/net/ghttp"
)

var Context = contextService{}

type contextService struct {
}

func (s contextService) Init(r *ghttp.Request, context *model.Context) {
	r.SetCtxVar(model.ContextKey, context)
}

func (s contextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

func (s contextService) SetUser(ctx context.Context, ctxUsr *model.ContextUser) {
	s.Get(ctx).User = ctxUsr
}

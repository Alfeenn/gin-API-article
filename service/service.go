package service

import (
	"context"

	"github.com/Alfeenn/article/model/web"
)

type Service interface {
	Create(ctx context.Context, req web.CatRequest) web.CatResp
	Update(ctx context.Context, req web.UpdateRequest) web.CatResp
	Delete(ctx context.Context, id string)
	Find(ctx context.Context, id string) web.CatResp
	FindAll(ctx context.Context, limit int, offset int) []web.CatResp
}

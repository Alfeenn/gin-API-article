package helper

import (
	"github.com/Alfeenn/article/model"
	"github.com/Alfeenn/article/model/web"
)

func ConvertModel(req model.Article) web.CatResp {
	return web.CatResp{
		Id:         req.Id,
		Name:       req.Name,
		Status:     req.Status,
		Visibility: req.Visibility,
		Details:    req.Details,
	}
}

package helper

import (
	"github.com/Alfeenn/article/model"
	"github.com/Alfeenn/article/model/web"
)

func ConvertModel(req model.Article) web.CatResp {

	categoryslice := GetCategoryObj(req)
	slugname := DashString(req.Name)

	return web.CatResp{
		Id:         req.Id,
		Name:       req.Name,
		Slug:       slugname,
		Category:   categoryslice,
		Status:     req.Status,
		Visibility: req.Visibility,
		Details:    req.Details,
	}
}

func GetCategoryObj(req model.Article) []model.CategoryArticle {

	var sliceCategory []model.CategoryArticle
	categoryArticle := model.CategoryArticle{
		Name: req.Category,
		Url:  req.Url,
	}

	catName := LowerAndDash(categoryArticle.Name)

	for v := range catName {
		categoryArticle.Name = catName[v]
		categoryArticle.Slug = catName[v]
		// url, _ := url.JoinPath("http://localhost:8000/api/categories", catName[v])
		// categoryArticle.Url = url

		sliceCategory = append(sliceCategory, categoryArticle)
	}
	return sliceCategory
}

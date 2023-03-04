package service

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/article/helper"
	"github.com/Alfeenn/article/model"
	"github.com/Alfeenn/article/model/web"
	"github.com/Alfeenn/article/repository"
)

type ServiceImpl struct {
	Rep repository.Repository
	DB  *sql.DB
}

func NewService(c repository.Repository, DB *sql.DB) Service {
	return &ServiceImpl{
		Rep: c,
		DB:  DB,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, req web.CatRequest) web.CatResp {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	request := model.Article{
		Id:         req.Id,
		Name:       req.Name,
		Status:     req.Status,
		Visibility: req.Visibility,
		Details:    req.Details,
	}

	article := s.Rep.Create(ctx, tx, request)

	return helper.ConvertModel(article)

}

func (s *ServiceImpl) Update(ctx context.Context, req web.UpdateRequest) web.CatResp {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	id := req.Id
	findId, err := s.Rep.Find(ctx, tx, id)
	helper.PanicIfErr(err)
	updateArticle := s.Rep.Update(ctx, tx, findId)
	return helper.ConvertModel(updateArticle)
}

func (s *ServiceImpl) Delete(ctx context.Context, id string) {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	req, err := s.Rep.Find(ctx, tx, id)
	helper.PanicIfErr(err)
	s.Rep.Delete(ctx, tx, req.Id)

}

func (s *ServiceImpl) Find(ctx context.Context, id string) web.CatResp {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	model, err := s.Rep.Find(ctx, tx, id)
	if err != nil {
		panic(err)
	}
	return helper.ConvertModel(model)

}

func (s *ServiceImpl) FindAll(ctx context.Context) []web.CatResp {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	slicemodel := s.Rep.FindAll(ctx, tx)

	var webResp []web.CatResp

	for _, v := range slicemodel {
		webResp = append(webResp, helper.ConvertModel(v))
	}
	return webResp
}

package web

import "github.com/Alfeenn/article/model"

type CatResp struct {
	Id         string                  `json:"id"`
	Name       string                  `json:"name"`
	Slug       string                  `json:"slug"`
	Status     string                  `json:"status"`
	Category   []model.CategoryArticle `json:"category"`
	Visibility string                  `json:"visibility"`
	Details    string                  `json:"details,omitempty"`
}

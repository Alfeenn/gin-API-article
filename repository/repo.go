package repository

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/article/model"
)

type Repository interface {
	Create(ctx context.Context, tx *sql.Tx, category model.Article) model.Article
	Update(ctx context.Context, tx *sql.Tx, category model.Article) model.Article
	Delete(ctx context.Context, tx *sql.Tx, id string)
	FindAll(ctx context.Context, tx *sql.Tx) []model.Article
	Find(ctx context.Context, tx *sql.Tx, id string) (model.Article, error)
}

package repository

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/article/helper"
	"github.com/Alfeenn/article/model"
	"github.com/google/uuid"
)

type RepoImpl struct{}

func NewRepository() *RepoImpl {
	return &RepoImpl{}
}

func (r *RepoImpl) Create(ctx context.Context, tx *sql.Tx, category model.Article) model.Article {
	SQL := "INSERT INTO article(id,name,category,status,visibility) VALUES(?,?,?,?,?)"
	category.Id = uuid.NewString()
	_, err := tx.ExecContext(ctx, SQL, category.Id, category.Name, category.Category, category.Status, category.Visibility)
	helper.PanicIfErr(err)
	return category
	// TODO: Implement
}

func (r *RepoImpl) Update(ctx context.Context, tx *sql.Tx, category model.Article) model.Article {
	SQL := "UPDATE article SET name=? WHERE id=?"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfErr(err)

	return category

}

func (r *RepoImpl) Delete(ctx context.Context, tx *sql.Tx, id string) {
	SQL := "DELETE FROM article WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfErr(err)
}

func (r *RepoImpl) FindAll(ctx context.Context, tx *sql.Tx, limit int, offset int) []model.Article {
	sql := "SELECT *FROM article LIMIT(?,?)"
	rows, err := tx.QueryContext(ctx, sql, limit, offset)
	helper.PanicIfErr(err)
	defer rows.Close()

	var sliceArticle []model.Article

	for rows.Next() {
		article := model.Article{}
		err := rows.Scan(&article.Id, &article.Name, &article.Category,
			&article.Status, &article.Visibility, &article.Visibility)
		helper.PanicIfErr(err)
		sliceArticle = append(sliceArticle, article)
	}
	return sliceArticle
}

func (r *RepoImpl) Find(ctx context.Context, tx *sql.Tx, id string) (model.Article, error) {
	SQL := "SELECT *FROM article WHERE id=?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer rows.Close()
	article := model.Article{}
	if rows.Next() {
		err := rows.Scan(&article.Id, &article.Name, &article.Category,
			&article.Status, &article.Visibility, &article.Visibility)
		helper.PanicIfErr(err)
		return article, nil
	} else {
		return article, err
	}

}

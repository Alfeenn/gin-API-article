package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Alfeenn/article/helper"
	"github.com/Alfeenn/article/model"
	"github.com/google/uuid"
)

type RepoImpl struct{}

func NewRepository() Repository {
	return &RepoImpl{}
}

func (r *RepoImpl) Create(ctx context.Context, tx *sql.Tx, category model.Article) model.Article {
	SQL := "INSERT INTO article(id,name,status,visibility,details) VALUES(?,?,?,?,?)"
	category.Id = uuid.NewString()
	_, err := tx.ExecContext(ctx, SQL, category.Id, category.Name, category.Status, category.Visibility, category.Details)
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

func (r *RepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.Article {
	sql := "SELECT a.id,a.name,a.status," +
		"GROUP_CONCAT(DISTINCT c.name ORDER BY c.name) AS category," +
		"GROUP_CONCAT(DISTINCT c.url ORDER BY c.url)AS URL,a.visibility FROM wordpress AS w " +
		"JOIN category_article AS c ON c.id = w.id_category " +
		"JOIN article AS a ON a.id= w.id_article " +
		"GROUP BY w.id_article"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfErr(err)
	defer rows.Close()

	var sliceArticle []model.Article

	for rows.Next() {
		article := model.Article{}
		err := rows.Scan(&article.Id, &article.Name,
			&article.Status, &article.Category, &article.Url, &article.Visibility)
		helper.PanicIfErr(err)
		sliceArticle = append(sliceArticle, article)
	}
	return sliceArticle
}

func (r *RepoImpl) Find(ctx context.Context, tx *sql.Tx, id string) (model.Article, error) {
	SQL := "SELECT *FROM article WHERE id =?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer rows.Close()
	article := model.Article{}
	if rows.Next() {
		rows.Scan(&article.Id, &article.Name,
			&article.Status, &article.Visibility, &article.Details)

		return article, nil
	} else {
		return article, errors.New("no category found")
	}

}

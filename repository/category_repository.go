package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ugik-dev/go-restful-api.git/helper"
	"github.com/ugik-dev/go-restful-api.git/models/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx, category domain.Category) []domain.Category
	Filter(ctx context.Context, tx *sql.Tx, category domain.Category) []domain.Category
}

// Category Repo Implementation

type CategoryRepositoryImpl struct {
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO category(name) values(?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category
}
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "Update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	query := "delete category where id = ?"
	_, err := tx.ExecContext(ctx, query, categoryId)
	helper.PanicIfError(err)
}
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	query := "select * from category where id = ?"
	result, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err)
	category := domain.Category{}
	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Data tidak ditemukan")
	}
}
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "select * from category"
	result, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	var categories []domain.Category
	for result.Next() {
		category := domain.Category{}
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
func (repository *CategoryRepositoryImpl) Filter(ctx context.Context, tx *sql.Tx, category domain.Category) []domain.Category {
	query := "select * from category"
	result, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	var categories []domain.Category
	for result.Next() {
		category := domain.Category{}
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}

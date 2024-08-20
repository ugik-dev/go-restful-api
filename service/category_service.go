package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/ugik-dev/go-restful-api.git/exception"
	"github.com/ugik-dev/go-restful-api.git/helper"
	"github.com/ugik-dev/go-restful-api.git/models/domain"
	"github.com/ugik-dev/go-restful-api.git/models/web"
	"github.com/ugik-dev/go-restful-api.git/repository"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context, filterKey string) []web.CategoryResponse
	Filter(ctx context.Context, filterKey string) []web.CategoryResponse
}

type CategoryServiceImp struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepo repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImp{
		CategoryRepository: categoryRepo,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImp) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ComitOrRollback(tx)
	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Create(ctx, tx, category)
	response := helper.ToCategoryResponse(category)
	return response
}

func (service *CategoryServiceImp) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ComitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	// helper.PanicIfError(err)

	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImp) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ComitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImp) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ComitOrRollback(tx)
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	response := helper.ToCategoryResponse(category)
	return response
}

func (service *CategoryServiceImp) FindAll(ctx context.Context, filterKey string) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ComitOrRollback(tx)
	categories := service.CategoryRepository.FindAll(ctx, tx, filterKey)
	return helper.ToCategoryResponses(categories)
}

func (service *CategoryServiceImp) Filter(ctx context.Context, filterKey string) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.ComitOrRollback(tx)

	categories := service.CategoryRepository.Filter(ctx, tx, filterKey)
	return helper.ToCategoryResponses(categories)
}

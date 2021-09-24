package service

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/exception"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
	"github.com/go-playground/validator"
)

type CategoryService struct {
	DB       *sql.DB
	Validate *validator.Validate
}

var (
	//Register Repo
	categoryRepo = repository.NewCategoryRepository()
)

func NewCategoryService(db *sql.DB, v *validator.Validate) CategoryServiceInterface {
	return &CategoryService{
		DB:       db,
		Validate: v,
	}
}

type CategoryServiceInterface interface {
	CreateCategory(ctx context.Context, category web.CategoryCreateRequest)
	GetCategory(ctx context.Context, categoryId int) web.CategoryResponse
	ListCategory(ctx context.Context, userId int) []web.CategoryResponse
	UpdateCategory(ctx context.Context, category web.CategoryUpdateRequest)
	DeleteCategory(ctx context.Context, categoryId int)
}

func (c *CategoryService) CreateCategory(ctx context.Context, categoryRequest web.CategoryCreateRequest) {
	err := c.Validate.Struct(categoryRequest)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()
	err = categoryRepo.Store(ctx, tx, helper.ToCategoryCreate(categoryRequest))
	helper.PanicIfError(err)
}
func (c *CategoryService) GetCategory(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	category, err := categoryRepo.GetByID(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (c *CategoryService) ListCategory(ctx context.Context, userId int) []web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	categories, err := categoryRepo.ListByUser(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponses(categories)
}

func (c *CategoryService) UpdateCategory(ctx context.Context, categoryRequest web.CategoryUpdateRequest) {
	err := c.Validate.Struct(categoryRequest)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	_, err = categoryRepo.GetByID(ctx, tx, int(categoryRequest.Id))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = categoryRepo.Update(ctx, tx, helper.ToCategoryUpdate(categoryRequest))
	helper.PanicIfError(err)
}

func (c *CategoryService) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	_, err = categoryRepo.GetByID(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = categoryRepo.Delete(ctx, tx, categoryId)
	helper.PanicIfError(err)
}

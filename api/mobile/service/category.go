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
	CategoryRepo repository.CategoryRepository
	DB           *sql.DB
	Validate     *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, v *validator.Validate) CategoryServiceInterface {
	return &CategoryService{
		CategoryRepo: categoryRepository,
		DB:           db,
		Validate:     v,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, categoryRequest web.CategoryCreateRequest) {
	err := c.Validate.Struct(categoryRequest)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()
	err = c.CategoryRepo.Store(ctx, tx, helper.ToCategoryCreate(categoryRequest))
	helper.PanicIfError(err)
}
func (c *CategoryService) GetCategory(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	category, err := c.CategoryRepo.GetByID(ctx, tx, categoryId)
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

	categories, err := c.CategoryRepo.ListByUser(ctx, tx, userId)
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

	_, err = c.CategoryRepo.GetByID(ctx, tx, int(categoryRequest.Id))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = c.CategoryRepo.Update(ctx, tx, helper.ToCategoryUpdate(categoryRequest))
	helper.PanicIfError(err)
}

func (c *CategoryService) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	_, err = c.CategoryRepo.GetByID(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = c.CategoryRepo.Delete(ctx, tx, categoryId)
	helper.PanicIfError(err)
}

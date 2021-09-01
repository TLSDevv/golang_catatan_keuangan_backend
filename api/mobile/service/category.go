package service

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

type CategoryService struct {
	CategoryRepo repository.CategoryRepository
	DB           *sql.DB
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB) CategoryServiceInterface {
	return &CategoryService{
		CategoryRepo: categoryRepository,
		DB:           db,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, categoryRequest web.CategoryCreateRequest) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()
	err = c.CategoryRepo.Store(ctx, tx, helper.ToCategory(categoryRequest))
	helper.PanicIfError(err)
}
func (c *CategoryService) GetCategory(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	category, err := c.CategoryRepo.GetByID(ctx, tx, categoryId)
	helper.PanicIfError(err)

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

func (c *CategoryService) UpdateCategory(ctx context.Context, categoryId int, categoryRequest web.CategoryCreateRequest) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()
	err = c.CategoryRepo.Update(ctx, tx, categoryId, helper.ToCategory(categoryRequest))
	helper.PanicIfError(err)
}
func (c *CategoryService) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()
	err = c.CategoryRepo.Delete(ctx, tx, categoryId)
	helper.PanicIfError(err)
}

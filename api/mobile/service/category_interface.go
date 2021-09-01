package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

type CategoryServiceInterface interface {
	CreateCategory(ctx context.Context, category web.CategoryCreateRequest)
	GetCategory(ctx context.Context, categoryId int) web.CategoryResponse
	ListCategory(ctx context.Context, userId int) []web.CategoryResponse
	UpdateCategory(ctx context.Context, categoryId int, category web.CategoryCreateRequest)
	DeleteCategory(ctx context.Context, categoryId int)
}

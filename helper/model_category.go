package helper

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

func ToCategoryCreate(category web.CategoryCreateRequest) domain.Category {
	return domain.Category{
		UserId:       category.UserId,
		NameCategory: category.NameCategory,
		Description:  category.Description,
		IconName:     category.IconName,
		IconColor:    category.IconColor,
	}
}

func ToCategoryUpdate(category web.CategoryUpdateRequest) domain.Category {
	return domain.Category{
		Id:           category.Id,
		UserId:       category.UserId,
		NameCategory: category.NameCategory,
		Description:  category.Description,
		IconName:     category.IconName,
		IconColor:    category.IconColor,
	}
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:           category.Id,
		UserId:       category.UserId,
		NameCategory: category.NameCategory,
		Description:  category.Description,
		IconName:     category.IconName,
		IconColor:    category.IconColor,
	}
}

func ToCategoryResponses(categories []domain.Category) (categoryResponses []web.CategoryResponse) {
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return
}

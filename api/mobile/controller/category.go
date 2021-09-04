package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/service"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
	"github.com/go-chi/chi"
)

type CategoryController struct {
	CategoryService service.CategoryServiceInterface
}

func NewCategoryController(categoryService service.CategoryServiceInterface) CategoryControllerInterface {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

func (c *CategoryController) CreateCategory(writer http.ResponseWriter, request *http.Request) {
	categoryRequest := web.CategoryCreateRequest{}

	helper.ReadFromRequestBody(request, &categoryRequest)

	c.CategoryService.CreateCategory(request.Context(), categoryRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "SUCCESS CREATE CATEGORY",
	}

	writer.WriteHeader(http.StatusAccepted)

	helper.WriterToResponseBody(writer, webResponse)
}
func (c *CategoryController) GetCategory(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		err := errors.New("Params Is Empty")
		helper.PanicIfError(err)
	}

	categoryId, _ := strconv.Atoi(id)

	category := c.CategoryService.GetCategory(request.Context(), categoryId)

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "SUCCESS GET DATA CATEGORY",
		Data:   category,
	}
	writer.WriteHeader(http.StatusAccepted)

	helper.WriterToResponseBody(writer, webResponse)
}
func (c *CategoryController) ListCategory(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		err := errors.New("Params Is Empty")
		helper.PanicIfError(err)
	}

	userId, _ := strconv.Atoi(id)

	categories := c.CategoryService.ListCategory(request.Context(), userId)

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "SUCCESS UPDATE CATEGORY",
		Data:   categories,
	}
	writer.WriteHeader(http.StatusAccepted)

	helper.WriterToResponseBody(writer, webResponse)
}
func (c *CategoryController) UpdateCategory(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		err := errors.New("Params Is Empty")
		helper.PanicIfError(err)
	}

	categoryId, _ := strconv.Atoi(id)

	categoryRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryRequest)
	categoryRequest.Id = uint8(categoryId)

	c.CategoryService.UpdateCategory(request.Context(), categoryRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "SUCCESS UPDATE CATEGORY",
	}
	writer.WriteHeader(http.StatusAccepted)

	helper.WriterToResponseBody(writer, webResponse)
}
func (c *CategoryController) DeleteCategory(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	if id == "" {
		err := errors.New("Params Is Empty")
		helper.PanicIfError(err)
	}

	categoryId, _ := strconv.Atoi(id)

	c.CategoryService.DeleteCategory(request.Context(), categoryId)

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "SUCCESS DELETE CATEGORY",
	}
	writer.WriteHeader(http.StatusAccepted)

	helper.WriterToResponseBody(writer, webResponse)
}

package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/ugik-dev/go-restful-api.git/helper"
	"github.com/ugik-dev/go-restful-api.git/models/web"
	"github.com/ugik-dev/go-restful-api.git/service"
)

type CategoryControllerImpl struct {
	CatergoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CatergoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Crate(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadRequestBody(request, &categoryCreateRequest)
	categoryResponse := controller.CatergoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteResponse(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadRequestBody(request, &categoryUpdateRequest)

	categoryResponse := controller.CatergoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteResponse(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CatergoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteResponse(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryResponse := controller.CatergoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteResponse(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	key := request.URL.Query().Get("key")
	categoryResponses := controller.CatergoryService.FindAll(request.Context(), key)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	helper.WriteResponse(w, webResponse)
}

func (controller *CategoryControllerImpl) Filter(w http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// key := request.URL.Query().Get("key")
	// fmt.Println(key)
	// categoryResponses := controller.CatergoryService.FindAll(request.Context())
	// webResponse := web.WebResponse{
	// 	Code:   200,
	// 	Status: "OK",
	// 	Data:   categoryResponses,
	// }
	// helper.WriteResponse(w, webResponse)
}

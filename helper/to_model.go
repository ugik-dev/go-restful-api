package helper

import (
	"github.com/ugik-dev/go-restful-api.git/models/domain"
	"github.com/ugik-dev/go-restful-api.git/models/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponse []web.CategoryResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryResponse(category))
	}
	return categoryResponse
}

package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/ugik-dev/go-restful-api.git/controller"
	"github.com/ugik-dev/go-restful-api.git/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = exception.PanicHandler
	// router.NotFound = http.HandlerFunc(utils.NotFound)

	// router.GET("/api/categories/s/", categoryController.Filter)
	router.GET("/api/categories", categoryController.FindAll)
	router.POST("/api/categories", categoryController.Crate)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	return router
}

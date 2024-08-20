package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ugik-dev/go-restful-api.git/app"
	"github.com/ugik-dev/go-restful-api.git/controller"
	"github.com/ugik-dev/go-restful-api.git/middleware"
	"github.com/ugik-dev/go-restful-api.git/repository"
	"github.com/ugik-dev/go-restful-api.git/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	middleware := middleware.NewAuth(router)

	app.StartServer(middleware)
}

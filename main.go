// main.go
package main

import (
	"adiubaidah/simple-crud/app"
	"adiubaidah/simple-crud/controller"
	"adiubaidah/simple-crud/helper"
	"adiubaidah/simple-crud/repository"
	"adiubaidah/simple-crud/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.SetupRouter(&app.CategoryController{CategoryController: categoryController})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

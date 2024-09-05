package app

import (
	"adiubaidah/simple-crud/controller"
	"adiubaidah/simple-crud/exception"

	"github.com/julienschmidt/httprouter"
)

type CategoryController struct {
	controller.CategoryController //name of the interface and the struct is the same
}

func (c *CategoryController) RegisterRoutes(router *httprouter.Router) {
	router.GET("/api/categories", c.FindAll)
	router.GET("/api/categories/:categoryId", c.FindById)
	router.POST("/api/categories", c.Create)
	router.PUT("/api/categories/:categoryId", c.Update)
	router.DELETE("/api/categories/:categoryId", c.Delete)
}

func SetupRouter(controllers ...controller.Controller) *httprouter.Router {
	router := httprouter.New()

	for _, ctrl := range controllers {
		ctrl.RegisterRoutes(router)
	}

	router.PanicHandler = exception.ErrorHandler

	return router
}

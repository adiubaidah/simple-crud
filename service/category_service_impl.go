package service

import (
	"adiubaidah/simple-crud/exception"
	"adiubaidah/simple-crud/helper"
	"adiubaidah/simple-crud/model/category"
	"adiubaidah/simple-crud/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate //validate memiliki fungsi Struct, dan memvalidasi berdasarkan tag yang didefinisikan pada struct
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request category.CategoryCreateRequest) category.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	result := category.Category{
		Name: request.Name,
	}

	result = service.CategoryRepository.Save(ctx, tx, result)

	return category.ToCategoryResponse(result)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request category.CategoryUpdateRequest) category.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	result, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	result.Name = request.Name

	result = service.CategoryRepository.Update(ctx, tx, result)

	return category.ToCategoryResponse(result)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) category.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	result, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return category.ToCategoryResponse(result)
}
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []category.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	results := service.CategoryRepository.FindAll(ctx, tx)

	return category.ToCategoryResponses(results)
}

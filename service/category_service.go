package service

import (
	"adiubaidah/simple-crud/model/category"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request category.CategoryCreateRequest) category.CategoryResponse
	Update(ctx context.Context, request category.CategoryUpdateRequest) category.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) category.CategoryResponse
	FindAll(ctx context.Context) []category.CategoryResponse
}

package repository

import (
	"adiubaidah/simple-crud/model/category"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category category.Category) category.Category
	Update(ctx context.Context, tx *sql.Tx, category category.Category) category.Category
	Delete(ctx context.Context, tx *sql.Tx, category category.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (category.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []category.Category
}

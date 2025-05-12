package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xoesae/cid-api/internal/domain/entity"
	"github.com/xoesae/cid-api/pkg/fault"
	"github.com/xoesae/cid-api/pkg/logger"
)

type CategoryRepository interface {
	Insert(ctx context.Context, category entity.Category) (int, error)
	GetByCode(code string) (entity.Category, error)
}

type categoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{DB: db}
}

func (r *categoryRepository) Insert(ctx context.Context, category entity.Category) (int, error) {
	query := "INSERT INTO categories (group_id, code, name) VALUES ($1, $2, $3) RETURNING id;"

	var id int
	err := r.DB.QueryRowContext(ctx, query, category.GroupID, category.Code, category.Name).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("could not insert category: %w", err)
	}

	return id, nil
}

func (r *categoryRepository) GetByCode(code string) (entity.Category, error) {
	var category entity.Category

	query := `
		SELECT 
			id,
			group_id,
			code,
			name 
		FROM 
		    categories 
		WHERE
		    code = $1
		LIMIT 1
	`

	err := r.DB.Get(&category, query, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Get().Debug("cid not found")

			return category, fault.NewNotFoundFault("could not find cid")
		}

		logger.Get().Error(err.Error())
		return category, fmt.Errorf("could not get paginated subcategories: %w", err)
	}

	return category, nil
}

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

type SubcategoryRepository interface {
	Insert(ctx context.Context, subcategory entity.Subcategory) (int, error)
	GetPaginated(limit, offset int, search string) ([]entity.Subcategory, error)
	GetByCode(code string) (entity.Subcategory, error)
}

type subcategoryRepository struct {
	DB *sqlx.DB
}

func NewSubcategoryRepository(db *sqlx.DB) SubcategoryRepository {
	return &subcategoryRepository{DB: db}
}

func (r *subcategoryRepository) Insert(ctx context.Context, subcategory entity.Subcategory) (int, error) {
	query := "INSERT INTO subcategories (category_id, code, name) VALUES ($1, $2, $3) RETURNING id;"

	var id int
	err := r.DB.QueryRowContext(ctx, query, subcategory.CategoryID, subcategory.Code, subcategory.Name).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("could not insert subcategory: %w", err)
	}

	return id, nil
}

func (r *subcategoryRepository) GetPaginated(limit, offset int, search string) ([]entity.Subcategory, error) {
	subcategories := make([]entity.Subcategory, 0)

	query := `
		SELECT 
			id,
			category_id,
			code,
			name 
		FROM 
		    subcategories 
		WHERE
		    code ILIKE $1
			OR name ILIKE $1
		ORDER BY id 
		LIMIT $2
		OFFSET $3
	`

	if search != "" {
		search = "%" + search + "%"
	}

	err := r.DB.Select(&subcategories, query, search, limit, offset)
	if err != nil {
		logger.Get().Error(err.Error())
		return nil, fmt.Errorf("could not get paginated subcategories: %w", err)
	}

	return subcategories, nil
}

func (r *subcategoryRepository) GetByCode(code string) (entity.Subcategory, error) {
	var subcategory entity.Subcategory

	query := `
		SELECT 
			id,
			category_id,
			code,
			name 
		FROM 
		    subcategories 
		WHERE
		    code = $1
		LIMIT 1
	`

	err := r.DB.Get(&subcategory, query, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Get().Debug("cid not found")

			return subcategory, fault.NewNotFoundFault("could not find cid")
		}

		logger.Get().Error(err.Error(), "err")
		return subcategory, fmt.Errorf("could not get by code: %w", err)
	}

	return subcategory, nil
}

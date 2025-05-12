package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xoesae/cid-api/internal/domain/entity"
)

type GroupRepository interface {
	Insert(ctx context.Context, group entity.Group) (int, error)
}

type groupRepository struct {
	DB *sqlx.DB
}

func NewGroupRepository(db *sqlx.DB) GroupRepository {
	return &groupRepository{DB: db}
}

func (r *groupRepository) Insert(ctx context.Context, group entity.Group) (int, error) {
	query := "INSERT INTO groups (chapter_id, code_start, code_end, name) VALUES ($1, $2, $3, $4) RETURNING id;"

	var id int
	err := r.DB.QueryRowContext(ctx, query, group.ChapterID, group.CodeStart, group.CodeEnd, group.Name).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("could not insert group: %w", err)
	}

	return id, nil
}

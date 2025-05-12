package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xoesae/cid-api/internal/domain/entity"
)

type ChapterRepository interface {
	Insert(ctx context.Context, chapter entity.Chapter) (int, error)
}

type chapterRepository struct {
	DB *sqlx.DB
}

func NewChapterRepository(db *sqlx.DB) ChapterRepository {
	return &chapterRepository{DB: db}
}

func (r *chapterRepository) Insert(ctx context.Context, chapter entity.Chapter) (int, error) {
	query := "INSERT INTO chapters (code_start, code_end, roman, name) VALUES ($1, $2, $3, $4) RETURNING id;"

	var id int
	err := r.DB.QueryRowContext(ctx, query, chapter.CodeStart, chapter.CodeEnd, chapter.Roman, chapter.Name).Scan(&id)
	if err != nil {
		return id, fmt.Errorf("could not insert chapter: %w", err)
	}

	return id, nil
}

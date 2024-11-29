package repository

import (
	"context"
	"errors"
	"github.com/GeorgiyGusev/hack-backend/internal/news/entity"
	"github.com/GeorgiyGusev/hack-backend/internal/news/storage/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, news *entity.News) error {
	query := `
		INSERT INTO news (id, organization_id, title, description, media, created_at, updated_at)
		VALUES (:id, :organization_id, :title, :description, :media, :created_at, :updated_at)
	`
	newsDB := models.MapFromDomain(news)
	_, err := r.db.NamedExecContext(ctx, query, newsDB)
	return err
}

func (r *Repository) GetAllNews(ctx context.Context) (*[]entity.News, error) {
	query := `
		SELECT id, organization_id, title, description, media, created_at, updated_at
		FROM news
		ORDER BY created_at DESC
	`
	var newsDBs []models.NewsDB
	err := r.db.SelectContext(ctx, &newsDBs, query)
	if err != nil {
		return nil, err
	}

	var newsList []entity.News
	for _, newsDB := range newsDBs {
		newsList = append(newsList, *newsDB.MapToDomain())
	}
	return &newsList, nil
}

func (r *Repository) GetByID(ctx context.Context, id string) (*entity.News, error) {
	query := `
		SELECT id, organization_id, title, description, media, created_at, updated_at
		FROM news
		WHERE id = $1
	`
	newsDB := &models.NewsDB{}
	err := r.db.GetContext(ctx, newsDB, query, id)
	if err != nil {
		return nil, err
	}
	return newsDB.MapToDomain(), nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM news WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("news not found")
	}
	return nil
}

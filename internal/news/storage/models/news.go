package models

import (
	"github.com/GeorgiyGusev/hack-backend/internal/news/entity"
	"time"
)

type NewsDB struct {
	Id             string    `db:"id"`
	OrganizationId string    `db:"organization_id"`
	Title          string    `db:"title"`
	Description    string    `db:"description"`
	Media          []string  `db:"media"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func (n *NewsDB) MapToDomain() *entity.News {
	return &entity.News{
		Id:             n.Id,
		Title:          n.Title,
		Description:    n.Description,
		OrganizationId: n.OrganizationId,
		Media:          n.Media,
		CreatedAt:      n.CreatedAt,
	}
}

func MapFromDomain(news *entity.News) *NewsDB {
	return &NewsDB{
		Id:             news.Id,
		Title:          news.Title,
		Description:    news.Description,
		OrganizationId: news.OrganizationId,
		Media:          news.Media,
		CreatedAt:      news.CreatedAt,
		UpdatedAt:      time.Now(),
	}
}

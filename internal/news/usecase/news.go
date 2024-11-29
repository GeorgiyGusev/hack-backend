package usecase

import (
	"context"
	"github.com/GeorgiyGusev/hack-backend/internal/news/entity"
	"github.com/GeorgiyGusev/hack-backend/internal/news/gen"
	"log/slog"
)

type Usecase interface {
	Create(ctx context.Context, news *gen.CreateNews) error
	GetAllNews(ctx context.Context) (*[]entity.News, error)
	GetByID(ctx context.Context, id string) (*entity.News, error)
}

type Repository interface {
	Create(ctx context.Context, news *entity.News) error
	GetAllNews(ctx context.Context) (*[]entity.News, error)
	GetByID(ctx context.Context, id string) (*entity.News, error)
}

type Impl struct {
	repo   Repository
	logger *slog.Logger
}

func NewImpl(repo Repository, logger *slog.Logger) *Impl {
	return &Impl{repo: repo, logger: logger}
}

func (i *Impl) Create(ctx context.Context, dto *gen.CreateNews) error {
	news := entity.NewNews(&entity.NewsInput{
		Title:          dto.Title,
		Description:    dto.Description,
		OrganizationId: dto.OrganizationId,
		Media:          dto.Media,
	})
	return i.repo.Create(ctx, news)
}

func (i *Impl) GetAllNews(ctx context.Context) (*[]entity.News, error) {
	return i.repo.GetAllNews(ctx)
}

func (i *Impl) GetByID(ctx context.Context, id string) (*entity.News, error) {
	return i.repo.GetByID(ctx, id)
}

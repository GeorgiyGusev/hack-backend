package entity

import (
	"github.com/google/uuid"
	"time"
)

type News struct {
	Id             string
	Title          string
	Description    string
	CreatedAt      time.Time
	OrganizationId string
	Media          []string
}

type NewsInput struct {
	Title          string
	Description    string
	OrganizationId string
	Media          []string
}

func NewNews(inpt *NewsInput) *News {
	return &News{
		Id:             uuid.NewString(),
		CreatedAt:      time.Now(),
		Title:          inpt.Title,
		Description:    inpt.Description,
		OrganizationId: inpt.OrganizationId,
		Media:          inpt.Media,
	}
}

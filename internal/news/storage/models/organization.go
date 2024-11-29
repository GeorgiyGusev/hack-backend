package models

import (
	"github.com/GeorgiyGusev/hack-backend/internal/organizations/entity"
)

type OrganizationDB struct {
	Id          string  `db:"id"`
	OwnerId     string  `db:"owner_id"`
	PhotoId     *string `db:"photo_id"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
	Phone       string  `db:"phone"`
	Email       string  `db:"email"`
	Status      string  `db:"status"`
	Longtitude  float32 `db:"longtitude"`
	Latitude    float32 `db:"latitude"`
}

func (o *OrganizationDB) MapToDomain() *entity.Organization {
	return &entity.Organization{
		Id:          o.Id,
		OwnerId:     o.OwnerId,
		PhotoId:     o.PhotoId,
		Title:       o.Title,
		Description: o.Description,
		Phone:       o.Phone,
		Email:       o.Email,
		Status:      o.Status,
		Longtitude:  o.Longtitude,
		Latitude:    o.Latitude,
	}
}

func MapFromDomain(o *entity.Organization) *OrganizationDB {
	return &OrganizationDB{
		Id:          o.Id,
		OwnerId:     o.OwnerId,
		PhotoId:     o.PhotoId,
		Title:       o.Title,
		Description: o.Description,
		Phone:       o.Phone,
		Email:       o.Email,
		Status:      o.Status,
		Longtitude:  o.Longtitude,
		Latitude:    o.Latitude,
	}
}

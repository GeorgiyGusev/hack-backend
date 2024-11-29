package entity

import "github.com/google/uuid"

const (
	PendingStatus  = "pending"
	ApprovedStatus = "approved"
	RejectedStatus = "rejected"
)

type Organization struct {
	Id          string
	OwnerId     string
	PhotoId     *string
	Title       string
	Description string
	Phone       string
	Email       string
	Status      string
	Longtitude  float32
	Latitude    float32
}

type CreateOrganizationInput struct {
	OwnerId     string
	PhotoId     *string
	Title       string
	Description string
	Phone       string
	Email       string
	Status      string
	Longtitude  float32
	Latitude    float32
}

func NewOrganization(inpt *CreateOrganizationInput) *Organization {
	return &Organization{
		Id:          uuid.NewString(),
		OwnerId:     inpt.OwnerId,
		PhotoId:     inpt.PhotoId,
		Title:       inpt.Title,
		Description: inpt.Description,
		Phone:       inpt.Phone,
		Email:       inpt.Email,
		Status:      inpt.Status,
		Longtitude:  inpt.Longtitude,
		Latitude:    inpt.Latitude,
	}
}

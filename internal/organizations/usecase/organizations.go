package usecase

import (
	"context"
	"errors"
	"github.com/GeorgiyGusev/hack-backend/internal/organizations/entity"
	"github.com/GeorgiyGusev/hack-backend/internal/organizations/gen"
	"log/slog"
)

type Usecase interface {
	GetAllOrganizations(ctx context.Context) (*[]entity.Organization, error)
	GetAllOrganizationsForUser(ctx context.Context, userId string) (*[]entity.Organization, error)
	CreateOrganization(ctx context.Context, dto *gen.OrganizationCreate, userId string) error
	UpdateOrganization(ctx context.Context, dto *gen.OrganizationUpdate, id string, userId string) error
	DeleteOrganization(ctx context.Context, id string, userId string) error
	GetOrganizationById(ctx context.Context, id string) (*entity.Organization, error)
	GetOrganizationsByStatus(ctx context.Context, dto *gen.GetOrganizationsByStatus) ([]entity.Organization, error)
	ApproveOrganization(ctx context.Context, dto *gen.ModerationApprove) error
	RejectOrganization(ctx context.Context, dto *gen.ModerationReject) error
}

type Repository interface {
	GetAllOrganizations(ctx context.Context) (*[]entity.Organization, error)
	GetAllOrganizationsByOwnerId(ctx context.Context, userId string) (*[]entity.Organization, error)
	GetOrganizationById(ctx context.Context, id string) (*entity.Organization, error)
	GetOrganizationsByStatus(ctx context.Context, status string) ([]entity.Organization, error)
	CreateOrganization(ctx context.Context, organization *entity.Organization) error
	UpdateOrganization(ctx context.Context, organization *entity.Organization) error
	DeleteOrganization(ctx context.Context, id string) error
}

type Impl struct {
	repo   Repository
	logger *slog.Logger
}

func NewImpl(repo Repository, logger *slog.Logger) *Impl {
	return &Impl{repo: repo, logger: logger}
}

func (i *Impl) GetAllOrganizationsForUser(ctx context.Context, id string) (*[]entity.Organization, error) {
	return i.repo.GetAllOrganizationsByOwnerId(ctx, id)
}

func (i *Impl) GetAllOrganizations(ctx context.Context) (*[]entity.Organization, error) {
	return i.repo.GetAllOrganizations(ctx)
}

func (i *Impl) CreateOrganization(ctx context.Context, dto *gen.OrganizationCreate, userId string) error {
	organization := entity.NewOrganization(&entity.CreateOrganizationInput{
		OwnerId:     userId,
		PhotoId:     *dto.PhotoId,
		Title:       dto.Title,
		Description: dto.Description,
		Email:       dto.Email,
		Phone:       dto.Phone,
		Status:      string(gen.Pending),
		Longtitude:  dto.Longitude,
		Latitude:    dto.Latitude,
	})
	return i.repo.CreateOrganization(ctx, organization)
}

func (i *Impl) UpdateOrganization(ctx context.Context, dto *gen.OrganizationUpdate, id string, userId string) error {
	organization, err := i.repo.GetOrganizationById(ctx, id)
	if err != nil {
		i.logger.Info("cannot find organization to update", "orgId", id)
		return err
	}

	if organization.OwnerId != userId {
		return errors.New("not authorized to update")
	}

	organization.Title = dto.Title
	organization.Description = dto.Description
	organization.Phone = dto.Phone
	organization.Email = string(dto.Email)
	organization.PhotoId = *dto.PhotoId
	organization.Longtitude = dto.Longitude
	organization.Latitude = dto.Latitude
	err = i.repo.UpdateOrganization(ctx, organization)
	if err != nil {
		i.logger.Info("cannot update organization")
		return err
	}
	return nil
}

func (i *Impl) DeleteOrganization(ctx context.Context, id string, userId string) error {
	organization, err := i.repo.GetOrganizationById(ctx, id)
	if err != nil {
		return err
	}
	if organization.OwnerId != userId {
		i.logger.Error("Not authorized to delete", "userId", userId, "ownerId", organization.OwnerId)
		return errors.New("not authorized to delete organization")
	}

	return i.repo.DeleteOrganization(ctx, id)
}

func (i *Impl) GetOrganizationById(ctx context.Context, id string) (*entity.Organization, error) {
	return i.repo.GetOrganizationById(ctx, id)
}

func (i *Impl) GetOrganizationsByStatus(ctx context.Context, dto *gen.GetOrganizationsByStatus) ([]entity.Organization, error) {
	return i.repo.GetOrganizationsByStatus(ctx, dto.Status)
}

func (i *Impl) ApproveOrganization(ctx context.Context, dto *gen.ModerationApprove) error {
	organization, err := i.repo.GetOrganizationById(ctx, dto.Id)
	if err != nil {
		return err
	}
	organization.Status = string(gen.Approved)
	err = i.repo.UpdateOrganization(ctx, organization)
	if err != nil {
		return err
	}
	return nil
}

func (i *Impl) RejectOrganization(ctx context.Context, dto *gen.ModerationReject) error {
	organization, err := i.repo.GetOrganizationById(ctx, dto.Id)
	if err != nil {
		return err
	}
	organization.Status = string(gen.Rejected)
	err = i.repo.UpdateOrganization(ctx, organization)
	if err != nil {
		return err
	}
	return nil
}

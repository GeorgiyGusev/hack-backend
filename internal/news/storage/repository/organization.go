package repository

import (
	"context"
	"errors"
	"github.com/GeorgiyGusev/hack-backend/internal/organizations/storage/models"

	"github.com/GeorgiyGusev/hack-backend/internal/organizations/entity"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllOrganizationsByOwnerId(ctx context.Context, ownerId string) (*[]entity.Organization, error) {
	query := `
		SELECT id, owner_id, photo_id, title, description, phone, email, status, longtitude, latitude
		FROM organizations
		WHERE owner_id = $1
	`
	var orgDBs []models.OrganizationDB
	err := r.db.SelectContext(ctx, &orgDBs, query, ownerId)
	if err != nil {
		return nil, err
	}

	// Маппинг в доменные сущности
	var organizations []entity.Organization
	for _, orgDB := range orgDBs {
		organizations = append(organizations, *orgDB.MapToDomain())
	}
	return &organizations, nil
}

func (r *Repository) GetAllOrganizations(ctx context.Context) (*[]entity.Organization, error) {
	query := `
		SELECT id, owner_id,photo_id, title, description, phone, email, status, longtitude, latitude
		FROM organizations
	`
	var orgDBs []models.OrganizationDB
	err := r.db.SelectContext(ctx, &orgDBs, query)
	if err != nil {
		return nil, err
	}

	// Маппинг в доменные сущности
	var organizations []entity.Organization
	for _, orgDB := range orgDBs {
		organizations = append(organizations, *orgDB.MapToDomain())
	}
	return &organizations, nil
}

func (r *Repository) GetOrganizationById(ctx context.Context, id string) (*entity.Organization, error) {
	query := `
		SELECT id,owner_id, photo_id, title, description, phone, email, status, longtitude, latitude
		FROM organizations
		WHERE id = $1
	`
	orgDB := &models.OrganizationDB{}
	err := r.db.GetContext(ctx, orgDB, query, id)
	if err != nil {
		return nil, err
	}
	return orgDB.MapToDomain(), nil
}

func (r *Repository) GetOrganizationsByStatus(ctx context.Context, status string) ([]entity.Organization, error) {
	query := `
		SELECT id, owner_id, photo_id, title, description, phone, email, status, longtitude, latitude
		FROM organizations
		WHERE status = $1
	`
	var orgDBs []models.OrganizationDB
	err := r.db.SelectContext(ctx, &orgDBs, query, status)
	if err != nil {
		return nil, err
	}

	// Маппинг в доменные сущности
	var organizations []entity.Organization
	for _, orgDB := range orgDBs {
		organizations = append(organizations, *orgDB.MapToDomain())
	}
	return organizations, nil
}

func (r *Repository) CreateOrganization(ctx context.Context, organization *entity.Organization) error {
	query := `
		INSERT INTO organizations (id, owner_id, photo_id, title, description, phone, email, status, longtitude, latitude)
		VALUES (:id, :owner_id, :photo_id, :title, :description, :phone, :email, :status, :longtitude, :latitude)
	`
	orgDB := models.MapFromDomain(organization)
	_, err := r.db.NamedExecContext(ctx, query, orgDB)
	return err
}

func (r *Repository) UpdateOrganization(ctx context.Context, organization *entity.Organization) error {
	query := `
		UPDATE organizations
		SET photo_id = :photo_id, title = :title, description = :description,
		    phone = :phone, email = :email, status = :status,
		    longtitude = :longtitude, latitude = :latitude
		WHERE id = :id
	`
	orgDB := models.MapFromDomain(organization)
	_, err := r.db.NamedExecContext(ctx, query, orgDB)
	return err
}

func (r *Repository) DeleteOrganization(ctx context.Context, id string) error {
	query := `DELETE FROM organizations WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("organization not found")
	}
	return nil
}

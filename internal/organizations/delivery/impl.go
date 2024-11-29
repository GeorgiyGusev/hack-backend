package delivery

import (
	"context"
	"github.com/GeorgiyGusev/auth-library/models"
	"github.com/GeorgiyGusev/auth-library/provider"
	"github.com/GeorgiyGusev/hack-backend/internal/organizations/gen"
	"github.com/GeorgiyGusev/hack-backend/internal/organizations/usecase"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

//go:generate oapi-codegen --config=../gen/gen.cfg.yaml https://raw.githubusercontent.com/GeorgiyGusev/hack-api/refs/heads/main/service-api.yaml
type Handlers struct {
	logger  *slog.Logger
	usecase usecase.Usecase
}

func Register(logger *slog.Logger, usecase usecase.Usecase, srv *echo.Echo, authProvider provider.AuthProvider) {
	impl := &Handlers{logger: logger, usecase: usecase}
	gen.RegisterHandlers(srv, impl)

	authProvider.AddEndpointSecurity("/organizations", "user")
	authProvider.AddEndpointSecurity("/organizations/user", "user")
	authProvider.AddEndpointSecurity("/organizations/moderation", "moderator")
}

func (h *Handlers) GetAllOrganizationsForUser(ctx echo.Context) error {
	userDetails := ctx.Get(provider.UserDetailsKey).(models.UserDetails)
	orgs, err := h.usecase.GetAllOrganizationsForUser(context.Background(), userDetails.UserId)
	if err != nil {
		h.logger.Error("Cannot get organizations", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	resp := []gen.Organization{}
	for _, o := range *orgs {
		resp = append(resp, gen.Organization{
			Id:          o.Id,
			Title:       o.Title,
			Description: o.Description,
			Email:       o.Email,
			Phone:       o.Phone,
			PhotoId:     &o.PhotoId,
			Status:      gen.OrganizationStatus(o.Status),
		})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handlers) GetAllOrganizations(ctx echo.Context) error {
	orgs, err := h.usecase.GetAllOrganizations(context.Background())
	if err != nil {
		h.logger.Error("Cannot get organizations", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	resp := []gen.Organization{}
	for _, o := range *orgs {
		resp = append(resp, gen.Organization{
			Id:          o.Id,
			Title:       o.Title,
			Description: o.Description,
			Email:       o.Email,
			Phone:       o.Phone,
			PhotoId:     &o.PhotoId,
			Status:      gen.OrganizationStatus(o.Status),
		})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handlers) AddNewOrganization(ctx echo.Context) error {
	var req gen.OrganizationCreate
	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("Cannot bind adding new organization", "error", err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	userDetails := ctx.Get(provider.UserDetailsKey).(models.UserDetails)

	err := h.usecase.CreateOrganization(context.Background(), &req, userDetails.UserId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusCreated)
}

func (h *Handlers) ApproveOrganization(ctx echo.Context) error {
	var req gen.ModerationApprove
	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("Cannot bind adding new organization", "error", err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	err := h.usecase.ApproveOrganization(context.Background(), &req)
	if err != nil {
		h.logger.Error("Cannot approve organization", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusOK)
}

func (h *Handlers) RejectOrganization(ctx echo.Context) error {
	var req gen.ModerationReject
	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("Cannot bind adding new organization", "error", err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	err := h.usecase.RejectOrganization(context.Background(), &req)
	if err != nil {
		h.logger.Error("Cannot approve organization", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusOK)
}

func (h *Handlers) GetOrganizationByStatus(ctx echo.Context) error {
	var req gen.GetOrganizationsByStatus
	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("Cannot bind adding new organization", "error", err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	orgs, err := h.usecase.GetOrganizationsByStatus(context.Background(), &req)
	if err != nil {
		h.logger.Error("Cannot get organizations", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	resp := []gen.Organization{}
	for _, v := range orgs {

		resp = append(resp, gen.Organization{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Email:       v.Email,
			Phone:       v.Phone,
			PhotoId:     &v.PhotoId,
			Status:      gen.OrganizationStatus(v.Status),
			Longitude:   v.Longtitude,
			Latitude:    v.Latitude,
		})
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (h *Handlers) DeleteOrganization(ctx echo.Context, id string) error {
	userDetails := ctx.Get(provider.UserDetailsKey).(models.UserDetails)

	err := h.usecase.DeleteOrganization(context.Background(), id, userDetails.UserId)
	if err != nil {
		h.logger.Error("Cannot delete organization", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusOK)
}

func (h *Handlers) GetOrganizationById(ctx echo.Context, id string) error {
	org, err := h.usecase.GetOrganizationById(context.Background(), id)
	if err != nil {
		h.logger.Error("Cannot get organization", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, gen.Organization{
		Id:          org.Id,
		Title:       org.Title,
		Description: org.Description,
		Email:       org.Email,
		Phone:       org.Phone,
		PhotoId:     &org.PhotoId,
		Status:      gen.OrganizationStatus(org.Status),
		Longitude:   org.Longtitude,
		Latitude:    org.Latitude,
	})
}

func (h *Handlers) UpdateOrganization(ctx echo.Context, id string) error {
	var req gen.OrganizationUpdate
	if err := ctx.Bind(&req); err != nil {
		h.logger.Error("Cannot bind adding new organization", "error", err.Error())
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	userDetails := ctx.Get(provider.UserDetailsKey).(models.UserDetails)
	err := h.usecase.UpdateOrganization(context.Background(), &req, id, userDetails.UserId)
	if err != nil {
		h.logger.Error("Cannot update organization", "error", err.Error())
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusOK)
}

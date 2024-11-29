// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package gen

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for OrganizationStatus.
const (
	Approved OrganizationStatus = "approved"
	Pending  OrganizationStatus = "pending"
	Rejected OrganizationStatus = "rejected"
)

// GetOrganizationsByStatus defines model for GetOrganizationsByStatus.
type GetOrganizationsByStatus struct {
	Status string `json:"status"`
}

// ModerationApprove defines model for ModerationApprove.
type ModerationApprove struct {
	// Id ID организации
	Id string `json:"id"`
}

// ModerationReject defines model for ModerationReject.
type ModerationReject struct {
	// Id ID организации
	Id string `json:"id"`

	// Reason Причина отклонения
	Reason string `json:"reason"`
}

// Organization defines model for Organization.
type Organization struct {
	// Description Подробное описание организации
	Description string `json:"description"`

	// Email Электронная почта
	Email openapi_types.Email `json:"email"`

	// Id Уникальный идентификатор организации
	Id openapi_types.UUID `json:"id"`

	// Latitude Географическая широта
	Latitude float32 `json:"latitude"`

	// Longitude Географическая долгота
	Longitude float32 `json:"longitude"`

	// OwnerId Уникальный идентификатор организации
	OwnerId *openapi_types.UUID `json:"owner_id,omitempty"`

	// Phone Контактный телефон
	Phone string `json:"phone"`

	// PhotoId ID фотографии
	PhotoId *openapi_types.UUID `json:"photo_id"`

	// Status Статус организации
	Status OrganizationStatus `json:"status"`

	// Title Название организации
	Title string `json:"title"`
}

// OrganizationStatus Статус организации
type OrganizationStatus string

// OrganizationCreate defines model for OrganizationCreate.
type OrganizationCreate struct {
	// Description Подробное описание организации
	Description string `json:"description"`

	// Email Электронная почта
	Email openapi_types.Email `json:"email"`

	// Latitude Географическая широта
	Latitude float32 `json:"latitude"`

	// Longitude Географическая долгота
	Longitude float32 `json:"longitude"`

	// Phone Контактный телефон
	Phone string `json:"phone"`

	// PhotoId ID фотографии
	PhotoId *openapi_types.UUID `json:"photo_id"`

	// Title Название организации
	Title string `json:"title"`
}

// OrganizationUpdate defines model for OrganizationUpdate.
type OrganizationUpdate struct {
	// Description Подробное описание организации
	Description string `json:"description"`

	// Email Электронная почта
	Email openapi_types.Email `json:"email"`

	// Latitude Географическая широта
	Latitude float32 `json:"latitude"`

	// Longitude Географическая долгота
	Longitude float32 `json:"longitude"`

	// Phone Контактный телефон
	Phone string `json:"phone"`

	// PhotoId ID фотографии
	PhotoId *openapi_types.UUID `json:"photo_id"`

	// Title Название организации
	Title string `json:"title"`
}

// AddNewOrganizationJSONRequestBody defines body for AddNewOrganization for application/json ContentType.
type AddNewOrganizationJSONRequestBody = OrganizationCreate

// ApproveOrganizationJSONRequestBody defines body for ApproveOrganization for application/json ContentType.
type ApproveOrganizationJSONRequestBody = ModerationApprove

// RejectOrganizationJSONRequestBody defines body for RejectOrganization for application/json ContentType.
type RejectOrganizationJSONRequestBody = ModerationReject

// GetOrganizationByStatusJSONRequestBody defines body for GetOrganizationByStatus for application/json ContentType.
type GetOrganizationByStatusJSONRequestBody = GetOrganizationsByStatus

// UpdateOrganizationJSONRequestBody defines body for UpdateOrganization for application/json ContentType.
type UpdateOrganizationJSONRequestBody = OrganizationUpdate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получить список всех организаций
	// (GET /organizations)
	GetAllOrganizations(ctx echo.Context) error
	// Добавить новую организацию
	// (POST /organizations)
	AddNewOrganization(ctx echo.Context) error
	// Одобрить организацю
	// (POST /organizations/moderation/approve)
	ApproveOrganization(ctx echo.Context) error
	// Отклонить организацию
	// (POST /organizations/moderation/reject)
	RejectOrganization(ctx echo.Context) error
	// Получить организации по типу верификации
	// (GET /organizations/status)
	GetOrganizationByStatus(ctx echo.Context) error
	// Получить список всех организаций для пользователя
	// (GET /organizations/user)
	GetAllOrganizationsForUser(ctx echo.Context) error
	// Удалить организацию
	// (DELETE /organizations/{id})
	DeleteOrganization(ctx echo.Context, id string) error
	// Получить информацию об организации по ID
	// (GET /organizations/{id})
	GetOrganizationById(ctx echo.Context, id string) error
	// Обновить информацию об организации
	// (PUT /organizations/{id})
	UpdateOrganization(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAllOrganizations converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllOrganizations(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAllOrganizations(ctx)
	return err
}

// AddNewOrganization converts echo context to params.
func (w *ServerInterfaceWrapper) AddNewOrganization(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AddNewOrganization(ctx)
	return err
}

// ApproveOrganization converts echo context to params.
func (w *ServerInterfaceWrapper) ApproveOrganization(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ApproveOrganization(ctx)
	return err
}

// RejectOrganization converts echo context to params.
func (w *ServerInterfaceWrapper) RejectOrganization(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.RejectOrganization(ctx)
	return err
}

// GetOrganizationByStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetOrganizationByStatus(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetOrganizationByStatus(ctx)
	return err
}

// GetAllOrganizationsForUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllOrganizationsForUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAllOrganizationsForUser(ctx)
	return err
}

// DeleteOrganization converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteOrganization(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteOrganization(ctx, id)
	return err
}

// GetOrganizationById converts echo context to params.
func (w *ServerInterfaceWrapper) GetOrganizationById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetOrganizationById(ctx, id)
	return err
}

// UpdateOrganization converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateOrganization(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateOrganization(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/organizations", wrapper.GetAllOrganizations)
	router.POST(baseURL+"/organizations", wrapper.AddNewOrganization)
	router.POST(baseURL+"/organizations/moderation/approve", wrapper.ApproveOrganization)
	router.POST(baseURL+"/organizations/moderation/reject", wrapper.RejectOrganization)
	router.GET(baseURL+"/organizations/status", wrapper.GetOrganizationByStatus)
	router.GET(baseURL+"/organizations/user", wrapper.GetAllOrganizationsForUser)
	router.DELETE(baseURL+"/organizations/:id", wrapper.DeleteOrganization)
	router.GET(baseURL+"/organizations/:id", wrapper.GetOrganizationById)
	router.PUT(baseURL+"/organizations/:id", wrapper.UpdateOrganization)

}

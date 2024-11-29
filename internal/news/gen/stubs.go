// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package gen

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// CreateNews defines model for CreateNews.
type CreateNews struct {
	// Description Полное описание новости
	Description string `json:"description"`

	// Media Список медиа файлов
	Media []string `json:"media"`

	// OrganizationId Id организации, которая выпустила пост
	OrganizationId string `json:"organization_id"`

	// Title Заголовок новости
	Title string `json:"title"`
}

// News defines model for News.
type News struct {
	// Description Полное описание новости
	Description string `json:"description"`

	// Id Id новости
	Id string `json:"id"`

	// Media Список медиа файлов
	Media []string `json:"media"`

	// OrganizationId Id организации, которая выпустила пост
	OrganizationId string `json:"organization_id"`

	// PublishData Дата публикации
	PublishData string `json:"publish_data"`

	// Title Заголовок новости
	Title string `json:"title"`
}

// CreateNewsJSONRequestBody defines body for CreateNews for application/json ContentType.
type CreateNewsJSONRequestBody = CreateNews

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// GetAllNews
	// (GET /news)
	GetAllNews(ctx echo.Context) error
	// CreateNews
	// (POST /news)
	CreateNews(ctx echo.Context) error
	// GetAllNewsById
	// (GET /news/{id})
	GetNewsById(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAllNews converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllNews(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAllNews(ctx)
	return err
}

// CreateNews converts echo context to params.
func (w *ServerInterfaceWrapper) CreateNews(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateNews(ctx)
	return err
}

// GetNewsById converts echo context to params.
func (w *ServerInterfaceWrapper) GetNewsById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetNewsById(ctx, id)
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

	router.GET(baseURL+"/news", wrapper.GetAllNews)
	router.POST(baseURL+"/news", wrapper.CreateNews)
	router.GET(baseURL+"/news/:id", wrapper.GetNewsById)

}

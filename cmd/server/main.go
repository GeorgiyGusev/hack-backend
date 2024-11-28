package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hack-boilerplate/internal/todo/delivery"
	"net/http"
	"os"
)

func main() {
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	server.Use(middleware.Logger())

	serverId := uuid.NewString()

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		panic("service-name not set")
	}
	serviceVersion := os.Getenv("SERVICE_VERSION")
	if serviceVersion == "" {
		panic("service-version not set")
	}

	baseGroup := server.Group(fmt.Sprintf("/%s/%s", serviceName, serviceVersion))

	baseGroup.GET("/id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"id": serverId})
	})

	delivery.Register(baseGroup)

	err := server.Start(":8080")
	if err != nil {
		return
	}
}

package api

import (
	"cacahuete-api/configuration"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ApiHandler struct {
	pg   *gorm.DB
	conf *configuration.Configuration
}

func NewApiHandler(db *gorm.DB, conf *configuration.Configuration) *ApiHandler {
	handler := ApiHandler{
		pg:   db,
		conf: conf,
	}
	return &handler
}

func (api *ApiHandler) Register(v1 *echo.Group) {

	health := v1.Group("/health")
	health.GET("/alive", api.getAliveStatus)
	health.GET("/live", api.getAliveStatus)
	health.GET("/ready", api.getReadyStatus)

}

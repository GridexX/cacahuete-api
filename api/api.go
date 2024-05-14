package api

import (
	"cacahuete-api/configuration"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
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

func (api *ApiHandler) Register(v1 *echo.Group, conf *configuration.Configuration) {

	health := v1.Group("/health")
	health.GET("/alive", api.getAliveStatus)
	health.GET("/live", api.getAliveStatus)
	health.GET("/ready", api.getReadyStatus)

	app := v1.Group("/api")
	app.POST("/login", api.login)
	app.POST("/signup", api.signup)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(conf.JWTSecret),
	}
	app.Use(echojwt.WithConfig(config))
	app.POST("/logout", api.logout)
	app.GET("/restricted", api.extractUser(api.restricted))

}

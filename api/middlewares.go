package api

import (
	"cacahuete-api/db"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (api *ApiHandler) extractUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		logger.Debugln(claims.ID)
		t, err := db.GetTokenUser(api.pg, user.Raw, claims.UserID)
		if err != nil {
			return NewUnauthorizedError(err)
		}
		if t == nil {
			return NewUnauthorizedError(err)
		}

		c.Set("user", user) // Set the user into the context
		return next(c)
	}
}

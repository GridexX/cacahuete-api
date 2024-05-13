package api

import (
	"cacahuete-api/db"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var logger = logrus.WithField("context", "api/routes")

func (api *ApiHandler) getAliveStatus(c echo.Context) error {
	l := logger.WithField("request", "getAliveStatus")
	status := NewHealthResponse(LiveStatus)
	if err := c.Bind(status); err != nil {
		FailOnError(l, err, "Response binding failed")
		return NewInternalServerError(err)
	}
	l.WithFields(logrus.Fields{
		"action": "getStatus",
		"status": status,
	}).Debug("Health Status ping")

	return c.JSON(http.StatusOK, &status)
}

func (api *ApiHandler) getReadyStatus(c echo.Context) error {
	l := logger.WithField("request", "getReadyStatus")
	db, err := api.pg.DB()
	if err != nil {
		WarnOnError(l, err, "Unable to ping database to check connection.")
		return c.JSON(http.StatusServiceUnavailable, NewHealthResponse(NotReadyStatus))
	}

	err = db.Ping()
	if err != nil {
		FailOnError(l, err, "Error when trying to check database connection")
		l.WithError(err).Error("Error when trying to check database connection.")
		return c.JSON(http.StatusServiceUnavailable, NewHealthResponse(NotReadyStatus))
	}

	return c.JSON(http.StatusOK, NewHealthResponse(ReadyStatus))
}

func (api *ApiHandler) login(c echo.Context) error {

	l := logger.WithField("request", "login")

	u := new(UserConnectionRequest)
	if err := c.Bind(u); err != nil {
		FailOnError(l, err, "Body param failed")
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	user, err := db.GetUsername(api.pg, u.Username)

	if err != nil {
		return NewInternalServerError(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))

	if err != nil {
		return NewNotFoundError(err)
	}

	expirationDate := time.Now().Add(time.Hour * 72)
	// Set custom claims
	claims := &jwtCustomClaims{
		user.Username,
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationDate),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(api.conf.JWTSecret))
	if err != nil {
		return err
	}

	tokenDb := db.Token{
		Value:          t,
		ExpirationDate: expirationDate,
	}
	//Insert the new token in the DB
	db.UpsertToken(api.pg, tokenDb)

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (api *ApiHandler) restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Username
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func (api *ApiHandler) signup(c echo.Context) error {
	l := logger.WithField("request", "sign-up")

	u := new(UserCreationRequest)
	if err := c.Bind(u); err != nil {
		FailOnError(l, err, "Body param failed")
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error during hash generation:", err)
		return err
	}

	user := db.User{
		Email:     u.Email,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.FirstName,
		Password:  string(hashedPassword[:]),
	}

	db.CreateUser(api.pg, user)
	return c.NoContent(http.StatusCreated)
}

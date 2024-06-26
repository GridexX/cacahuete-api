package configuration

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

var logger = logrus.WithFields(logrus.Fields{
	"context": "configuration/configuration",
})

type Configuration struct {
	ListenPort          string
	ListenAddress       string
	ListenRoute         string
	LogLevel            string
	DBHost              string
	DBPort              string
	DBUser              string
	DBName              string
	DBSSLMode           string
	DBTimezone          string
	DBPassword          string
	TranslateValidation bool
	JWTSecret           string
}

func New() *Configuration {

	conf := Configuration{}
	var err error

	logLevel := os.Getenv("LOG_LEVEL")
	if len(logLevel) < 1 || logLevel != "debug" && logLevel != "error" && logLevel != "info" && logLevel != "trace" && logLevel != "warn" {
		logrus.WithFields(logrus.Fields{
			"logLevel": logLevel,
		}).Info("logLevel not conform, use `info` ")
		conf.LogLevel = "info"
	} else {
		conf.LogLevel = logLevel
	}

	conf.ListenPort = os.Getenv("API_PORT")
	conf.ListenAddress = os.Getenv("API_ADDRESS")
	conf.ListenRoute = os.Getenv("API_ROUTE")

	conf.DBHost = os.Getenv("POSTGRES_HOST")
	conf.DBPort = os.Getenv("POSTGRES_PORT")
	conf.DBUser = os.Getenv("POSTGRES_USER")
	conf.DBName = os.Getenv("POSTGRES_DB")
	conf.DBPassword = os.Getenv("POSTGRES_PASSWORD")
	conf.DBSSLMode = "disable"       //os.Getenv("POSTGRES_SSL")
	conf.DBTimezone = "Europe/Paris" //os.Getenv("POSTGRES_TIMEZONE")

	conf.TranslateValidation, err = strconv.ParseBool(os.Getenv("TRANSLATE_VALIDATION"))

	if err != nil {
		logger.Error("Failed to parse bool for TRANSLATE_VALIDATION")
		os.Exit(1)
	}

	conf.JWTSecret = os.Getenv("JWT_SECRET")

	return &conf
}

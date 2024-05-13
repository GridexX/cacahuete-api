package main

import (
	"cacahuete-api/api"
	"cacahuete-api/configuration"
	"cacahuete-api/db"
	"cacahuete-api/validation"
	"fmt"

	"github.com/sirupsen/logrus"
)

var logger = logrus.WithFields(logrus.Fields{
	"context": "main",
})

func main() {
	logger.Info("Cacahuete API Starting...")

	conf := configuration.New()

	pg, err := db.New(conf)

	if err != nil {
		return
	}

	err = db.AutoMigrate(pg)

	val := validation.New(conf)
	r := api.New(val)
	v1 := r.Group(conf.ListenRoute)

	if err != nil {
		return
	}
	h := api.NewApiHandler(pg, conf)

	h.Register(v1, conf)
	r.Logger.Fatal(r.Start(fmt.Sprintf("%v:%v", conf.ListenAddress, conf.ListenPort)))
}

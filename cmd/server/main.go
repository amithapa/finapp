package main

import (
	"net/http"

	"github.com/amithapa/finapp/internal/api"
	"github.com/amithapa/finapp/internal/config"
	"github.com/sirupsen/logrus"
)

//Create Server object and start listener
func main() {

	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithField("version", config.Version).Debug("Starting Server.")

	router, err := api.NewRouter()

	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}

	const addr = "0.0.0.0:8088"
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}

	// Starting server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("Server failed.")
	}

}

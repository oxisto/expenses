package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/oxisto/expenses/common"
	"github.com/oxisto/expenses/routes"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	// Set log level to debug
	// TODO: make this configurable some how
	logrus.SetLevel(logrus.DebugLevel)

	log = logrus.WithField("component", "main")
}

func main() {
	listen := "0.0.0.0:8080"

	log.Infof("Starting HTTP server @ %s...", listen)

	router := handlers.LoggingHandler(&common.DebugLogWriter{Component: "http"}, routes.NewRouter())
	err := http.ListenAndServe(listen, router)

	log.Errorf("An error occured: %v", err)
}

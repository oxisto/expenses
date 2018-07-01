/*
Copyright 2018 Christian Banse

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

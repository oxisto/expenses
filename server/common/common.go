package common

// TODO: extract this into a common library

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// DebugLogWriter implements io.Writer and writes all incoming text out to log level info.
type DebugLogWriter struct {
	Component string
}

func (d DebugLogWriter) Write(p []byte) (n int, err error) {
	logrus.WithField("component", d.Component).Debug(strings.TrimRight(string(p), "\n"))

	return len(p), nil
}

func JsonResponse(w http.ResponseWriter, r *http.Request, object interface{}, err error) {
	// uh-uh, we have an error
	if err != nil {
		logrus.Errorf("An error occured during processing of a REST request: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return not found if object is nil
	if object == nil {
		http.NotFound(w, r)
		return
	}

	// otherwise, lets try to decode the JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(object); err != nil {
		// uh-uh we couldn't decode the JSON
		logrus.Errorf("An error occured during encoding of the JSON response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

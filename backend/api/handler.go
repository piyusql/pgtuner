package api

import (
	"encoding/json"
	"net/http"
)

// PGSettingHandler :: handler for returning the pg settings list
func PGSettingHandler(w http.ResponseWriter, req *http.Request) {
	pgsettings := allPGSettings()

	if err := json.NewEncoder(w).Encode(pgsettings); err != nil {
		panic(err)
	}
}

// HealthCheckHandler :: handler for sharing app health check data
func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	health := doHealthCheck(req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

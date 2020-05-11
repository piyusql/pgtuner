package api

import (
	"encoding/json"
	"net/http"
)

func PGSettingHandler(w http.ResponseWriter, req *http.Request) {
	pgsettings := AllPGSettings()

	if err := json.NewEncoder(w).Encode(pgsettings); err != nil {
		panic(err)
	}
}

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	health := doHealthCheck(req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

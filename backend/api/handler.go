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

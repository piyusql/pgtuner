package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PGSettingHandler :: handler for returning the pg settings list
func PGSettingHandler(w http.ResponseWriter, req *http.Request) {
	pgsettings := allPGSettings()
	if err := json.NewEncoder(w).Encode(pgsettings); err != nil {
		panic(err)
	}
}

// PGTableHandler :: handler for returning the pg user table list
func PGTableHandler(w http.ResponseWriter, req *http.Request) {
	tables := allDBTables()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tables)
}

// HealthCheckHandler :: handler for sharing app health check data
func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	health := doHealthCheck(req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

// GetResourceMetricsHandler :: return list of timeseries by cleintID, chartName
func GetResourceMetricsHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	// start_time and end_time are GET parameters by URL string.
	clientID, _ := strconv.Atoi(params["clientID"])
	chartInfo := getChartInfo(params["chartName"])
	var data []TimeSeriesData
	var columns []string
	if chartInfo.ChartType != "" {
		data = getResourceMetrics(clientID, chartInfo.ChartID)
		if len(data) > 0 {
			for key := range data[0].Data {
				columns = append(columns, key)
			}
		}
	}
	properties := map[string]interface{}{
		"parameters": map[string]interface{}{
			"client_id":  clientID,
			"chart_name": params["chartName"],
			"start_time": nil,
			"end_time":   nil,
		},
		"current_page": 1,
		// placeholder for future need for pagination handler
		"next_page_link":   fmt.Sprintf("%s?_page=%d", req.URL.Path, 2),
		"data_columns":     columns,
		"data_time_bucket": 60, // data is bucketed for each per 60 seconds
		"chart_type":       chartInfo.ChartType,
	}
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"metadata":  properties,
		"resources": data,
	}
	json.NewEncoder(w).Encode(response)
}

package api

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/piyusgupta/pgtuner/backend/dba"
)

// ChartInfo :: data type to hold the chartType for stats
type ChartInfo struct {
	ChartID   int    `json:"chart_id"`
	ChartType string `json:"chart_type"`
	ChartName string `json:"chart_name"`
}

// TimeSeriesData :: data type to hold the stats for server DB
type TimeSeriesData struct {
	Timestamp time.Time  `json:"timestamp"`
	Data      JsonFields `json:"data"`
}

type JsonFields map[string]interface{}

func (f JsonFields) Value() (driver.Value, error) {
	j, err := json.Marshal(f)
	return j, err
}
func (f *JsonFields) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*f, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}
	return nil
}

// getChartInfo :: get a chart detail by its name
func getChartInfo(chartName string) ChartInfo {
	// return list of pg user table
	var chart ChartInfo
	db := dba.GetConnection()
	q := fmt.Sprintf(`SELECT
	chart_id AS ChartID,
	type AS ChartType,
	name AS ChartName
FROM
    charts
WHERE
    name = '%s';`, chartName)

	db.Get(&chart, q)
	return chart
}

// getResourceMetrics :: return a list of timeseries data
func getResourceMetrics(clientID, chartID int) []TimeSeriesData {
	// return list of pg user table
	var data []TimeSeriesData
	db := dba.GetConnection()
	q := fmt.Sprintf(`SELECT
	timestamp AS Timestamp,
	data AS Data
FROM
	metrics m
WHERE
	m.client_id = %d AND m.chart_id=%d;`, clientID, chartID)

	db.Select(&data, q)
	return data
}

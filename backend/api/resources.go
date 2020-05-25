package api

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/piyusgupta/pgtuner/backend/dba"
)

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

// getResourceMetrics :: return a list of timeseries data
func getResourceMetrics(clientID int, chartName string) []TimeSeriesData {
	// return list of pg user table
	var data []TimeSeriesData
	db := dba.GetConnection()
	q := fmt.Sprintf(`SELECT
	timestamp AS Timestamp,
	data AS Data
FROM
    metrics m
    JOIN charts c USING (chart_id)
WHERE
    c.name = '%s'
	AND m.client_id = %d;`, chartName, clientID)

	db.Select(&data, q)
	return data
}

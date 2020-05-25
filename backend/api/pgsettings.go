package api

import (
	"github.com/piyusgupta/pgtuner/backend/dba"
)

// PGSetting :: data type to hold postgres settings
type PGSetting struct {
	Name             string `json:"name"`
	Setting          string `json:"value"`
	Category         string `json:"category"`
	ShortDescription string `json:"description"`
	Context          string `json:"context"`
	ValueType        string `json:"type"`
}

func allPGSettings() []PGSetting {
	// return all pg setting values
	var pgsettings []PGSetting
	db := dba.GetConnection()

	err := db.Select(&pgsettings, dba.QueryDBSettings)
	dba.CheckErr(err)
	return pgsettings
}

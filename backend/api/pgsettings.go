package api

import (
	"github.com/piyusgupta/pgtuner/backend/dba"
)

// PGSetting :: data type to hold postgres settings
type PGSetting struct {
	Name             string
	Setting          string
	Category         string
	ShortDescription string
	Context          string
	ValueType        string
}

func allPGSettings() []PGSetting {
	// return all pg setting values
	var pgsettings []PGSetting
	db := dba.GetConnection()
	q := `SELECT
    name as Name,
    setting as Setting,
    category as Category,
    short_desc as ShortDescription,
	context as Context,
	vartype as ValueType
FROM
    pg_settings
ORDER BY
    category,
    name;`

	rows, err := db.Queryx(q)
	dba.CheckErr(err)
	for rows.Next() {
		setting := new(PGSetting)
		rows.StructScan(&setting)
		pgsettings = append(pgsettings, *setting)
	}
	if err := rows.Err(); err != nil {
		// make sure that there was no issue during the process
		dba.CheckErr(err)
	}
	return pgsettings
}

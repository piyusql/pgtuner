package api

import (
	"fmt"

	dba "github.com/piyusgupta/pgtuner/backend/dba"
)

type PGSetting struct {
	name      string
	setting   string
	category  string
	shortDesc string
	context   string
}

func AllPGSettings() []PGSetting {
	// return all pg setting values
	var pgsettings []PGSetting
	db := dba.GetConnection()
	q := `
SELECT name,
       setting,
       category,
       short_desc,
       context
FROM pg_settings
ORDER BY category,
         name LIMIT 5;`

	rows, err := db.Queryx(q)
	dba.CheckErr(err)
	for rows.Next() {
		setting := new(PGSetting)
		rows.StructScan(&setting)
		pgsettings = append(pgsettings, *setting)
		fmt.Printf("%#v\n", *setting)
	}
	if err := rows.Err(); err != nil {
		// make sure that there was no issue during the process
		dba.CheckErr(err)
	}
	return pgsettings
}

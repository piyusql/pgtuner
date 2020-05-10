package api

import (
    "dbutils"
)

type struct PGSettings{
    name string `name`
    settings string
    category string
    short_desc string
};

func AllPGSettings(dbname string) PGSettings[] {
    db = db.GetConnection(dbname)
    defer db.Close()
	pgsettings := PGSettings[]
    q := `
SELECT name,
       setting,
       category,
       short_desc,
       context
FROM pg_settings
ORDER BY category,
         name;`

    rows, err := db.Query(q)
    checkErr(err)
    for rows.Next() {
        var settings PGSettings
        rows.Scan(&settings)
		append(pgsettings, settings)
    }
	return pgsettings
}

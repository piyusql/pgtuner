package api

import (
	"github.com/piyusgupta/pgtuner/backend/dba"
)

// PGSetting :: data type to hold postgres settings
type DBTable struct {
	Name     string
	RowCount string
	Size     uint64
	SizeTxt  string
}

func allDBTables() []DBTable {
	// return list of pg user table
	var tables []DBTable
	db := dba.GetConnection()

	rows, err := db.Queryx(dba.QueryDBTables)
	dba.CheckErr(err)
	for rows.Next() {
		table := new(DBTable)
		rows.StructScan(&table)
		tables = append(tables, *table)
	}
	if err := rows.Err(); err != nil {
		// make sure that there was no issue during the process
		dba.CheckErr(err)
	}
	return tables
}

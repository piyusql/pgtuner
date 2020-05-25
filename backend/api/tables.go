package api

import (
	"github.com/piyusgupta/pgtuner/backend/dba"
)

// DBTable :: data type to hold user table properties
type DBTable struct {
	Name     string `json:"name"`
	RowCount string `json:"row_count"`
	Size     uint64 `json:size"`
	SizeTxt  string `json:"size_text"`
}

func allDBTables() []DBTable {
	// return list of pg user table
	var tables []DBTable
	db := dba.GetConnection()

	err := db.Select(&tables, dba.QueryDBTables)
	dba.CheckErr(err)
	return tables
}

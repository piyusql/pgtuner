package dba

const (
	QueryDBSettings = `SELECT
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

	QueryDBTables = `SELECT
    table_name AS Name,
    n_live_tup AS RowCount,
    pg_relation_size(table_name) AS Size,
    pg_size_pretty(pg_relation_size(table_name)) AS SizeTxt
FROM
    information_schema.tables
    INNER JOIN pg_stat_user_tables ON table_name = relname
WHERE
    table_schema = 'public'
ORDER BY
    3 DESC`
)

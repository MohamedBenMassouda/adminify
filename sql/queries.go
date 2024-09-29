package sql_queries

func ListQuery(tableName string) string {
	return "SELECT * FROM " + tableName
}

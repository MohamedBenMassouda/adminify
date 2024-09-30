package sql_queries

import (
	"fmt"
	"strings"
)

func ListQuery(tableName string, fields []string) string {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ", "), tableName)

	return query
}

func ListQuerWithPagination(tableName string, fields []string, limit, offset int) string {
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT %d OFFSET %d", strings.Join(fields, ", "), tableName, limit, offset)

	return query
}

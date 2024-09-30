package sql_queries

import (
	"fmt"
	"strings"
)

func ListQuery(tableName string, fields []string) string {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ", "), tableName)

	return query
}

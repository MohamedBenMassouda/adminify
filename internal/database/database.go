package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/MohamedBenMassouda/adminify/internal/model"
)

// DB wraps the sql.DB object and provides custom methods
type DB struct {
	*sql.DB
}

// NewDB creates a new DB instance
func NewDB(db *sql.DB) *DB {
	return &DB{db}
}

// Query executes a query and returns the results
func (db *DB) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				entry[col] = string(b)
			} else {
				entry[col] = val
			}
		}

		result = append(result, entry)
	}

	return result, nil
}

// Exec executes a query without returning any rows
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

// Insert inserts a new record into the database
func (db *DB) Insert(model *model.Model, data map[string]interface{}) error {
	var columns []string
	var placeholders []string
	var values []interface{}

	for _, field := range model.Fields {
		if value, ok := data[field.Name]; ok {
			columns = append(columns, field.Name)
			placeholders = append(placeholders, "?")
			values = append(values, value)
		}
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		model.TableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := db.Exec(query, values...)
	return err
}

// Update updates an existing record in the database
func (db *DB) Update(model *model.Model, id interface{}, data map[string]interface{}) error {
	var setStatements []string
	var values []interface{}

	for _, field := range model.Fields {
		if value, ok := data[field.Name]; ok && field.Name != "id" {
			setStatements = append(setStatements, fmt.Sprintf("%s = ?", field.Name))
			values = append(values, value)
		}
	}

	values = append(values, id)

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = ?",
		model.TableName,
		strings.Join(setStatements, ", "),
	)

	_, err := db.Exec(query, values...)
	return err
}

// Delete removes a record from the database
func (db *DB) Delete(model *model.Model, id interface{}) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", model.TableName)
	_, err := db.Exec(query, id)
	return err
}

// GetByID retrieves a single record by its ID
func (db *DB) GetByID(model *model.Model, id interface{}) (map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", model.TableName)
	results, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, sql.ErrNoRows
	}

	return results[0], nil
}

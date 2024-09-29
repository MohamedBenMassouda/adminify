package model

import (
	"errors"
	"reflect"
)

type Model struct {
	TableName string
	Name      string
	Fields    []Field
}

type Field struct {
	SQLType string
	Name    string
	Type    reflect.Type
}

func New(modelStruct interface{}, tableName string) (*Model, error) {
	t := reflect.TypeOf(modelStruct)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("Model must be a struct")
	}

	model := &Model{
		Name:      t.Name(),
		TableName: tableName,
	}
	fields := make([]Field, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.PkgPath != "" {
			continue
		}

		fields = append(fields, Field{
			SQLType: getSQLType(field.Type),
			Name:    field.Name,
			Type:    field.Type,
		})
	}

	model.Fields = fields

	return model, nil
}

func getSQLType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "INTEGER"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "INTEGER UNSIGNED"
	case reflect.Float32, reflect.Float64:
		return "FLOAT"
	case reflect.String:
		return "TEXT"
	case reflect.Bool:
		return "BOOLEAN"
	default:
		return "TEXT"
	}
}

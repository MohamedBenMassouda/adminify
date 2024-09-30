package model

import (
	"errors"
	"log"
	"reflect"
)

type Model struct {
	TableName string
	Name      string
	Fields    []Field
}

type Field struct {
	SQLType    string
	Name       string
	Type       reflect.Type
	ColumnName string
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

	var checkTags []string = []string{"admin", "json"}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.PkgPath != "" {
			continue
		}

		var columnName string

		if field.Type.Kind() == reflect.Slice {
			log.Printf("Ignoring slice type field: %s\n", field.Name)
			continue
		}

		if field.Type.Kind() == reflect.Ptr {
			log.Printf("Ignoring pointer type field: %s\n", field.Name)
			continue
		}

		if field.Type.Kind() == reflect.Struct {
			if field.Name == "Model" || field.Name == "gorm.Model" {
				for j := 0; j < field.Type.NumField(); j++ {
					f := field.Type.Field(j)

					if f.PkgPath != "" {
						continue
					}

					columnName = f.Tag.Get("json")

					if columnName == "-" {
						continue
					}

					fields = append(fields, Field{
						SQLType:    getSQLType(f.Type),
						Name:       f.Name,
						Type:       f.Type,
						ColumnName: columnName,
					})
				}

				continue
			} else {
				continue
			}
		}

		for _, tag := range checkTags {
			if tagValue := field.Tag.Get(tag); tagValue != "" {
				columnName = tagValue
				break
			}
		}

		if columnName == "" {
			columnName = field.Name
		}

		if columnName == "-" {
			continue
		}

		fields = append(fields, Field{
			SQLType:    getSQLType(field.Type),
			Name:       field.Name,
			Type:       field.Type,
			ColumnName: columnName,
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

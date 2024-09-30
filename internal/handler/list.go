package handler

import (
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/MohamedBenMassouda/adminify/internal/database"
	"github.com/MohamedBenMassouda/adminify/internal/model"
	sql_queries "github.com/MohamedBenMassouda/adminify/sql"
)

const ITEMS_PER_PAGE = 100

func List(w http.ResponseWriter, r *http.Request, db *database.DB, models map[string]*model.Model, tmpl *template.Template) {
	modelName := r.URL.Query().Get("model")
	model, ok := models[modelName]

	if !ok {
		http.Error(w, "Model not found", http.StatusNotFound)
		return
	}

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * ITEMS_PER_PAGE

	fieldNames := make([]string, 0)

	for _, field := range model.Fields {
		fieldNames = append(fieldNames, field.ColumnName)
	}

	// Get total count of items
	totalItems, err := db.Count(model.TableName)

	if err != nil {
		log.Println("Error counting items", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / ITEMS_PER_PAGE))

	data, err := db.Query(sql_queries.ListQuerWithPagination(model.TableName, fieldNames, ITEMS_PER_PAGE, offset))
	if err != nil {
		log.Println("Error querying database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	startIndex := (page - 1) * ITEMS_PER_PAGE
	endIndex := int(math.Min(float64(page*ITEMS_PER_PAGE), float64(totalItems)))

	err = tmpl.ExecuteTemplate(w, "list.html", map[string]interface{}{
		"Model":        model,
		"Data":         data,
		"CurrentPage":  page,
		"PreviousPage": page - 1,
		"NextPage":     page + 1,
		"TotalPages":   totalPages,
		"StartIndex":   startIndex,
		"EndIndex":     endIndex,
		"TotalItems":   totalItems,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

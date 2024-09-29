package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/MohamedBenMassouda/adminify/internal/database"
	"github.com/MohamedBenMassouda/adminify/internal/model"
	sql_queries "github.com/MohamedBenMassouda/adminify/sql"
)

func List(w http.ResponseWriter, r *http.Request, db *database.DB, models map[string]*model.Model, tmpl *template.Template) {
	modelName := r.URL.Query().Get("model")
	model, ok := models[modelName]
	log.Println(model.Fields)

	if !ok {
		http.Error(w, "Model not found", http.StatusNotFound)
		return
	}

	data, err := db.Query(sql_queries.ListQuery(model.TableName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "list.html", map[string]interface{}{
		"Model": model,
		"Data":  data,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
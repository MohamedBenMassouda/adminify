package handler

import (
	"html/template"
	"net/http"

	"github.com/MohamedBenMassouda/adminify/internal/model"
)

func Home(w http.ResponseWriter, r *http.Request, models map[string]*model.Model, tmpl *template.Template) {
	err := tmpl.ExecuteTemplate(w, "base.html", map[string]interface{}{
		"Models": models,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MohamedBenMassouda/adminify/internal/database"
	"github.com/MohamedBenMassouda/adminify/internal/model"
)

func Delete(w http.ResponseWriter, r *http.Request, db *database.DB, models map[string]*model.Model, path string) {
	modelName := r.URL.Query().Get("model")
	model, ok := models[modelName]

	if !ok {
		http.Error(w, "Model not found", http.StatusNotFound)
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	err := db.Delete(model, id)
	if err != nil {
		log.Println("Error deleting from database", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf(path+"/list?model=%s", model.TableName), http.StatusSeeOther)
}

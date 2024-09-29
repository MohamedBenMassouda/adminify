package admin

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/MohamedBenMassouda/adminify/internal/database"
	"github.com/MohamedBenMassouda/adminify/internal/handler"
	"github.com/MohamedBenMassouda/adminify/internal/model"
)

type Panel struct {
	models    map[string]*model.Model
	db        *database.DB
	templates *template.Template
}

func NewPanel(db *sql.DB) (*Panel, error) {
	templates, err := template.ParseGlob("templates/*.html")
	if err != nil {
		return nil, err
	}

	return &Panel{
		models:    make(map[string]*model.Model),
		db:        database.NewDB(db),
		templates: templates,
	}, nil
}

func (p *Panel) RegisterModel(modelStruct interface{}, tableName string) error {
	model, err := model.New(modelStruct, tableName)
	if err != nil {
		return err
	}
	p.models[model.TableName] = model
	return nil
}

func (p *Panel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	switch r.URL.Path {
	case "/":
		handler.Home(w, r, p.models, p.templates)
	case "/list":
		handler.List(w, r, p.db, p.models, p.templates)
	}
}

package admin

import (
	"database/sql"
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/MohamedBenMassouda/adminify/internal/database"
	"github.com/MohamedBenMassouda/adminify/internal/handler"
	"github.com/MohamedBenMassouda/adminify/internal/model"
)

//go:embed templates
var templatesFS embed.FS

//go:emned static
var staticFS embed.FS

type Panel struct {
	models    map[string]*model.Model
	db        *database.DB
	templates *template.Template
	Path      string
}

func NewPanel(db *sql.DB, driver, path string) (*Panel, error) {
	templates, err := template.ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		return nil, err
	}

	return &Panel{
		models:    make(map[string]*model.Model),
		db:        database.NewDB(db, driver),
		templates: templates,
		Path:      path,
	}, nil
}

func (p *Panel) GetStaticFS() http.FileSystem {
	return http.FS(staticFS)
}

func (p *Panel) GetModels() map[string]*model.Model {
	return p.models
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
	case p.Path + "/":
		handler.Home(w, r, p.models, p.templates)
	case p.Path + "/list":
		handler.List(w, r, p.db, p.models, p.templates)
	case p.Path + "/delete":
		handler.Delete(w, r, p.db, p.models, p.Path)
	}
}

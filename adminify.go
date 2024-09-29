package adminify

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MohamedBenMassouda/adminify/pkg/admin"
)

type Adminify struct {
	Panel *admin.Panel
}

func NewAdminify(db *sql.DB) Adminify {
	panel, err := admin.NewPanel(db)

	if err != nil {
		log.Fatal(err)
	}

	return Adminify{
		Panel: panel,
	}
}

func (a *Adminify) RegisterModel(modelStruct interface{}, tableName string) {
	err := a.Panel.RegisterModel(modelStruct, tableName)

	if err != nil {
		log.Fatal(err)
	}
}

func (a *Adminify) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Panel.ServeHTTP(w, r)
}

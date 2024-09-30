package adminify

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MohamedBenMassouda/adminify/pkg/admin"
	"github.com/gin-gonic/gin"
)

type Adminify struct {
	Panel *admin.Panel
}

func NewAdminify(db *sql.DB, path string) Adminify {
	panel, err := admin.NewPanel(db, path)

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

// RegisterGinRoute registers the admin panel to a Gin route
func (a *Adminify) RegisterGinRoute(r *gin.Engine) {
	r.StaticFS(a.Panel.Path+"/static", a.Panel.GetStaticFS()) // Serve the static files
	r.Any(a.Panel.Path+"/*any", gin.WrapH(a.Panel))           // Wrap the Adminify panel in a Gin handler
}

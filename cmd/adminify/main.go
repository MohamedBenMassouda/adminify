package main

import (
	"log"
	"net/http"

	"github.com/MohamedBenMassouda/adminify/config"
	"github.com/MohamedBenMassouda/adminify/internal/middleware"
	"github.com/MohamedBenMassouda/adminify/pkg/admin"
	_ "github.com/lib/pq" // or your preferred database driver
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func init() {
	config.LoadEnvVariables()
}

func main() {
	db, err := config.ConnectDB("postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	panel, err := admin.NewPanel(db)
	if err != nil {
		log.Fatal(err)
	}

	err = panel.RegisterModel(User{}, "users")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/admin/", http.StripPrefix("/admin", panel))
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	loggedMux := middleware.Logger(mux)

	log.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", loggedMux)
}

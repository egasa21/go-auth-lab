package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/egasa21/go-auth-lab/basic-auth/internal/httpapi"
	"github.com/egasa21/go-auth-lab/basic-auth/internal/user"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	store := user.NewStore(db)

	handler := httpapi.NewHandler(store)
	router := httpapi.NewRouter(handler)

	s := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on :8080")
	log.Fatal(s.ListenAndServe())
}

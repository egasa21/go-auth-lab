package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/egasa21/go-auth-lab/basic-auth/internal/user"
)

func main() {

	db, err := sql.Open("pgx", "postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	store := user.NewStore(db)

	_ = store

	s := http.Server{
		Addr: ":8080",
	}

	log.Println("Starting server on :8080")
	log.Fatal(s.ListenAndServe())
}

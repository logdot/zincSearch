package main

import (
	"net/http"
	"os"
	"zincSearchProject/zincindex"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", index)
	r.Get("/IndexDB", indexDB)

	http.ListenAndServe(":3000", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func indexDB(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Indexing"))

	username := os.Getenv("ZincUsername")
	password := os.Getenv("ZincPassword")
	zincindex.Authenticate("http://127.0.0.1:4080/", "enronmail", username, password).IndexDB(`.\resources\enron_mail_20110402\`)
}

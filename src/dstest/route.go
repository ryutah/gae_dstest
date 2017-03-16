package dstest

import (
	"net/http"

	"dstest/gzip"
	"dstest/index"
	"dstest/middleware"
	"dstest/user"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetRoutes() {
	m := mux.NewRouter()

	m.HandleFunc("/", index.Get)

	m.HandleFunc("/user", user.Post).Methods(http.MethodPost)
	m.HandleFunc("/user/{id}", user.Get).Methods(http.MethodGet)

	m.HandleFunc("/gzip", gzip.Post).Methods(http.MethodPost)
	m.HandleFunc("/gzip/{id}", gzip.Get).Methods(http.MethodGet)

	n := negroni.New(
		middleware.NewRequestLogger(),
		middleware.NewContext(),
	)
	n.UseHandler(m)

	http.Handle("/", n)
}

package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type API struct{}

func NewAPI() *API {
	return &API{}
}

func (a *API) Routes(r *mux.Router) {
	// Initialize routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
}

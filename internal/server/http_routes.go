package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *HTTP) setRoutes(r *mux.Router) {
	api := r.PathPrefix("/api/v1/users-api").Subrouter()

	api.HandleFunc("/users",
		h.getUsersByID,
	).Methods(http.MethodGet)
}

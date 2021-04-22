package store

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type StoreController struct {
	service *StoreService
}

func (c *StoreController) GetStore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c.service.GetStoreByName(id)
}

func (c *StoreController) initRoutes(r *mux.Router) {
	r.HandleFunc("/store/{id:[A-z]+}", c.GetStore).Methods(http.MethodGet)
}

func StoreInit(db *sql.DB, r *mux.Router) {
	service := StoreService{db: db}
	controller := StoreController{service: &service}
	controller.initRoutes(r)
}

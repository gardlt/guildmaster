package store

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type StoreController struct {
	service *StoreService
}

func (c *StoreController) getStore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	c.service.getStoreByName(id)
}

func (c *StoreController) createStore(w http.ResponseWriter, r *http.Request) {

	if err := c.service.createStore("potato"); err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(nil)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(nil)
}

func (c *StoreController) initRoutes(r *mux.Router) {
	r.HandleFunc("/store", c.createStore).Methods(http.MethodPost)
	r.HandleFunc("/store/{id:[A-z]+}", c.getStore).Methods(http.MethodGet)
}

func StoreInit(db *sql.DB, r *mux.Router) {
	service := StoreService{db: db}
	controller := StoreController{service: &service}
	controller.initRoutes(r)
	log.Info("Initializing the Store Controller")
}

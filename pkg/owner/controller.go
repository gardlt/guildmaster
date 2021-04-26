package owner

import (
	"database/sql"
	"encoding/json"
	"net/http"

	common "baiten.io/genbo/pkg/common"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type OwnerController struct {
	service *OwnerService
}

func (c *OwnerController) createOwner(w http.ResponseWriter, r *http.Request) {
	var o Owner
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&o); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := c.service.createOwner(o.Name); err != nil {
		common.RespondWithError(w, http.StatusBadRequest, "Failed to create new owner")
		return
	}

	common.RespondWithJSON(w, http.StatusOK, "Created a new Owner")
}

func (c *OwnerController) initRoutes(r *mux.Router) {
	r.HandleFunc("/owner", c.createOwner).Methods(http.MethodPost)
}

func OwnerInit(db *sql.DB, r *mux.Router) {
	service := OwnerService{db: db}
	controller := OwnerController{service: &service}
	controller.initRoutes(r)
	log.Info("Initializing the Owner Controller")
}

package store

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type StoreService struct {
	db *sql.DB
}

func (s *StoreService) getStoreByName(name string) Store {

	store := Store{}

	if err := s.db.QueryRow("SELECT * FROM  genbo.stores WHERE id=$1", name).Scan(&store.Name); err != nil {
		log.Error(err)
	}

	return store
}

func (s *StoreService) createStore(name string) error {
	return nil
}

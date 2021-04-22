package store

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type StoreService struct {
	db *sql.DB
}

func (s *StoreService) GetStoreByName(name string) Store {

	store := Store{}

	if err := s.db.QueryRow("SELECT * FROM  public.store WHERE id=$1", name).Scan(&store.Name); err != nil {
		log.Error(err)
	}

	return store
}

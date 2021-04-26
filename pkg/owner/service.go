package owner

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type OwnerService struct {
	db *sql.DB
}

func (o *OwnerService) getOwnerByName(name string) (Owner, error) {

	owner := Owner{}

	if err := o.db.QueryRow("SELECT * FROM  genbo.owner WHERE name=$1", name).Scan(&owner.Name); err != nil {
		log.Error(err)
		return owner, err
	}

	return owner, nil
}

func (o *OwnerService) createOwner(name string) error {
	if err := o.db.QueryRow("INSERT INTO genbo.owner(name) VALUES($1);", name).Err(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

package service

import (
	"database/sql"
	"log"

	"github.com/charakoba-com/fault_info/backend/db"
	"github.com/charakoba-com/fault_info/backend/model"
	"github.com/pkg/errors"
)

// TypeService type
type TypeService struct{}

// Listup types
func (svc *TypeService) Listup(tx *sql.Tx) ([]model.Type, error) {
	log.Printf("service.Type.Listup")

	var typeList db.TypeList
	if err := typeList.Listup(tx); err != nil {
		return nil, errors.Wrap(err, `loading type list`)
	}
	l := make(model.TypeList, len(typeList))
	for i, typ := range typeList {
		if err := l[i].FromDB(&typ); err != nil {
			return nil, errors.Wrap(err, `converting db.Type to model.Type`)
		}
	}

	return l, nil
}

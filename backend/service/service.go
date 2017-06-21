package service

import (
	"database/sql"
	"log"

	"github.com/charakoba-com/fault_info/backend/db"
	"github.com/charakoba-com/fault_info/backend/model"
	"github.com/pkg/errors"
)

// ServiceService type
type ServiceService struct{}

// Listup services
func (svc *ServiceService) Listup(tx *sql.Tx) ([]model.Service, error) {
	log.Printf("service.Service.Listup")

	var serviceList db.ServiceList
	if err := serviceList.Listup(tx); err != nil {
		return nil, errors.Wrap(err, `loading service list`)
	}
	l := make(model.ServiceList, len(serviceList))
	for i, service := range serviceList {
		if err := l[i].FromDB(&service); err != nil {
			return nil, errors.Wrap(err, `converting db.Service to model.Service`)
		}
	}

	return l, nil
}

package model

import (
	"log"

	"github.com/charakoba-com/fault_info/backend/db"
)

// Service model
type Service struct {
	Name string `json:"name"`
}

// ServiceList is array of service
type ServiceList []Service

// FromDB binds db.Service to model.Service
func (service *Service) FromDB(ds *db.Service) error {
	log.Printf("model.Service.FromDB")

	service.Name = ds.Name
	return nil
}

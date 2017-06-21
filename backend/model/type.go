package model

import (
	"log"

	"github.com/charakoba-com/fault_info/backend/db"
)

// Type model
type Type struct {
	Type string `json:"type"`
}

// TypeList is array of type
type TypeList []Type

// FromDB binds db.Type to model.Type
func (typ *Type) FromDB(dt *db.Type) error {
	log.Printf("model.Type.FromDB")

	typ.Type = dt.Type
	return nil
}

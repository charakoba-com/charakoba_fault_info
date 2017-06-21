package model

import (
	"database/sql"
	"log"
	"time"

	"github.com/charakoba-com/fault_info/backend/db"
	"github.com/pkg/errors"
)

// Info model
type Info struct {
	ID      int       `json:"id"`
	Type    string    `json:"type"`
	Service string    `json:"service"`
	Begin   time.Time `json:"begin"`
	End     time.Time `json:"end,omitempty"`
	Detail  string    `json:"detail"`
}

// InfoList is array of info
type InfoList []Info

// Load info from DB
func (info *Info) Load(tx *sql.Tx, id int) error {
	log.Printf("model.Info.Load %d", id)
	di := db.Info{}
	if err := di.Load(tx, id); err != nil {
		return errors.Wrap(err, "loading db.Info")
	}

	if err := info.FromDB(&di); err != nil {
		return errors.Wrap(err, "scanning db.Info")
	}

	return nil
}

// FromDB binds db.Info to model.Info
func (info *Info) FromDB(di *db.Info) error {
	log.Printf("model.Info.FromDB")

	info.ID = di.ID
	info.Type = di.Type
	info.Service = di.Service
	info.Begin = di.Begin
	if di.End.Valid {
		info.End = di.End.Time
	}
	info.Detail = di.Detail

	return nil
}

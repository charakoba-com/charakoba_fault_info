package db

import (
	"bytes"
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

// Service database connector
type Service struct {
	Name string
}

// ServiceList is array of Service
type ServiceList []Service

const serviceSelectColumns = `name`
const serviceTable = `services`

// Scan database row
func (ds *Service) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&ds.Name)
}

// Listup Service List
func (l *ServiceList) Listup(tx *sql.Tx) error {
	log.Printf("db.Service.Listup")

	stmt := bytes.Buffer{}
	stmt.WriteString(`SELECT `)
	stmt.WriteString(serviceSelectColumns)
	stmt.WriteString(` FROM `)
	stmt.WriteString(serviceTable)

	log.Printf("SQL: %s", stmt.String())

	rows, err := tx.Query(stmt.String())
	if err != nil {
		return errors.Wrap(err, `querying select`)
	}
	if err := l.FromRows(rows); err != nil {
		return errors.Wrap(err, `scanning rows`)
	}
	return nil
}

// FromRows scanning rows
func (l *ServiceList) FromRows(rows *sql.Rows) error {
	log.Printf("db.Service.FromRows")

	res := ServiceList{}
	for rows.Next() {
		service := Service{}
		if err := service.Scan(rows); err != nil {
			return errors.Wrap(err, `scannning row`)
		}
		res = append(res, service)
	}
	*l = res
	return nil
}

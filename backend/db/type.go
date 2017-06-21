package db

import (
	"bytes"
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

// Type database connector
type Type struct {
	Type string
}

// TypeList is array of Type
type TypeList []Type

const typeSelectColumns = `infotype`
const typeTable = `infotype`

// Scan database row
func (dt *Type) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&dt.Type)
}

// Listup Type List
func (l *TypeList) Listup(tx *sql.Tx) error {
	log.Printf("db.Type.Listup")

	stmt := bytes.Buffer{}
	stmt.WriteString(`SELECT `)
	stmt.WriteString(typeSelectColumns)
	stmt.WriteString(` FROM `)
	stmt.WriteString(typeTable)

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
func (l *TypeList) FromRows(rows *sql.Rows) error {
	log.Printf("db.Type.FromRows")

	res := TypeList{}
	for rows.Next() {
		typ := Type{}
		if err := typ.Scan(rows); err != nil {
			return errors.Wrap(err, `scanning row`)
		}
		res = append(res, typ)
	}
	*l = res
	return nil
}

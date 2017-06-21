package db

import (
	"bytes"
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const infoSelectColumns = `id, infotype, service, begin, end, detail`
const infoTable = `faultinfo`

// Info database connector
type Info struct {
	ID      int
	Type    string
	Service string
	Begin   time.Time
	End     mysql.NullTime
	Detail  string
}

// InfoList is array of info
type InfoList []Info

// Scan database row
func (di *Info) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&di.ID, &di.Type, &di.Service, &di.Begin, &di.End, &di.Detail)
}

// Create new record
func (di *Info) Create(tx *sql.Tx) error {
	log.Printf("db.Info.Create")

	stmt := bytes.Buffer{}
	stmt.WriteString(`INSERT INTO `)
	stmt.WriteString(infoTable)
	stmt.WriteString(` (infotype, service, begin, end, detail) VALUES (?, ?, ?, ?, ?)`)

	result, err := tx.Exec(stmt.String(), di.Type, di.Service, di.Begin, di.End, di.Detail)
	if err != nil {
		return errors.Wrap(err, `exec query`)
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return err
	}

	di.ID = int(lii)

	return nil
}

// Load from database
func (di *Info) Load(tx *sql.Tx, id int) error {
	log.Printf("db.Info.Load %d", id)

	stmt := bytes.Buffer{}
	stmt.WriteString(`SELECT `)
	stmt.WriteString(infoSelectColumns)
	stmt.WriteString(` FROM `)
	stmt.WriteString(infoTable)
	stmt.WriteString(` WHERE id = ?`)

	row := tx.QueryRow(stmt.String(), id)

	if err := di.Scan(row); err != nil {
		return errors.Wrap(err, `scanning row`)
	}

	return nil
}

// Update Info
func (di *Info) Update(tx *sql.Tx) error {
	log.Printf("db.Info.Update")

	stmt := bytes.Buffer{}
	stmt.WriteString(`UPDATE `)
	stmt.WriteString(infoTable)
	stmt.WriteString(` SET infotype = ?, service = ?, begin = ?, end = ?, detail = ? WHERE id = ?`)

	log.Printf("SQL: %s", stmt.String())
	_, err := tx.Exec(stmt.String(), di.Type, di.Service, di.Begin, di.End, di.Detail, di.ID)
	if err != nil {
		return errors.Wrap(err, `exec query`)
	}

	return nil
}

// Listup Info List
func (l *InfoList) Listup(tx *sql.Tx) error {
	log.Printf("db.Info.Listup")

	stmt := bytes.Buffer{}
	stmt.WriteString(`SELECT `)
	stmt.WriteString(infoSelectColumns)
	stmt.WriteString(` FROM `)
	stmt.WriteString(infoTable)

	log.Printf("SQL: %s", stmt.String())

	rows, err := tx.Query(stmt.String())
	if err != nil {
		return errors.Wrap(err, `querying select`)
	}
	if err := l.FromRows(rows); err != nil {
		return errors.Wrap(err, "scanning rows")
	}
	return nil
}

// FromRows scanning rows
func (l *InfoList) FromRows(rows *sql.Rows) error {
	log.Printf("db.Info.FromRows")

	res := InfoList{}
	for rows.Next() {
		info := Info{}
		if err := info.Scan(rows); err != nil {
			return errors.Wrap(err, `scanning row`)
		}
		res = append(res, info)
	}
	*l = res
	return nil
}

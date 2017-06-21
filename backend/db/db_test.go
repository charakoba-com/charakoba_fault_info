package db_test

import (
	"testing"

	"github.com/charakoba-com/fault_info/backend/db"
)

func TestDBConnection(t *testing.T) {
	err := db.Init(nil)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	tx, err := db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	defer tx.Rollback()
}

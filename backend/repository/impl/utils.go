package impl

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func rowExists(db *sqlx.DB, query string, args ...interface{}) (bool, error) {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := db.QueryRow(query, args...).Scan(&exists)
	return exists, err
}

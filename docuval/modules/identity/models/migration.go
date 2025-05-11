package models

import (
	"github.com/Villsource/Docuval/modules/shared/databases"
)

func Migrate(db *sqliteDB.SQLiteService) error {
	CreateAuthenticationTable(db)
	return nil
}
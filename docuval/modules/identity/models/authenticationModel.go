package models

import (
	"fmt"

	"github.com/Villsource/Docuval/modules/shared/databases"
)

type AuthenticationModel struct {
	Challenge       string `json:"code_challenge"`
	ChallengeMethod string `json:"code_challenge_method"`
	AuthorizationCode   string `json:"authorization_code"`
}

func CreateAuthenticationTable(db *sqliteDB.SQLiteService) error {
	query := `
	CREATE TABLE IF NOT EXISTS authentication (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code_challenge TEXT NOT NULL,
		code_challenge_method TEXT NOT NULL,
		authorization_code TEXT NOT NULL
	);`
	err := db.CreateTable(query)
	if err != nil {
		return fmt.Errorf("failed to create authentication table: %w", err)
	}
	return nil
}

package main

import (
	"github.com/Villsource/Docuval/modules/identity"

	sqliteDB "github.com/Villsource/Docuval/modules/shared/databases"
	"github.com/Villsource/Docuval/modules/identity/models"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	sqliteDBService, err := sqliteDB.NewSQLiteService("docuval.db")
	if err != nil {
		color.Red("[ERROR] Failed to initialize SQLite service: %v", err)
		return
	}

	err = models.Migrate(sqliteDBService)
	if err != nil {
		color.Red("[ERROR] Failed to migrate database: %v", err)
		return
	}
	color.Green("[INFO] Database migration completed successfully")


	app.Hooks().OnShutdown(func() error {
		color.Yellow("[INFO] Application is shutting down...")
		if err := sqliteDBService.Close(); err != nil {
			color.Red("[ERROR] Failed to close SQLite service: %v", err)
		} else {
			color.Green("[INFO] SQLite service closed successfully")
		}
		return nil
	})


	identity.RegisterOAuth(app)

	// Start the server
	color.Magenta("[INFO] Starting OAuth 2.0 server on :3000")
	app.Listen(":3000")
}

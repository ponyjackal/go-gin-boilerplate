package migrations

import (
	"github.com/ponyjackal/go-gin-boilerplate/internal/domain/models"

	"github.com/ponyjackal/go-gin-boilerplate/internal/adapters/database"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() {
	var migrationModels = []interface{}{&models.Example{}}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}

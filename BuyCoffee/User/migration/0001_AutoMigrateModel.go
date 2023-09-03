package migration

import (
	"BuyCoffee/User"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func init() {
	id := "2023-09-02-15-16"
	MigrationsList = append(MigrationsList, &gormigrate.Migration{
		ID: id,
		Migrate: func(db *gorm.DB) error {
			return errors.Wrap(db.AutoMigrate(
				&User.Profile{},
			), id)
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable(
				&User.Profile{},
			)
		},
	})

}

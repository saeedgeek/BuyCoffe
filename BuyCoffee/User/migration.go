package User

import (
	"BuyCoffee/User/migration"
	"BuyCoffee/db"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
)

func (u *User) Migration() error {
	err := errors.Wrap(gormigrate.New(db.DB, nil, migration.MigrationsList).Migrate(), "User Migrations")
	migration.MigrationsList = nil
	return err
}

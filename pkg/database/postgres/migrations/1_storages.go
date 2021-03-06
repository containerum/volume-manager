package migrations

import (
	"git.containerum.net/ch/volume-manager/pkg/models"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg/orm"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		if _, err := db.Exec( /* language=sql */ `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); err != nil {
			return err
		}

		if _, err := orm.CreateTable(db, &model.Storage{}, &orm.CreateTableOptions{IfNotExists: true, FKConstraints: true}); err != nil {
			return err
		}

		return nil
	}, func(db migrations.DB) error {
		if _, err := orm.DropTable(db, &model.Storage{}, &orm.DropTableOptions{IfExists: true}); err != nil {
			return err
		}

		_, err := db.Exec( /* language=sql */ `DROP EXTENSION IF EXISTS "uuid-ossp"`)
		return err
	})
}

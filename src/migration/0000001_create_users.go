package migration

import (
	global "github.com/Rawipass/chat-service/global_variable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var createUsersTableMigration = &Migration{
	Number: 1,
	Name:   "create users table",
	Forwards: func(db *gorm.DB) error {

		const dropSql = `
			DROP TABLE IF EXISTS "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_USERS + `" CASCADE;
			DROP EXTENSION IF EXISTS "uuid-ossp";
		`
		err := db.Exec(dropSql).Error
		if err != nil {
			return errors.Wrap(err, "unable to drop users table")
		}

		const sql = `
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
			CREATE TABLE "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_USERS + `" (
				"id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
   			    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
			);
			
		`

		err = db.Exec(sql).Error
		if err != nil {
			return errors.Wrap(err, "unable to create users table")
		}

		return nil
	},
}

func init() {
	Migrations = append(Migrations, createUsersTableMigration)
}

package migration

import (
	global "github.com/Rawipass/chat-service/global_variable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var createRoomsTableMigration = &Migration{
	Number: 2,
	Name:   "create rooms table",
	Forwards: func(db *gorm.DB) error {

		const dropSql = `
			DROP TABLE IF EXISTS "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_ROOMS + `" CASCADE;
		`
		err := db.Exec(dropSql).Error
		if err != nil {
			return errors.Wrap(err, "unable to drop rooms table")
		}

		const sql = `
			CREATE TABLE "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_ROOMS + `" (
			"id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    		"name" varchar NOT NULL,
   		    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
   			"updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
   			"updated_by" text NOT NULL
			);
			
		`

		err = db.Exec(sql).Error
		if err != nil {
			return errors.Wrap(err, "unable to create rooms table")
		}

		return nil
	},
}

func init() {
	Migrations = append(Migrations, createRoomsTableMigration)
}

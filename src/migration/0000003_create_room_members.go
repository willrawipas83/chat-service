package migration

import (
	global "github.com/Rawipass/chat-service/global_variable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var createRoomMembersTableMigration = &Migration{
	Number: 3,
	Name:   "create room members table",
	Forwards: func(db *gorm.DB) error {

		const dropSql = `
			DROP TABLE IF EXISTS "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_ROOM_MEMBERS + `" CASCADE;
		`
		err := db.Exec(dropSql).Error
		if err != nil {
			return errors.Wrap(err, "unable to drop room members table")
		}

		const sql = `
			CREATE TABLE "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_ROOM_MEMBERS + `" (
			 "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
   			 "user_id" uuid NOT NULL,
   			 "room_id" uuid NOT NULL,
   			 "joined_at" timestamp DEFAULT CURRENT_TIMESTAMP,
   			 "leaved_at" timestamp,
   			 FOREIGN KEY ("user_id") REFERENCES "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_USERS + `"("id"),
   			 FOREIGN KEY ("room_id") REFERENCES "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_ROOMS + `"("id")
			);
			
		`

		err = db.Exec(sql).Error
		if err != nil {
			return errors.Wrap(err, "unable to create room members table")
		}

		return nil
	},
}

func init() {
	Migrations = append(Migrations, createRoomMembersTableMigration)
}

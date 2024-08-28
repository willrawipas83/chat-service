package migration

import (
	global "github.com/Rawipass/chat-service/global_variable"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var createRoomMessagesTableMigration = &Migration{
	Number: 4,
	Name:   "create room messages table",
	Forwards: func(db *gorm.DB) error {

		const dropSql = `
			DROP TABLE IF EXISTS "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_ROOM_MESSAGES + `" CASCADE;
		`
		err := db.Exec(dropSql).Error
		if err != nil {
			return errors.Wrap(err, "unable to drop room messages table")
		}

		const sql = `
			CREATE TABLE "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_ROOM_MESSAGES + `" (
			  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    		  "room_id" uuid NOT NULL,
              "user_id" uuid NOT NULL,
              "message" text NOT NULL,
              "status" text NOT NULL,
              "mentioned_user_ids" text,
              "reply_messages_id" uuid,
              "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
              FOREIGN KEY ("user_id") REFERENCES "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_USERS + `"("id"),
              FOREIGN KEY ("room_id") REFERENCES "` + global.SCHEMA_NAME_PUBLIC + `"."` + global.TABLE_NAME_ROOMS + `"("id")
			);
			
		`

		err = db.Exec(sql).Error
		if err != nil {
			return errors.Wrap(err, "unable to create room messages table")
		}

		return nil
	},
}

func init() {
	Migrations = append(Migrations, createRoomMessagesTableMigration)
}

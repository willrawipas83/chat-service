package repository

import (
	"context"
	"database/sql"

	"github.com/Rawipass/chat-service/config"
	"github.com/Rawipass/chat-service/logger"
	"github.com/Rawipass/chat-service/models"
	"go.uber.org/zap"
)

type ChatRepository interface {
	CreateChatRoom(room models.Room) error
	JoinChatRoom(roomMember models.RoomMember) error
	LeaveChatRoom(roomMember models.RoomMember) error
	UnsentMessage(message models.Message) error
	ListMessagesOnChat(roomID string, page int, per_page int) ([]models.RoomMessage, error)
	ListMembersOnChat(roomID string) ([]models.RoomMember, error)
	SendMessage(message *models.RoomMessage) error
}

type chatRepository struct {
	DB        *sql.DB
	RequestId string
	UserId    string
	Logger    *zap.SugaredLogger
}

func NewChatRepository(requestId string, userId string) ChatRepository {
	repoObj := chatRepository{
		RequestId: requestId,
		UserId:    userId,
	}

	repoObj.Logger = logger.Logger.With(
		"request_id", requestId,
		"user_id", userId,
		"part", "usecase",
	)
	return &repoObj
}

func (r *chatRepository) CreateChatRoom(room models.Room) error {

	r.Logger.Infof("start create chatroom")
	query :=
		`INSERT INTO 
        rooms 
        (id, name, created_at, updated_at, updated_by) 
        VALUES 
        (uuid_generate_v4(), $1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $2)`
	_, err := config.DB.Exec(context.Background(), query, room.Name, room.UpdatedBy)
	if err != nil {
		r.Logger.Errorf("create chatroom failed cuz: %s", err)
		return err
	}
	return nil
}

func (r *chatRepository) JoinChatRoom(roomMember models.RoomMember) error {
	r.Logger.Infof("start join chatroom")
	query :=
		`INSERT INTO 
        room_members 
        (id, user_id, room_id, joined_at)
        VALUES 
        (uuid_generate_v4(), $1, $2, CURRENT_TIMESTAMP)`
	_, err := config.DB.Exec(context.Background(), query, roomMember.UserId, roomMember.RoomId)
	if err != nil {
		r.Logger.Errorf("join chatroom room id %d failed cuz: %s", roomMember.RoomId, err)
		return err
	}
	return nil
}

func (r *chatRepository) LeaveChatRoom(roomMember models.RoomMember) error {
	r.Logger.Infof("start leave chatroom")
	query :=
		`UPDATE room_members
         SET leaved_at = CURRENT_TIMESTAMP
         WHERE user_id = $1 and room_id = $2;
`
	_, err := config.DB.Exec(context.Background(), query, roomMember.UserId, roomMember.RoomId)
	if err != nil {
		r.Logger.Errorf("leave chatroom room id %d failed cuz: %s", roomMember.RoomId, err)
		return err
	}
	return nil
}

func (r *chatRepository) UnsentMessage(message models.Message) error {
	r.Logger.Infof("start unsent message")
	query :=
		`UPDATE room_messages
         SET status = 'unsent'
         WHERE id = $1;
        `
	_, err := config.DB.Exec(context.Background(), query, message.ID)
	if err != nil {
		r.Logger.Errorf("unsent message id %d failed cuz: %s", message.ID, err)
		return err
	}
	return nil
}

func (r *chatRepository) ListMessagesOnChat(roomID string, limit int, offset int) ([]models.RoomMessage, error) {
	r.Logger.Infof("start list message")
	cal_offset := (limit - 1) * offset
	query := `
        SELECT id, room_id, user_id, message, status, mentioned_user_ids, reply_messages_id, created_at
        FROM room_messages
        WHERE room_id = $1
        ORDER BY created_at ASC
        LIMIT $2 OFFSET $3
    `

	rows, err := config.DB.Query(context.Background(), query, roomID, limit, cal_offset)
	if err != nil {
		r.Logger.Errorf("list message failed cuz: %s", err)
	}
	defer rows.Close()

	var messages []models.RoomMessage
	for rows.Next() {
		var msg models.RoomMessage
		if err := rows.Scan(&msg.ID, &msg.RoomID, &msg.UserID, &msg.Message, &msg.Status, &msg.MentionedUserIDs, &msg.ReplyMessagesID, &msg.CreatedAt); err != nil {
			r.Logger.Errorf("scan list message failed cuz: %s", err)
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		r.Logger.Errorf("get rows in list message failed cuz: %s", err)
	}

	return messages, nil
}

func (r *chatRepository) ListMembersOnChat(roomID string) ([]models.RoomMember, error) {
	r.Logger.Infof("start list member")
	query := `
   	SELECT id, room_id, user_id, joined_at, leaved_at FROM room_members WHERE room_id = $1
    `
	rows, err := config.DB.Query(context.Background(), query, roomID)
	if err != nil {
		r.Logger.Errorf("list member failed cuz: %s", err)
	}
	defer rows.Close()

	var members []models.RoomMember
	for rows.Next() {
		var member models.RoomMember
		if err := rows.Scan(&member.ID, &member.RoomId, &member.UserId, &member.JoinedAt, &member.LeavedAt); err != nil {
			r.Logger.Errorf("scan member failed cuz: %s", err)
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		r.Logger.Errorf("get rows in list member failed cuz: %s", err)
	}

	return members, nil
}

func (r *chatRepository) SendMessage(message *models.RoomMessage) error {
	r.Logger.Infof("start send message")
	query := `
    INSERT INTO room_messages 
	(id, room_id, user_id, message, status, reply_messages_id, mentioned_user_ids, created_at)
    VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)
    `
	_, err := config.DB.Exec(context.Background(), query, message.RoomID, message.UserID, message.Message, message.Status, message.ReplyMessagesID, message.MentionedUserIDs)
	if err != nil {
		r.Logger.Errorf("send message failed cuz: %s", err)
		return err
	}
	return nil
}

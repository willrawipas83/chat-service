package usecase

import (
	"github.com/Rawipass/chat-service/custom_error"
	"github.com/Rawipass/chat-service/internal/repository"
	"github.com/Rawipass/chat-service/logger"
	"github.com/Rawipass/chat-service/models"
	"go.uber.org/zap"
)

type ChatUseCase struct {
	repo      repository.ChatRepository
	RequestId string
	UserId    string
	Logger    *zap.SugaredLogger
}

func NewChatUseCase(requestId string, userId string) *ChatUseCase {
	usecaseObj := ChatUseCase{
		RequestId: requestId,
		UserId:    userId,
	}

	usecaseObj.Logger = logger.Logger.With(
		"request_id", requestId,
		"user_id", userId,
		"part", "usecase",
	)
	return &usecaseObj
}

type ListMessagesOnChatRequest struct {
	RoomID string `json:"room_id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func (uc *ChatUseCase) ListMessagesOnChat(req ListMessagesOnChatRequest) ([]models.RoomMessage, error) {
	uc.Logger.Info("start list message")
	uc.Logger.Debugf("list message input data: %v", req)
	repo := repository.NewChatRepository(uc.RequestId, uc.UserId)
	messages, err := repo.ListMessagesOnChat(req.RoomID, req.Limit, req.Offset)
	if err != nil {
		uc.Logger.Errorf("list messsage failed cuz: %s", err.Error())
		customError := custom_error.CustomError{
			Code:    custom_error.DatabaseError,
			Message: err.Error(),
		}
		return messages, customError
	}
	return messages, nil
}

func (uc *ChatUseCase) ListMembersOnChat(req ListMessagesOnChatRequest) ([]models.RoomMember, error) {
	uc.Logger.Info("start list member")
	uc.Logger.Debugf("list member input data: %v", req)
	repo := repository.NewChatRepository(uc.RequestId, uc.UserId)
	members, err := repo.ListMembersOnChat(req.RoomID)
	if err != nil {
		uc.Logger.Errorf("list member failed cuz: %s", err.Error())
		customError := custom_error.CustomError{
			Code:    custom_error.DatabaseError,
			Message: err.Error(),
		}
		return members, customError
	}
	return members, nil
}

func (uc *ChatUseCase) CreateChatRoom(room models.Room) error {
	uc.Logger.Info("start create chatroom")
	uc.Logger.Debugf("create chatroom input data: %v", room)
	repo := repository.NewChatRepository(uc.RequestId, uc.UserId)
	err := repo.CreateChatRoom(room)
	if err != nil {
		uc.Logger.Errorf("create chatroom failed cuz: %s", err.Error())
		customError := custom_error.CustomError{
			Code:    custom_error.DatabaseError,
			Message: err.Error(),
		}
		return customError
	}
	return nil
}

func (uc *ChatUseCase) JoinChatRoom(roomMember models.RoomMember) error {
	uc.Logger.Info("start join chatroom")
	uc.Logger.Debugf("join chatroom input data: %v", roomMember)
	err := uc.repo.JoinChatRoom(roomMember)
	if err != nil {
		uc.Logger.Errorf("join chatroom failed cuz: %s", err.Error())
		customError := custom_error.CustomError{
			Code:    custom_error.DatabaseError,
			Message: err.Error(),
		}
		return customError
	}
	return nil
}

func (uc *ChatUseCase) LeaveChatRoom(roomMember models.RoomMember) error {
	uc.Logger.Info("start leave chatroom")
	uc.Logger.Debugf("leave chatroom input data: %v", roomMember)
	repo := repository.NewChatRepository(uc.RequestId, uc.UserId)
	err := repo.LeaveChatRoom(roomMember)
	if err != nil {
		uc.Logger.Errorf("leave chatroom failed cuz: %s", err.Error())
		customError := custom_error.CustomError{
			Code:    custom_error.DatabaseError,
			Message: err.Error(),
		}
		return customError
	}
	return nil
}

func (uc *ChatUseCase) UnsentMessage(message models.Message) error {
	uc.Logger.Info("start unsent message")
	uc.Logger.Debugf("unsent message input data: %v", message)
	repo := repository.NewChatRepository(uc.RequestId, uc.UserId)
	err := repo.UnsentMessage(message)
	if err != nil {
		uc.Logger.Errorf("unsent message failed cuz: %s", err.Error())
		customError := custom_error.CustomError{
			Code:    custom_error.DatabaseError,
			Message: err.Error(),
		}
		return customError
	}
	return nil
}

func (uc *ChatUseCase) SendMessage(message models.RoomMessage) error {
	uc.Logger.Info("start send message")
	uc.Logger.Debugf("send message input data: %v", message)
	repo := repository.NewChatRepository(uc.RequestId, uc.UserId)
	err := repo.SendMessage(&message)
	if err != nil {
		uc.Logger.Errorf("send message failed cuz: %s", err.Error())
		customError := custom_error.CustomError{
			Code:    custom_error.DatabaseError,
			Message: err.Error(),
		}
		return customError
	}
	return nil
}

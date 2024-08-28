package routes

import (
	"github.com/Rawipass/chat-service/internal/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(http.Logger())
	router.Use(http.SetRequestId())
	//chat rooms
	router.POST("/chatrooms/create", http.CreateChatRoomHandler)
	router.POST("/chatroom/join", http.JoinChatRoomHandler)
	router.PUT("/chatrooms/leave", http.LeaveChatRoomHandler)
	// message
	router.PUT("/message/unsent", http.UnsentMessageHandler)
	router.GET("/message/list", http.ListMessagesOnChatHandler)
	router.POST("/message/send", http.SendMessage)
	//member
	router.GET("/member/list", http.ListMembersOnChatHandler)

	return router
}

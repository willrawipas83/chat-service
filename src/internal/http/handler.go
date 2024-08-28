package http

import (
	"log"
	"net/http"
	"sync"

	global "github.com/Rawipass/chat-service/global_variable"
	"github.com/Rawipass/chat-service/internal/usecase"
	"github.com/Rawipass/chat-service/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func CreateChatRoomHandler(c *gin.Context) {
	contextLogger, exists := c.Get(global.KEY_LOGGER)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logger not found"})
		return
	}
	apiLogger := contextLogger.(*zap.SugaredLogger)
	apiLogger.Info("call handler to create chatroom")
	var req models.Room
	if err := c.ShouldBindJSON(&req); err != nil {
		apiLogger.Errorf("could not bind json body to create chatroom cuz: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestId := c.GetString(global.KEY_REQUEST_ID)
	userId := c.GetString(global.KEY_USER_ID)

	uc := usecase.NewChatUseCase(requestId, userId)
	err := uc.CreateChatRoom(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "chat room created"})
}

func JoinChatRoomHandler(c *gin.Context) {
	contextLogger, exists := c.Get(global.KEY_LOGGER)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logger not found"})
		return
	}
	apiLogger := contextLogger.(*zap.SugaredLogger)
	apiLogger.Info("call handler to join chatroom")
	var req models.RoomMember
	if err := c.ShouldBindJSON(&req); err != nil {
		apiLogger.Errorf("could not bind json body to join chatroom cuz: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestId := c.GetString(global.KEY_REQUEST_ID)
	userId := c.GetString(global.KEY_USER_ID)
	uc := usecase.NewChatUseCase(requestId, userId)
	err := uc.JoinChatRoom(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "join chat room success"})
}

func LeaveChatRoomHandler(c *gin.Context) {
	contextLogger, exists := c.Get(global.KEY_LOGGER)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logger not found"})
		return
	}
	apiLogger := contextLogger.(*zap.SugaredLogger)
	apiLogger.Info("call handler to leave chatroom")
	var req models.RoomMember
	if err := c.ShouldBindJSON(&req); err != nil {
		apiLogger.Errorf("could not bind json body to leave chatroom cuz: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestId := c.GetString(global.KEY_REQUEST_ID)
	userId := c.GetString(global.KEY_USER_ID)
	uc := usecase.NewChatUseCase(requestId, userId)
	err := uc.LeaveChatRoom(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "leave chat room success"})
}

func UnsentMessageHandler(c *gin.Context) {
	contextLogger, exists := c.Get(global.KEY_LOGGER)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logger not found"})
		return
	}
	apiLogger := contextLogger.(*zap.SugaredLogger)
	apiLogger.Info("call handler to unsent message")
	var req models.Message
	if err := c.ShouldBindJSON(&req); err != nil {
		apiLogger.Errorf("could not bind json body to unsent message cuz: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestId := c.GetString(global.KEY_REQUEST_ID)
	userId := c.GetString(global.KEY_USER_ID)
	uc := usecase.NewChatUseCase(requestId, userId)
	err := uc.UnsentMessage(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "unsent message success"})
}

func ListMessagesOnChatHandler(c *gin.Context) {
	contextLogger, exists := c.Get(global.KEY_LOGGER)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logger not found"})
		return
	}
	apiLogger := contextLogger.(*zap.SugaredLogger)
	apiLogger.Info("call handler to list message")
	var req usecase.ListMessagesOnChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiLogger.Errorf("could not bind json body to list message cuz: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestId := c.GetString(global.KEY_REQUEST_ID)
	userId := c.GetString(global.KEY_USER_ID)
	uc := usecase.NewChatUseCase(requestId, userId)
	messages, err := uc.ListMessagesOnChat(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func ListMembersOnChatHandler(c *gin.Context) {
	contextLogger, exists := c.Get(global.KEY_LOGGER)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logger not found"})
		return
	}
	apiLogger := contextLogger.(*zap.SugaredLogger)
	apiLogger.Info("call handler to list member")
	var req usecase.ListMessagesOnChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiLogger.Errorf("could not bind json body to list member cuz: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestId := c.GetString(global.KEY_REQUEST_ID)
	userId := c.GetString(global.KEY_USER_ID)
	uc := usecase.NewChatUseCase(requestId, userId)
	members, err := uc.ListMembersOnChat(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": members})
}

func SendMessage(c *gin.Context) {
	contextLogger, exists := c.Get(global.KEY_LOGGER)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logger not found"})
		return
	}
	apiLogger := contextLogger.(*zap.SugaredLogger)
	apiLogger.Info("call handler to send message")
	var req models.RoomMessage

	if err := c.ShouldBindJSON(&req); err != nil {
		apiLogger.Errorf("could not bind json body to send message cuz: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestId := c.GetString(global.KEY_REQUEST_ID)
	userId := c.GetString(global.KEY_USER_ID)
	uc := usecase.NewChatUseCase(requestId, userId)
	err := uc.SendMessage(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message sent"})
}

type Message struct {
	UserID    string `json:"user_id"`
	RoomID    string `json:"room_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

// Notification struct for notifications
type Notification struct {
	Type  string      `json:"type"`
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

// Map เพื่อเก็บการเชื่อมต่อแยกตามห้อง
var rooms = make(map[string]map[*websocket.Conn]bool)
var roomsMutex = sync.Mutex{}
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// อ่าน RoomID จาก WebSocket query parameters
	roomID := r.URL.Query().Get(viper.GetString("Http.Room_id"))
	if roomID == "" {
		log.Println("room_id is required")
		ws.Close()
		return
	}
	// อ่าน UserID จาก WebSocket query parameters
	userID := r.URL.Query().Get(viper.GetString("Http.User_id"))
	if userID == "" {
		log.Println("room_id is required")
		ws.Close()
		return
	}

	roomsMutex.Lock()
	if rooms[roomID] == nil {
		rooms[roomID] = make(map[*websocket.Conn]bool)
	}
	rooms[roomID][ws] = true
	roomsMutex.Unlock()
	var msg Message
	// ส่ง notification ว่ามีผู้ใช้ใหม่เข้าห้อง
	sendNotificationToRoom(roomID, "user_joined", map[string]string{"user_id": userID})

	for {
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			roomsMutex.Lock()
			delete(rooms[roomID], ws)
			roomsMutex.Unlock()
			break
		}
		// ส่งข้อความและ notification
		broadcastToRoom(roomID, msg)
		sendNotificationToRoom(roomID, "new_message", map[string]string{"message": msg.Message, "user_id": msg.UserID})
	}
}

func broadcastToRoom(roomID string, msg Message) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	// ส่งข้อความไปยังการเชื่อมต่อในห้องที่กำหนดเท่านั้น
	for client := range rooms[roomID] {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(rooms[roomID], client)
		}
	}
}

func sendNotificationToRoom(roomID, event string, data interface{}) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	notification := Notification{
		Type:  "notification",
		Event: event,
		Data:  data,
	}

	// ส่ง notification ไปยังผู้ใช้ทุกคนในห้องที่กำหนด
	for client := range rooms[roomID] {
		err := client.WriteJSON(notification)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(rooms[roomID], client)
		}
	}
}

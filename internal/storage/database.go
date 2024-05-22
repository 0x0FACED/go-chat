package storage

import (
	"go-chat/internal/models"
)

type Database interface {
	Connect() error
	Disconnect() error
	GetConnectionString() string

	UserDB
	MessageDB
	ChatDB
}

type UserDB interface {
	Login(u *models.User) (*models.User, error)
	Register(u *models.User) (*models.User, error)
	GetUserIDByUsername(username string) (int, error)
	GetUserByID(id int) (*models.User, error)
}

type MessageDB interface {
	SaveMessages(mes []models.Message) error
	SaveMessage(mes *models.Message) (*models.Message, error)
	GetMessageByID(id int) (*models.Message, error)
}

type ChatDB interface {
	CreateChat(firstUserID int, secondUserID int) (int, error)
	FindChatByUserIDs(firstUserID int, secondUserID int) (int, error)
	GetChatHistory(senderID int, recipientID int) ([]models.Message, error)
}

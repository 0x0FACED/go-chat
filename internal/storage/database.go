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
}

type UserDB interface {
	Login(u *models.User) (*models.User, error)
	Register(u *models.User) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}

type MessageDB interface {
	SaveMessages(mes []models.Message) error
	GetMessageByID(id int) (*models.Message, error)
	GetChatHistory(senderID int, recipientID int) ([]models.Message, error)
}

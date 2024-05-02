package redis

import "go-chat/internal/models"

type Repository interface {
	SaveMessage(message *models.Message) error
	GetMessages(senderID int, recipientID int) (*[]models.Message, error)
	DeleteMessage(id int) error
}

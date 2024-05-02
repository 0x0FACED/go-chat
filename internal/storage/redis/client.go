package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-chat/internal/models"
)

type Client struct {
	Redis *redis.Client
}

func (r *Client) SaveMessage(message *models.Message) error {
	ctx := context.Background()
	key := fmt.Sprintf("messages:%d", message.ID)

	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = r.Redis.Set(ctx, key, jsonData, 0).Result()
	return err
}

func (r *Client) GetMessages(senderID int, recipientID int) ([]*models.Message, error) {
	ctx := context.Background()
	key := fmt.Sprintf("messages:%d:%d", senderID, recipientID)

	results, err := r.Redis.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var messages []*models.Message
	for _, data := range results {
		var message models.Message
		err := json.Unmarshal([]byte(data), &message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

func (r *Client) DeleteMessage(id int) error {
	ctx := context.Background()
	key := fmt.Sprintf("messages:%d", id)

	_, err := r.Redis.Del(ctx, key).Result()
	return err
}

package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-chat/config"
	"go-chat/internal/models"
	"strconv"
)

type Client struct {
	Redis *redis.Client
}

func New(cfg *config.RedisConfig) *Client {
	r := redis.NewClient(&redis.Options{
		Addr:        cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Network:     cfg.Network,
		Username:    cfg.Username,
		Password:    cfg.Password,
		DialTimeout: cfg.DialTimeout,
		MaxRetries:  cfg.MaxRetries,
	})
	return &Client{
		Redis: r,
	}
}

func (c *Client) SaveMessage(message *models.Message) error {
	ctx := context.Background()
	key := fmt.Sprintf("messages:%d", message.ID)

	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = c.Redis.Set(ctx, key, jsonData, 0).Result()
	return err
}

func (c *Client) GetMessages(senderID int, recipientID int) ([]*models.Message, error) {
	ctx := context.Background()
	key := fmt.Sprintf("messages:%d:%d", senderID, recipientID)

	results, err := c.Redis.LRange(ctx, key, 0, -1).Result()
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

func (c *Client) DeleteMessage(id int) error {
	ctx := context.Background()
	key := fmt.Sprintf("messages:%d", id)

	_, err := c.Redis.Del(ctx, key).Result()
	return err
}

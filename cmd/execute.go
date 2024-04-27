package cmd

import "go-chat/internal/http"

func Execute() error {
	s := http.NewServer()
	err := s.PrepareServer()
	if err != nil {
		return err
	}
	s.StartServer()
	return nil
}

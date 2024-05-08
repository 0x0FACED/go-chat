package cmd

import (
	"go-chat/config"
	"go-chat/internal/http"
	"log"
)

func Execute() {
	cfg, err := config.Load()
	s := http.NewServer(*cfg)
	err = s.StartServer()
	if err != nil {
		log.Fatalln("cant start server:", err)
	}
}

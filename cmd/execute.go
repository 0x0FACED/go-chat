package cmd

import (
	"go-chat/config"
	"go-chat/internal/http"
	"log"
)

func Execute() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("cant load config, exit: ", err)
	}
	s := http.NewServer(*cfg)
	err = s.StartServer()
	if err != nil {
		log.Fatalln("cant start server:", err)
	}
}

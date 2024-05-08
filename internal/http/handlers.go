package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) prepareRoutes() {
	s.r.Handle(http.MethodPost, "/register", s.handleRegister)
	s.r.Handle(http.MethodPost, "/login", s.handleLogin)
	s.r.Handle(http.MethodPost, "/send_message", s.handleSendMessage)
	s.r.Handle(http.MethodGet, "/get_messages", s.handleGetMessages)
}

func (s *Server) handleRegister(ctx *gin.Context) {
	// ...
}

func (s *Server) handleLogin(ctx *gin.Context) {
	// ...
}

func (s *Server) handleSendMessage(ctx *gin.Context) {
	// ...
}

func (s *Server) handleGetMessages(ctx *gin.Context) {
	// ...
}

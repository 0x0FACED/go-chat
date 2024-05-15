package http

import (
	"go-chat/internal/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) prepareRoutes() {
	s.r.Handle(http.MethodPost, "/register", s.handleRegister)
	s.r.Handle(http.MethodPost, "/login", s.handleLogin)
	s.r.Handle(http.MethodPost, "/send_message", s.handleSendMessage)
	s.r.Handle(http.MethodGet, "/get_messages", s.handleGetMessages)
}

func (s *Server) handleRegister(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	u, err := s.db.Register(&newUser)
	if err != nil {
		s.logger.Errorln("cant register u:", u, ", err:", err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"mes": "registered"})
}

func (s *Server) handleLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	token := session.Get("token")
	s.logger.Println("Token: ", token)
	if token != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "already authorized", "token": token})
		return
	}

	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	u, err := s.db.Login(&user)
	if err != nil {
		s.logger.Errorln("cant login:", err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"err": err.Error()})
		return
	}
	uuid := uuid.NewString()
	session.Set("token", uuid)
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{"token": session.Get("token"), "username": u.Username})
}

func (s *Server) handleSendMessage(ctx *gin.Context) {
	var mes models.Message
	if err := ctx.BindJSON(&mes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	savedMes, err := s.db.SaveMessage(&mes)
	if err != nil {
		s.logger.Println("err save mes:", err)
		return
	}

	err = s.redis.SaveMessage(savedMes)
	if err != nil {
		s.logger.Println("err save mes:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"your_mes": savedMes})
}

func (s *Server) handleGetMessages(ctx *gin.Context) {
	// ...
}

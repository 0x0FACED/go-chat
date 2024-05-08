package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"go-chat/config"
	"go-chat/internal/storage"
	"go-chat/internal/storage/redis"
	"go-chat/internal/utils"
	"net/http"
)

type Server struct {
	r      *gin.Engine
	db     storage.Database
	config config.Config
	logger *logrus.Logger
	redis  *redis.Client
}

func NewServer(cfg config.Config) *Server {
	return &Server{
		r:      gin.Default(),
		logger: logrus.New(),
		config: cfg,
		redis:  redis.NewRedis(&cfg.Redis),
	}
}

func (s *Server) PrepareServer() error {
	err := s.initDatabase()
	if err != nil {
		s.logger.Fatalln("Cannot InitDatabase(): ", err)
		return err
	}
	s.prepareRoutes()
	return nil
}

func (s *Server) initDatabase() error {
	db := &storage.Postgres{
		Config: s.config.Database,
	}
	err := db.Connect()
	if err != nil {
		s.logger.Errorln("Failed to connect to database:", err)
		return err
	}
	s.db = db
	return nil
}

func (s *Server) prepareRoutes() {
	r := gin.Default()
	r.Handle(http.MethodPost, "/register", s.handleRegister)
	r.Handle(http.MethodPost, "/login", s.handleLogin)
	r.Handle(http.MethodPost, "/send_message", s.handleSendMessage)
	r.Handle(http.MethodGet, "/get_messages", s.handleGetMessages)
	s.r = r
}

func (s *Server) StartServer() {
	// ...
}

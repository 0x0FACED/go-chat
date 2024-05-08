package http

import (
	"errors"
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

func (s *Server) prepareServer() error {
	err := s.initDatabase()
	if err != nil {
		s.logger.Fatalln("Cannot InitDatabase(): ", err)
		return err
	}
	s.prepareRoutes()
	return nil
}

func (s *Server) initDatabase() error {
	var db storage.Database
	switch s.config.Database.DriverName {
	case "postgres":
		db = &storage.Postgres{Config: s.config.Database}
	case "mysql":
		db = &storage.MySQL{Config: s.config.Database}
	default:
		return errors.New(utils.ErrDriverName)
	}
	err := db.Connect()
	if err != nil {
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

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
	"strconv"
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
	s.r.Handle(http.MethodPost, "/register", s.handleRegister)
	s.r.Handle(http.MethodPost, "/login", s.handleLogin)
	s.r.Handle(http.MethodPost, "/send_message", s.handleSendMessage)
	s.r.Handle(http.MethodGet, "/get_messages", s.handleGetMessages)
}

func (s *Server) StartServer() error {
	err := s.prepareServer()
	if err != nil {
		return err
	}
	addr := s.config.Server.Host + strconv.Itoa(s.config.Server.Port)
	err = s.r.Run(addr)
	if err != nil {
		return err
	}
	return nil
}

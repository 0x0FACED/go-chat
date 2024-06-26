package http

import (
	"errors"
	"go-chat/config"
	"go-chat/internal/storage"
	"go-chat/internal/storage/redis"
	"go-chat/internal/utils"
	"go-chat/migrations"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Server struct {
	r         *gin.Engine
	db        storage.Database
	config    config.Config
	logger    *logrus.Logger
	redis     *redis.Client
	clients   map[*websocket.Conn]bool
	broadcast chan models.Message
	upgrader  websocket.Upgrader
}

func NewServer(cfg config.Config) *Server {
	return &Server{
		r:      gin.Default(),
		logger: logrus.New(),
		config: cfg,
		redis:  redis.New(&cfg.Redis),
	}
}

func (s *Server) prepareServer() error {
	s.prepareCookiesStore()
	err := s.initDatabase()
	if err != nil {
		s.logger.Fatalln("Cannot InitDatabase(): ", err)
		return err
	}
	migrations.Up(s.db.GetConnectionString())
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

func (s *Server) prepareCookiesStore() {
	store := cookie.NewStore([]byte(s.config.Server.SessionKey))
	s.r.Use(sessions.Sessions(s.config.Server.StoreName, store))
}

func (s *Server) StartServer() error {
	err := s.prepareServer()
	if err != nil {
		return err
	}
	addr := s.config.Server.Host + ":" + strconv.Itoa(s.config.Server.Port)
	err = s.r.Run(addr)
	if err != nil {
		return err
	}
	return nil
}

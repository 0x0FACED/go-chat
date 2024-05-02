package storage

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-chat/config"
	"go-chat/internal/models"
)

type MySQL struct {
	db     *sql.DB
	Config config.DatabaseConfig
	logger *logrus.Logger
}

func (m *MySQL) Connect() error {
	m.logger = logrus.New()
	db, err := sql.Open(m.Config.DriverName, m.GetConnectionString())
	if err != nil {
		m.logger.Fatalln("Postgres: Cannot Open() database with driver name ", m.Config.DriverName, ", connection string ", m.GetConnectionString())
		return err
	}
	err = db.Ping()
	if err != nil {
		m.logger.Errorln("Postgres: Cannot Ping() database with driver name ", m.Config.DriverName, ", connection string ", m.GetConnectionString())
		return err
	}
	m.db = db
	return nil
}

func (m *MySQL) Disconnect() error {
	//TODO implement me
	panic("implement me")
}

func (m *MySQL) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		m.Config.Host, m.Config.Port, m.Config.Username, m.Config.Password, m.Config.DBName)
}

func (m *MySQL) Login(u *models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySQL) Register(u *models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySQL) SaveMessages(mes *models.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m *MySQL) GetMessageByID(id int) (*models.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySQL) GetChatHistory(senderID int, recipientID int) ([]models.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySQL) GetUserByUsername(username string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MySQL) GetUserByID(id int) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

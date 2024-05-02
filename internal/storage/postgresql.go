package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-chat/config"
	"go-chat/internal/models"
	"go-chat/internal/utils"
)

type Postgres struct {
	db     *sql.DB
	Config config.DatabaseConfig
	logger *logrus.Logger
}

func (p *Postgres) Connect() error {
	p.logger = logrus.New()
	db, err := sql.Open(p.Config.DriverName, p.GetConnectionString())
	if err != nil {
		p.logger.Fatalln("Postgres: Cannot Open() database with driver name ", p.Config.DriverName, ", connection string ", p.GetConnectionString())
		return err
	}
	err = db.Ping()
	if err != nil {
		p.logger.Errorln("Postgres: Cannot Ping() database with driver name ", p.Config.DriverName, ", connection string ", p.GetConnectionString())
		return err
	}
	p.db = db
	return nil
}

func (p *Postgres) Disconnect() error {
	if p.db != nil {
		err := p.db.Close()
		if err != nil {
			p.logger.Fatalln("Postgres: Cannot Close() database: ", err)
			return err
		}
	}
	return nil
}

func (p *Postgres) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.Config.Host, p.Config.Port, p.Config.Username, p.Config.Password, p.Config.DBName)
}

func (p *Postgres) Login(u *models.User) (*models.User, error) {
	err := utils.ValidateUser(u)
	if err != nil {
		return nil, err
	}

	hashedPass := u.HashPassword(u.Password)

	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(utils.QueryLoginTx)
	if err != nil {
		return nil, err
	}

	var id int
	var username string
	err = stmt.QueryRow(u.Username, hashedPass).Scan(&id, &username)
	if err != nil {
		if tx.Rollback() != nil {
			return nil, errors.New(utils.ErrRollbackTx + err.Error())
		}
		return nil, errors.New(utils.ErrIncorrectUsernameOrPass)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New(utils.ErrCommitTx + err.Error())
	}
	return u, nil
}

func (p *Postgres) Register(u *models.User) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) SaveMessages(mes *models.Message) error {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) GetMessageByID(id int) (*models.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) GetChatHistory(senderID int, recipientID int) ([]models.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) GetUserByUsername(username string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) GetUserByID(id int) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

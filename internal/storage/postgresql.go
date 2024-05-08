package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-chat/config"
	"go-chat/internal/models"
	"go-chat/internal/utils"
	"time"
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
		return errors.New(utils.ErrOpenDB + err.Error())
	}
	err = db.Ping()
	if err != nil {
		return errors.New(utils.ErrPingDB + err.Error())
	}
	p.db = db
	return nil
}

func (p *Postgres) Disconnect() error {
	if p.db != nil {
		err := p.db.Close()
		if err != nil {
			return errors.New(utils.ErrCloseDB + err.Error())
		}
	}
	return nil
}

func (p *Postgres) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.Config.Username, p.Config.Password, p.Config.Host, p.Config.Port, p.Config.DBName)
}

func (p *Postgres) Login(u *models.User) (*models.User, error) {
	err := utils.ValidateUser(u)
	if err != nil {
		return nil, err
	}

	hashedPass := u.HashPassword(u.Password)

	tx, err := p.db.Begin()
	if err != nil {
		return nil, errors.New(utils.ErrBeginTx + err.Error())
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(utils.QueryLoginTx)
	if err != nil {
		return nil, errors.New(utils.ErrPrepareTx + err.Error())
	}

	var id int
	var name string
	var username string
	var regDate time.Time
	var desc string
	err = stmt.QueryRow(u.Username, hashedPass).Scan(&id, &name, &username, &regDate, &desc)
	if err != nil {
		return nil, errors.New(utils.ErrIncorrectUsernameOrPass)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New(utils.ErrCommitTx + err.Error())
	}
	u.Description = desc
	u.RegistrationDate = regDate
	return u, nil
}

func (p *Postgres) Register(u *models.User) (*models.User, error) {
	if err := utils.ValidateUser(u); err != nil {
		return nil, err
	}

	if _, err := p.GetUserByUsername(u.Username); err != nil {
		return nil, err
	}

	hashedPass := u.HashPassword(u.Password)
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(utils.QueryRegisterTx)
	if err != nil {
		return nil, errors.New(utils.ErrPrepareTx + err.Error())
	}

	_, err = stmt.Exec(u.Name, u.Username, hashedPass, u.Description)
	if err != nil {
		return nil, errors.New(utils.ErrExecQueryTx + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New(utils.ErrCommitTx + err.Error())
	}

	return u, nil
}

func (p *Postgres) SaveMessages(mes []models.Message) error {
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

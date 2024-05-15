package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"go-chat/config"
	"go-chat/internal/models"
	"go-chat/internal/utils"
	"log"
	"time"

	"github.com/sirupsen/logrus"
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

	tx, err := p.db.Begin()
	if err != nil {
		return nil, errors.New(utils.ErrBeginTx + err.Error())
	}
	defer tx.Rollback()

	stmtPass, err := tx.Prepare(utils.QueryGetPasswordTx)
	if err != nil {
		return nil, errors.New(utils.ErrPrepareTx + err.Error())
	}

	var hashedPass string
	err = stmtPass.QueryRow(u.Username).Scan(&hashedPass)
	if err != nil {
		return nil, errors.New(utils.ErrExecQueryTx + err.Error())
	}

	err = u.CompareHashPassword(u.Password, hashedPass)
	if err != nil {
		return nil, errors.New(utils.ErrIncorrectUsernameOrPass)
	}

	stmtLogin, err := tx.Prepare(utils.QueryLoginTx)
	if err != nil {
		return nil, errors.New(utils.ErrPrepareTx + err.Error())
	}

	var id int
	var name string
	var username string
	var regDate time.Time
	var desc string
	err = stmtLogin.QueryRow(u.Username, hashedPass).Scan(&id, &name, &username, &regDate, &desc)
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

	/* 	if _, err := p.GetUserByUsername(u.Username); err != nil {
		return nil, err
	} */

	hashedPass, err := u.HashPassword(u.Password)
	fmt.Println("Hashed pass register:", hashedPass)
	if err != nil {
		return nil, err
	}
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

func (p *Postgres) SaveMessage(mes *models.Message) (*models.Message, error) {
	chatID, err := p.FindChatByUserIDs(mes.SenderID, mes.ReceiverID)
	if err != nil && err != sql.ErrNoRows {
		log.Println("FindChat")
		return nil, err
	}

	if chatID == -1 {
		chatID, err = p.CreateChat(mes.SenderID, mes.ReceiverID)
		if err != nil {
			log.Println("CreateChat")
			return nil, err
		}
	}

	mes.ChatID = chatID

	tx, err := p.db.Begin()
	if err != nil {
		log.Println("Begin")
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(utils.QuerySaveMessageTx)
	if err != nil {
		log.Println("Prepare")
		return nil, err
	}

	var id int
	var ts time.Time
	err = stmt.QueryRow(mes.SenderID, mes.ReceiverID, mes.Text, mes.ChatID).Scan(&id, &ts)
	if err != nil {
		log.Println("Query")
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("Commit")
		return nil, err
	}

	mes.ID = id
	mes.Timestamp = ts

	return mes, nil
}

func (p *Postgres) GetMessageByID(id int) (*models.Message, error) {
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

func (p *Postgres) CreateChat(firstUserID int, secondUserID int) (int, error) {
	var chatID int
	err := p.db.QueryRow(utils.QueryCreateChatTx, firstUserID, secondUserID).Scan(&chatID)
	if err != nil {
		return -1, err
	}
	return chatID, nil
}

func (p *Postgres) FindChatByUserIDs(firstUserID int, secondUserID int) (int, error) {
	var chatID int
	err := p.db.QueryRow(utils.QueryFindChatTx, firstUserID, secondUserID).Scan(&chatID)
	if err != nil {
		return -1, err
	}
	return chatID, nil
}

func (p *Postgres) GetChatHistory(senderID int, recipientID int) ([]models.Message, error) {
	//TODO implement me
	panic("implement me")
}

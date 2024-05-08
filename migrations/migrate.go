package migrations

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

func Up(dbURL string) {
	m, err := migrate.New(
		"file://./migrations/",
		dbURL)
	if err != nil {
		logrus.Errorln("cant create migration:", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logrus.Errorln("cant migrate:", err)
	}
}

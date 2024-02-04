package storage

import (
	"db-service-homework-kodzimo/internal/config"
	"db-service-homework-kodzimo/internal/storage/postgres"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Storage struct {
	Logger   *log.Logger
	Postgres *postgres.PostgresDB
}

func NewStorage(cfg *config.Config, log *log.Logger) (*Storage, error) {
	// Создаётся новый сторедж
	res := &Storage{}
	// Создаётся новое подключение
	db, err := postgres.NewPostgres(cfg, log)
	if err != nil {
		return res, errors.WithStack(err)
	}

	res.Logger = log
	res.Postgres = db

	return res, nil
}

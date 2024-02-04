package postgres

import (
	"db-service-homework-kodzimo/pkg/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type PostgresDB struct {
	Logger *log.Logger
	// Само подключение к базе данных
	DB *sqlx.DB
}

// PostgresOperator Хорошим тоном является подключение к базе не методом, а интерфесом
type PostgresOperator interface {
	//GenerateStudents(n int) error
}

const DefaultSchema = "blog"

func NewPostgres(cfg *config.Config, logger *log.Logger) (*PostgresDB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	ommitesStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s ",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, "[ommited]", cfg.DBName)

	// Принтим connectionString без пароля
	log.Info("current connection string: ", connectionString)

	// Создаём подключение к бд. Одно соединение менеджит множество подключений к бд в горутинах
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		logger.Errorf("could not establish connection to %s %s", ommitesStr, err)
		return nil, errors.WithStack(err)
	}

	//https://www.alexedwards.net/blog/configuring-sqldb
	//db.SetMaxOpenConns(100) // Sane default
	//db.SetMaxIdleConns(30)
	//db.SetConnMaxLifetime(30 * time.Second)

	// Устанавливаем дефолтную схему, в которую будет ходит данное приложение
	_, err = db.Exec(fmt.Sprintf("set search_path='%s'", DefaultSchema))
	if err != nil {
		logger.Errorf("set search_path to %s %s", DefaultSchema, err)
		return nil, errors.WithStack(err)
	}

	// Заполняем структуру postgres и возвращаем её в качестве нашего соединения
	res := &PostgresDB{
		Logger: logger,
		DB:     db,
	}

	return res, nil
}

package main

import (
	"db-service-homework-kodzimo/internal/config"
	"db-service-homework-kodzimo/internal/storage"
	"github.com/sirupsen/logrus"
)

func databaseConnection() {

}

type ServerError struct {
	Message  string
	HTTPCode int
}

func main() {
	cfg := config.ReadEnv("")

	ll := logrus.New()

	st, err := storage.NewStorage(cfg, ll)
	if err != nil {
		ll.Fatal(err)
	}

	_ = st
}

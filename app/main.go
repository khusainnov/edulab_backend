package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/khusainnov/edulab"
	"github.com/khusainnov/edulab/pkg/handler"
	"github.com/khusainnov/edulab/pkg/repository"
	"github.com/khusainnov/edulab/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	logrus.Infoln(".env config loading")
	if err := godotenv.Load("./config/.env"); err != nil {
		logrus.Errorf("Cannot load .env config, due to error: %s", err.Error())
	}

	logrus.Infoln(".yml config loading")
	if err := initConfig(); err != nil {
		logrus.Errorf("Cannot load .yml config, due to error: %s", err.Error())
	}

	logrus.Infoln("Initializing DB")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Errorf("Cannot initialize db, due to error: %s", err.Error())
	}

	logrus.Infoln("Repository initializing")
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	s := new(edulab.Server)

	logrus.Infof("Starting server on port: %s", os.Getenv("PORT"))
	if err = s.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Errorf("Cannot run the server, due to error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

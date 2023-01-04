package main

import (
	"levanidus/todo-app"
	"levanidus/todo-app/pkg/handler"
	"levanidus/todo-app/pkg/repository"
	"levanidus/todo-app/pkg/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	/*if err := initConfig(); err != nil {
		logrus.Fatalf("config init error: %s", err.Error())
	}*/

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("load env variables error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	})

	if err != nil {
		logrus.Fatalf("database initializaion error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

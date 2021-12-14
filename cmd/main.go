package main

import (
	"article"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"article/pkg/handler"
	"article/pkg/repository"
	"article/pkg/service"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	if err := initConfigs(); err != nil {
		log.Fatalf("initConfigs %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("load envs from .env  %s", err.Error())
	}

	db, err := article.NewPdb(&article.ConfigDB{
		Username: viper.GetString("db.username"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Dbname:   viper.GetString("db.dbname"),
		Sslmode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf(" Postgres DB init %s", err.Error())
	}

	// repo -> service -> handler
	repos := repository.NewRepostory(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	serv := article.NewServer(viper.GetString("webport"), handlers.NewRouter())
	go func() {
		if err := serv.Run(); err != nil {
			log.Fatalf("Server init %s", err.Error())
		}
	}()
	log.Println("Server has been started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Graceful shutdown")
	if err := serv.Shutdown(context.Background()); err != nil {
		log.Errorf("Server Shutdown %s", err.Error())
	}
	if err := db.Close(); err != nil {
		log.Errorf("db connection close: %s", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

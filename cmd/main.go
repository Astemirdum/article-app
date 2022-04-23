package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Astemirdum/article-app"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Astemirdum/article-app/pkg/handler"
	"github.com/Astemirdum/article-app/pkg/repository"
	"github.com/Astemirdum/article-app/pkg/service"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfigs(); err != nil {
		log.Fatalf("initConfigs %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("load envs from .env  %s", err.Error())
	}

	db, err := article.NewPdb(&article.ConfigDB{
		Username: viper.GetString("db.username"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Dbname:   viper.GetString("db.dbname"),
		Sslmode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf(" Postgres DB init %s", err.Error())
	}
	// repo -> service -> handler
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, log)

	serv := article.NewServer(handlers.NewRouter(),
		article.AddrHttp{
			Addr: viper.GetString("service.addr"),
			Port: viper.GetInt("service.webport"),
		})
	go func() {
		if err := serv.Run(); err != nil {
			log.Fatalf("Server init %v", err)
		}
	}()
	log.Println("Server has been started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Graceful shutdown")
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
	defer cancelFn()
	if err := serv.Shutdown(ctx); err != nil {
		log.Errorf("Server Shutdown %v", err)
	}
	if err := db.Close(); err != nil {
		log.Errorf("db connection close: %v", err)
	}
}

func initConfigs() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

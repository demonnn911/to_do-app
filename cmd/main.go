package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	todo "todo-app/app-models"
	ssogrpc "todo-app/clients/sso/grpc"
	"todo-app/pkg/config"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

// TODO refactor initializing ssoClient, remove hardcode, from srv.Run method
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	dbConfig := config.NewDBConfig()
	logrus.Info(dbConfig)
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		logrus.Fatalf("cannot initialize db: %s", err.Error())
	}

	logrus.Print("initializing grpc service")
	ssoConfig := ssogrpc.NewSSOConfig()
	logrus.Info(ssoConfig)
	ssoClient, err := ssogrpc.New(
		logrus.New(),
		*ssoConfig,
	)
	if err != nil {
		logrus.Fatal("failde to init sso client", err)
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos, ssoClient)
	handlers := handler.NewHandler(service)
	srv := new(todo.Server)
	srvConfig := config.NewHTTPServerConfig()
	logrus.Info(srvConfig)
	go func() {
		if err := srv.Run(*srvConfig, handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("cannot start server %s", err.Error())
		}
	}()
	logrus.Print("todo app started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Todo app shutting down")
	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("couldn't shut down an app %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("couldn't close db connection %s", err.Error())
	}
	logrus.Print("Todo app shutted down")
}

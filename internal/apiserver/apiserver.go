package apiserver

import (
	"WB0/internal/apiserver/orderservice"
	"WB0/internal/config"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Start(config *config.Config) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log, err := newLogger(config)
	if err != nil {
		return err
	}
	db, err := newDb(config.DbDriverName, config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	service := orderservice.NewOrderService(db, log)
	srv := newHttpServer(log, service)

	go func() {
		http.ListenAndServe(config.BindAddr, srv)
	}()

	////***
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Fatalf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Fatalf("ctx.Done: %v", done)
	}

	log.Println("Server Exited Property")

	return nil
}

func newDb(driverName string, dbURL string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dbURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func newLogger(config *config.Config) (*logrus.Logger, error) {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}
	logger := logrus.New()
	logger.SetLevel(level)

	return logger, nil
}

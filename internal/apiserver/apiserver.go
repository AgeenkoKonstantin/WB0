package apiserver

import (
	"WB0/internal/apiserver/nats"
	"WB0/internal/apiserver/orderservice"
	"WB0/internal/apiserver/repository"
	"WB0/internal/cache"
	"WB0/internal/config"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type apiserver struct {
	log      *logrus.Logger
	config   *config.Config
	natsConn stan.Conn
	db       *sql.DB
	cache    *cache.Cache
}

func NewApiServer(log *logrus.Logger, config *config.Config, natsConn stan.Conn, db *sql.DB) *apiserver {
	return &apiserver{
		log:      log,
		config:   config,
		natsConn: natsConn,
		db:       db,
	}
}

func (s *apiserver) Start() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := orderservice.NewOrderService(
		repository.NewCacheRepository(),
		repository.NewSqlRepository(s.db),
		s.log)
	srv := newHttpServer(s.log, service, ctx)

	go func() {
		http.ListenAndServe(s.config.BindAddr, srv)
	}()

	go func() {
		sub := nats.NewSubscriber(s.log, s.natsConn, service)
		sub.Run(ctx)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		log.Fatalf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Fatalf("ctx.Done: %v", done)
	}
	return nil
}

func NewDb(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open(config.DbDriverName, config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewLogger(config *config.Config) (*logrus.Logger, error) {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return nil, err
	}
	logger := logrus.New()
	logger.SetLevel(level)
	return logger, nil
}

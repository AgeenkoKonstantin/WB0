package producer

import (
	"WB0/internal/config"
	"context"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type producerServer struct {
	logger   *logrus.Logger
	config   *config.Config
	natsConn stan.Conn
}

func NewProducerServer(
	logger *logrus.Logger,
	config *config.Config,
	natsConn stan.Conn,
) *producerServer {
	return &producerServer{
		logger:   logger,
		config:   config,
		natsConn: natsConn,
	}
}

func (ps *producerServer) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		producer := newProducer(ps.natsConn, ps.logger)
		producer.Run()
	}()

	////***
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		ps.logger.Fatalf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		ps.logger.Fatalf("ctx.Done: %v", done)
	}

	return nil

}

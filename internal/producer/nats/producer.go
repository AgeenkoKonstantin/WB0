package nats

import (
	"WB0/internal/models"
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"time"
)

type producer struct {
	stanConn stan.Conn
	logger   *logrus.Logger
}

func NewProducer(stanConn stan.Conn, logger *logrus.Logger) *producer {
	return &producer{
		stanConn: stanConn,
		logger:   logger,
	}
}

func (p *producer) Publish(subject string, data []byte) error {
	p.logger.Infof("Publish data: %v to subject: %v", string(data), subject)
	return p.stanConn.Publish(subject, data)
}

func (p *producer) Run() {
	for {
		order := &models.Order{}
		if err := faker.FakeData(order, options.WithRandomMapAndSliceMaxSize(2)); err != nil {
			p.logger.Infof("producer.Run error while creating fake order: %+v", err)
		}
		orderBytes, _ := json.Marshal(*order)

		p.logger.Info("Publish new random order")
		err := p.Publish("ordersChannel", orderBytes)

		if err != nil {
			p.logger.Infof("send to ordersChannel failed with err = %+v", err)
		}

		time.Sleep(3000 * time.Millisecond)
	}

}

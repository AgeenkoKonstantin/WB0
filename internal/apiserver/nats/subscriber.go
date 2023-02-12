package nats

import (
	"WB0/internal/apiserver/orderservice"
	"WB0/internal/models"
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	workersNum = 1
)

type Subscriber struct {
	logger   *logrus.Logger
	stanConn stan.Conn
	service  *orderservice.OrderService
}

func NewSubscriber(logger *logrus.Logger, stanConn stan.Conn, service *orderservice.OrderService) *Subscriber {
	s := &Subscriber{
		logger:   logger,
		stanConn: stanConn,
		service:  service,
	}
	return s
}

func (s *Subscriber) Subscribe(subject, qgroup string, workerNum int, cb stan.MsgHandler) {
	s.logger.Printf("Subscribing to Subject: %v, group: %v", subject, qgroup)
	wg := &sync.WaitGroup{}

	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go s.runWorker(
			wg,
			i,
			s.stanConn,
			subject,
			qgroup,
			cb,
		)
	}
	wg.Wait()
}

func (s *Subscriber) runWorker(
	wg *sync.WaitGroup,
	workerID int,
	conn stan.Conn,
	subject string,
	qgroup string,
	cb stan.MsgHandler,
	opts ...stan.SubscriptionOption,
) {
	s.logger.Printf("Subscribing worker: %v, subject: %v, qgroup: %v", workerID, subject, qgroup)
	defer wg.Done()

	_, err := conn.QueueSubscribe(subject, qgroup, cb, opts...)
	if err != nil {
		s.logger.Printf("WorkerID: %v, QueueSubscribe: %v", workerID, err)
		if err := conn.Close(); err != nil {
			s.logger.Printf("WorkerID: %v, conn.Close error: %v", workerID, err)
		}
	}

}

func (s *Subscriber) Run(ctx context.Context) {
	go s.Subscribe("order:create", "subscriber", workersNum, s.processCreateOrder(ctx))
}

func (s *Subscriber) processCreateOrder(ctx context.Context) stan.MsgHandler {
	return func(msg *stan.Msg) {
		s.logger.Printf("subscriber process Create Order: %s", msg.String())

		var m models.Order

		if err := json.Unmarshal(msg.Data, &m); err != nil {
			s.logger.Printf("json.Unmarshal: %v", err)
			return
		}

		if err := s.service.SaveOrder(&m, ctx); err != nil {
			s.logger.Printf("Subscriber.processCreateOrder : %v", err)
			return
		}
	}
}

package nats

import (
	"WB0/internal/config"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	connectWait        = time.Second * 30
	pubAckWait         = time.Second * 30
	interval           = 10
	maxOut             = 5
	maxPubAcksInflight = 25
)

func NewNatsConnect(config *config.Config, logger *logrus.Logger, clientID string) (stan.Conn, error) {

	return stan.Connect(
		config.NatsClusterId,
		clientID,
		stan.ConnectWait(connectWait),
		stan.PubAckWait(pubAckWait),
		stan.NatsURL("nats://"+config.NatsHostname+":4222"),
		stan.Pings(interval, maxOut),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			logger.Info(fmt.Sprintf("Connection lost, reason: %v", reason))
		}),
		stan.MaxPubAcksInflight(maxPubAcksInflight),
	)
}

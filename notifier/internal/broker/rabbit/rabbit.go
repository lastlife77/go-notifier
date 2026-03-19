// Package rabbit provides helpers for working with a rabbitmq.
package rabbit

import (
	"fmt"
	"time"

	"github.com/wb-go/wbf/rabbitmq"
	"github.com/wb-go/wbf/retry"
)

type Rabbitmq struct {
	client       *rabbitmq.RabbitClient
	exchange     string
	dlx          string
	deadQueue    string
	deadQueueKey string
}

func New(user, pass, port string) (*Rabbitmq, error) {
	r := &Rabbitmq{
		exchange:     "notifier",
		dlx:          "dlx",
		deadQueue:    "dead_queue",
		deadQueueKey: "dead_key",
	}

	strategy := retry.Strategy{
		Attempts: 3,
		Delay:    3 * time.Second,
		Backoff:  2,
	}
	url, err := getURL(user, pass, port)
	if err != nil {
		return nil, err
	}

	cfg := rabbitmq.ClientConfig{
		URL:            url,
		ConnectionName: "notifier",
		ConnectTimeout: 5 * time.Second,
		Heartbeat:      10 * time.Second,
		ReconnectStrat: strategy,
		ProducingStrat: strategy,
		ConsumingStrat: strategy,
	}

	r.client, err = rabbitmq.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// main exchange
	err = r.client.DeclareExchange(r.exchange, "direct", true, false, false, nil)
	if err != nil {
		return nil, err
	}

	// dlx exchange
	err = r.client.DeclareExchange(r.dlx, "direct", true, false, false, nil)
	if err != nil {
		return nil, err
	}

	// dlx queue
	err = r.client.DeclareQueue(r.deadQueue, r.dlx, r.deadQueueKey, true, false, false, nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func getURL(user, pass, port string) (string, error) {
	url := "amqp://"
	url += user + ":"
	url += pass
	url += fmt.Sprintf("@rabbitmq:%v/", port)

	return url, nil
}

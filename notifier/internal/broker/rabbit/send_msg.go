package rabbit

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lastlife77/go-notifier/internal/broker"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wb-go/wbf/rabbitmq"
)

func (r *Rabbitmq) SendMsg(id string, msg string, t time.Time) error {
	timeout := time.Until(t).Milliseconds()

	args := amqp.Table{
		"x-message-ttl":             timeout,
		"x-dead-letter-exchange":    r.dlx,
		"x-dead-letter-routing-key": r.deadQueueKey,
	}

	err := r.client.DeclareQueue(id, r.exchange, id, true, false, false, args)
	if err != nil {
		var amqpErr *amqp091.Error
		if errors.As(err, &amqpErr) {
			return broker.NewError(amqpErr.Code, amqpErr.Reason)
		}
		return broker.NewError(500, err.Error())
	}

	publisher := rabbitmq.NewPublisher(r.client, r.exchange, "application/json")

	ctx := context.Background()

	bodyMsg := fmt.Appendf([]byte{}, `{"msg":"%v"}`, msg)
	routingKey := id

	err = publisher.Publish(
		ctx,
		bodyMsg,
		routingKey,
		rabbitmq.WithExpiration(5*time.Minute),
		rabbitmq.WithHeaders(amqp.Table{"x-service": "auth"}),
	)
	if err != nil {
		var amqpErr *amqp091.Error
		if errors.As(err, &amqpErr) {
			return broker.NewError(amqpErr.Code, amqpErr.Reason)
		}
		return broker.NewError(500, err.Error())
	}

	return nil
}

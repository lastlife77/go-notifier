package rabbit

import (
	"errors"

	"github.com/lastlife77/go-notifier/internal/broker"
	"github.com/rabbitmq/amqp091-go"
)

func (r *Rabbitmq) GetStatus(id string) (string, error) {
	ch, err := r.client.GetChannel()
	if err != nil {
		var amqpErr *amqp091.Error
		if errors.As(err, &amqpErr) {
			return "", broker.NewError(amqpErr.Code, amqpErr.Reason)
		}
		return "", broker.NewError(500, err.Error())
	}

	q, err := ch.QueueDeclarePassive(id, true, false, false, false, nil)
	if err != nil {
		var amqpErr *amqp091.Error
		if errors.As(err, &amqpErr) {
			return "", broker.NewError(amqpErr.Code, amqpErr.Reason)
		}
		return "", broker.NewError(500, err.Error())
	}

	if q.Messages == 0 {
		return "The message has been sent", nil
	}

	return "The message is waiting to be sent", nil
}

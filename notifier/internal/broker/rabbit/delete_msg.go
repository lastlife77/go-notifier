package rabbit

import (
	"errors"

	"github.com/lastlife77/go-notifier/internal/broker"
	"github.com/rabbitmq/amqp091-go"
)

func (r *Rabbitmq) DeleteMsg(id string) error {
	ch, err := r.client.GetChannel()
	if err != nil {
		var amqpErr *amqp091.Error
		if errors.As(err, &amqpErr) {
			return broker.NewError(amqpErr.Code, amqpErr.Reason)
		}
		return broker.NewError(500, err.Error())
	}

	_, err = ch.QueueDelete(id, false, false, false)
	if err != nil {
		var amqpErr *amqp091.Error
		if errors.As(err, &amqpErr) {
			return broker.NewError(amqpErr.Code, amqpErr.Reason)
		}
		return broker.NewError(500, err.Error())
	}

	return nil
}

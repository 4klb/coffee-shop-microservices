package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/4klb/coffee-shop-microservices/dish/config"
	"github.com/streadway/amqp"
)

type Publisher struct {
	channel *amqp.Channel
}

func NewPublisher(rmq *RabbitMQ) (*Publisher, error) {
	return &Publisher{
		channel: rmq.Channel,
	}, nil
}

func (p *Publisher) Created(ctx context.Context, routingKey string, msg []byte) error {
	return p.publish(ctx, routingKey, msg)
}

func (p *Publisher) publish(ctx context.Context, routingKey string, msg []byte) error {
	log.Printf("Msg %v published into RabbitMQ", string(msg))

	if err := p.channel.Publish(
		config.GetConfig().RabbitMQ.ExchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/x-encoding-job",
			Body:        msg,
			Timestamp:   time.Now(),
		},
	); err != nil {
		return fmt.Errorf("ch.Publish %w", err)
	}
	return nil
}

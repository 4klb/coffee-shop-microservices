package broker

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/4klb/coffee-shop-microservices/dish/config"
)

type Consumer struct {
	rmq  *RabbitMQ
	done chan struct{}
}

func (c *Consumer) RunConsumer(rmq *RabbitMQ) (<-chan error, error) {
	consumer := &Consumer{
		rmq:  rmq,
		done: make(chan struct{}),
	}

	errC := make(chan error, 1)
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()
		log.Println("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer func() {
			rmq.Close()
			stop()
			cancel()
			close(errC)
		}()

		if err := c.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}
	}()

	go func() {
		if err := consumer.Consumer(); err != nil {
			errC <- err
		}
	}()

	return errC, nil
}

func (c *Consumer) Consumer() error {
	queue, err := c.rmq.Channel.QueueDeclare(
		config.GetConfig().RabbitMQ.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("ch.ExchangeDeclare %w", err)
	}

	err = c.rmq.Channel.QueueBind(
		queue.Name,
		config.GetConfig().RabbitMQ.RoutingKey,
		config.GetConfig().RabbitMQ.ExchangeName,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Channel.QueueBind %w", err)
	}
	msgs, err := c.rmq.Channel.Consume(
		config.GetConfig().RabbitMQ.QueueName,
		config.GetConfig().RabbitMQ.ConsumerName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Channel.Consume %w", err)
	}

	go func() {
		for msg := range msgs {
			var nack bool
			var msgBody []byte

			switch msg.RoutingKey {
			case "created":
				msgBody = msg.Body
				log.Printf("Consumer got %v msg from RabbitMQ", string(msgBody))
				//TODO
			case "updated":
				//TODO
			case "deleted":
				//TODO
			default:
				nack = true
			}

			if nack {
				msg.Nack(false, nack)
			} else {
				msg.Ack(false)
			}
		}

		c.done <- struct{}{}
	}()
	return nil
}

func (c *Consumer) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server")

	c.rmq.Channel.Cancel(config.GetConfig().RabbitMQ.ConsumerName, false)

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context.Done: %w", ctx.Err())
		case <-c.done:
			return nil
		}
	}
}

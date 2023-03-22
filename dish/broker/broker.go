package broker

import (
	"fmt"

	"github.com/4klb/coffee-shop-microservices/dish/config"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func GetRabbitConnection() (*RabbitMQ, error) {
	uri := fmt.Sprintf("amqp://%v:%v@%v:%v/",
		config.GetConfig().RabbitMQ.User,
		config.GetConfig().RabbitMQ.Password,
		config.GetConfig().RabbitMQ.Host,
		config.GetConfig().RabbitMQ.Port)

	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("amqp.Dial %w", err)
	}
	fmt.Println("Successfully connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("conn.Channel %w", err)
	}

	if err = ch.ExchangeDeclare(
		config.GetConfig().RabbitMQ.ExchangeName,
		config.GetConfig().RabbitMQ.ExchangeType,
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("ch.ExchangeDeclare %w", err)
	}

	// if err = ch.Qos(
	// 	0,
	// 	1,
	// 	false,
	// ); err != nil {
	// 	return nil, fmt.Errorf("ch.Qos %w", err)
	// }

	return &RabbitMQ{
		Connection: conn,
		Channel:    ch,
	}, err
}

func (r *RabbitMQ) Close() {
	r.Connection.Close()
}

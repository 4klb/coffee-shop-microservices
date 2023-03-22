package server

// type Server struct {
// 	rmq  *broker.RabbitMQ
// 	done chan struct{}
// }

// func Run(rmq *broker.RabbitMQ) (<-chan error, error) {
// 	svr := &Server{
// 		rmq:  rmq,
// 		done: make(chan struct{}),
// 	}

// 	errC := make(chan error, 1)
// 	ctx, stop := signal.NotifyContext(context.Background(),
// 		os.Interrupt,
// 		syscall.SIGTERM,
// 		syscall.SIGQUIT)

// 	go func() {
// 		<-ctx.Done()
// 		logger.Info("Shutdown signal received")
// 		errC <- err
// 	}()

// 	return errC, nil
// }

// func (s *Server) ListenAndServe() {
// 	queue, err := s.rmq.Channel.QueueDeclare(
// 		config.GetConfig().RabbitMQ.QueueName,
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("ch.ExchangeDeclare %w", err)
// 	}

// 	err = s.rmq.Channel.QueueBind(
// 		queue.Name,
// 		config.GetConfig().RabbitMQ.RoutingKey,
// 		config.GetConfig().RabbitMQ.ExchangeName,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// }

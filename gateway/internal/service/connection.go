package service

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetConnection(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if ctx.Err() == context.DeadlineExceeded {
		return nil, errors.New("service unavailable addr: " + addr)
	}
	if err != nil {
		return nil, err
	}

	return conn, nil
}

package grpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	ssov1 "github.com/dm1tl/protos/gen/go/sso"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	authAPI ssov1.AuthClient
	userAPI ssov1.UserClient
}

func New(
	log *logrus.Logger,
	addr string,
	timeout time.Duration,
	retriesCount int,
) (*Client, error) {
	const op = "clients.sso.grpc.New()"
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}
	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}
	cc, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Client{
		authAPI: ssov1.NewAuthClient(cc),
		userAPI: ssov1.NewUserClient(cc),
	}, nil
}

func InterceptorLogger(l *logrus.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, level grpclog.Level, msg string, fields ...any) {
		l.Log(logrus.Level(level), msg)
	})
}

func (c *Client) Login(ctx context.Context,
	email string,
	password string) (string, error) {
	const op = "clients.sso.grpc.Login()"
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	resp, err := c.authAPI.Login(ctx, &ssov1.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return resp.Token, nil
}

func (c *Client) Register(ctx context.Context,
	email string,
	password string,
) (int64, error) {
	const op = "clients.sso.grpc.Register()"
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	resp, err := c.authAPI.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return resp.UserId, nil
}

func (c *Client) ValidateToken(ctx context.Context,
	token string) (int64, error) {
	const op = "clients.sso.grpc.ValidateToken()"
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	resp, err := c.authAPI.ValidateToken(ctx, &ssov1.ValidateTokenRequest{
		Token: token,
	})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return resp.Id, nil
}

func (c *Client) Delete(ctx context.Context,
	id int64) (err error) {
	const op = "clients.sso.grpc.Delete()"
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	resp, err := c.userAPI.Delete(ctx, &ssov1.DeleteRequest{
		Id: id,
	})
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if resp.ErrorMessage != "success" {
		return errors.New("couldn't delete user from grpc db")
	}
	return nil
}

package grpc

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	ssov1 "github.com/dm1tl/protos/gen/go/sso"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type SSOServiceClient struct {
	authAPI ssov1.AuthClient
	userAPI ssov1.UserClient
}

type SSOClientWrapper struct {
	SSOProvider
}

func NewSSOClientWrapper(provider SSOProvider) *SSOClientWrapper {
	return &SSOClientWrapper{
		SSOProvider: provider,
	}
}

type SSOProvider interface {
	Register(ctx context.Context, email string, password string) (int64, error)
	Login(ctx context.Context, email string, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (int64, error)
	Delete(ctx context.Context, id int64) error
}

func NewSSOServiceClient(
	log *logrus.Logger,
	cfg SSOConfig,
) (*SSOServiceClient, error) {
	const op = "clients.sso.grpc.New()"
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(cfg.RetriesCount),
		grpcretry.WithPerRetryTimeout(cfg.Timeout),
	}
	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}
	cc, err := grpc.NewClient(cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &SSOServiceClient{
		authAPI: ssov1.NewAuthClient(cc),
		userAPI: ssov1.NewUserClient(cc),
	}, nil
}

func InterceptorLogger(l *logrus.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, level grpclog.Level, msg string, fields ...any) {
		l.Log(logrus.Level(level), msg)
	})
}

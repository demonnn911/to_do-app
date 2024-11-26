package service

import (
	"context"
	"fmt"
	todo "todo-app/app-models"
	"todo-app/clients/sso/grpc"
	"todo-app/pkg/repository"
)

type AuthService struct {
	repo      repository.Authorization
	ssoClient grpc.SSOProvider
}

func NewAuthService(repo repository.Authorization, ssoclient grpc.SSOProvider) *AuthService {
	return &AuthService{
		repo:      repo,
		ssoClient: ssoclient,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, user todo.User) (int64, error) {
	const op = "pkg.service.CreateUser()(grpc)"
	id, err := s.ssoClient.Register(ctx, user.Email, user.Password)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	err = s.repo.CreateUser(ctx, id)
	if err != nil {
		if rollback := s.ssoClient.Delete(ctx, id); rollback != nil {
			return 0, fmt.Errorf("failed to rollback user in grpc after DB error %s: %w", op, rollback)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, err
}

func (s *AuthService) Login(ctx context.Context, input todo.SignInInput) (string, error) {
	const op = "pkg.service.Login()(grpc)"
	token, err := s.ssoClient.Login(ctx, input.Email, input.Password)
	if err != nil {

		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) (int64, error) {
	const op = "pkg.service.ValidateToken()(grpc)"

	id, err := s.ssoClient.ValidateToken(ctx, token)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

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
	ssoClient *grpc.Client
}

func NewAuthService(repo repository.Authorization, ssoclient *grpc.Client) *AuthService {
	return &AuthService{
		repo:      repo,
		ssoClient: ssoclient,
	}
}

func (s *AuthService) CreateUser(user todo.User) (int64, error) {
	const op = "pkg.service.CreateUser()(grpc)"
	ctx := context.Background()
	id, err := s.ssoClient.Register(ctx, user.Email, user.Password)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	err = s.repo.CreateUser(id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, err
}

func (s *AuthService) Login(input todo.SignInInput) (string, error) {
	const op = "pkg.service.Login()(grpc)"
	ctx := context.Background()
	token, err := s.ssoClient.Login(ctx, input.Email, input.Password)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (s *AuthService) ValidateToken(token string) (int64, error) {
	const op = "pkg.service.ValidateToken()(grpc)"
	ctx := context.Background()
	id, err := s.ssoClient.ValidateToken(ctx, token)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

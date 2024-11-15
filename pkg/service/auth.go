package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	todo "todo-app/app-models"
	"todo-app/clients/sso/grpc"
	"todo-app/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "^#&^@#^(^*#qwprjffxlsnm;kv)"
	tokenTTL   = 12 * time.Hour
	signingKey = "hfdjsfjdskhfjsdhfsl*^&(*&^#@%$)(@*@*#NZM)"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

//func (s *AuthService) CreateUser(user todo.User) (int, error) {
//	user.Password = generatePasswordHash(user.Password)
//	return s.repo.CreateUser(user)
//}

// First try to connect grpc service
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	const op = "pkg.service.CreateUser()(grpc)"
	ctx := context.Background()
	id, err := s.ssoClient.Register(ctx, user.Email, user.Password)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return int(id), nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("incorrect signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("invalid type of token: need *tokenClaims")
	}
	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

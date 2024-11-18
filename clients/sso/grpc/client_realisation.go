package grpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	ssov1 "github.com/dm1tl/protos/gen/go/sso"
)

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

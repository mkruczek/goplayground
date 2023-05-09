package something_like_httpserver

import (
	"context"
)

type UserRepository interface {
	GetUserById(ctx context.Context, ID string) (User, error)
	GetUsersByGroup(ctx context.Context, group string) ([]User, error)
	CreateUSer(ctx context.Context, user User) error
}

type UserRepo struct {
}

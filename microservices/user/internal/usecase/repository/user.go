package repository

import (
	"context"

	"github.com/hizzuu/grpc-sample-user/internal/domain"
)

type UserRepository interface {
	Get(ctx context.Context, id int64) (*domain.User, error)
}
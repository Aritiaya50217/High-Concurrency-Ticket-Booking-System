package usecase

import "context"

type UserService interface {
	GetUser(ctx context.Context, userID uint64) (bool, error)
}

package user

import (
	"context"

	mUser "github/fadlinux/edot/internal/model/user"
)

type (
	// Param : param struct that sent to repository
	Param struct {
		ID       int64
		Name     string
		Phone    string
		Email    string
		Password string
		Page     int64
		Size     int64
	}
)

// Repository : contract of functions in repository layer
type Repository interface {
	AddUser(ctx context.Context, params Param) (lastID int64, err error)
	FetchUserByPhoneEmail(ctx context.Context, params Param) (result []mUser.User, err error)
	CountUserByPhoneEmail(ctx context.Context, params Param) (total int64, err error)
}

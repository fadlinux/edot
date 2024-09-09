package user

import (
	"context"
	mUser "github/fadlinux/edot/internal/model/user"
)

type (
	// Param : param struct that sent to usecase
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

// Usecase : contract of functions in Usecase layer
type Usecase interface {
	FetchUser(ctx context.Context, params mUser.User) (result mUser.Response, err error)
	AddUser(ctx context.Context, params mUser.User) (lastID int64, err error)
	GetUserLogin(ctx context.Context, params mUser.User) (totalData int64, err error)
}

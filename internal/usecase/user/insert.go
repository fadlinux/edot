package user

import (
	"context"
	mUser "github/fadlinux/edot/internal/model/user"

	rUser "github/fadlinux/edot/internal/repository/user"
)

func (a *userUC) AddUser(ctx context.Context, params mUser.User) (lastID int64, err error) {
	lastID, err = a.mysqlUserRepo.AddUser(ctx, rUser.Param{
		Name:     params.Name,
		Phone:    params.Phone,
		Email:    params.Email,
		Password: params.PasswordHash,
	})
	return
}

package user

import (
	"context"
	mUser "github/fadlinux/edot/internal/model/user"
	"log"

	rUser "github/fadlinux/edot/internal/repository/user"
)

func (a *userUC) FetchUser(ctx context.Context, params mUser.User) (result mUser.Response, err error) {
	userData, err := a.mysqlUserRepo.FetchUserByPhoneEmail(ctx, rUser.Param{
		Name:  params.Name,
		Email: params.Email,
		Phone: params.Phone,
	})

	if err != nil {
		log.Println("[User Repository] Fail to fetch data", err)
		return
	}

	for _, v := range userData {
		result.Data = append(result.Data, mUser.Data{
			ID:       v.ID,
			Name:     v.Name,
			Email:    v.Email,
			Phone:    v.Phone,
			Password: v.PasswordHash,
		})
	}

	return
}

func (a *userUC) GetUserLogin(ctx context.Context, params mUser.User) (totalData int64, err error) {
	totalData, err = a.mysqlUserRepo.CountUserByPhoneEmail(ctx, rUser.Param{
		Email: params.Email,
		Phone: params.Phone,
	})

	return
}

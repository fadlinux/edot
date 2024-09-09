// // Usecase : contract of functions in Usecase layer
// type Usecase interface {
// 	FetchUser(ctx context.Context, params mUser.User) (result mUser.Response, err error)
// 	AddUser(ctx context.Context, params mUser.User) (lastID int64, err error)
// 	GetUserLogin(ctx context.Context, params mUser.User) (err error)
// }

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

	// totalData, err := a.mysqlUserRepo.CountUserByPhoneEmail(ctx, rUser.Param{
	// 	Email: params.Email,
	// })

	// if err != nil || totalData <= 0 {
	// 	totalData = 0
	// 	log.Println("[User Repository] Fail to count data", err)
	// 	return
	// }

	return
}

func (a *userUC) AddUser(ctx context.Context, params mUser.User) (lastID int64, err error) {
	lastID, err = a.mysqlUserRepo.AddUser(ctx, rUser.Param{
		Name:     params.Name,
		Phone:    params.Phone,
		Email:    params.Email,
		Password: params.PasswordHash,
	})
	return
}

func (a *userUC) GetUserLogin(ctx context.Context, params mUser.User) (totalData int64, err error) {
	totalData, err = a.mysqlUserRepo.CountUserByPhoneEmail(ctx, rUser.Param{
		Email: params.Email,
		Phone: params.Phone,
	})

	return
}

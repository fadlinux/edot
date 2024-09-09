package mysql

import (
	"context"

	mUser "github/fadlinux/edot/internal/model/user"
	rUser "github/fadlinux/edot/internal/repository/user"
)

func (a *mysqlUserRepo) AddUser(ctx context.Context, params rUser.Param) (lastID int64, err error) {
	err = addUserPrepareStmt.QueryRow(params.Name, params.Email, params.Phone, params.Password).Scan(&lastID)
	if err != nil {
		lastID = 0
		return
	}

	return
}

func (a *mysqlUserRepo) FetchUserByPhoneEmail(ctx context.Context, params rUser.Param) (result []mUser.User, err error) {
	rows, err := fetchUserPrepareStmt.Query(params.Email, params.Phone)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		item := mUser.User{}
		rows.Scan(
			&item.ID,
			&item.Name,
			&item.Email,
			&item.Phone,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.PasswordHash,
		)
		if item.ID <= 0 || item.Name == "" {
			continue
		}
		result = append(result, item)
	}

	return
}

func (a *mysqlUserRepo) CountUserByPhoneEmail(ctx context.Context, params rUser.Param) (total int64, err error) {
	err = countUserLoginPrepareStmt.QueryRow(params.Email, params.Phone).Scan(&total)
	return
}

func (a *mysqlUserRepo) CountUserLoginPrepareStmt(ctx context.Context, params rUser.Param) (total int64, err error) {
	err = countUserLoginPrepareStmt.QueryRow(params.Email, params.Phone).Scan(&total)
	return
}

func (a *mysqlUserRepo) FetchUserLogin(ctx context.Context, params rUser.Param) (total int64, err error) {
	err = countUserPrepareStmt.QueryRow("%"+params.Email+"%", "%"+params.Phone+"%").Scan(&total)
	return
}

package mysql

import (
	"context"

	mShop "github/fadlinux/edot/internal/model/shop"
	rShop "github/fadlinux/edot/internal/repository/shop"
)

func (a *mysqlShopRepo) AddShop(ctx context.Context, params rShop.Param) (lastID int64, err error) {
	err = addShopPrepareStmt.QueryRow(params.Name).Scan(&lastID)
	if err != nil {
		lastID = 0
		return
	}

	return
}

func (a *mysqlShopRepo) FetchShopByPhoneEmail(ctx context.Context, params rShop.Param) (result []mShop.Shop, err error) {
	rows, err := fetchShopPrepareStmt.Query(params.Name)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		item := mShop.Shop{}
		rows.Scan(
			&item.ID,
			&item.Name,
		)
		if item.ID <= 0 || item.Name == "" {
			continue
		}
		result = append(result, item)
	}

	return
}

func (a *mysqlShopRepo) CountShopByPhoneEmail(ctx context.Context, params rShop.Param) (total int64, err error) {
	err = countShopLoginPrepareStmt.QueryRow(params.Email, params.Phone).Scan(&total)
	return
}

func (a *mysqlShopRepo) CountShopLoginPrepareStmt(ctx context.Context, params rShop.Param) (total int64, err error) {
	err = countShopLoginPrepareStmt.QueryRow(params.Email, params.Phone).Scan(&total)
	return
}

func (a *mysqlShopRepo) FetchShopLogin(ctx context.Context, params rShop.Param) (total int64, err error) {
	err = countShopPrepareStmt.QueryRow("%"+params.Email+"%", "%"+params.Phone+"%").Scan(&total)
	return
}

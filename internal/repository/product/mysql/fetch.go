package mysql

import (
	"context"
	"time"

	mProduct "github/fadlinux/edot/internal/model/product"
	rProduct "github/fadlinux/edot/internal/repository/product"
)

func (a *mysqlProductRepo) AddProduct(ctx context.Context, params rProduct.Param) (lastID int64, err error) {
	var datetime = time.Now()
	dt := datetime.Format("2006-01-02 15:04:05")

	err = addProductPrepareStmt.QueryRow(params.Name, params.Description, params.Price, params.Stock, params.WarehouseId, dt).Scan(&lastID)
	if err != nil {
		lastID = 0
		return
	}

	return
}

func (a *mysqlProductRepo) FetchSearchProduct(ctx context.Context, params mProduct.SearchParam) (result []mProduct.Product, err error) {
	rows, err := fetchProductPrepareStmt.Query("%"+params.Q+"%", (params.Page-1)*params.Size, params.Size)

	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		item := mProduct.Product{}
		rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Stock,
			&item.Price,
			&item.WarehouseId,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if item.ID <= 0 || item.Name == "" || item.Stock <= 0 {
			continue
		}

		result = append(result, item)
	}

	return
}

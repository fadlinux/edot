package product

import (
	"database/sql"
	"github/fadlinux/edot/common/util/config"
	httpAdmin "github/fadlinux/edot/internal/delivery/product/http"
	productRepo "github/fadlinux/edot/internal/repository/product"
	mysqlProductRepository "github/fadlinux/edot/internal/repository/product/mysql"
	productUC "github/fadlinux/edot/internal/usecase/product"
)

var (
	mysqlProduct     *sql.DB
	mysqlProductRepo productRepo.Repository
	productUsecase   productUC.Usecase

	HTTPDelivery httpAdmin.Delivery
)

// Initialize : initialize function for product
func Initialize() {
	mysqlProduct = newDBConnection("mysql", config.GetString("mysql.host"))
	mysqlProductRepo = mysqlProductRepository.NewMySqlProductRepo(mysqlProduct)

	productUsecase = productUC.NewProductUC(mysqlProductRepo)
	HTTPDelivery = httpAdmin.NewProductHTTP(productUsecase)
}

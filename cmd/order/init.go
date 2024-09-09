package order

import (
	"database/sql"
	"github/fadlinux/edot/common/util/config"
	httpAdmin "github/fadlinux/edot/internal/delivery/order/http"
	orderRepo "github/fadlinux/edot/internal/repository/order"
	mysqlOrderRepository "github/fadlinux/edot/internal/repository/order/mysql"
	orderUC "github/fadlinux/edot/internal/usecase/order"

	productRepo "github/fadlinux/edot/internal/repository/product"
	mysqlProductRepository "github/fadlinux/edot/internal/repository/product/mysql"
)

var (
	mysqlOrder     *sql.DB
	mysqlOrderRepo orderRepo.Repository

	mysqlProduct     *sql.DB
	mysqlProductRepo productRepo.Repository

	orderUsecase orderUC.Usecase

	HTTPDelivery httpAdmin.Delivery
)

// Initialize : initialize function for order
func Initialize() {
	mysqlOrder = newDBConnection("mysql", config.GetString("mysql.host"))
	mysqlProduct = newDBConnection("mysql", config.GetString("mysql.host"))

	mysqlOrderRepo = mysqlOrderRepository.NewMySqlOrderRepo(mysqlOrder)
	mysqlProductRepo = mysqlProductRepository.NewMySqlProductRepo(mysqlProduct)

	orderUsecase = orderUC.NewOrderUC(mysqlOrderRepo, mysqlProductRepo)
	HTTPDelivery = httpAdmin.NewOrderHTTP(orderUsecase)
}

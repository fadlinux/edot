package shop

import (
	"database/sql"
	"github/fadlinux/edot/common/util/config"
	httpAdmin "github/fadlinux/edot/internal/delivery/shop/http"
	shopRepo "github/fadlinux/edot/internal/repository/shop"
	mysqlShopRepository "github/fadlinux/edot/internal/repository/shop/mysql"
	shopUC "github/fadlinux/edot/internal/usecase/shop"
)

var (
	mysqlShop     *sql.DB
	mysqlShopRepo shopRepo.Repository
	shopUsecase   shopUC.Usecase

	HTTPDelivery httpAdmin.Delivery
)

// Initialize : initialize function for shop
func Initialize() {
	mysqlShop = newDBConnection("mysql", config.GetString("mysql.host"))
	mysqlShopRepo = mysqlShopRepository.NewMySqlShopRepo(mysqlShop)

	shopUsecase = shopUC.NewShopUC(mysqlShopRepo)
	HTTPDelivery = httpAdmin.NewShopHTTP(shopUsecase)
}

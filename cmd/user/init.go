package user

import (
	"database/sql"
	"github/fadlinux/edot/common/util/config"
	httpAdmin "github/fadlinux/edot/internal/delivery/user/http"
	userRepo "github/fadlinux/edot/internal/repository/user"
	mysqlUserRepository "github/fadlinux/edot/internal/repository/user/mysql"
	userUC "github/fadlinux/edot/internal/usecase/user"
)

var (
	mysqlUser     *sql.DB
	mysqlUserRepo userRepo.Repository
	userUsecase   userUC.Usecase

	HTTPDelivery httpAdmin.Delivery
)

// Initialize : initialize function for user
func Initialize() {
	mysqlUser = newDBConnection("mysql", config.GetString("mysql.host"))
	mysqlUserRepo = mysqlUserRepository.NewMySqlUserRepo(mysqlUser)

	userUsecase = userUC.NewUserUC(mysqlUserRepo)
	HTTPDelivery = httpAdmin.NewUserHTTP(userUsecase)
}

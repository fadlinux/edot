package user

import (
	rUser "github/fadlinux/edot/internal/repository/user"
)

type (
	userUC struct {
		mysqlUserRepo rUser.Repository
	}
)

// NewUserUC : create new instance for User usecase
func NewUserUC(mysqlUser rUser.Repository) Usecase {
	return &userUC{
		mysqlUserRepo: mysqlUser,
	}
}

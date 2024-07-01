package usecase

import (
	"github.com/TheManhattanProject/crypt_pg_backend/user/domain/config"
	userdomain "github.com/TheManhattanProject/crypt_pg_backend/user/domain/user"
	"github.com/TheManhattanProject/crypt_pg_backend/user/repository"
	userusecase "github.com/TheManhattanProject/crypt_pg_backend/user/usecase/user"
)

type Usecases struct {
	User userdomain.Usecase
}

func InitUsecases(
	config *config.Config,
	repos repository.Repositories,
	caches repository.Caches,
) *Usecases {
	userUsecase := userusecase.NewUserUsecaseImplementation(config, repos.User, caches.User)

	return &Usecases{
		User: userUsecase,
	}
}

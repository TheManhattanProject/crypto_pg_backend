package userusecase

import (
	"context"

	"github.com/TheManhattanProject/crypt_pg_backend/user/domain/config"
	userdomain "github.com/TheManhattanProject/crypt_pg_backend/user/domain/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	cfg  *config.Config
	repo userdomain.Repository
}

func (impl *UserUsecaseImpl) CreateOrUpdateUser(
	ctx context.Context,
	user userdomain.User,
) (*userdomain.User, error) {
	if user.ID != uuid.Nil {
		// Update
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		impl.cfg.HashRounds,
	)
	if err != nil {
		return nil, err
	}

	user.Password = string(encryptedPassword)

	// TODO impl country mapping
	user.Country = "INDIA"

	return impl.repo.Create(ctx, user)
}

func (impl *UserUsecaseImpl) GetByID(ctx context.Context, id uuid.UUID) (*userdomain.User, error) {
	user, err := impl.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, err
}

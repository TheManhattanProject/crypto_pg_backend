package userusecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/TheManhattanProject/crypt_pg_backend/user/domain/config"
	userdomain "github.com/TheManhattanProject/crypt_pg_backend/user/domain/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	cfg   *config.Config
	repo  userdomain.Repository
	cache userdomain.Cache
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

func (impl *UserUsecaseImpl) SendOTP(ctx context.Context, id uuid.UUID) error {

	// Check user exists
	_, err := impl.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err // TODO: impl error codes
		}

		return err
	}

	otp, err := impl.cache.CreateAndSetOTP(ctx, id)
	if err != nil {
		return err
	}

	//TODO: implement sms/email
	fmt.Printf("sending otp for user %v ->  %v", id, otp)

	return nil
}

func (impl *UserUsecaseImpl) ValidateOTP(
	ctx context.Context,
	id uuid.UUID,
	otp string,
) (bool, error) {

	_, err := impl.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, err // TODO: impl error codes
		}
		return false, err
	}

	cachedOTP, err := impl.cache.GetOTP(ctx, id)
	if err != nil {
		// Check against not found error code
		return false, err
	}

	if cachedOTP != otp {
		return false, fmt.Errorf("error: invalid otp")
	}

	return true, nil
}

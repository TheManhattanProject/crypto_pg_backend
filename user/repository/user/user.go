package userrepository

import (
	"context"
	"database/sql"

	"github.com/TheManhattanProject/crypt_pg_backend/user/domain/config"
	userdomain "github.com/TheManhattanProject/crypt_pg_backend/user/domain/user"
	"github.com/google/uuid"
)

type UserRepositoryImpl struct {
	config *config.Config
	db     *sql.DB
}

func (repo *UserRepositoryImpl) Create(
	ctx context.Context,
	user userdomain.User,
) (*userdomain.User, error) {
	var err error

	user.ID, err = uuid.NewV7()
	if err != nil {
		// TODO: impl logger
		return nil, err
	}

	_, err = repo.db.ExecContext(
		ctx,
		"INSERT INTO users (id, name, phone_number, email, password, country) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID,
		user.UserName,
		user.PhoneNumber,
		user.Email,
		user.Password,
		user.Country,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepositoryImpl) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*userdomain.User, error) {
	var user userdomain.User
	user.ID = id

	err := repo.db.QueryRowContext(
		ctx,
		"SELECT name, phone_number, email, country FROM users WHERE id = $1",
		id,
	).Scan(&user.UserName, &user.PhoneNumber, &user.Email, &user.Country)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepositoryImpl) Update(
	ctx context.Context,
	user userdomain.User,
) (*userdomain.User, error) {
	return nil, nil
}

package userdomain

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	CountryCode string    `json:"-"` // as in the mobile number code for the country
	Country     string    `json:"country"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	UserName    string    `json:"user_name"` // The Full name of the User
	Password    string    `json:"-"`

	//TODO add other details
}

type Repository interface {
	Create(context.Context, User) (*User, error)
	GetByID(context.Context, uuid.UUID) (*User, error)
	Update(context.Context, User) (*User, error)
}

type Cache interface {
	CreateAndSetOTP(
		context.Context,
		uuid.UUID,
	) (string, error) // Create and set the OTP (can be revised)
	GetOTP(context.Context, uuid.UUID) (string, error)
}

type Usecase interface { // the service implementation
	CreateOrUpdateUser(context.Context, User) (*User, error)
	GetByID(context.Context, uuid.UUID) (*User, error)
	SendOTP(context.Context, uuid.UUID) error
	ValidateOTP(context.Context, uuid.UUID, string) (bool, error)
}

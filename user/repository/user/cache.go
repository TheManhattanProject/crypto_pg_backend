package userrepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/TheManhattanProject/crypt_pg_backend/user/domain/config"
	"github.com/TheManhattanProject/crypto_pg_backend/common/utils"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type UserCacheImpl struct {
	config *config.Config
	client *redis.Client
}

func NewUserCache(cfg *config.Config, client *redis.Client) *UserCacheImpl {
	return &UserCacheImpl{
		config: cfg,
		client: client,
	}
}

func (cache *UserCacheImpl) CreateAndSetOTP(ctx context.Context, id uuid.UUID) (string, error) {
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return "", err
	}

	if err := cache.client.Set(ctx, fmt.Sprintf("%s:otp", id.String()), otp, cache.config.OTPExpiryTime).Err(); err != nil {
		return "", err
	}

	return otp, nil
}

func (cache *UserCacheImpl) GetOTP(ctx context.Context, id uuid.UUID) (string, error) {
	otp, err := cache.client.Get(ctx, fmt.Sprintf("%s:otp", id.String())).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", err // TODO make separate error codes
		}
		return "", err
	}

	return otp, err
}

package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
	"time"
)

// GetUserSessionFromCache  - get refresh token from cache
func (r *DatastoreRepository) GetUserSessionFromCache(ctx context.Context, accountNo string) (string, error) {
	redisKey := fmt.Sprintf(model.RedisKeyRefresh, accountNo)
	return r.redisClient.Get(ctx, redisKey).Result()
}

// UpdateUserSessionIntoCache  - safe refresh token into cache
func (r *DatastoreRepository) UpdateUserSessionIntoCache(ctx context.Context, accountNo, refreshData string) error {
	redisKey := fmt.Sprintf(model.RedisKeyRefresh, accountNo)
	return r.redisClient.Set(ctx, redisKey, refreshData, r.cfg.Token.RefreshExpiry).Err()
}

func (r *DatastoreRepository) DeleteUserSessionFromCache(ctx context.Context, accountNo string) error {
	redisKey := fmt.Sprintf(model.RedisKeyRefresh, accountNo)
	return r.redisClient.Del(ctx, redisKey).Err()
}

func (r *DatastoreRepository) InsertUserInfoIntoDB(ctx context.Context, payload model.RegistrationPayload) error {
	// begin query TX
	trx, err := r.dbClient.Beginx()
	if err != nil {
		log.Err(err)
		return err
	}

	// commit & rolback if err
	defer func() {
		if err != nil {
			err = trx.Rollback()
		} else {
			err = trx.Commit()
		}
	}()

	_, err = r.queries.RegisterUserInfo.ExecContext(ctx, map[string]interface{}{
		"account_number": payload.AccountNumber,
		"created_at":     time.Now(),
		"email":          payload.Email,
		"fullname":       payload.Fullname,
		"status":         payload.Status,
	})
	if err != nil {
		log.Error().Msgf("RegisterUserInfo err %#v", err)
	}

	_, err = r.queries.InsertUserAuth.ExecContext(ctx, map[string]interface{}{
		"account_number": payload.AccountNumber,
		"username":       payload.Username,
		"hash":           payload.Hash,
		"created_at":     time.Now(),
	})

	if err != nil {
		log.Error().Msgf("InsertUserAuth err %#v", err)
	}

	return nil
}

func (r *DatastoreRepository) GetUserInfoFromDB(ctx context.Context, accountNumber string) (data model.UserProfileInfo, err error) {
	err = r.queries.GetUserInfo.GetContext(ctx, &data, map[string]interface{}{
		"account_number": accountNumber,
	})
	if err != nil {
		fmt.Printf("err : %#v\n", err)
	}
	return
}

func (r *DatastoreRepository) GetUserByUsername(ctx context.Context, username string) (model.UserAuthInfo, error) {
	var (
		err  error
		data model.UserAuthInfo
	)
	err = r.queries.GetUserInfoByUname.GetContext(ctx, &data, map[string]interface{}{
		"username": username,
	})
	if err != nil {
		fmt.Printf("err : %#v\n", err)
		return data, err
	}
	return data, nil
}

// redis datastore handler
func (r *DatastoreRepository) GetUserInfoFromCache(ctx context.Context, accountNumber string) (model.UserProfileInfo, error) {
	// fastcampus:user:info:[account_number]
	// fastcampus:user:info:001
	redisKey := fmt.Sprintf("fastcampus:user:info:%s", accountNumber)
	data, err := r.redisClient.Get(ctx, redisKey).Result()
	if err != nil {
		return model.UserProfileInfo{}, err
	}
	// if no error then parse redis data
	var result model.UserProfileInfo
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return result, err
	}
	return result, err
}

func (r *DatastoreRepository) UpdateUserInfoOnCache(ctx context.Context, data model.UserProfileInfo) error {
	// set into fastcampus:user:info:[account_number]
	// TTL 1 hour
	// fastcampus:user:info:001
	payload, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	redisKey := fmt.Sprintf("fastcampus:user:info:%s", data.AccountNumber)
	return r.redisClient.Set(ctx, redisKey, payload, model.DefaultUserInfoTTL).Err()
}

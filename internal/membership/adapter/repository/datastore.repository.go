package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/membership/core/model"
)

func (r *DatastoreRepository) GetUserSessionFromCache(ctx context.Context) {

}

func (r *DatastoreRepository) UpdateUserSessionIntoCache(ctx context.Context) {

}

func (r *DatastoreRepository) InsertUserInfoIntoDB(ctx context.Context, payload model.RegistrationPayload) error {
	// begin query TX

	// commit & rolback if err

	// insert user info

	// insert user auth
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

func (r *DatastoreRepository) GetUserByUsername(ctx context.Context) {

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

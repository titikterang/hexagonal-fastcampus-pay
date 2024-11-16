package repository

import (
	"context"
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

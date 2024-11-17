package repository

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
	"io"
	"net/http"
)

func (r *BankingRepository) GetAccountInfo(ctx context.Context, accountNumber string) (model.UserProfileInfo, error) {
	// get account info from membership service

	url := r.cfg.ExternalAPI.MembershipService + "v1/membership/info?account_number=" + accountNumber
	method := "GET"

	var returnData model.UserProfileInfo
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Err(err).Msg("GetAccountInfo.http.NewRequest")
		return returnData, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("GetAccountInfo.http.DoRequest")
		return returnData, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Err(err).Msg("GetAccountInfo.io.ReadAll")
		return returnData, err
	}

	var result model.UserProfileAPIResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Err(err).Msg("GetAccountInfo.Unmarshal")
		return model.UserProfileInfo{}, err
	}

	return model.UserProfileInfo{
		AccountNumber: result.AccountNumber,
		Email:         result.Email,
		Fullname:      result.Fullname,
		Status:        result.Status,
	}, nil
}

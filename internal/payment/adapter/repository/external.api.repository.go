package repository

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"io"
	"net/http"
)

func (r *PaymentRepository) GetAccountInfo(ctx context.Context, accountNumber string) (types.UserProfileInfo, error) {
	// get account info from membership service
	url := r.cfg.ExternalAPI.MembershipService + "v1/membership/info?account_number=" + accountNumber
	method := "GET"

	var returnData types.UserProfileInfo
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Err(err).Msg("GetAccountInfo.http.NewRequest")
		return returnData, err
	}

	res, err := r.httpClient.Do(req)
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
		return types.UserProfileInfo{}, err
	}

	return types.UserProfileInfo{
		AccountNumber: result.AccountNumber,
		Email:         result.Email,
		Fullname:      result.Fullname,
		Status:        result.Status,
	}, nil
}

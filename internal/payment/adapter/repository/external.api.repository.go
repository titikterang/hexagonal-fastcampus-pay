package repository

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/banking/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/common"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/membership"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/money"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"io"
	"net/http"
)

func (r *PaymentRepository) GetAccountInfo(ctx context.Context, accountNumber string) (types.UserProfileInfo, error) {
	result, err := r.membershipClient.GetUserInfo(ctx, &membership.UserInfoPayload{
		AccountNumber: accountNumber,
	})
	if err != nil {
		return types.UserProfileInfo{}, err
	}

	return types.UserProfileInfo{
		AccountNumber: result.AccountNumber,
		Email:         result.Email,
		Fullname:      result.Fullname,
		Status:        result.Status,
	}, nil
}

func (r *PaymentRepository) GetAccountInfoHttp(ctx context.Context, accountNumber string) (types.UserProfileInfo, error) {
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

func (r *PaymentRepository) GetUserBalanceInfo(ctx context.Context, accountNumber string) (decimal.Decimal, error) {
	var (
		balance decimal.Decimal
		err     error
	)

	result, err := r.moneyClient.GetUserBalancePrivate(ctx, &money.UserBalancePayload{
		AccountNumber: accountNumber,
	})

	if err != nil {
		log.Err(err).Msgf("failed GetUserBalancePrivate %#v", err)
		return balance, err
	}

	balanceMoney := result.GetBalance()
	balanceStr := common.MoneyToString(balanceMoney)
	return decimal.NewFromString(balanceStr)
}

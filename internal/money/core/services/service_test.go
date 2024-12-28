package services

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/model"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/money/core/ports/mocks"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/common"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/types"
	"google.golang.org/genproto/googleapis/type/money"
	"testing"
	"time"
)

func TestMoneyService_UpdateHistory(t *testing.T) {
	mockTime := time.Now()
	cfg := &config.Config{}
	ctx := context.Background()
	type fields struct {
		config     *config.Config
		repository *mocks.MockMoneyRepositoryAdapter
	}
	type args struct {
		ctx  context.Context
		info model.CashMovementInfo
	}
	tests := []struct {
		name    string
		wantErr bool
		args    args
		fields  fields
		mock    func(fields fields, args args)
	}{
		{
			name: "test_UpdateHistory",
			args: args{
				ctx: ctx,
				info: model.CashMovementInfo{
					RequestID:     "001",
					AccountNumber: "acc001",
					Date:          time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
					Amount:        decimal.NewFromInt(100_000),
					MovementType:  "payment",
				},
			},
			fields: fields{
				config: cfg,
			},
			wantErr: false,
			mock: func(fields fields, args args) {
				fields.repository.EXPECT().AppendCashMovementIntoDatastore(args.ctx, args.info).Return(nil)
				fields.repository.EXPECT().AppendCashMovementInfoIntoCache(args.ctx, args.info).Return(nil)
				// update snapshoot
				finalAmount := common.DecimalToMoney(decimal.NewFromInt(0))
				fields.repository.EXPECT().GetCashMovementFromCache(args.ctx, args.info.AccountNumber).Return([]model.CashMovementInfo{}, nil)
				fields.repository.EXPECT().UpdateSnapshot(args.ctx, args.info.AccountNumber, "0").Return(nil)

				fields.repository.EXPECT().ConstructBalanceInfo(args.ctx, model.UserCashInfo{
					AccountNumber: args.info.AccountNumber,
					LastUpdate:    mockTime,
					BalanceAmount: common.MoneyToDecimal(finalAmount),
				})
			},
		},
		{
			name: "test_UpdateHistory_err_append_datastore",
			args: args{
				ctx: ctx,
				info: model.CashMovementInfo{
					RequestID:     "001",
					AccountNumber: "acc001",
					Date:          time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
					Amount:        decimal.NewFromInt(100_000),
					MovementType:  "payment",
				},
			},
			fields: fields{
				config: cfg,
			},
			wantErr: true,
			mock: func(fields fields, args args) {
				fields.repository.EXPECT().AppendCashMovementIntoDatastore(args.ctx, args.info).Return(errors.New("failed"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mocks.NewMockMoneyRepositoryAdapter(ctrl)
			tt.fields.repository = mockRepo

			tt.mock(tt.fields, tt.args)
			s := &MoneyService{
				config:     tt.fields.config,
				repository: tt.fields.repository,
				currentTime: types.CurrentTime{
					Mock: true,
					Now:  mockTime,
				},
			}
			if err := s.UpdateHistory(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMoneyService_UpdateUserBalance(t *testing.T) {
	cfg := &config.Config{}
	ctx := context.Background()
	//amount := common.DecimalToMoney(decimal.NewFromInt(100_000))
	type fields struct {
		config     *config.Config
		repository *mocks.MockMoneyRepositoryAdapter
	}
	type args struct {
		ctx           context.Context
		requestID     string
		accountNumber string
		amount        *money.Money
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(fields fields, args args)
	}{
		{
			name: "test_update_balance",
			fields: fields{
				config: cfg,
			},
			args: args{
				ctx:           ctx,
				requestID:     "req001",
				accountNumber: "user001",
				amount:        nil,
			},
			wantErr: true,
			mock: func(fields fields, args args) {
				fields.repository.EXPECT().ReqIDExists(args.ctx, args.accountNumber, args.requestID).Return(false)
				fields.repository.EXPECT().SaveReqID(args.ctx, args.accountNumber, args.requestID)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mocks.NewMockMoneyRepositoryAdapter(ctrl)
			tt.fields.repository = mockRepo

			tt.mock(tt.fields, tt.args)

			s := &MoneyService{
				config:     tt.fields.config,
				repository: tt.fields.repository,
			}
			if err := s.UpdateUserBalance(tt.args.ctx, tt.args.requestID, tt.args.accountNumber, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

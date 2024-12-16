package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	model2 "github.com/titikterang/hexagonal-fastcampus-pay/internal/payment/core/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type PaymentHistoryPayload struct {
	DateTime time.Time `bson:"date_time"`
	Status   string    `bson:"status"`
}

type PaymentDBPayload struct {
	InvoiceID     string                  `bson:"invoice_id"`
	MerchantID    string                  `bson:"merchant_id"`
	AccountNumber string                  `bson:"account_number"`
	Message       string                  `bson:"message"`
	Amount        string                  `bson:"amount"`
	DateTime      time.Time               `bson:"date_time"`
	Status        string                  `bson:"status"`
	UpdateHistory []PaymentHistoryPayload `bson:"history"`
}

func (r *PaymentRepository) InsertPaymentHistory(ctx context.Context, data model2.PaymentInfo) {
	payload := PaymentDBPayload{
		InvoiceID:     data.InvoiceID,
		MerchantID:    data.MerchantID,
		AccountNumber: data.AccountNumber,
		Message:       data.Message,
		Amount:        data.Amount.String(),
		DateTime:      data.DateTime,
		Status:        string(data.Status),
		UpdateHistory: []PaymentHistoryPayload{
			{
				DateTime: data.DateTime,
				Status:   string(data.Status),
			},
		},
	}
	_, err := r.DB.Collection("payment_history").InsertOne(ctx, payload)
	if err != nil {
		log.Error().Msgf("failed insert history : %#v", err)
	}
}

func (r *PaymentRepository) UpdatePaymentStatus(ctx context.Context, data model2.PaymentInfo) {
	filter := bson.D{{"invoice_id", data.InvoiceID}}
	update := bson.D{{"$set",
		bson.D{
			{"status", data.Status},
		},
	}}
	_, err := r.DB.Collection("payment_history").UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error().Msgf("failed update history : %#v", err)
	}
	//_, err = r.DB.Collection("payment_history.history").InsertOne(ctx, []PaymentHistoryPayload{{
	//	DateTime: data.DateTime,
	//	Status:   string(data.Status),
	//},
	//})
	//if err != nil {
	//	log.Error().Msgf("failed insert new history : %#v", err)
	//}
}

func (r *PaymentRepository) InsertPaymentInfo(ctx context.Context) {

}

func (r *PaymentRepository) EventIDExists(ctx context.Context, eventType string, accountNo, id string) bool {
	return false
}

func (r *PaymentRepository) SaveEventID(ctx context.Context, eventType string, accountNo, id string) error {
	return nil
}

func (r *PaymentRepository) GetPaymentInfo(ctx context.Context, invoiceID string) (model2.PaymentInfo, []model2.PaymentHistory, error) {
	var (
		data    PaymentDBPayload
		info    model2.PaymentInfo
		history = make([]model2.PaymentHistory, 0)
	)

	filter := bson.D{{
		"invoice_id", invoiceID},
	}
	err := r.DB.Collection("payment_history").FindOne(ctx, filter).Decode(&data)
	amt, err := decimal.NewFromString(data.Amount)
	info = model2.PaymentInfo{
		InvoiceID:     data.InvoiceID,
		MerchantID:    data.MerchantID,
		AccountNumber: data.AccountNumber,
		Message:       data.Message,
		Amount:        amt,
		DateTime:      data.DateTime,
		Status:        model2.PaymentStatus(data.Status),
	}

	for _, v := range data.UpdateHistory {
		history = append(history, model2.PaymentHistory{
			DateTime: v.DateTime.Format(time.DateTime),
			Status:   model2.PaymentStatus(v.Status),
		})
	}
	return info, history, err
}

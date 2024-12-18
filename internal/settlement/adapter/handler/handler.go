package handler

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/settlement"
)

func (h *Handler) GetSettlementReport(ctx context.Context, data *settlement.SettlementReportRequest) (*settlement.SettlementReportResponse, error) {
	result, err := h.settlementService.HandleSettlementReport(ctx, data.GetAccountNo(), data.GetSettlementDate())
	if err != nil {
		return &settlement.SettlementReportResponse{}, err
	}
	var settlementList = settlement.SettlementReportResponse{
		SettlementDate: data.GetSettlementDate(),
		AccountNo:      data.GetAccountNo(),
		Report:         make([]*settlement.SettlementReport, 0),
	}
	for _, v := range result {
		settlementList.Report = append(settlementList.Report, &settlement.SettlementReport{
			Id:               v.TransactionID,
			AccountNo:        v.AccountNo,
			CashMovementType: string(v.SettlementType),
			FeePercentage:    v.FeePercentage.InexactFloat64(),
			FeeAmount:        v.FeeAmount.InexactFloat64(),
		})
	}

	return &settlementList, nil
}

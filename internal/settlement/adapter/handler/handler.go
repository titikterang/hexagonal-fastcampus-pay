package handler

import (
	"context"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/settlement"
)

func (h *Handler) GetSettlementReport(ctx context.Context, data *settlement.SettlementReportRequest) (*settlement.SettlementReportResponse, error) {
	err := h.settlementService.HandleSettlementReport(ctx, data.GetAccountNo(), data.GetSettlementDate())
	if err != nil {
		return &settlement.SettlementReportResponse{}, err
	}
	result := settlement.SettlementReportResponse{
		SettlementDate: "",
		AccountNo:      "",
		Report:         nil,
	}

	return &result, nil
}

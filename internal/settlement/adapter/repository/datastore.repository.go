package repository

import "context"

// to datastore
func (r *SettlementRepository) SaveSettlementReport(ctx context.Context) error {
	return nil
}
func (r *SettlementRepository) GetSettlementReport(ctx context.Context, accountNo, date string) {

}

// idempotence
func (r *SettlementRepository) EventIDExists(ctx context.Context, eventType string, accountNo, id string) bool {
	return false
}

func (r *SettlementRepository) SaveEventID(ctx context.Context, eventType string, accountNo, id string) error {
	return nil
}

package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const (
	prepareSchema = `create table if not exists money_cash_movement
						(
							request_id         varchar(100) not null,
							account_number     varchar(100),
							cash_movement_date date         not null,
							amount             numeric(20, 6),
							cash_movement_type varchar(100),
							constraint money_cash_movement_pk
								primary key (request_id, cash_movement_date)
						) partition by LIST (cash_movement_date);`

	insertCashMovement = `INSERT INTO public.money_cash_movement (request_id, account_number, cash_movement_date, amount, cash_movement_type) 
						VALUES (:request_id, :account_number, :cash_movement_date, :amount, :cash_movement_type);`
)

type statementQueries struct {
	InsertCashMovement *sqlx.NamedStmt
}

func (r *MoneyRepository) prepareStatements() error {
	stmtNamed, err := r.dbClientMaster.PrepareNamed(insertCashMovement)
	if err != nil {
		log.Error().Msgf("prepare statement GetUserInfo err %#v", err.Error())
		return err
	}
	r.queries.InsertCashMovement = stmtNamed

	return nil
}

func (r *MoneyRepository) initDBSchema(ctx context.Context) error {
	res := r.dbClientMaster.MustExecContext(ctx, prepareSchema)
	if _, err := res.RowsAffected(); err != nil {
		log.Error().Msgf("failed prepare schema, %#v", err.Error())
		return err
	}
	return nil
}

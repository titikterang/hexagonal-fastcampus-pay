package repository

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jmoiron/sqlx"
)

const (
	CreateTableHistory = `create table if not exists banking_cash_history
							(
								id                  bigserial
									constraint id_pk
										primary key,
								account_source      varchar(100)            not null,
								account_destination varchar(100)            not null,
								amount              numeric(20, 6),
								cash_movement_type  varchar(20),
								created_at          timestamp default now() not null
							);
`
	InsertCashHistory = `insert into banking_cash_history(account_source,
															 account_destination,
															 amount,
															 cash_movement_type)
							values (:account_source, :account_destination, :amount, :cash_movemet_type)`
)

type statementQueries struct {
	InsertCashHistory *sqlx.NamedStmt
}

func (r *BankingRepository) prepareStatements() error {
	stmtNamed, err := r.dbClientMaster.PrepareNamed(InsertCashHistory)
	if err != nil {
		log.Error("prepare statement InsertCashHistory err ", err.Error())
		return err
	}
	r.queries.InsertCashHistory = stmtNamed

	return nil
}

func (r *BankingRepository) initDBSchema(ctx context.Context) error {
	res := r.dbClientMaster.MustExecContext(ctx, CreateTableHistory)
	if _, err := res.RowsAffected(); err != nil {
		log.Error("failed prepare schema ", err.Error())
		return err
	}
	return nil
}

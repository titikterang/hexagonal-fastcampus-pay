package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const (
	createTableHistory = `create table if not exists public.transfer_history
					(
						transaction_id             varchar(100),
						transfer_date              date,
						amount                     decimal(20),
						source_account_number      varchar(100),
						destination_account_number varchar(100),
						transfer_type              varchar(20),
						bank_code                  varchar(4),
						status                     varchar(20),
						created_at                 timestamp,
						updated_at                 timestamp,
						message                    varchar(100),
						constraint transfer_history_pk
							primary key (transaction_id, transfer_date)
					)`
	queryInsertHistory = `INSERT INTO transfer_history (
									transaction_id, 
                                     transfer_date, 
                                     amount, 
                                     source_account_number,
                                     destination_account_number, 
                                     transfer_type, 
                                     bank_code, 
                                     status, 
                                     created_at,
                                     updated_at, 
                                     message)
					VALUES (:transaction_id,
							:transfer_date,
							:amount,
							:source_account_number,
							:destination_account_number,
							:transfer_type,
							:bank_code,
							:status,
							:created_at,
							:updated_at,
							:message);`
	queryUpdateStatus  = `update transfer_history set status = :status, updated_at = :updated where transaction_id = :id`
	queryGetAllHistory = `select transaction_id,
								 transfer_date,
								 amount,
								 source_account_number,
								 destination_account_number,
								 transfer_type,
								 bank_code,
								 status,
								 created_at,
								 updated_at,
								 message from transfer_history where transfer_date = :transfer_date`
)

type statementQueries struct {
	InsertTransferHistory *sqlx.NamedStmt
	UpdateTransferStatus  *sqlx.NamedStmt
	GetTransferHistory    *sqlx.NamedStmt
}

func (r *TransferRepository) prepareStatements() error {
	stmtNamed, err := r.dbClientMaster.PrepareNamed(queryInsertHistory)
	if err != nil {
		log.Error().Msgf("prepare statement InsertTransferHistory err %s", err.Error())
		return err
	}
	r.queries.InsertTransferHistory = stmtNamed

	stmtNamed, err = r.dbClientMaster.PrepareNamed(queryUpdateStatus)
	if err != nil {
		log.Error().Msgf("prepare statement UpdateTransferStatus err %s", err.Error())
		return err
	}
	r.queries.UpdateTransferStatus = stmtNamed

	stmtNamed, err = r.dbClientSlave.PrepareNamed(queryGetAllHistory)
	if err != nil {
		log.Error().Msgf("prepare statement GetTransferHistory err %s", err.Error())
		return err
	}
	r.queries.GetTransferHistory = stmtNamed

	return nil
}

func (r *TransferRepository) initDBSchema(ctx context.Context) error {
	res := r.dbClientMaster.MustExecContext(ctx, createTableHistory)
	if _, err := res.RowsAffected(); err != nil {
		log.Error().Msgf("failed prepare schema %s", err.Error())
		return err
	}
	return nil
}

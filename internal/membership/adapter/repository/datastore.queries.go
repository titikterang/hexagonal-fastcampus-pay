package repository

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jmoiron/sqlx"
)

const (
	prepareSchema = `create table if not exists user_auth
			(
				account_number varchar(100) not null
					constraint user_auth_pk
						primary key,
				username       varchar(100) not null,
				hash           varchar(200) not null,
				last_login     timestamp    not null,
				created_at     timestamp    not null
			);
			
			create table if not exists user_profile
			(
				account_number varchar(100) not null
					constraint user_profile_pk
						primary key,
				created_at     timestamp,
				email          varchar(100),
				fullname       varchar(150) not null,
				status         varchar(10)  not null
			);`

	GetUserInfo = `select account_number, email, fullname, status
					from user_profile p
					where account_number = :account_number`
	GetUserInfoByUname = "select username, hash from user_auth where username = :username"
)

type statementQueries struct {
	GetUserInfo        *sqlx.NamedStmt
	GetUserInfoByUname *sqlx.NamedStmt
}

func (r *DatastoreRepository) prepareStatements() error {
	stmtNamed, err := r.dbClient.PrepareNamed(GetUserInfo)
	if err != nil {
		log.Error("prepare statement GetUserInfo err ", err.Error())
		return err
	}
	r.queries.GetUserInfo = stmtNamed

	stmtNamed, err = r.dbClient.PrepareNamed(GetUserInfoByUname)
	if err != nil {
		log.Error("prepare statement GetUserInfo err ", err.Error())
		return err
	}
	r.queries.GetUserInfoByUname = stmtNamed
	return nil
}

func (r *DatastoreRepository) initDBSchema(ctx context.Context) error {
	res := r.dbClient.MustExecContext(ctx, prepareSchema)
	if _, err := res.RowsAffected(); err != nil {
		log.Error("failed prepare schema ", err.Error())
		return err
	}
	return nil
}

package postgre

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
)

type Database struct {
	cfg      *config.Config
	DBClient DBInterface
}

func InitDBConnection(cfg *config.Config) *Database {
	fmt.Printf("%#v\n\n\n", cfg)
	return &Database{
		cfg: cfg,
	}
}

func (d *Database) InitiateConnection() (DBInterface, error) {
	cfg := d.cfg.Postgre
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Address, cfg.Port, cfg.Username, cfg.Password, cfg.Dbname, cfg.Sslmode)
	return d.connect(connectionStr)
}

func (d *Database) connect(connStr string) (DBInterface, error) {
	cfg := d.cfg.Postgre
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		// log err
		return nil, err
	}

	if cfg.Maxopenconn != -1 {
		db.SetMaxOpenConns(cfg.Maxopenconn)
	}

	if cfg.Maxidleconn != -1 {
		db.SetMaxIdleConns(cfg.Maxidleconn)
	}

	return db, nil
}

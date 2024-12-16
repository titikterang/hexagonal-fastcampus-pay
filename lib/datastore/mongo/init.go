package mongo

import (
	"fmt"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Database struct {
	cfg      *config.Config
	DBClient DBInterface
}

func InitDBConnection(cfg *config.Config) (*Database, error) {
	db := Database{
		cfg: cfg,
	}

	var err error
	connStr := db.getConnStr()
	db.DBClient, err = mongo.Connect(options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	return &db, nil
}

func (d *Database) getConnStr() string {
	connStr := "mongodb://"
	if d.cfg.MongoDB.DBUser != "" && d.cfg.MongoDB.DBPass != "" {
		connStr += "%s:%s@%s/?maxPoolSize=%d"
		return fmt.Sprintf(connStr, d.cfg.MongoDB.DBUser, d.cfg.MongoDB.DBPass, d.cfg.MongoDB.Address, d.cfg.MongoDB.MaxPoolSize)
	}

	connStr += "%s/?maxPoolSize=%d"
	return fmt.Sprintf(connStr, d.cfg.MongoDB.Address, d.cfg.MongoDB.MaxPoolSize)
}

func (d *Database) GetDB() *mongo.Database {
	return d.DBClient.Database(d.cfg.MongoDB.DBName)
}

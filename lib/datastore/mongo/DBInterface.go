package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type DBInterface interface {
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	Database(name string, opts ...options.Lister[options.DatabaseOptions]) *mongo.Database
	StartSession(opts ...options.Lister[options.SessionOptions]) (*mongo.Session, error)
}

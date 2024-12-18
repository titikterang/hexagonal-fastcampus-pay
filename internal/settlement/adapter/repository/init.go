package repository

import (
	"github.com/redis/go-redis/v9"
	"github.com/titikterang/hexagonal-fastcampus-pay/internal/settlement/core/ports"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/config"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/datastore/mongo"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/money"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/payment"
	"github.com/titikterang/hexagonal-fastcampus-pay/lib/protos/v1/transfer"
	mongo2 "go.mongodb.org/mongo-driver/v2/mongo"
)

type SettlementRepository struct {
	cfg         *config.Config
	dbClient    mongo.DBInterface
	DB          *mongo2.Database
	redisClient *redis.Client
	// grpc client
	moneyClient    money.MoneyServiceClient
	transferClient transfer.TransferServiceClient
	paymentClient  payment.PaymentServiceClient
}

func NewSettlementRepository(cfg *config.Config,
	redisClient *redis.Client,
	dbClient mongo.DBInterface,
	moneyClient money.MoneyServiceClient,
	transferClient transfer.TransferServiceClient,
	paymentClient payment.PaymentServiceClient) ports.SettlementRepository {

	repo := &SettlementRepository{
		cfg:         cfg,
		dbClient:    dbClient,
		DB:          dbClient.Database(cfg.MongoDB.DBName),
		redisClient: redisClient,
		// grpc client
		moneyClient:    moneyClient,
		transferClient: transferClient,
		paymentClient:  paymentClient,
	}

	return repo
}

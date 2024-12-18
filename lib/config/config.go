package config

import "time"

type Config struct {
	App           AppConfig         `json:"app" mapstructure:"app"`
	WorkerPool    WorkerPoolConfig  `json:"worker_pool" mapstructure:"worker_pool"`
	Http          HttpConfig        `json:"http" mapstructure:"http"`
	Grpc          GRPCConfig        `json:"grpc" mapstructure:"grpc"`
	Redis         RedisConfig       `json:"redis" mapstructure:"redis"`
	MongoDB       MongoDBConfig     `json:"mongo_db" mapstructure:"mongo_db"`
	Postgre       PostgreConfig     `json:"postgre" mapstructure:"postgre"`
	PostgreMaster PostgreConfig     `json:"postgre_master" mapstructure:"postgre_master"`
	PostgreSlave  PostgreConfig     `json:"postgre_slave" mapstructure:"postgre_slave"`
	ExternalAPI   ExternalAPIConfig `json:"external_api" mapstructure:"external_api"`
	Kafka         KafkaConfig       `json:"kafka_config" mapstructure:"kafka_config"`
}

type WorkerPoolConfig struct {
	Limit int `json:"limit" mapstructure:"limit"`
}

type MongoDBConfig struct {
	Address     string `json:"address" mapstructure:"address"`
	DBName      string `json:"db_name" mapstructure:"db_name"`
	DBUser      string `json:"db_user" mapstructure:"db_user"`
	DBPass      string `json:"db_pass" mapstructure:"db_pass"`
	MaxPoolSize int    `json:"max_pool_size" mapstructure:"max_pool_size"`
}

type KafkaConfig struct {
	Hosts          []string          `json:"hosts" mapstructure:"hosts"`
	MaxPollRecord  int               `json:"max_poll_record" mapstructure:"max_poll_record"`
	ConsumerGroup  string            `json:"consumer_group" mapstructure:"consumer_group"`
	ConsumerTopics map[string]string `json:"consumer_topics" mapstructure:"consumer_topics"`
	ProducerTopics map[string]string `json:"producer_topics" mapstructure:"producer_topics"`
}

type ExternalAPIConfig struct {
	MembershipService string `json:"membership" mapstructure:"membership"`
	MembershipGrpc    string `json:"membership_grpc" mapstructure:"membership_grpc"`
	MoneyGrpc         string `json:"money_grpc" mapstructure:"money_grpc"`
	TransferGrpc      string `json:"transfer_grpc" mapstructure:"transfer_grpc"`
	PaymentGrpc       string `json:"payment_grpc" mapstructure:"payment_grpc"`
}

type PostgreConfig struct {
	Address         string `json:"address" mapstructure:"address"`
	Port            string `json:"port" mapstructure:"port"`
	Username        string `json:"username" mapstructure:"username"`
	Password        string `json:"password" mapstructure:"password"`
	Dbname          string `json:"dbname" mapstructure:"dbname"`
	Sslmode         string `json:"sslmode" mapstructure:"sslmode"`
	Maxopenconn     int    `json:"maxopenconn" mapstructure:"maxopenconn"`
	Maxidleconn     int    `json:"maxidleconn" mapstructure:"maxidleconn"`
	Connmaxidletime int    `json:"connmaxidletime" mapstructure:"connmaxidletime"`
	Connmaxlifetime int    `json:"connmaxlifetime" mapstructure:"connmaxlifetime"`
}

type AppConfig struct {
	Address string `json:"address" mapstructure:"address"`
	Label   string `json:"label" mapstructure:"label"`
}

type GRPCConfig struct {
	GrpcAddress string        `json:"address" mapstructure:"address"`
	Timeout     time.Duration `json:"timeout" mapstructure:"timeout"`
}

type HttpConfig struct {
	Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
}

type RedisConfig struct {
	Address  string `json:"address" mapstructure:"address"`
	PoolSize int    `json:"poolsize" mapstructure:"poolsize"`
}

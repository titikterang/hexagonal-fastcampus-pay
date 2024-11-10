package config

import "time"

type Config struct {
	App     AppConfig     `json:"app" mapstructure:"app"`
	Http    HttpConfig    `json:"http" mapstructure:"http"`
	Redis   RedisConfig   `json:"redis" mapstructure:"redis"`
	Postgre PostgreConfig `json:"postgre" mapstructure:"postgre"`
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

type HttpConfig struct {
	Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
}

type RedisConfig struct {
	Address  string `json:"address" mapstructure:"address"`
	PoolSize int    `json:"poolsize" mapstructure:"poolsize"`
}

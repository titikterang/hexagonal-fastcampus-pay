package config

type SecretConfig struct {
	PostgreMaster PostgreConfig `json:"postgre_master"`
	PostgreSlave  PostgreConfig `json:"postgre_slave"`
	Postgre       PostgreConfig `json:"postgre"`
	TokenConfig   TokenConfig   `json:"token"`
}

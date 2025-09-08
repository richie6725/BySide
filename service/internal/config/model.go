package config

import "go.uber.org/dig"

type BysideServer struct {
	dig.Out
	DBMS           DatabaseManageSystem `mapstructure:"DatabaseManageSystem"`
	ServiceAddress ServiceAddress       `mapstructure:"service_address"`
}

type ServiceAddress struct {
	Byside string `mapstructure:"Byside"`
}

type DatabaseManageSystem struct {
	MongoDBSystem map[string]MongoDB `mapstructure:"MongoDB"`
	RedisServer   map[string]Redis   `mapstructure:"Redis"`
}

type MongoDB struct {
	Host     string `mapstructure:"Host"`
	Port     string `mapstructure:"Port"`
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	Database string `mapstructure:"Database"`
}

type Redis struct {
	Host     string `mapstructure:"Host"`
	Port     string `mapstructure:"Port"`
	Password string `mapstructure:"Password"`
	Database int    `mapstructure:"Database"`
}

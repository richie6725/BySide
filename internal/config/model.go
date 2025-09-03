package config

type Server struct {
	DatabaseManage
}

type DatabaseManage struct {
	MongoDBSystem map[string]MongoDB `json:"MongoDB"`
	RedisServer   map[string]Redis   `json:"Redis"`
}

type MongoDB struct {
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
	Database string `json:"Database"`
}

type Redis struct {
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	Password string `json:"Password"`
	Database int    `json:"Database"`
}

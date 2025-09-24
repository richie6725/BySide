package database

import (
	"Byside/service/internal/config"
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

const (
	mongoLocal = "RichieMongo"
	redisLocal = "RichieRedis"
	mariaLocal = "RichieMaria"
)

type BysideOut struct {
	dig.Out
	MongoLocal *mongo.Database `name:"mongo_byside"`
	RedisLocal *redis.Client   `name:"redis_byside"`
	MariaLocal *gorm.DB        `name:"maria_byside"`
}

func NewByside(ctx context.Context, dbms config.DatabaseManageSystem) BysideOut {
	return BysideOut{
		MongoLocal: newMongoDB(ctx, mongoLocal, dbms.MongoDBSystem[mongoLocal]),

		RedisLocal: newRedis(ctx, redisLocal, dbms.RedisServer[redisLocal]),
		MariaLocal: newMariaDB(mariaLocal, dbms.MariaDBServer[mariaLocal]),
	}
}

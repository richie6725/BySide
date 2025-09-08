package database

import (
	"Byside/service/internal/config"
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

const (
	mongoLocal = "RichieMongo"
	redisLocal = "RichieRedis"
)

type BysideOut struct {
	dig.Out
	MongoLocal *mongo.Database `name:"mongo_byside"`
	RedisLocal *redis.Client   `name:"redis_byside"`
}

func NewByside(ctx context.Context, dbms config.DatabaseManageSystem) BysideOut {
	return BysideOut{
		MongoLocal: newMongoDB(ctx, mongoLocal, dbms.MongoDBSystem[mongoLocal]),

		RedisLocal: newRedis(ctx, redisLocal, dbms.RedisServer[redisLocal]),
	}
}

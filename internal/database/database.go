package database

import (
	"Richie_tester/internal/config"
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

const (
	mongoLocal = "RichieMongo"
	redisLocal = "RichieRedis"
)

type DBOut struct {
	dig.Out
	MongoLocal *mongo.Database
	RedisLocal *redis.Client
}

func NewDB(ctx context.Context, dbms config.DatabaseManage) DBOut {
	return DBOut{
		MongoLocal: newMongoDB(ctx, mongoLocal, dbms.MongoDBSystem[mongoLocal]),

		RedisLocal: newRedis(ctx, redisLocal, dbms.RedisServer[redisLocal]),
	}
}

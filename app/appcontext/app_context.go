package appcontext

import (
	"github.com/Uttamnath64/logger"
	"github.com/Uttamnath64/quixzap/app/config"
	"github.com/Uttamnath64/quixzap/app/storage"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext struct {
	Logger *logger.Logger
	MySQL  *config.MySQL
	Env    *config.AppEnv
	Mongo  *mongo.Database
	Redis  *storage.RedisClient
	Kafka  *kafka.Writer
}

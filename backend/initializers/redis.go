package initializers

import (
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/vinit-chauhan/tic-tac-toe/config"
	"github.com/vinit-chauhan/tic-tac-toe/utils/logger"
)

var RedisClient *redis.Client

func ConnectRedis(conf config.RedisConfig) {
	logger.Debug("connecting to Redis", "ConnectRedis")
	RedisClient = redis.NewClient(&redis.Options{
		Addr: conf.Host + ":" + strconv.Itoa(conf.Port),
	})

	logger.Debug("pinging Redis server", "ConnectRedis")
	_, err := RedisClient.Ping(RedisClient.Context()).Result()
	if err != nil {
		logger.Panic("unable to connect to Redis", "ConnectRedis", err)
	}
	logger.Debug("connected to Redis", "ConnectRedis")
}

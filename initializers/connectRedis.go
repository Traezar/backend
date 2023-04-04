package initializers

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	res, err := Redis.Ping(ctx).Result()
	println("REdis response " + res)
	if err != nil {
		println("jere")
		log.Fatal(err)
	}

}

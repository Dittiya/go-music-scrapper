package db

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v9"
)

type Redis struct {
	Client  *redis.Client
	Options *redis.Options
}

var ctx = context.Background()

func (r *Redis) InitRedis() {
	r.Client = redis.NewClient(r.Options)
	_, err := r.Client.Ping(ctx).Result()

	if err != nil {
		log.Fatal(err)
	}
}

func (r *Redis) Save(item string) (string, error) {
	fmt.Println("Saving", item)

	return "Saved", nil
}

func (r *Redis) Load(item string) (string, error) {
	fmt.Println("Loading", item)

	return "Loaded", nil
}

package db

import (
	"context"
	"log"
	"time"

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

func (r *Redis) Save(key string, item string, duration int) (string, error) {
	err := r.Client.Set(ctx, key, item, time.Duration(duration)*time.Second).Err()
	if err != nil {
		return "", err
	}

	return "Saved", nil
}

func (r *Redis) Load(key string) (string, error) {
	item, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return item, nil
}

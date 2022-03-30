package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedis(opt *redis.Options) (*redis.Client, error) {
	rdb := redis.NewClient(opt)
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

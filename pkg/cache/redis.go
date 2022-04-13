package cache

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var (
	once sync.Once
	Rdb  *redis.Client
	err  error
)

func SetupRedis(opt *redis.Options) (*redis.Client, error) {
	Rdb = new(redis.Client)
	once.Do(func() {
		Rdb = redis.NewClient(opt)
		err = Rdb.Ping(ctx).Err()
		if err != nil {
			panic(err)
		}
	})

	return Rdb, err
}

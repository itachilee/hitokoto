package logredis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

const (
	Prefix = "logredis:"
)

// HookConfig stores configuration needed to setup the hook
type HookConfig struct {
	Addr     string
	Password string
	DB       int    //redis dbnum
	Prefix   string // set redis prefix
}

// RedisHook to sends logs to Redis server
type RedisHook struct {
	RedisClient *redis.Client
}

// Levels returns the available logging levels.
func (hook *RedisHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.TraceLevel,
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (hook *RedisHook) Fire(e *logrus.Entry) error {
	hook.RedisClient.Set(ctx, "s", e.Message, 0)
	return nil
}

var ctx = context.Background()

func NewRedis(opt *redis.Options) (*redis.Client, error) {
	rdb := redis.NewClient(opt)
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

func DefaultConfig() *HookConfig {
	return &HookConfig{
		Addr:   "localhost:6379",
		DB:     0,
		Prefix: Prefix,
	}
}

// NewHook creates a hook to be added to an instance of logger
func NewHook(config HookConfig) (*RedisHook, error) {
	rdb, err := NewRedis(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	if err != nil {
		return nil, fmt.Errorf("unable to connect to REDIS: %s", err)
	}
	// defer rdb.Close()

	return &RedisHook{
		RedisClient: rdb,
	}, nil

}

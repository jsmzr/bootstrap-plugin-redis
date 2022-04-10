package redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/jsmzr/bootstrap-config/config"
	"github.com/jsmzr/bootstrap-plugin-redis/connection"
	"github.com/jsmzr/bootstrap-plugin/plugin"
)

type RedisPlugin struct{}

func (r *RedisPlugin) Order() int {
	return 100
}
func (r *RedisPlugin) Load() error {
	var options redis.Options
	if err := config.Resolve("bootstrap.redis", &options); err != nil {
		return err
	}
	instance := redis.NewClient(&options)
	pong, err := instance.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	if pong != "PONG" {
		return errors.New(pong)
	}
	connection.SetClient(instance)
	return nil
}

func init() {
	plugin.Register("redis", &RedisPlugin{})
}

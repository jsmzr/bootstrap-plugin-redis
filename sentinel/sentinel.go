package sentinel

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/jsmzr/bootstrap-config/config"
	"github.com/jsmzr/bootstrap-plugin-redis/connection"
	"github.com/jsmzr/bootstrap-plugin/plugin"
)

type RedisSentinelPlugin struct {
}

func (r *RedisSentinelPlugin) Order() int {
	return 100
}

func (r *RedisSentinelPlugin) Load() error {
	var options redis.FailoverOptions

	if err := config.Resolve("bootstrap.redis.sentinel", &options); err != nil {
		return err
	}
	instance := redis.NewFailoverClient(&options)
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
	plugin.Register("redis-sentinel", &RedisSentinelPlugin{})
}

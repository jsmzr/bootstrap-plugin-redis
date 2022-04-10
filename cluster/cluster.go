package cluster

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/jsmzr/bootstrap-config/config"
	"github.com/jsmzr/bootstrap-plugin-redis/connection"
	"github.com/jsmzr/bootstrap-plugin/plugin"
)

type RedisClusterPlugin struct{}

func (r *RedisClusterPlugin) Order() int {
	return 100
}

func (r *RedisClusterPlugin) Load() error {
	var options redis.ClusterOptions

	if err := config.Resolve("bootstrap.redis.cluster", &options); err != nil {
		return err
	}
	instance := redis.NewClusterClient(&options)
	pong, err := instance.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	if pong != "PONG" {
		return errors.New(pong)
	}
	connection.SetClusterClient(instance)
	return nil
}

func init() {
	plugin.Register("redis-cluster", &RedisClusterPlugin{})
}

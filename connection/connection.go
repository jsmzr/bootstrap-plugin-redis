package connection

import "github.com/go-redis/redis/v8"

// 集群模式下使用
var clusterClient *redis.ClusterClient

// 单机和 sentinel 模式下使用该客户端
var client *redis.Client

func GetClutserClient() *redis.ClusterClient {
	return clusterClient
}

func GetClient() *redis.Client {
	return client
}

func SetClient(instance *redis.Client) {
	client = instance
}

func SetClusterClient(instance *redis.ClusterClient) {
	clusterClient = instance
}

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var ctx = context.Background()

// 定义集群节点
var clusterAddrs = []string{
	"10.59.16.113:16379",
	"10.59.16.113:16380",
	"10.59.16.114:16379",
	"10.59.16.114:16380",
	"10.59.16.115:16379",
	"10.59.16.115:16380",
}

// 创建Redis Cluster客户端
func createRedisClusterClient() *redis.ClusterClient {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       clusterAddrs,
		Password:    "bvKXSvujqeNLPrJ._ZAa1",
		DialTimeout: 5 * time.Second,
	})

	// Ping一下，确保连接正常
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis cluster: %v", err)
	} else {
		fmt.Println("Connect to Redis cluster successfully")
	}

	return client
}

// SCAN查找特定前缀的键
func scanKeysWithPrefix(client *redis.ClusterClient, prefix string) []string {
	var cursor uint64
	var keys []string

	for {
		// SCAN操作
		var err error
		var newKeys []string
		newKeys, cursor, err = client.Scan(ctx, cursor, prefix+"*", 1000).Result()
		if err != nil {
			log.Fatalf("Error scanning keys: %v", err)
		}
		keys = append(keys, newKeys...)

		if cursor == 0 {
			break
		}
	}
	return keys
}

var isDelete bool

func init() {
	flag.BoolVar(&isDelete, "isDelete", false, "Delete all keys if set to true")
	flag.Parse()
}
func main() {
	client := createRedisClusterClient()
	defer client.Close()

	// 查询前缀为 s:cs:ext_cache: 的键
	fmt.Println("Scanning for keys with prefix 's:cs:ext_cache:'...")
	keys := scanKeysWithPrefix(client, "s:cs:ext_cache:")

	if len(keys) == 0 {
		fmt.Println("No keys found.")
	} else {
		fmt.Printf("Found %d keys:\n", len(keys))
	}
	if isDelete {
		//deleteKeysPipeline(client, keys)
		tempKeys := chunkKeys(keys, 10000)
		for _, v := range tempKeys {
			deleteKeysPipeline(client, v)
		}
	}
}

func chunkKeys(keys []string, chunkSize int) [][]string {
	var chunks [][]string
	for chunkSize < len(keys) {
		keys, chunks = keys[chunkSize:], append(chunks, keys[0:chunkSize:chunkSize])
	}
	return append(chunks, keys)
}

// 删除键（使用管道）
func deleteKeysPipeline(client *redis.ClusterClient, keys []string) {
	pipe := client.Pipeline()
	for _, key := range keys {
		pipe.Del(ctx, key)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Fatalf("Failed to delete keys in pipeline: %v", err)
	}
	fmt.Printf("Deleted %d keys using pipeline\n", len(keys))
}

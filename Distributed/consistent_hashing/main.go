package main

import (
	"blog-example/Distributed/consistent_hashing/chash"
	"blog-example/Distributed/consistent_hashing/hash"
	"blog-example/Distributed/consistent_hashing/kv_system"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	hash := hash.NewKVSystem(3, nil)
	runKvSystem(hash)
	cHash := chash.NewKVSystem(3)
	runKvSystem(cHash)
}

func runKvSystem(kv kv_system.KVSystem) {
	allKey := make([]string, 0)
	for i := 0; i < 100000; i++ {
		key := generateRandomString(10)
		value := generateRandomString(10)
		kv.Set(key, value)
		allKey = append(allKey, key)
	}
	fmt.Println(kv.CountKeys())
	nodes := kv.GetAllNode()
	kv.DeleteNode(nodes[0])
	fmt.Println(kv.CountKeys())
	for _, v := range allKey {
		_, ok := kv.Get(v)
		if !ok {
			fmt.Println("key not found")
		}
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

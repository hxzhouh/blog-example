package chash

import (
	"blog-example/Distributed/consistent_hashing/kvstorage"
	"fmt"
	"hash/crc32"
)

type KVSystem struct {
	hashRing *HashRing
	kvStores map[string]*kvstorage.KVStore
}

func NewKVSystem(nodes int) *KVSystem {
	hashRing := NewHashRing(crc32.ChecksumIEEE)
	for i := 0; i < nodes; i++ {
		node := &Node{
			ID:   fmt.Sprintf("Node%d", i),
			Addr: fmt.Sprintf("192.168.1.%d", i+1),
		}
		hashRing.AddNode(node)
	}
	kvStores := make(map[string]*kvstorage.KVStore)
	for id := range hashRing.Nodes {
		kvStores[id] = kvstorage.NewKVStore()
	}
	return &KVSystem{
		hashRing: hashRing,
		kvStores: kvStores,
	}
}

func (kv *KVSystem) Get(key string) (string, bool) {
	node := kv.hashRing.GetNode(key)
	return kv.kvStores[node.ID].Get(key)
}

func (kv *KVSystem) Set(key string, value string) {
	node := kv.hashRing.GetNode(key)
	kv.kvStores[node.ID].Set(key, value)
}

func (kv *KVSystem) Delete(key string) {
	node := kv.hashRing.GetNode(key)
	kv.kvStores[node.ID].Delete(key)
}
func (kv *KVSystem) DeleteNode(nodeID string) {
	allData := kv.kvStores[nodeID].GetAll()
	kv.hashRing.RemoveNode(nodeID)
	delete(kv.kvStores, nodeID)
	for key, value := range allData {
		kv.Set(key, value)
	}
}

func (kv *KVSystem) AddNode() {
	node := &Node{
		ID:   fmt.Sprintf("Node%d", len(kv.hashRing.Nodes)),
		Addr: fmt.Sprintf("192.168.1.%d", len(kv.hashRing.Nodes)+1),
	}
	kv.hashRing.AddNode(node)
	kv.kvStores[node.ID] = kvstorage.NewKVStore()
}

func (kv *KVSystem) GetAll() map[string]string {
	allData := make(map[string]string)
	for _, store := range kv.kvStores {
		for key, value := range store.GetAll() {
			allData[key] = value
		}
	}
	return allData
}

func (kv *KVSystem) CountKeys() map[string]int {
	result := make(map[string]int)
	for id := range kv.hashRing.Nodes {
		result[id] = len(kv.kvStores[id].GetAll())
	}
	return result
}

func (kv *KVSystem) GetAllNode() []string {
	nodes := make([]string, 0)
	for id := range kv.hashRing.Nodes {
		nodes = append(nodes, id)
	}
	return nodes
}

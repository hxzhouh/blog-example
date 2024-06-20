package hash

import (
	"blog-example/Distributed/consistent_hashing/kvstorage"
	"hash/crc32"
	"strconv"
)

type HashFunc func(data []byte) uint32

type KVSystem struct {
	kvStores map[string]*kvstorage.KVStore
	hash     HashFunc
}

func NewKVSystem(nodes int, hash HashFunc) *KVSystem {

	kvStores := make(map[string]*kvstorage.KVStore)
	if hash == nil {
		hash = crc32.ChecksumIEEE
	}
	for i := 0; i < nodes; i++ {
		kvStores[strconv.Itoa(i)] = kvstorage.NewKVStore()
	}
	return &KVSystem{
		kvStores: kvStores,
		hash:     hash,
	}
}

func (kv *KVSystem) Get(key string) (string, bool) {
	index := kv.hash([]byte(key)) % uint32(len(kv.kvStores))
	return kv.kvStores[strconv.Itoa(int(index))].Get(key)
}

func (kv *KVSystem) Set(key string, value string) {
	index := kv.hash([]byte(key)) % uint32(len(kv.kvStores))
	kv.kvStores[strconv.Itoa(int(index))].Set(key, value)
}

func (kv *KVSystem) Delete(key string) {
	index := kv.hash([]byte(key)) % uint32(len(kv.kvStores))
	kv.kvStores[strconv.Itoa(int(index))].Delete(key)
}

func (kv *KVSystem) DeleteNode(nodeID string) {

	allData := kv.GetAll()
	delete(kv.kvStores, nodeID)
	nodeLength := len(kv.kvStores)
	kv.kvStores = make(map[string]*kvstorage.KVStore)
	for i := 0; i < nodeLength; i++ {
		kv.kvStores[strconv.Itoa(i)] = kvstorage.NewKVStore()
	}
	for key, value := range allData {
		kv.Set(key, value)
	}
}

func (kv *KVSystem) AddNode() {
	index := uint32(len(kv.kvStores))
	kv.kvStores[strconv.Itoa(int(index))] = kvstorage.NewKVStore()
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
	for id := range kv.kvStores {
		result[id] = len(kv.kvStores[id].GetAll())
	}
	return result
}

func (kv *KVSystem) GetAllNode() []string {
	nodes := make([]string, 0)
	for key := range kv.kvStores {
		nodes = append(nodes, key)
	}
	return nodes
}

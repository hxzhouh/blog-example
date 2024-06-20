package kv_system

type KVSystem interface {
	Get(key string) (string, bool)
	Set(key string, value string)
	Delete(key string)
	DeleteNode(nodeID string)
	AddNode()
	GetAll() map[string]string
	CountKeys() map[string]int
	GetAllNode() []string
}

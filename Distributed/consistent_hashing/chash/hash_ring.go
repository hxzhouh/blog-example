package chash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

const (
	VirtualNodesPerNode = 2048
)

type HashFunc func(data []byte) uint32

type Node struct {
	ID   string
	Addr string
}

type VirtualNode struct {
	Hash uint32
	Node *Node
}

type HashRing struct {
	Nodes        map[string]*Node
	VirtualNodes []VirtualNode
	mu           sync.Mutex
	hash         HashFunc
}

func NewHashRing(hash HashFunc) *HashRing {
	if hash == nil {
		hash = crc32.ChecksumIEEE
	}
	return &HashRing{
		Nodes:        make(map[string]*Node),
		VirtualNodes: make([]VirtualNode, 0),
		hash:         hash,
	}
}

func (hr *HashRing) AddNode(node *Node) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	hr.Nodes[node.ID] = node
	for i := 0; i < VirtualNodesPerNode; i++ {
		virtualNodeID := fmt.Sprintf("%s-%d", node.ID, i)
		hash := hr.hash([]byte(virtualNodeID))
		hr.VirtualNodes = append(hr.VirtualNodes, VirtualNode{Hash: hash, Node: node})
	}
	sort.Slice(hr.VirtualNodes, func(i, j int) bool {
		return hr.VirtualNodes[i].Hash < hr.VirtualNodes[j].Hash
	})
}

func (hr *HashRing) RemoveNode(nodeID string) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	delete(hr.Nodes, nodeID)
	virtualNodes := make([]VirtualNode, 0)
	for _, vn := range hr.VirtualNodes {
		if vn.Node.ID != nodeID {
			virtualNodes = append(virtualNodes, vn)
		}
	}
	hr.VirtualNodes = virtualNodes
}

func (hr *HashRing) GetNode(key string) *Node {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	if len(hr.VirtualNodes) == 0 {
		return nil
	}

	hash := hr.hash([]byte(key))
	idx := sort.Search(len(hr.VirtualNodes), func(i int) bool {
		return hr.VirtualNodes[i].Hash >= hash
	})
	if idx == len(hr.VirtualNodes) {
		idx = 0
	}

	return hr.VirtualNodes[idx].Node
}

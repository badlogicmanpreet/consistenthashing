package consistenthashing

import (
	"hash/crc32"
	"sort"
)

// NewRing ... creates a new unit circle of nodes
func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}

// NewNode ... creates new node
func NewNode(id string) *Node {
	return &Node{ID: id, HashID: hashID(id)}
}

// AddNode ... adds new node to the ring
func (r *Ring) AddNode(id string) {
	node := NewNode(id)
	r.Nodes = append(r.Nodes, node)
	sort.Sort(r.Nodes)
}

// RemoveNode ... removes node from the ring
func (r *Ring) RemoveNode(id string) error {
	i := r.Search(id)

	r.Nodes = append(r.Nodes[:i], r.Nodes[i+1:]...)

	return nil
}

// Get ... gets the node for given id
func (r *Ring) Get(id string) string {
	i := r.Search(id)
	return r.Nodes[i].ID
}

// Search ... searches the ring for given id or the greater one
func (r *Ring) Search(id string) int {
	searchF := func(i int) bool {
		return r.Nodes[i].HashID >= hashID(id)
	}

	return sort.Search(r.Nodes.Len(), searchF)
}

func hashID(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

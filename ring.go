package consistenthashing

//Ring ... is a unit circle of nodes
type Ring struct {
	Nodes Nodes
}

// Nodes ... is an array of node(s)
type Nodes []*Node

// Node ... is the actual server with id(ip-address) and hashed id
type Node struct {
	ID     string
	HashID uint32
}

// Len ... gives number of nodes
func (n Nodes) Len() int {
	return len(n)
}

// Swap ... swap nodes
func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// Less ... compare nodes
func (n Nodes) Less(i, j int) bool {
	return n[i].HashID < n[j].HashID
}

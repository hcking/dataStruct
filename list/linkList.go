package list

// Node 单链表
type Node struct {
	data *Data
	next *Node
}

type Data struct {
	no   uint64
	name string
	age  int
}

func NewNode() *Node {
	l := &Node{}
	return l
}

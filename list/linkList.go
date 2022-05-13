package list

// Node 单链表
type ListNode struct {
	data *Data
	next *ListNode
}

type Data struct {
	no   uint64
	name string
	age  int
}

func NewListNode() *ListNode {
	l := &ListNode{}
	return l
}

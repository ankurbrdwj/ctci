package node

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func New() Node {
	return Node{
		value: nil,
		prev:  nil,
		next:  nil,
	}
}

func NewNode(data interface{}) Node {
	return Node{
		value: data,
		prev:  nil,
		next:  nil,
	}
}
func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) SetValue(data interface{}) {
	n.value = data
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) SetPrev(prev Node) {
	n.prev = &prev
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetNext(next Node) {
	n.next = &next
}

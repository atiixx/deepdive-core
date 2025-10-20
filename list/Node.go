package list

type Node[T any] struct {
	value T
	next  *Node[T]
}

func (n Node[T]) GetValue() T {
	return n.value
}

func (n *Node[T]) SetValue(value T) {
	n.value = value
}

func (n Node[T]) GetNext() *Node[T] {
	return n.next
}

func (n *Node[T]) ReplaceNext(node *Node[T]) {
	var temp *Node[T]
	if n.next != nil {
		temp.next = n.next.next
	}
	n.next = node
	node.next = temp
}
